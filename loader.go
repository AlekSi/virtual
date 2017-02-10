// Copyright 2017 The Virtual Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package virtual

import (
	"fmt"
	"go/token"

	"github.com/cznic/ir"
	"github.com/cznic/mathutil"
)

var (
	builtins   = map[ir.NameID]Opcode{}
	nonReturns = map[Opcode]struct{}{
		abort: {},
		exit:  {},
		Panic: {},
	}
)

func registerBuiltins(m map[int]Opcode) {
	for k, v := range m {
		nm := ir.NameID(k)
		if _, ok := builtins[nm]; ok {
			panic("internal error")
		}

		builtins[nm] = v
	}
}

type PCInfo struct {
	PC   int
	Line int
	C    int       // Column of # of arguments.
	Name ir.NameID // File name of func name.
}

func (l *PCInfo) Position() token.Position {
	return token.Position{Line: l.Line, Column: l.C, Filename: string(dict.S(int(l.Name)))}
}

type Binary struct {
	BSS       int
	Code      []Operation
	Data      []byte
	Functions []PCInfo
	Lines     []PCInfo
	Model     string
	Text      []byte
}

func newBinary(model string) *Binary {
	return &Binary{Model: model}
}

type nfo struct {
	off int
	sz  int
}

type loader struct {
	bss        int
	m          map[int]int // Object #: {BSS,Code,Data,Text} index.
	model      ir.MemoryModel
	objects    []ir.Object
	out        *Binary
	prev       Operation
	stackAlign int
	strings    map[ir.StringID]int
	tc         ir.TypeCache
}

func newLoader(modelName string, objects []ir.Object) *loader {
	model, ok := ir.MemoryModels[modelName]
	if !ok {
		panic(fmt.Errorf("unknown memory model %q", modelName))
	}

	ptrItem, ok := model[ir.Pointer]
	if !ok {
		panic(fmt.Errorf("invalid memory model %q, missing item for pointer", modelName))
	}

	return &loader{
		m:          map[int]int{},
		model:      model,
		objects:    objects,
		out:        newBinary(modelName),
		prev:       Operation{Opcode: -1},
		stackAlign: int(ptrItem.Align),
		strings:    map[ir.StringID]int{},
		tc:         ir.TypeCache{},
	}
}

func (l *loader) loadDataDefinition(d *ir.DataDefinition) int {
	switch {
	case d.Value != nil:
		panic("TODO")
	default:
		panic("TODO")
	}
}

func (l *loader) emitOne(op Operation) {
	prev := l.prev
	if _, ok := nonReturns[prev.Opcode]; ok {
		switch op.Opcode {
		case Func, Label:
		default:
			return
		}
	}

	l.prev = op
	switch op.Opcode {
	case AddSP:
		if prev.Opcode == AddSP {
			i := len(l.out.Code) - 1
			l.out.Code[i].N += op.N
			if l.out.Code[i].N == 0 {
				l.out.Code = l.out.Code[:i]
			}
			break
		}

		l.out.Code = append(l.out.Code, op)
	case Return:
		switch {
		case prev.Opcode == AddSP:
			l.out.Code[len(l.out.Code)-1] = op
		default:
			l.out.Code = append(l.out.Code, op)
		}
	case Label:
		// nop
	default:
		l.out.Code = append(l.out.Code, op)
	}
}

func (l *loader) emit(li PCInfo, op ...Operation) {
	if li.Line != 0 {
		li.C = 1
		if n := len(l.out.Lines); n == 0 || l.out.Lines[n-1].Line != li.Line || l.out.Lines[n-1].Name != li.Name {
			l.out.Lines = append(l.out.Lines, li)
		}
	}
	for _, v := range op {
		l.emitOne(v)
	}
}

func (l *loader) sizeof(tid ir.TypeID) int {
	sz := l.model.Sizeof(l.tc.MustType(tid))
	if sz > mathutil.MaxInt {
		panic(fmt.Errorf("size of %s out of limits", tid))
	}

	return int(sz)
}

