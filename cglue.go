package wren

/*
extern void f0(void* vm);
extern void f1(void* vm);
extern void f2(void* vm);
extern void f3(void* vm);
extern void f4(void* vm);
extern void f5(void* vm);
extern void f6(void* vm);
extern void f7(void* vm);
extern void f8(void* vm);
extern void f9(void* vm);
extern void f10(void* vm);
extern void f11(void* vm);
extern void f12(void* vm);
extern void f13(void* vm);
extern void f14(void* vm);
extern void f15(void* vm);
extern void f16(void* vm);
extern void f17(void* vm);
extern void f18(void* vm);
extern void f19(void* vm);
extern void f20(void* vm);
extern void f21(void* vm);
extern void f22(void* vm);
extern void f23(void* vm);
extern void f24(void* vm);
extern void f25(void* vm);
extern void f26(void* vm);
extern void f27(void* vm);
extern void f28(void* vm);
extern void f29(void* vm);
extern void f30(void* vm);
extern void f31(void* vm);
extern void f32(void* vm);
extern void f33(void* vm);
extern void f34(void* vm);
extern void f35(void* vm);
extern void f36(void* vm);
extern void f37(void* vm);
extern void f38(void* vm);
extern void f39(void* vm);
extern void f40(void* vm);
extern void f41(void* vm);
extern void f42(void* vm);
extern void f43(void* vm);
extern void f44(void* vm);
extern void f45(void* vm);
extern void f46(void* vm);
extern void f47(void* vm);
extern void f48(void* vm);
extern void f49(void* vm);
extern void f50(void* vm);
extern void f51(void* vm);
extern void f52(void* vm);
extern void f53(void* vm);
extern void f54(void* vm);
extern void f55(void* vm);
extern void f56(void* vm);
extern void f57(void* vm);
extern void f58(void* vm);
extern void f59(void* vm);
extern void f60(void* vm);
extern void f61(void* vm);
extern void f62(void* vm);
extern void f63(void* vm);
extern void f64(void* vm);
extern void f65(void* vm);
extern void f66(void* vm);
extern void f67(void* vm);
extern void f68(void* vm);
extern void f69(void* vm);
extern void f70(void* vm);
extern void f71(void* vm);
extern void f72(void* vm);
extern void f73(void* vm);
extern void f74(void* vm);
extern void f75(void* vm);
extern void f76(void* vm);
extern void f77(void* vm);
extern void f78(void* vm);
extern void f79(void* vm);
extern void f80(void* vm);
extern void f81(void* vm);
extern void f82(void* vm);
extern void f83(void* vm);
extern void f84(void* vm);
extern void f85(void* vm);
extern void f86(void* vm);
extern void f87(void* vm);
extern void f88(void* vm);
extern void f89(void* vm);
extern void f90(void* vm);
extern void f91(void* vm);
extern void f92(void* vm);
extern void f93(void* vm);
extern void f94(void* vm);
extern void f95(void* vm);
extern void f96(void* vm);
extern void f97(void* vm);
extern void f98(void* vm);
extern void f99(void* vm);
extern void f100(void* vm);
extern void f101(void* vm);
extern void f102(void* vm);
extern void f103(void* vm);
extern void f104(void* vm);
extern void f105(void* vm);
extern void f106(void* vm);
extern void f107(void* vm);
extern void f108(void* vm);
extern void f109(void* vm);
extern void f110(void* vm);
extern void f111(void* vm);
extern void f112(void* vm);
extern void f113(void* vm);
extern void f114(void* vm);
extern void f115(void* vm);
extern void f116(void* vm);
extern void f117(void* vm);
extern void f118(void* vm);
extern void f119(void* vm);
extern void f120(void* vm);
extern void f121(void* vm);
extern void f122(void* vm);
extern void f123(void* vm);
extern void f124(void* vm);
extern void f125(void* vm);
extern void f126(void* vm);
extern void f127(void* vm);

static inline void* get_f(int i) {
	switch (i) {
		case 0: return f0;
		case 1: return f1;
		case 2: return f2;
		case 3: return f3;
		case 4: return f4;
		case 5: return f5;
		case 6: return f6;
		case 7: return f7;
		case 8: return f8;
		case 9: return f9;
		case 10: return f10;
		case 11: return f11;
		case 12: return f12;
		case 13: return f13;
		case 14: return f14;
		case 15: return f15;
		case 16: return f16;
		case 17: return f17;
		case 18: return f18;
		case 19: return f19;
		case 20: return f20;
		case 21: return f21;
		case 22: return f22;
		case 23: return f23;
		case 24: return f24;
		case 25: return f25;
		case 26: return f26;
		case 27: return f27;
		case 28: return f28;
		case 29: return f29;
		case 30: return f30;
		case 31: return f31;
		case 32: return f32;
		case 33: return f33;
		case 34: return f34;
		case 35: return f35;
		case 36: return f36;
		case 37: return f37;
		case 38: return f38;
		case 39: return f39;
		case 40: return f40;
		case 41: return f41;
		case 42: return f42;
		case 43: return f43;
		case 44: return f44;
		case 45: return f45;
		case 46: return f46;
		case 47: return f47;
		case 48: return f48;
		case 49: return f49;
		case 50: return f50;
		case 51: return f51;
		case 52: return f52;
		case 53: return f53;
		case 54: return f54;
		case 55: return f55;
		case 56: return f56;
		case 57: return f57;
		case 58: return f58;
		case 59: return f59;
		case 60: return f60;
		case 61: return f61;
		case 62: return f62;
		case 63: return f63;
		case 64: return f64;
		case 65: return f65;
		case 66: return f66;
		case 67: return f67;
		case 68: return f68;
		case 69: return f69;
		case 70: return f70;
		case 71: return f71;
		case 72: return f72;
		case 73: return f73;
		case 74: return f74;
		case 75: return f75;
		case 76: return f76;
		case 77: return f77;
		case 78: return f78;
		case 79: return f79;
		case 80: return f80;
		case 81: return f81;
		case 82: return f82;
		case 83: return f83;
		case 84: return f84;
		case 85: return f85;
		case 86: return f86;
		case 87: return f87;
		case 88: return f88;
		case 89: return f89;
		case 90: return f90;
		case 91: return f91;
		case 92: return f92;
		case 93: return f93;
		case 94: return f94;
		case 95: return f95;
		case 96: return f96;
		case 97: return f97;
		case 98: return f98;
		case 99: return f99;
		case 100: return f100;
		case 101: return f101;
		case 102: return f102;
		case 103: return f103;
		case 104: return f104;
		case 105: return f105;
		case 106: return f106;
		case 107: return f107;
		case 108: return f108;
		case 109: return f109;
		case 110: return f110;
		case 111: return f111;
		case 112: return f112;
		case 113: return f113;
		case 114: return f114;
		case 115: return f115;
		case 116: return f116;
		case 117: return f117;
		case 118: return f118;
		case 119: return f119;
		case 120: return f120;
		case 121: return f121;
		case 122: return f122;
		case 123: return f123;
		case 124: return f124;
		case 125: return f125;
		case 126: return f126;
		case 127: return f127;
		default: return (void*)(0);
	}
}
*/
import "C"
import (
	"errors"
	"unsafe"
)

