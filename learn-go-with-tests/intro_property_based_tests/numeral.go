package intropropertybasedtests

import (
	"strings"
)

// [1]
/* func ConvertToRoman(arabic int) string {
	if arabic == 3 {
		return "III"
	}

	if arabic == 2 {
		return "II"
	}

	return "I"
} */

// [2]
/* func ConvertToRoman(arabic int) string {
	var result strings.Builder

	if arabic == 4 {
		return "IV"
	}

	for i := 0; i < arabic; i++ {
		result.WriteString("I")
	}

	return result.String()
} */

/* func ConvertToRoman(arabic int) string {
	var result strings.Builder

	for i := arabic; i > 0; i-- {
		if arabic == 5 {
			result.WriteString("V")

			break
		}

		if arabic == 4 {
			result.WriteString("IV")

			break
		}

		result.WriteString("I")
	}

	return result.String()
} */

/* func ConvertToRoman(arabic int) string {
	var result strings.Builder

	for arabic > 0 {
		switch {
		case arabic > 9:
			result.WriteString("X")
			arabic -= 10
		case arabic > 8:
			result.WriteString("IX")
			arabic -= 9
		case arabic > 4:
			result.WriteString("V")
			arabic -= 5
		case arabic > 3:
			result.WriteString("IV")
			arabic -= 4
		default:
			result.WriteString("I")
			arabic--
		}
	}

	return result.String()
} */

type RomanNumeral struct {
	Value  uint16
	Symbol string
}

type RomanNumerals []RomanNumeral

/* func (r RomanNumerals) ValueOf(symbol string) int {
	for _, s := range r {
		if s.Symbol == symbol {
			return s.Value
		}
	}

	return 0
}
*/

func (r RomanNumerals) ValueOf(symbols ...byte) uint16 {
	symbol := string(symbols)
	for _, s := range r {
		if s.Symbol == symbol {
			return s.Value
		}
	}

	return 0
}

var allRomanNumerals = RomanNumerals{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{500, "D"},
	{100, "C"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(arabic uint16) string {
	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String()
}

// [4]
/* func ConvertToArabic(roman string) int {
	total := 0

	for i := 0; i < len(roman); i++ {
		symbol := roman[i]

		if i+1 < len(roman) && symbol == 'I' {
			nextSymbol := roman[i+1]

			potentialNumber := string([]byte{symbol, nextSymbol})

			value := allRomanNumerals.ValueOf(potentialNumber)

			if value != 0 {
				total += value
				i++
			} else {
				total++
			}
		} else {
			total++
		}
	}

	return total
}

func couldBeSubtractive(index int, currentSymbol uint8, roman string) bool {
	return index+1 < len(roman) && currentSymbol == 'I'
}

func ConvertToArabic(roman string) int {
	total := 0

	for i := 0; i < len(roman); i++ {
		symbol := roman[i]

		// look ahead to next symbol if we can and, the current symbol is base 10 (only valid subtractors)
		if couldBeSubtractive(i, symbol, roman) {
			nextSymbol := roman[i+1]

			// build the two character string
			potentialNumber := string([]byte{symbol, nextSymbol})

			if value := allRomanNumerals.ValueOf(potentialNumber); value != 0 {
				total += value
				i++ // move past this character too for the next loop
			} else {
				total += allRomanNumerals.ValueOf(string([]byte{symbol}))
			}
		} else {
			total += allRomanNumerals.ValueOf(string([]byte{symbol}))
		}
	}
	return total
} */

func ConvertToArabic(roman string) (total uint16) {
	for _, symbols := range windowedRoman(roman).Symbols() {
		total += allRomanNumerals.ValueOf(symbols...)
	}
	return
}

func (r RomanNumerals) Exists(symbols ...byte) bool {
	symbol := string(symbols)

	for _, s := range r {
		if s.Symbol == symbol {
			return true
		}
	}

	return false
}

type windowedRoman string

func (w windowedRoman) Symbols() (symbols [][]byte) {
	for i := 0; i < len(w); i++ {
		symbol := w[i]
		notAtEnd := i+1 < len(w)

		if notAtEnd && isSubstractive(symbol) && allRomanNumerals.Exists(symbol, w[i+1]) {
			symbols = append(symbols, []byte{symbol, w[i+1]})
			i++
		} else {
			symbols = append(symbols, []byte{symbol})
		}

	}

	return
}

func isSubstractive(symbol uint8) bool {
	return symbol == 'I' || symbol == 'X' || symbol == 'C'
}