func (l *loader) stackSize(tid ir.TypeID) int { return roundup(l.sizeof(tid), l.stackAlign) }

func (l *loader) text(b []byte) int {
	p := len(l.out.Text)
	l.out.Text = append(l.out.Text, b...)
	sz := roundup(len(l.out.Text)+1, mallocAlign)
	l.out.Text = append(l.out.Text, make([]byte, sz-len(l.out.Text))...)
	return p
}

func (l *loader) pos(op ir.Operation) PCInfo {
	p := op.Pos()
	if !p.IsValid() {
		return PCInfo{}
	}

	return PCInfo{PC: len(l.out.Code), Line: p.Line, C: p.Column, Name: ir.NameID(dict.SID(p.Filename))}
}

func (l *loader) ip() int { return len(l.out.Code) }

func (l *loader) loadFunctionDefinition(f *ir.FunctionDefinition) {
	var (
		arguments []nfo
		calls     []int
		labels    = map[int]int{}
		results   []nfo
		variables []nfo
	)

	t := l.tc.MustType(f.TypeID).(*ir.FunctionType)
	for _, v := range t.Arguments {
		arguments = append(arguments, nfo{sz: l.sizeof(v.ID())})
	}
	off := 0
	for i := range arguments {
		off -= roundup(arguments[i].sz, l.stackAlign)
		arguments[i].off = off
	}

	for _, v := range t.Results {
		results = append(results, nfo{sz: l.sizeof(v.ID())})
	}
	off = 0
	for i := len(results) - 1; i >= 0; i-- {
		results[i].off = off
		off += roundup(results[i].sz, stackAlign)
	}

	for _, v := range f.Body {
		switch x := v.(type) {
		case *ir.VariableDeclaration:
			variables = append(variables, nfo{sz: l.sizeof(x.TypeID)})
		}
	}
	off = 0
	for i := range variables {
		off -= roundup(variables[i].sz, l.stackAlign)
		variables[i].off = off
	}

	n := 0
	if m := len(variables); m != 0 {
		n = variables[m-1].off
	}
	fp := f.Position
	fi := PCInfo{PC: len(l.out.Code), Line: fp.Line, C: len(arguments), Name: ir.NameID(f.NameID)}
	l.out.Functions = append(l.out.Functions, fi)
	l.emit(l.pos(f.Body[0]), Operation{Opcode: Func, N: n})
	ip0 := l.ip()
	for ip, v := range f.Body {
		switch x := v.(type) {
		case *ir.Add:
			switch l.sizeof(x.TypeID) {
			case 4:
				l.emit(l.pos(x), Operation{Opcode: AddI32})
			default:
				panic(fmt.Errorf("internal error %s", x.TypeID))
			}
		case *ir.AllocResult:
			l.emit(l.pos(x), Operation{Opcode: AddSP, N: -l.stackSize(x.TypeID)})
		case *ir.Argument:
			switch {
			case x.Address:
				panic("TODO")
			default:
				switch val := arguments[x.Index]; val.sz {
				case 4:
					l.emit(l.pos(x), Operation{Opcode: Argument32, N: val.off})
				case 8:
					l.emit(l.pos(x), Operation{Opcode: Argument64, N: val.off})
				default:
					panic(fmt.Errorf("internal error %v %v", t.Arguments[x.Index].ID(), val))
				}
			}
		case *ir.Arguments:
			l.emit(l.pos(x), Operation{Opcode: Arguments})
		case *ir.BeginScope:
			// nop
		case *ir.Call:
			fn := calls[len(calls)-1]
			calls = calls[:len(calls)-1]
			if fn < 0 { // fn ptr
				panic("TODO")
				break
			}

			if opcode, ok := builtins[l.objects[fn].(*ir.FunctionDefinition).NameID]; ok {
				l.emit(l.pos(x), Operation{Opcode: opcode})
				break
			}

			l.emit(l.pos(x), Operation{Opcode: Call, N: fn})
		case *ir.Drop:
			l.emit(l.pos(x), Operation{Opcode: AddSP, N: l.stackSize(x.TypeID)})
		case *ir.Dup:
			switch l.sizeof(x.TypeID) {
			case 4:
				l.emit(l.pos(x), Operation{Opcode: Dup32})
			default:
				panic(fmt.Errorf("internal error %s", x.TypeID))
			}
		case *ir.Element:
			t := l.tc.MustType(x.TypeID).(*ir.PointerType).Element
			sz := l.sizeof(t.ID())
			xt := l.tc.MustType(x.IndexType)
			switch xt.Kind() {
			case ir.Int32:
				l.emit(l.pos(x), Operation{Opcode: IndexI32, N: sz})
			default:
				panic(fmt.Errorf("TODO %v", xt.Kind()))
			}
			if !x.Address {
				switch sz {
				case 4:
					l.emit(l.pos(x), Operation{Opcode: Load32})
				default:
					panic(fmt.Errorf("TODO %v", sz))
				}
			}
		case *ir.EndScope:
			// nop
		case *ir.Eq:
			switch t := l.tc.MustType(x.TypeID); t.Kind() {
			case ir.Int32:
				l.emit(l.pos(x), Operation{Opcode: EqI32})
			default:
				panic(fmt.Errorf("TODO %v", t.Kind()))
			}
		case *ir.Extern:
			switch ex := l.objects[x.Index].(type) {
			case *ir.FunctionDefinition:
				if !x.Address {
					panic(fmt.Errorf("invalid IR"))
				}
				calls = append(calls, x.Index)
			default:
				panic(fmt.Errorf("TODO %T(%v)", ex, ex))
			}
		case *ir.Field:
			fields := l.model.Layout(l.tc.MustType(x.TypeID).(*ir.PointerType).Element.(*ir.StructOrUnionType))
			switch {
			case x.Address:
				l.emit(l.pos(x), Operation{Opcode: AddPtr, N: int(fields[x.Index].Offset)})
			default:
				switch fields[x.Index].Size {
				case 4:
					l.emit(l.pos(x), Operation{Opcode: Load32, N: int(fields[x.Index].Offset)})
				default:
					panic(fmt.Errorf("TODO %v", fields[x.Index].Size))
				}
			}
		case *ir.Int32Const:
			l.emit(l.pos(x), Operation{Opcode: Int32, N: int(x.Value)})
		case *ir.Jmp:
			n := int(x.NameID)
			if n == 0 {
				n = x.Number
			}
			l.emit(l.pos(x), Operation{Opcode: Jmp, N: n})
		case *ir.Jnz:
			n := int(x.NameID)
			if n == 0 {
				n = x.Number
			}
			l.emit(l.pos(x), Operation{Opcode: Jnz, N: n})
		case *ir.Jz:
			n := int(x.NameID)
			if n == 0 {
				n = x.Number
			}
			l.emit(l.pos(x), Operation{Opcode: Jz, N: n})
		case *ir.Label:
			n := -int(x.NameID)
			if n == 0 {
				n = x.Number
			}
			labels[n] = len(l.out.Code)
			l.emit(l.pos(x), Operation{Opcode: Label, N: n})
		case *ir.Leq:
			switch t := l.tc.MustType(x.TypeID); t.Kind() {
			case ir.Int32:
				l.emit(l.pos(x), Operation{Opcode: LeqI32})
			default:
				panic(fmt.Errorf("TODO %v", t.Kind()))
			}
		case *ir.Lt:
			switch t := l.tc.MustType(x.TypeID); t.Kind() {
			case ir.Int32:
				l.emit(l.pos(x), Operation{Opcode: LtI32})
			default:
				panic(fmt.Errorf("TODO %v", t.Kind()))
			}
		case *ir.Mul:
			switch l.sizeof(x.TypeID) {
			case 4:
				l.emit(l.pos(x), Operation{Opcode: MulI32})
			default:
				panic(fmt.Errorf("internal error %s", x.TypeID))
			}
		case *ir.Panic:
			l.emit(l.pos(x), Operation{Opcode: Panic})
		case *ir.PostIncrement:
			switch t := l.tc.MustType(x.TypeID); t.Kind() {
			case ir.Int32:
				l.emit(l.pos(x), Operation{Opcode: PostIncI32})
			default:
				panic(fmt.Errorf("TODO %v", t.Kind()))
			}
		case *ir.Result:
			var r nfo
			switch {
			case len(results) == 0 && x.Index == 0:
				// nop
			default:
				r = results[x.Index]
			}
			switch {
			case x.Address:
				l.emit(l.pos(x), Operation{Opcode: AP, N: r.off})
			default:
				panic("TODO")
			}
		case *ir.Return:
			l.emit(l.pos(x), Operation{Opcode: Return})
		case *ir.Store:
			switch l.sizeof(x.TypeID) {
			case 4:
				l.emit(l.pos(x), Operation{Opcode: Store32})
			default:
				panic(fmt.Errorf("internal error %s", x.TypeID))
			}
		case *ir.StringConst:
			p, ok := l.strings[x.Value]
			if !ok {
				p = l.text(dict.S(int(x.Value)))
				l.strings[x.Value] = p
			}
			l.emit(l.pos(x), Operation{Opcode: Text, N: p})
		case *ir.Sub:
			switch l.sizeof(x.TypeID) {
			case 4:
				l.emit(l.pos(x), Operation{Opcode: SubI32})
			default:
				panic(fmt.Errorf("internal error %s", x.TypeID))
			}
		case *ir.Variable:
			switch {
			case x.Address:
				l.emit(l.pos(x), Operation{Opcode: BP, N: variables[x.Index].off})
			default:
				switch val := variables[x.Index]; val.sz {
				case 4:
					l.emit(l.pos(x), Operation{Opcode: Variable32, N: val.off})
				default:
					panic(fmt.Errorf("internal error %v", val))
				}
			}
		case *ir.VariableDeclaration:
			switch v := x.Value.(type) {
			case nil:
				// nop
			case *ir.Int32Value:
				l.emit(l.pos(x),
					Operation{Opcode: BP, N: variables[x.Index].off},
					Operation{Opcode: Int32, N: int(v.Value)},
					Operation{Opcode: Store32},
				)
			default:
				panic(fmt.Errorf("TODO %T(%v)", v, v))
			}
		default:
			panic(fmt.Errorf("TODO %T\n\t%#05x\t%v", x, ip, x))
		}
	}
	for i, v := range l.out.Code[ip0:] {
		switch v.Opcode {
		case Jmp, Jnz, Jz:
			l.out.Code[ip0+i].N = labels[v.N]
		}
	}
}

func (l *loader) load() {
	for i, v := range l.objects {
		switch x := v.(type) {
		case *ir.DataDefinition:
			l.m[i] = l.loadDataDefinition(x)
		}
	}
	for i, v := range l.objects {
		switch x := v.(type) {
		case *ir.FunctionDefinition:
			if _, ok := builtins[x.NameID]; ok {
				break
			}

			l.m[i] = len(l.out.Code)
			l.loadFunctionDefinition(x)
		}
	}
	for i, v := range l.out.Code {
		switch v.Opcode {
		case Call:
			l.out.Code[i].N = l.m[v.N]
		}
	}
}

func Load(model string, objects []ir.Object) (_ *Binary, err error) {
	if !Testing {
		defer func() {
			switch x := recover().(type) {
			case nil:
				// nop
			default:
				err = fmt.Errorf("Load: %v", x)
			}
		}()
	}

	l := newLoader(model, objects)
	l.load()
	return l.out, nil
}
