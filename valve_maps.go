// all the lookup tables for tubes.go where they don't clutter up the source

package main

// source: https://en.wikipedia.org/wiki/Mullard%E2%80%93Philips_tube_designation

var ph_heater = map[string]string{
	"A": "4V AC",
	"B": "180mA DC",
	"C": "200mA AC/DC",
	"D": "1.4V or 1.4V/2.8V",
	"E": "6.3V or 6.3/12.6V",
	"F": "12.6V",
	"G": "5V or other",
	"H": "150mA AC/DC",
	"I": "20V",
	"K": "2V DC",
	"L": "450mA AC/DC",
	"M": "1.9V directly heated",
	"N": "12.6V indirectly heated",
	"O": "Cold cathode",
	"P": "300mA AC/DC",
	"Q": "2.4V indirectly heated",
	"S": "1.9V indirectly heated",
	"T": "7.4V",
	"U": "100mA AC/DC",
	"V": "50mA AC/DC",
	"X": "600mA AC/DC",
	"Y": "450mA AC/DC",
	"Z": "Cold cathode",
}

var ph_device = map[string]string{
	"A": "Low-current diode",
	"B": "Low-current double diode with common cathode",
	"C": "Small-signal triode",
	"D": "Power triode",
	"E": "Small-signal tetrode",
	"F": "Small-signal pentode",
	"H": "Hexode or Heptode (of the Hexode type)",
	"K": "Octode or Heptode (of the Octode type)",
	"L": "Power tetrode/pentode",
	"M": "Tuning indicator",
	"N": "Gas-filled triode or thyratron",
	"P": "Tube designed for secondary emission",
	"Q": "Nonode",
	"S": "TV sync oscillator",
	"T": "Beam tube, or miscellaneous",
	"W": "Gas-filled halfwave rectifier",
	"X": "Gas-filled fullwave rectifier",
	"Y": "Halfwave rectifier",
	"Z": "Fullwave rectifier",
}

var ph_base = map[int]string{
	0:  "Pinch-type",
	1:  "G8A",
	2:  "B8G loctal (mostly)",
	3:  "K8A octal",
	4:  "B8A rimlok",
	5:  "various",
	6:  "B9G or subminiature",
	7:  "Loctal Lorenz subminiatures",
	8:  "B9A noval",
	9:  "B7G button",
	10: "B7G (German)",
	11: "B8A rimlok or 8-pin German octal",
	13: "Octal",
	15: "German 10-pin with spigot or octal",
	16: "Flat wire subminiature or 8-pin German octal",
	17: "RFT 8-pin or RFT 11-pin",
	18: "B9A noval",
	19: "B7G",
	20: "B10B decal",
	23: "Octal",
	27: "RFT",
	28: "B9A noval",
}

var ph_z = map[string]string{
	"A": "Long-life amplifier tube",
	"B": "Binary counter or switching tube",
	"C": "Common-cathode Counter Dekatron",
	"E": "Electrometer tube",
	"G": "Amplifier tube",
	"M": "Optical indicator",
	"S": "Separate-cathode Counter/Selector Dekatron",
	"T": "Relay triode (a low-power triode thyratron)",
	"U": "Low-power tetrode thyratron",
}

// source: https://en.wikipedia.org/wiki/List_of_Mullard%E2%80%93Philips_vacuum_tubes