const MAX_REGISTRATIONS = 128

var (
	fMap = make(map[int]func())
	counter int
)


//export f0
func f0(vm unsafe.Pointer) {
	f := fMap[0]
	if f == nil {
		panic("function 0 not registered")
	}
	f()
}

//export f1
func f1(vm unsafe.Pointer) {
	f := fMap[1]
	if f == nil {
		panic("function 1 not registered")
	}
	f()
}

//export f2
func f2(vm unsafe.Pointer) {
	f := fMap[2]
	if f == nil {
		panic("function 2 not registered")
	}
	f()
}

//export f3
func f3(vm unsafe.Pointer) {
	f := fMap[3]
	if f == nil {
		panic("function 3 not registered")
	}
	f()
}

//export f4
func f4(vm unsafe.Pointer) {
	f := fMap[4]
	if f == nil {
		panic("function 4 not registered")
	}
	f()
}

//export f5
func f5(vm unsafe.Pointer) {
	f := fMap[5]
	if f == nil {
		panic("function 5 not registered")
	}
	f()
}

//export f6
func f6(vm unsafe.Pointer) {
	f := fMap[6]
	if f == nil {
		panic("function 6 not registered")
	}
	f()
}

//export f7
func f7(vm unsafe.Pointer) {
	f := fMap[7]
	if f == nil {
		panic("function 7 not registered")
	}
	f()
}

