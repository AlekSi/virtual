// Code generated by "stringer -type Opcode enum.go"; DO NOT EDIT

package virtual

import "fmt"

const _Opcode_name = "NopAPAddF32AddF64AddC64AddC128AddI32AddI64AddPtrAddPtrsAddSPAnd16And32And64And8ArgumentArgument16Argument32Argument64Argument8ArgumentsArgumentsFPBPBitfieldI8BitfieldI16BitfieldI32BitfieldI64BitfieldU8BitfieldU16BitfieldU32BitfieldU64BoolC128BoolF32BoolF64BoolI16BoolI32BoolI64BoolI8CallCallFPConvC64C128ConvF32C64ConvF32C128ConvF32F64ConvF32I32ConvF32I64ConvF32U32ConvF64F32ConvF64C128ConvF64I32ConvF64I64ConvF64I8ConvF64U16ConvF64U32ConvF64U64ConvI16I32ConvI16I64ConvI16U32ConvI32C128ConvI32C64ConvI32F32ConvI32F64ConvI32I16ConvI32I64ConvI32I8ConvI64F64ConvI64I16ConvI64I32ConvI64I8ConvI64U16ConvI8I16ConvI8I32ConvI8I64ConvI8U32ConvU16I32ConvU16I64ConvU16U32ConvU32F32ConvU32F64ConvU32I16ConvU32I64ConvU32U8ConvU8I16ConvU8I32ConvU8U32ConvU8U64CopyCpl32Cpl64Cpl8DSDSC128DSI16DSI32DSI64DSI8DSNDivF32DivC64DivC128DivF64DivI32DivI64DivU32DivU64Dup32Dup64Dup8EqF32EqF64EqI32EqI64EqI8ExtFPField8Field16Field64FuncGeqF32GeqF64GeqI32GeqI64GeqU32GeqU64GtF32GtF64GtI32GtI64GtU32GtU64IndexIndexI16IndexI32IndexI64IndexU32IndexU64IndexI8IndexU8JmpJmpPJnzJzLabelLeqF32LeqF64LeqI32LeqI64LeqU32LeqU64LoadLoad16Load32Load64Load8LshI16LshI32LshI64LshI8LtF32LtF64LtI32LtI64LtU32LtU64MulC64MulC128MulF32MulF64MulI32MulI64NegF32NegF64NegI16NegI32NegI64NegI8NegIndexI32NegIndexI64NegIndexU64NeqC128NeqC64NeqF32NeqF64NeqI32NeqI64NotOr32Or64PanicPostIncF64PostIncI16PostIncI32PostIncI64PostIncI8PostIncPtrPostIncU32BitsPostIncU64BitsPreIncI16PreIncI32PreIncI64PreIncI8PreIncPtrPreIncU32BitsPreIncU64BitsPtrDiffPush8Push16Push32Push64PushC128RemI32RemI64RemU32RemU64ReturnRshI16RshI32RshI64RshI8RshU16RshU32RshU64RshU8StoreStore16Store32Store64StoreC128Store8StoreBits16StoreBits32StoreBits64StoreBits8StrNCopySubF32SubF64SubI32SubI64SubPtrsTextVariableVariable16Variable32Variable64Variable8Xor32Xor64Zero8Zero16Zero32Zero64abortabsacosallocaasinatanbswap64callocceilcimagfclrsbclrsblclrsbllclzclzlclzllcopysigncoscoshcrealfctzctzlctzllexitexpfabsfcloseffsffslffsllfgetcfgetsfloorfopenfprintfframeAddressfreadfreefwriteisinfisinffisinflisprintloglog10mallocmemcmpmemcpymemsetopenparityparitylparityllpopcountpopcountlpopcountllpowprintfreadreturnAddressroundsign_bitsign_bitfsinsinhsprintfsqrtstrcatstrchrstrcmpstrcpystrlenstrncmpstrncpystrrchrtantanhtolowervfprintfvprintfwrite"

var _Opcode_index = [...]uint16{0, 3, 5, 11, 17, 23, 30, 36, 42, 48, 55, 60, 65, 70, 75, 79, 87, 97, 107, 117, 126, 135, 146, 148, 158, 169, 180, 191, 201, 212, 223, 234, 242, 249, 256, 263, 270, 277, 283, 287, 293, 304, 314, 325, 335, 345, 355, 365, 375, 386, 396, 406, 415, 425, 435, 445, 455, 465, 475, 486, 496, 506, 516, 526, 536, 545, 555, 565, 575, 584, 594, 603, 612, 621, 630, 640, 650, 660, 670, 680, 690, 700, 709, 718, 727, 736, 745, 749, 754, 759, 763, 765, 771, 776, 781, 786, 790, 793, 799, 805, 812, 818, 824, 830, 836, 842, 847, 852, 856, 861, 866, 871, 876, 880, 883, 885, 891, 898, 905, 909, 915, 921, 927, 933, 939, 945, 950, 955, 960, 965, 970, 975, 980, 988, 996, 1004, 1012, 1020, 1027, 1034, 1037, 1041, 1044, 1046, 1051, 1057, 1063, 1069, 1075, 1081, 1087, 1091, 1097, 1103, 1109, 1114, 1120, 1126, 1132, 1137, 1142, 1147, 1152, 1157, 1162, 1167, 1173, 1180, 1186, 1192, 1198, 1204, 1210, 1216, 1222, 1228, 1234, 1239, 1250, 1261, 1272, 1279, 1285, 1291, 1297, 1303, 1309, 1312, 1316, 1320, 1325, 1335, 1345, 1355, 1365, 1374, 1384, 1398, 1412, 1421, 1430, 1439, 1447, 1456, 1469, 1482, 1489, 1494, 1500, 1506, 1512, 1520, 1526, 1532, 1538, 1544, 1550, 1556, 1562, 1568, 1573, 1579, 1585, 1591, 1596, 1601, 1608, 1615, 1622, 1631, 1637, 1648, 1659, 1670, 1680, 1688, 1694, 1700, 1706, 1712, 1719, 1723, 1731, 1741, 1751, 1761, 1770, 1775, 1780, 1785, 1791, 1797, 1803, 1808, 1811, 1815, 1821, 1825, 1829, 1836, 1842, 1846, 1852, 1857, 1863, 1870, 1873, 1877, 1882, 1890, 1893, 1897, 1903, 1906, 1910, 1915, 1919, 1922, 1926, 1932, 1935, 1939, 1944, 1949, 1954, 1959, 1964, 1971, 1983, 1988, 1992, 1998, 2003, 2009, 2015, 2022, 2025, 2030, 2036, 2042, 2048, 2054, 2058, 2064, 2071, 2079, 2087, 2096, 2106, 2109, 2115, 2119, 2132, 2137, 2145, 2154, 2157, 2161, 2168, 2172, 2178, 2184, 2190, 2196, 2202, 2209, 2216, 2223, 2226, 2230, 2237, 2245, 2252, 2257}

func (i Opcode) String() string {
	if i < 0 || i >= Opcode(len(_Opcode_index)-1) {
		return fmt.Sprintf("Opcode(%d)", i)
	}
	return _Opcode_name[_Opcode_index[i]:_Opcode_index[i+1]]
}