var other_to_ph = map[string]string{
	"10CW5":   "LL86",
	"10DX8":   "LCL84",
	"10GV8":   "LCL85",
	"11Y9":    "LFL200",
	"12AL5":   "HAA91",
	"12AT6":   "HBC90",
	"12AT7":   "ECC81",
	"12AU6":   "HF94",
	"12AU7":   "ECC82",
	"12AV6":   "HBC91",
	"12AX7":   "ECC83",
	"12BA6":   "HF93",
	"12BY7":   "EL180",
	"12DW7":   "ECC832",
	"12FG6":   "UM84",
	"12HU8":   "PLL80",
	"12S7":    "UAF42",
	"13CM5":   "XL36",
	"13GB5":   "XL500",
	"14GW8":   "PCL86",
	"14K7":    "UCH42",
	"15A6":    "PL83",
	"15CW5":   "PL84",
	"15DQ8":   "PCL84",
	"16A5":    "PL82",
	"16A8":    "PCL82",
	"16AQ3":   "XY88",
	"16Y9":    "PFL200",
	"17C8":    "UBF80",
	"17EW8":   "HCC85",
	"17KW6":   "PL508",
	"18GB5":   "LL500",
	"18GV8":   "PCL85",
	"19BR5":   "UM80",
	"19BX6":   "UF80",
	"19BY7":   "UF85",
	"19D8":    "UCH81",
	"19FL8":   "UBF89",
	"19T8":    "HABC80",
	"1A7":     "DK32",
	"1AB6":    "DK96",
	"1AC6":    "DK92",
	"1AD4":    "DF62",
	"1AH5":    "DAF96",
	"1AJ4":    "DF96",
	"1AN5":    "DF97",
	"1E3":     "DC80",
	"1L4":     "DF92",
	"1M3":     "DM70",
	"1N3":     "DM71",
	"1R5":     "DK91",
	"1S2":     "DY86",
	"1S2A":    "DY87",
	"1S4":     "DL91",
	"1S5":     "DAF91",
	"1T4":     "DF91",
	"1U4":     "DF904",
	"1V6":     "DCF60",
	"20AQ3":   "LY88",
	"21A6":    "PL81",
	"25E5":    "PL36",
	"27AK8":   "UABC80",
	"28GB5":   "PL500",
	"29KQ6":   "PL521",
	"2B35":    "EA50",
	"2D21":    "EN91",
	"30A5":    "HL94",
	"30AE3":   "PY88",
	"31A3":    "UY41",
	"35W4":    "HY90",
	"38A3":    "UY85",
	"3A4":     "DL93",
	"3AL5":    "XAA91",
	"3AU6":    "XF94",
	"3B4":     "DL98",
	"3C4":     "DL96",
	"3D6":     "DL29",
	"3EH7":    "XF183",
	"3EJ7":    "XF184",
	"3Q4":     "DL95",
	"3Q5GT":   "DL33",
	"3S4":     "DL92",
	"3V4":     "DL94",
	"408A":    "EF95",
	"40KG6A":  "PL509",
	"42EC4A":  "PY500A",
	"45A5":    "UL41",
	"45B5":    "UL84",
	"4652":    "AX1",
	"4654":    "EL50",
	"4671":    "E1C",
	"4672":    "E1F",
	"4677":    "AM2",
	"4686":    "AC50",
	"4695":    "E2F",
	"4696":    "EE1",
	"4BL8":    "XCF80",
	"4CM4":    "PC86",
	"4DL4":    "PC88",
	"4EH7":    "LF183",
	"4EJ7":    "LF184",
	"4ER5":    "PC95",
	"4ES8":    "XCC189",
	"4FY5":    "PC97",
	"4GJ7":    "XCF801",
	"4HA5":    "PC900",
	"50BM8":   "UCL82",
	"50C5":    "HL92",
	"55N3":    "UY82",
	"5636":    "EF730",
	"5654":    "EF95",
	"5678":    "DF60",
	"5718":    "EC71",
	"5726":    "EAA901S",
	"5823":    "Z900T",
	"5861":    "EC55",
	"5899":    "EF71",
	"5902":    "EL71",
	"5910":    "DF904",
	"5915":    "EH900S",
	"5920":    "E90CC",
	"5AQ4":    "GZ32",
	"5AR4":    "GZ34",
	"5ES8":    "YCC189",
	"5GJ7":    "LCF801",
	"5HG8":    "LCF86",
	"5U9":     "LCF201",
	"5V4":     "GZ32",
	"5Z4-G":   "GZ30",
	"6007":    "DL67",
	"6008":    "DF67",
	"6021":    "ECC70",
	"6057":    "ECC83",
	"6064":    "EF91",
	"6067":    "ECC82",
	"6080":    "ECC230",
	"6084":    "E80F",
	"6085":    "ECC87",
	"6132":    "EL821",
	"6189":    "ECC802S",
	"6201":    "ECC801S",
	"6218":    "E80T",
	"6227":    "E80L",
	"6267":    "EF86",
	"62VP":    "EF41",
	"6375":    "DC70",
	"6574":    "EN32",
	"6681":    "E83CC",
	"6686":    "E81L",
	"6687":    "E91H",
	"6688":    "E180F",
	"6689":    "E83F",
	"6778":    "EC70",
	"6779":    "Z803U",
	"6922":    "E88CC",
	"6923":    "EA52",
	"6977":    "DM160",
	"6AB4":    "EC92",
	"6AB8":    "ECL80",
	"6AF4":    "EC94",
	"6AG5":    "EF96",
	"6AJ4":    "EC84",
	"6AJ8":    "ECH81",
	"6AK5":    "EF95",
	"6AK5W":   "E95F",
	"6AK8":    "EABC80",
	"6AL3":    "EY88",
	"6AL5":    "EAA91",
	"6AM5":    "EL91",
	"6AM6":    "EF91",
	"6AN7":    "ECH80",
	"6AQ4":    "EC91",
	"6AQ5":    "EL90",
	"6AQ8":    "ECC85",
	"6AS7G":   "ECC230",
	"6AT6":    "EBC90",
	"6AU6":    "EF94",
	"6AV6":    "EBC91",
	"6BA6":    "EF93",
	"6BD7A":   "EBC81",
	"6BE6":    "EK90",
	"6BE7":    "EQ80",
	"6BH6":    "E90F",
	"6BJ6":    "E99F",
	"6BK8":    "EF86",
	"6BL8":    "ECF80",
	"6BM8":    "ECL82",
	"6BN5":    "EL85",
	"6BQ5":    "EL84",
	"6BQ7A":   "ECC180",
	"6BR5":    "EM80",
	"6BX6":    "EF80",
	"6BY7":    "EF85",
	"6C4":     "EC90",
	"6CA4":    "EZ81",
	"6CA7":    "EL34",
	"6CD7":    "EM34",
	"6CH6":    "EL821",
	"6CJ5":    "EF41",
	"6CJ6":    "EL81",
	"6CK5":    "EL41",
	"6CK6":    "EL83",
	"6CM4":    "EC86",
	"6CM5":    "EL36",
	"6CN6":    "EL38",
	"6CQ6":    "EF92",
	"6CS6":    "EH90",
	"6CT7":    "EAF42",
	"6CU7":    "ECH42",
	"6CW5":    "EL86",
	"6CW7":    "ECC84",
	"6DA5":    "EM81",
	"6DA6":    "EF89",
	"6DC8":    "EBF89",
	"6DH7":    "EM84",
	"6DJ8":    "ECC88",
	"6DL4":    "EC88",
	"6DL5":    "EL95",
	"6DR8":    "EBF83",
	"6DS8":    "ECH83",
	"6DX8":    "ECL84",
	"6DY5":    "EL82",
	"6EC4A":   "EY500A",
	"6EH7":    "EF183",
	"6EJ7":    "EF184",
	"6ES6":    "EF97",
	"6ES8":    "ECC89",
	"6ET6":    "EF98",
	"6FG6":    "EM84",
	"6FN5":    "EL300",
	"6FY5":    "EC97",
	"6GB5":    "EL500",
	"6GJ7":    "ECF801",
	"6GM8":    "ECC86",
	"6GV8":    "ECL85",
	"6GW8":    "ECL86",
	"6GX8":    "EAM86",
	"6HG8":    "ECF86",
	"6HU6":    "EM87",
	"6HU8":    "ELL80",
	"6J1P":    "EF95",
	"6J6":     "ECC91",
	"6J7":     "EF37",
	"6JW8":    "ECF802",
	"6K7":     "EF39",
	"6KG6A":   "EL509",
	"6KX8":    "ECC808",
	"6L6":     "EL37",
	"6LN8":    "LCF80",
	"6LX8":    "LCF802",
	"6M5":     "EL80",
	"6N3":     "EY82",
	"6N8":     "EBF80",
	"6Q4":     "EC80",
	"6R3":     "EY81",
	"6R4":     "EC81",
	"6SL7":    "ECC35",
	"6SN7":    "ECC32",
	"6U8":     "ECF82",
	"6U9":     "ECF201",
	"6V4":     "EZ80",
	"6X2":     "EY51",
	"6X4":     "EZ90",
	"6X9":     "ECF200",
	"6Y9":     "EFL200",
	"6Ж1П":    "EF95",
	"7062":    "E180CC",
	"709":     "EL91",
	"7119":    "E182CC",
	"7125":    "EBF89",
	"7247":    "ECC832",
	"7308":    "E188CC",
	"7316":    "ECC186",
	"7320":    "E84L",
	"7534":    "E130L",
	"7643":    "E80CF",
	"7693":    "E90F",
	"7694":    "E99F",
	"7709":    "Z70W",
	"7710":    "Z70U",
	"7711":    "Z71U",
	"7713":    "Z804U",
	"7714":    "Z805U",
	"7722":    "E280F",
	"7737":    "E186F",
	"7751":    "E235L",
	"7788":    "E810F",
	"7AN7":    "PCC84",
	"7DJ8":    "PCC88",
	"7ES8":    "PCC189",
	"7GV7":    "PCF805",
	"7HG8":    "PCF86",
	"8108":    "EC157",
	"8223":    "E288CC",
	"8233":    "E55L",
	"8255":    "E88C",
	"8608":    "EL5070",
	"8CW5":    "XL86",
	"8D3":     "EF91",
	"8DX8":    "XCL84",
	"8GJ7":    "PCF801",
	"954":     "E1F",
	"955":     "E1C",
	"956":     "E2F",
	"9A8":     "PCF80",
	"9AB4":    "UC92",
	"9AK8":    "PABC80",
	"9AQ8":    "PCC85",
	"9GV8":    "XCL85",
	"9JW8":    "PCF802",
	"9U8A":    "PCF82",
	"9V9":     "PCH200",
	"B329":    "ECC82",
	"B339":    "ECC83",
	"CV10407": "EM87",
	"CV4010":  "EF95",
	"CV5724":  "E80T",
	"CV797":   "EN91",
	"EEP1":    "EE1",
	"LF183":   "YF183",
	"LF184":   "YF184",
	"M8082":   "EL91",
	"M8083":   "EF91",
	"M8136":   "ECC82",
	"M8137":   "ECC83",
	"M8161":   "EF92",
	"N150":    "EL41",
	"N309":    "PL83",
	"N709":    "EL84",
	"N727":    "EL90",
	"PL21":    "EN91",
	"PL2D21":  "EN91",
	"R243":    "EC55",
	"W719":    "EF85",
	"W727":    "EF93",
	"WD709":   "EBF80",
	"X719":    "ECH81",
	"YF183":   "LF183",
	"YF184":   "LF184",
	"Z152":    "EF80",
	"Z729":    "EF86",
	"Z77":     "EF91",
	"ИВ-15":   "DM160",
	// source: https://en.wikipedia.org/wiki/List_of_vacuum_tubes
	"7025": "ECC83",
}