//export f8
func f8(vm unsafe.Pointer) {
	f := fMap[8]
	if f == nil {
		panic("function 8 not registered")
	}
	f()
}

//export f9
func f9(vm unsafe.Pointer) {
	f := fMap[9]
	if f == nil {
		panic("function 9 not registered")
	}
	f()
}

//export f10
func f10(vm unsafe.Pointer) {
	f := fMap[10]
	if f == nil {
		panic("function 10 not registered")
	}
	f()
}

//export f11
func f11(vm unsafe.Pointer) {
	f := fMap[11]
	if f == nil {
		panic("function 11 not registered")
	}
	f()
}

//export f12
func f12(vm unsafe.Pointer) {
	f := fMap[12]
	if f == nil {
		panic("function 12 not registered")
	}
	f()
}

//export f13
func f13(vm unsafe.Pointer) {
	f := fMap[13]
	if f == nil {
		panic("function 13 not registered")
	}
	f()
}

//export f14
func f14(vm unsafe.Pointer) {
	f := fMap[14]
	if f == nil {
		panic("function 14 not registered")
	}
	f()
}

//export f15
func f15(vm unsafe.Pointer) {
	f := fMap[15]
	if f == nil {
		panic("function 15 not registered")
	}
	f()
}

//export f16
func f16(vm unsafe.Pointer) {
	f := fMap[16]
	if f == nil {
		panic("function 16 not registered")
	}
	f()
}

//export f17
func f17(vm unsafe.Pointer) {
	f := fMap[17]
	if f == nil {
		panic("function 17 not registered")
	}
	f()
}

//export f18
func f18(vm unsafe.Pointer) {
	f := fMap[18]
	if f == nil {
		panic("function 18 not registered")
	}
	f()
}

//export f19
func f19(vm unsafe.Pointer) {
	f := fMap[19]
	if f == nil {
		panic("function 19 not registered")
	}
	f()
}

//export f20
func f20(vm unsafe.Pointer) {
	f := fMap[20]
	if f == nil {
		panic("function 20 not registered")
	}
	f()
}

//export f21
func f21(vm unsafe.Pointer) {
	f := fMap[21]
	if f == nil {
		panic("function 21 not registered")
	}
	f()
}

//export f22
func f22(vm unsafe.Pointer) {
	f := fMap[22]
	if f == nil {
		panic("function 22 not registered")
	}
	f()
}

//export f23
func f23(vm unsafe.Pointer) {
	f := fMap[23]
	if f == nil {
		panic("function 23 not registered")
	}
	f()
}

//export f24
func f24(vm unsafe.Pointer) {
	f := fMap[24]
	if f == nil {
		panic("function 24 not registered")
	}
	f()
}

//export f25
func f25(vm unsafe.Pointer) {
	f := fMap[25]
	if f == nil {
		panic("function 25 not registered")
	}
	f()
}

//export f26
func f26(vm unsafe.Pointer) {
	f := fMap[26]
	if f == nil {
		panic("function 26 not registered")
	}
	f()
}

//export f27
func f27(vm unsafe.Pointer) {
	f := fMap[27]
	if f == nil {
		panic("function 27 not registered")
	}
	f()
}

//export f28
func f28(vm unsafe.Pointer) {
	f := fMap[28]
	if f == nil {
		panic("function 28 not registered")
	}
	f()
}

