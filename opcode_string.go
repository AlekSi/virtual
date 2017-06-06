// Code generated by "stringer -type Opcode enum.go"; DO NOT EDIT.

package virtual

import "fmt"

const _Opcode_name = "NopAPAddF32AddF64AddC64AddC128AddI32AddI64AddPtrAddPtrsAddSPAnd16And32And64And8ArgumentArgument16Argument32Argument64Argument8ArgumentsArgumentsFPBPBitfieldI8BitfieldI16BitfieldI32BitfieldI64BitfieldU8BitfieldU16BitfieldU32BitfieldU64BoolC128BoolF32BoolF64BoolI16BoolI32BoolI64BoolI8CallCallFPConvC64C128ConvF32C128ConvF32C64ConvF32F64ConvF32I32ConvF32I64ConvF32U32ConvF64C128ConvF64F32ConvF64I32ConvF64I64ConvF64I8ConvF64U16ConvF64U32ConvF64U64ConvI16I32ConvI16I64ConvI16U32ConvI32C128ConvI32C64ConvI32F32ConvI32F64ConvI32I16ConvI32I64ConvI32I8ConvI64ConvI64F64ConvI64I16ConvI64I32ConvI64I8ConvI64U16ConvI8I16ConvI8I32ConvI8I64ConvI8F64ConvI8U32ConvU16I32ConvU16I64ConvU16U32ConvU16U64ConvU32F32ConvU32F64ConvU32I16ConvU32I64ConvU32U8ConvU8I16ConvU8I32ConvU8U32ConvU8U64CopyCpl32Cpl64Cpl8DSDSC128DSI16DSI32DSI64DSI8DSNDivC128DivC64DivF32DivF64DivI32DivI64DivU32DivU64Dup32Dup64Dup8EqF32EqF64EqI32EqI64EqI8ExtFFIReturnFPField16Field64Field8FuncGeqF32GeqF64GeqI32GeqI64GeqI8GeqU32GeqU64GtF32GtF64GtI32GtI64GtU32GtU64IndexIndexI16IndexU16IndexI32IndexI64IndexI8IndexU32IndexU64IndexU8JmpJmpPJnzJzLabelLeqF32LeqF64LeqI32LeqI64LeqI8LeqU32LeqU64LoadLoad16Load32Load64Load8LshI16LshI32LshI64LshI8LtF32LtF64LtI32LtI64LtU32LtU64MulC128MulC64MulF32MulF64MulI32MulI64NegF32NegF64NegI16NegI32NegI64NegI8NegIndexI32NegIndexI64NegIndexU16NegIndexU32NegIndexU64NeqC128NeqC64NeqF32NeqF64NeqI32NeqI64NeqI8NotOr32Or64PanicPostIncF64PostIncI16PostIncI32PostIncI64PostIncI8PostIncPtrPostIncU32BitsPostIncU64BitsPreIncI16PreIncI32PreIncI64PreIncI8PreIncPtrPreIncU32BitsPreIncU64BitsPtrDiffPush16Push32Push64Push8PushC128RemI32RemI64RemU32RemU64ReturnRshI16RshI32RshI64RshI8RshU16RshU32RshU64RshU8StoreStore16Store32Store64Store8StoreBits16StoreBits32StoreBits64StoreBits8StoreC128StrNCopySubF32SubF64SubI32SubI64SubPtrsSwitchI32SwitchI64TextVariableVariable16Variable32Variable64Variable8Xor32Xor64Zero8Zero16Zero32Zero64__signbit__signbitfabortabsaccessacosallocaasinatanbswap64builtincallocceilcimagfclose_clrsbclrsblclrsbllclzclzlclzllcopysigncoscoshcrealfctzctzlctzlldlclosedlerrordlopendlsymerrno_locationexitexpfabsfchmodfchownfclosefcntlfflushffsffslffsllfgetcfgetsfloorfopen64fprintfframeAddressfreadfreefstat64fsyncftruncate64fwritegetcwdgetenvgeteuidgetpidgettimeofdayisinfisinffisinflisprintlocaltimeloglog10longjmplseek64lstat64mallocmemcmpmemcpymemmovemempcpymemsetmkdirmmap64munmapopen64parityparitylparityllpopcountpopcountlpopcountllpowprintfpthread_createpthread_equalpthread_joinpthread_mutex_destroypthread_mutex_initpthread_mutex_lockpthread_mutex_trylockpthread_mutex_unlockpthread_mutexattr_destroypthread_mutexattr_initpthread_mutexattr_settypepthread_selfqsortreadreadlinkreallocreturnAddressrmdirroundsetjmpsinsinhsleepsprintfsqrtstat64register_stdfilesstrcatstrchrstrcmpstrcpystrlenstrncmpstrncpystrrchrsysconftantanhtimetolowerunlinkusleeputimesvfprintfvprintfwrite_beginthreadex_endthreadex_msizeAreFileApisANSICloseHandleCreateFileMappingACreateFileMappingWCreateMutexWCreateFileACreateFileWDeleteCriticalSectionDeleteFileADeleteFileWEnterCriticalSectionFlushFileBuffersFlushViewOfFileFormatMessageAFormatMessageWFreeLibraryGetCurrentProcessIdGetCurrentThreadIdGetDiskFreeSpaceAGetDiskFreeSpaceWGetFileAttributesAGetFileAttributesWGetFileAttributesExWGetFileSizeGetFullPathNameAGetFullPathNameWGetLastErrorGetProcAddressGetProcessHeapGetSystemInfoGetSystemTimeGetSystemTimeAsFileTimeGetTempPathAGetTempPathWGetTickCountGetVersionExAGetVersionExWHeapAllocHeapCreateHeapCompactHeapDestroyHeapFreeHeapReAllocHeapSizeHeapValidateInitializeCriticalSectionInterlockedCompareExchangeLoadLibraryALoadLibraryWLocalFreeLockFileLockFileExLeaveCriticalSectionMapViewOfFileMultiByteToWideCharOutputDebugStringAOutputDebugStringWQueryPerformanceCounterReadFileSetEndOfFileSetFilePointerSleepSystemTimeToFileTimeUnlockFileUnlockFileExUnmapViewOfFileWaitForSingleObjectWaitForSingleObjectExWideCharToMultiByteWriteFile"

