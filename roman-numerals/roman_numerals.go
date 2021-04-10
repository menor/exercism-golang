package romannumerals

import "errors"

func ToRomanNumeral(arabic int) (string, error) {
	var roman string

	if arabic <= 0 || arabic > 3000 {
		return roman, errors.New("number must be between 1 and 3000")
	}

	thousands := arabic / 1000
	hundreds := (arabic % 1000) / 100
	tens := (arabic % 100) / 10
	units := (arabic % 10)

	roman += romanBaseThousand[thousands]
	roman += romanBaseHundred[hundreds]
	roman += romanBaseTen[tens]
	roman += romanBase[units]

	return roman, nil
}

var romanBase = []string{
	0: "",
	1: "I",
	2: "II",
	3: "III",
	4: "IV",
	5: "V",
	6: "VI",
	7: "VII",
	8: "VIII",
	9: "IX",
}

var romanBaseTen = map[int]string{
	0: "",
	1: "X",
	2: "XX",
	3: "XXX",
	4: "XL",
	5: "L",
	6: "LX",
	7: "LXX",
	8: "LXXX",
	9: "XC",
}

var romanBaseHundred = map[int]string{
	0: "",
	1: "C",
	2: "CC",
	3: "CCC",
	4: "CD",
	5: "D",
	6: "DC",
	7: "DCC",
	8: "DCCC",
	9: "CM",
}

var romanBaseThousand = map[int]string{
	0: "",
	1: "M",
	2: "MM",
	3: "MMM",
}