//export f29
func f29(vm unsafe.Pointer) {
	f := fMap[29]
	if f == nil {
		panic("function 29 not registered")
	}
	f()
}

//export f30
func f30(vm unsafe.Pointer) {
	f := fMap[30]
	if f == nil {
		panic("function 30 not registered")
	}
	f()
}

//export f31
func f31(vm unsafe.Pointer) {
	f := fMap[31]
	if f == nil {
		panic("function 31 not registered")
	}
	f()
}

//export f32
func f32(vm unsafe.Pointer) {
	f := fMap[32]
	if f == nil {
		panic("function 32 not registered")
	}
	f()
}

//export f33
func f33(vm unsafe.Pointer) {
	f := fMap[33]
	if f == nil {
		panic("function 33 not registered")
	}
	f()
}

//export f34
func f34(vm unsafe.Pointer) {
	f := fMap[34]
	if f == nil {
		panic("function 34 not registered")
	}
	f()
}

//export f35
func f35(vm unsafe.Pointer) {
	f := fMap[35]
	if f == nil {
		panic("function 35 not registered")
	}
	f()
}

//export f36
func f36(vm unsafe.Pointer) {
	f := fMap[36]
	if f == nil {
		panic("function 36 not registered")
	}
	f()
}

//export f37
func f37(vm unsafe.Pointer) {
	f := fMap[37]
	if f == nil {
		panic("function 37 not registered")
	}
	f()
}

//export f38
func f38(vm unsafe.Pointer) {
	f := fMap[38]
	if f == nil {
		panic("function 38 not registered")
	}
	f()
}

//export f39
func f39(vm unsafe.Pointer) {
	f := fMap[39]
	if f == nil {
		panic("function 39 not registered")
	}
	f()
}

//export f40
func f40(vm unsafe.Pointer) {
	f := fMap[40]
	if f == nil {
		panic("function 40 not registered")
	}
	f()
}

//export f41
func f41(vm unsafe.Pointer) {
	f := fMap[41]
	if f == nil {
		panic("function 41 not registered")
	}
	f()
}

//export f42
func f42(vm unsafe.Pointer) {
	f := fMap[42]
	if f == nil {
		panic("function 42 not registered")
	}
	f()
}

//export f43
func f43(vm unsafe.Pointer) {
	f := fMap[43]
	if f == nil {
		panic("function 43 not registered")
	}
	f()
}

//export f44
func f44(vm unsafe.Pointer) {
	f := fMap[44]
	if f == nil {
		panic("function 44 not registered")
	}
	f()
}

//export f45
func f45(vm unsafe.Pointer) {
	f := fMap[45]
	if f == nil {
		panic("function 45 not registered")
	}
	f()
}

//export f46
func f46(vm unsafe.Pointer) {
	f := fMap[46]
	if f == nil {
		panic("function 46 not registered")
	}
	f()
}

//export f47
func f47(vm unsafe.Pointer) {
	f := fMap[47]
	if f == nil {
		panic("function 47 not registered")
	}
	f()
}

//export f48
func f48(vm unsafe.Pointer) {
	f := fMap[48]
	if f == nil {
		panic("function 48 not registered")
	}
	f()
}

//export f49
func f49(vm unsafe.Pointer) {
	f := fMap[49]
	if f == nil {
		panic("function 49 not registered")
	}
	f()
}

//export f50
func f50(vm unsafe.Pointer) {
	f := fMap[50]
	if f == nil {
		panic("function 50 not registered")
	}
	f()
}

//export f51
func f51(vm unsafe.Pointer) {
	f := fMap[51]
	if f == nil {
		panic("function 51 not registered")
	}
	f()
}

//export f52
func f52(vm unsafe.Pointer) {
	f := fMap[52]
	if f == nil {
		panic("function 52 not registered")
	}
	f()
}

//export f53
func f53(vm unsafe.Pointer) {
	f := fMap[53]
	if f == nil {
		panic("function 53 not registered")
	}
	f()
}

