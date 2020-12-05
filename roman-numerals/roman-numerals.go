package numerals

import "strings"

// RomanNumeral is collection of arabic int and its represent in roman symbols (string)
type RomanNumeral struct {
	Arabic int
	Roman  string
}

// RomanNumerals is a collection of multiple roman numeral collection
type RomanNumerals []RomanNumeral

// ValueOf method for returning the Arabic representation of Roman symbols
func (r RomanNumerals) ValueOf(romans ...byte) int {
	roman := string(romans)

	for _, s := range r {
		if s.Roman == roman {
			return s.Arabic
		}
	}
	return 0
}

var allRomanNumerals = RomanNumerals{
	{Arabic: 1000, Roman: "M"},
	{Arabic: 900, Roman: "CM"},
	{Arabic: 500, Roman: "D"},
	{Arabic: 400, Roman: "CD"},
	{Arabic: 100, Roman: "C"},
	{Arabic: 90, Roman: "XC"},
	{Arabic: 50, Roman: "L"},
	{Arabic: 40, Roman: "XL"},
	{Arabic: 10, Roman: "X"},
	{Arabic: 9, Roman: "IX"},
	{Arabic: 5, Roman: "V"},
	{Arabic: 4, Roman: "IV"},
	{Arabic: 1, Roman: "I"},
}

// ConvertToRoman converts arabic int to roman strings
func ConvertToRoman(arabic int) string {
	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Arabic {
			result.WriteString(numeral.Roman)
			arabic -= numeral.Arabic
		}
	}

	return result.String()
}

// ConvertToArabic converts roman string to arabic int
func ConvertToArabic(roman string) int {
	total := 0

	for i := 0; i < len(roman); i++ {
		symbol := roman[i]

		// look ahead to next symbol if we can and, the current symbol is base 10 (only valid subtractors)
		if couldBeSubtractive(i, symbol, roman) {

			// get the value of the two character string
			if value := allRomanNumerals.ValueOf(symbol, roman[i+1]); value != 0 {
				total += value
				i++ // move past this character too for the next loop
			} else {
				total += allRomanNumerals.ValueOf(symbol)
			}
		} else {
			total += allRomanNumerals.ValueOf(symbol)
		}
	}
	return total
}

func couldBeSubtractive(index int, currentSymbol uint8, roman string) bool {
	isSubtractiveSymbol := currentSymbol == 'I' || currentSymbol == 'X' || currentSymbol == 'C'
	return index+1 < len(roman) && isSubtractiveSymbol
}