var _Opcode_index = [...]uint16{0, 3, 5, 11, 17, 23, 30, 36, 42, 48, 55, 60, 65, 70, 75, 79, 87, 97, 107, 117, 126, 135, 146, 148, 158, 169, 180, 191, 201, 212, 223, 234, 242, 249, 256, 263, 270, 277, 283, 287, 293, 304, 315, 325, 335, 345, 355, 365, 376, 386, 396, 406, 415, 425, 435, 445, 455, 465, 475, 486, 496, 506, 516, 526, 536, 545, 552, 562, 572, 582, 591, 601, 610, 619, 628, 637, 646, 656, 666, 676, 686, 696, 706, 716, 726, 735, 744, 753, 762, 771, 775, 780, 785, 789, 791, 797, 802, 807, 812, 816, 819, 826, 832, 838, 844, 850, 856, 862, 868, 873, 878, 882, 887, 892, 897, 902, 906, 909, 918, 920, 927, 934, 940, 944, 950, 956, 962, 968, 973, 979, 985, 990, 995, 1000, 1005, 1010, 1015, 1020, 1028, 1036, 1044, 1052, 1059, 1067, 1075, 1082, 1085, 1089, 1092, 1094, 1099, 1105, 1111, 1117, 1123, 1128, 1134, 1140, 1144, 1150, 1156, 1162, 1167, 1173, 1179, 1185, 1190, 1195, 1200, 1205, 1210, 1215, 1220, 1227, 1233, 1239, 1245, 1251, 1257, 1263, 1269, 1275, 1281, 1287, 1292, 1303, 1314, 1325, 1336, 1347, 1354, 1360, 1366, 1372, 1378, 1384, 1389, 1392, 1396, 1400, 1405, 1415, 1425, 1435, 1445, 1454, 1464, 1478, 1492, 1501, 1510, 1519, 1527, 1536, 1549, 1562, 1569, 1575, 1581, 1587, 1592, 1600, 1606, 1612, 1618, 1624, 1630, 1636, 1642, 1648, 1653, 1659, 1665, 1671, 1676, 1681, 1688, 1695, 1702, 1708, 1719, 1730, 1741, 1751, 1760, 1768, 1774, 1780, 1786, 1792, 1799, 1808, 1817, 1821, 1829, 1839, 1849, 1859, 1868, 1873, 1878, 1883, 1889, 1895, 1901, 1910, 1920, 1925, 1928, 1934, 1938, 1944, 1948, 1952, 1959, 1966, 1972, 1976, 1982, 1988, 1993, 1999, 2006, 2009, 2013, 2018, 2026, 2029, 2033, 2039, 2042, 2046, 2051, 2058, 2065, 2071, 2076, 2090, 2094, 2097, 2101, 2107, 2113, 2119, 2124, 2130, 2133, 2137, 2142, 2147, 2152, 2157, 2164, 2171, 2183, 2188, 2192, 2199, 2204, 2215, 2221, 2227, 2233, 2240, 2246, 2258, 2263, 2269, 2275, 2282, 2291, 2294, 2299, 2306, 2313, 2320, 2326, 2332, 2338, 2345, 2352, 2358, 2363, 2369, 2375, 2381, 2387, 2394, 2402, 2410, 2419, 2429, 2432, 2438, 2452, 2465, 2477, 2498, 2516, 2534, 2555, 2575, 2600, 2622, 2647, 2659, 2664, 2668, 2676, 2683, 2696, 2701, 2706, 2712, 2715, 2719, 2724, 2731, 2735, 2741, 2758, 2764, 2770, 2776, 2782, 2788, 2795, 2802, 2809, 2816, 2819, 2823, 2827, 2834, 2840, 2846, 2852, 2860, 2867, 2872, 2886, 2898, 2904, 2919, 2930, 2948, 2966, 2978, 2989, 3000, 3021, 3032, 3043, 3063, 3079, 3094, 3108, 3122, 3133, 3152, 3170, 3187, 3204, 3222, 3240, 3260, 3271, 3287, 3303, 3315, 3329, 3343, 3356, 3369, 3392, 3404, 3416, 3428, 3441, 3454, 3463, 3473, 3484, 3495, 3503, 3514, 3522, 3534, 3559, 3585, 3597, 3609, 3618, 3626, 3636, 3656, 3669, 3688, 3706, 3724, 3747, 3755, 3767, 3781, 3786, 3806, 3816, 3828, 3843, 3862, 3883, 3902, 3911}

func (i Opcode) String() string {
	if i < 0 || i >= Opcode(len(_Opcode_index)-1) {
		return fmt.Sprintf("Opcode(%d)", i)
	}
	return _Opcode_name[_Opcode_index[i]:_Opcode_index[i+1]]
}