//export f54
func f54(vm unsafe.Pointer) {
	f := fMap[54]
	if f == nil {
		panic("function 54 not registered")
	}
	f()
}

//export f55
func f55(vm unsafe.Pointer) {
	f := fMap[55]
	if f == nil {
		panic("function 55 not registered")
	}
	f()
}

//export f56
func f56(vm unsafe.Pointer) {
	f := fMap[56]
	if f == nil {
		panic("function 56 not registered")
	}
	f()
}

//export f57
func f57(vm unsafe.Pointer) {
	f := fMap[57]
	if f == nil {
		panic("function 57 not registered")
	}
	f()
}

//export f58
func f58(vm unsafe.Pointer) {
	f := fMap[58]
	if f == nil {
		panic("function 58 not registered")
	}
	f()
}

//export f59
func f59(vm unsafe.Pointer) {
	f := fMap[59]
	if f == nil {
		panic("function 59 not registered")
	}
	f()
}

//export f60
func f60(vm unsafe.Pointer) {
	f := fMap[60]
	if f == nil {
		panic("function 60 not registered")
	}
	f()
}

//export f61
func f61(vm unsafe.Pointer) {
	f := fMap[61]
	if f == nil {
		panic("function 61 not registered")
	}
	f()
}

//export f62
func f62(vm unsafe.Pointer) {
	f := fMap[62]
	if f == nil {
		panic("function 62 not registered")
	}
	f()
}

//export f63
func f63(vm unsafe.Pointer) {
	f := fMap[63]
	if f == nil {
		panic("function 63 not registered")
	}
	f()
}

//export f64
func f64(vm unsafe.Pointer) {
	f := fMap[64]
	if f == nil {
		panic("function 64 not registered")
	}
	f()
}

//export f65
func f65(vm unsafe.Pointer) {
	f := fMap[65]
	if f == nil {
		panic("function 65 not registered")
	}
	f()
}

//export f66
func f66(vm unsafe.Pointer) {
	f := fMap[66]
	if f == nil {
		panic("function 66 not registered")
	}
	f()
}

//export f67
func f67(vm unsafe.Pointer) {
	f := fMap[67]
	if f == nil {
		panic("function 67 not registered")
	}
	f()
}

//export f68
func f68(vm unsafe.Pointer) {
	f := fMap[68]
	if f == nil {
		panic("function 68 not registered")
	}
	f()
}

//export f69
func f69(vm unsafe.Pointer) {
	f := fMap[69]
	if f == nil {
		panic("function 69 not registered")
	}
	f()
}

//export f70
func f70(vm unsafe.Pointer) {
	f := fMap[70]
	if f == nil {
		panic("function 70 not registered")
	}
	f()
}

//export f71
func f71(vm unsafe.Pointer) {
	f := fMap[71]
	if f == nil {
		panic("function 71 not registered")
	}
	f()
}

//export f72
func f72(vm unsafe.Pointer) {
	f := fMap[72]
	if f == nil {
		panic("function 72 not registered")
	}
	f()
}

//export f73
func f73(vm unsafe.Pointer) {
	f := fMap[73]
	if f == nil {
		panic("function 73 not registered")
	}
	f()
}

//export f74
func f74(vm unsafe.Pointer) {
	f := fMap[74]
	if f == nil {
		panic("function 74 not registered")
	}
	f()
}

//export f75
func f75(vm unsafe.Pointer) {
	f := fMap[75]
	if f == nil {
		panic("function 75 not registered")
	}
	f()
}

//export f76
func f76(vm unsafe.Pointer) {
	f := fMap[76]
	if f == nil {
		panic("function 76 not registered")
	}
	f()
}

//export f77
func f77(vm unsafe.Pointer) {
	f := fMap[77]
	if f == nil {
		panic("function 77 not registered")
	}
	f()
}

//export f78
func f78(vm unsafe.Pointer) {
	f := fMap[78]
	if f == nil {
		panic("function 78 not registered")
	}
	f()
}

//export f79
func f79(vm unsafe.Pointer) {
	f := fMap[79]
	if f == nil {
		panic("function 79 not registered")
	}
	f()
}

//export f80
func f80(vm unsafe.Pointer) {
	f := fMap[80]
	if f == nil {
		panic("function 80 not registered")
	}
	f()
}

//export f81
func f81(vm unsafe.Pointer) {
	f := fMap[81]
	if f == nil {
		panic("function 81 not registered")
	}
	f()
}

//export f82
func f82(vm unsafe.Pointer) {
	f := fMap[82]
	if f == nil {
		panic("function 82 not registered")
	}
	f()
}

//export f83
func f83(vm unsafe.Pointer) {
	f := fMap[83]
	if f == nil {
		panic("function 83 not registered")
	}
	f()
}

//export f84
func f84(vm unsafe.Pointer) {
	f := fMap[84]
	if f == nil {
		panic("function 84 not registered")
	}
	f()
}

//export f85
func f85(vm unsafe.Pointer) {
	f := fMap[85]
	if f == nil {
		panic("function 85 not registered")
	}
	f()
}

//export f86
func f86(vm unsafe.Pointer) {
	f := fMap[86]
	if f == nil {
		panic("function 86 not registered")
	}
	f()
}

//export f87
func f87(vm unsafe.Pointer) {
	f := fMap[87]
	if f == nil {
		panic("function 87 not registered")
	}
	f()
}

//export f88
func f88(vm unsafe.Pointer) {
	f := fMap[88]
	if f == nil {
		panic("function 88 not registered")
	}
	f()
}

//export f89
func f89(vm unsafe.Pointer) {
	f := fMap[89]
	if f == nil {
		panic("function 89 not registered")
	}
	f()
}

//export f90
func f90(vm unsafe.Pointer) {
	f := fMap[90]
	if f == nil {
		panic("function 90 not registered")
	}
	f()
}

//export f91
func f91(vm unsafe.Pointer) {
	f := fMap[91]
	if f == nil {
		panic("function 91 not registered")
	}
	f()
}

//export f92
func f92(vm unsafe.Pointer) {
	f := fMap[92]
	if f == nil {
		panic("function 92 not registered")
	}
	f()
}

//export f93
func f93(vm unsafe.Pointer) {
	f := fMap[93]
	if f == nil {
		panic("function 93 not registered")
	}
	f()
}

//export f94
func f94(vm unsafe.Pointer) {
	f := fMap[94]
	if f == nil {
		panic("function 94 not registered")
	}
	f()
}

//export f95
func f95(vm unsafe.Pointer) {
	f := fMap[95]
	if f == nil {
		panic("function 95 not registered")
	}
	f()
}

//export f96
func f96(vm unsafe.Pointer) {
	f := fMap[96]
	if f == nil {
		panic("function 96 not registered")
	}
	f()
}

//export f97
func f97(vm unsafe.Pointer) {
	f := fMap[97]
	if f == nil {
		panic("function 97 not registered")
	}
	f()
}

//export f98
func f98(vm unsafe.Pointer) {
	f := fMap[98]
	if f == nil {
		panic("function 98 not registered")
	}
	f()
}

//export f99
func f99(vm unsafe.Pointer) {
	f := fMap[99]
	if f == nil {
		panic("function 99 not registered")
	}
	f()
}

//export f100
func f100(vm unsafe.Pointer) {
	f := fMap[100]
	if f == nil {
		panic("function 100 not registered")
	}
	f()
}

//export f101
func f101(vm unsafe.Pointer) {
	f := fMap[101]
	if f == nil {
		panic("function 101 not registered")
	}
	f()
}

//export f102
func f102(vm unsafe.Pointer) {
	f := fMap[102]
	if f == nil {
		panic("function 102 not registered")
	}
	f()
}

//export f103
func f103(vm unsafe.Pointer) {
	f := fMap[103]
	if f == nil {
		panic("function 103 not registered")
	}
	f()
}

//export f104
func f104(vm unsafe.Pointer) {
	f := fMap[104]
	if f == nil {
		panic("function 104 not registered")
	}
	f()
}

//export f105
func f105(vm unsafe.Pointer) {
	f := fMap[105]
	if f == nil {
		panic("function 105 not registered")
	}
	f()
}

//export f106
func f106(vm unsafe.Pointer) {
	f := fMap[106]
	if f == nil {
		panic("function 106 not registered")
	}
	f()
}

//export f107
func f107(vm unsafe.Pointer) {
	f := fMap[107]
	if f == nil {
		panic("function 107 not registered")
	}
	f()
}

//export f108
func f108(vm unsafe.Pointer) {
	f := fMap[108]
	if f == nil {
		panic("function 108 not registered")
	}
	f()
}

//export f109
func f109(vm unsafe.Pointer) {
	f := fMap[109]
	if f == nil {
		panic("function 109 not registered")
	}
	f()
}

//export f110
func f110(vm unsafe.Pointer) {
	f := fMap[110]
	if f == nil {
		panic("function 110 not registered")
	}
	f()
}

//export f111
func f111(vm unsafe.Pointer) {
	f := fMap[111]
	if f == nil {
		panic("function 111 not registered")
	}
	f()
}

//export f112
func f112(vm unsafe.Pointer) {
	f := fMap[112]
	if f == nil {
		panic("function 112 not registered")
	}
	f()
}

//export f113
func f113(vm unsafe.Pointer) {
	f := fMap[113]
	if f == nil {
		panic("function 113 not registered")
	}
	f()
}

//export f114
func f114(vm unsafe.Pointer) {
	f := fMap[114]
	if f == nil {
		panic("function 114 not registered")
	}
	f()
}

//export f115
func f115(vm unsafe.Pointer) {
	f := fMap[115]
	if f == nil {
		panic("function 115 not registered")
	}
	f()
}

//export f116
func f116(vm unsafe.Pointer) {
	f := fMap[116]
	if f == nil {
		panic("function 116 not registered")
	}
	f()
}

//export f117
func f117(vm unsafe.Pointer) {
	f := fMap[117]
	if f == nil {
		panic("function 117 not registered")
	}
	f()
}

//export f118
func f118(vm unsafe.Pointer) {
	f := fMap[118]
	if f == nil {
		panic("function 118 not registered")
	}
	f()
}

//export f119
func f119(vm unsafe.Pointer) {
	f := fMap[119]
	if f == nil {
		panic("function 119 not registered")
	}
	f()
}

//export f120
func f120(vm unsafe.Pointer) {
	f := fMap[120]
	if f == nil {
		panic("function 120 not registered")
	}
	f()
}

//export f121
func f121(vm unsafe.Pointer) {
	f := fMap[121]
	if f == nil {
		panic("function 121 not registered")
	}
	f()
}

//export f122
func f122(vm unsafe.Pointer) {
	f := fMap[122]
	if f == nil {
		panic("function 122 not registered")
	}
	f()
}

//export f123
func f123(vm unsafe.Pointer) {
	f := fMap[123]
	if f == nil {
		panic("function 123 not registered")
	}
	f()
}

//export f124
func f124(vm unsafe.Pointer) {
	f := fMap[124]
	if f == nil {
		panic("function 124 not registered")
	}
	f()
}

//export f125
func f125(vm unsafe.Pointer) {
	f := fMap[125]
	if f == nil {
		panic("function 125 not registered")
	}
	f()
}

//export f126
func f126(vm unsafe.Pointer) {
	f := fMap[126]
	if f == nil {
		panic("function 126 not registered")
	}
	f()
}

//export f127
func f127(vm unsafe.Pointer) {
	f := fMap[127]
	if f == nil {
		panic("function 127 not registered")
	}
	f()
}


func registerFunc(name string, f func()) (unsafe.Pointer, error) {
	if (counter+1) >= MAX_REGISTRATIONS {
		return nil, errors.New("maximum function registration reached")
	}

	// TODO: make this thread-safe
	fMap[counter] = f
	ptr := C.get_f(C.int(counter))
	counter++
	return ptr, nil
}
