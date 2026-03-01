package validator

import (
	"regexp"
	"strconv"
)

func IsValidDocument(doc string) bool {
	doc = cleanDocument(doc)

	if len(doc) == 11 {
		return isValidCPF(doc)
	}
	if len(doc) == 14 {
		return isValidCNPJ(doc)
	}
	return false
}

func cleanDocument(doc string) string {
	reg := regexp.MustCompile(`[^0-9]`)
	return reg.ReplaceAllString(doc, "")
}

func isValidCPF(cpf string) bool {
	if len(cpf) != 11 {
		return false
	}

	if allDigitsEqual(cpf) {
		return false
	}

	digits := make([]int, 11)
	for i, r := range cpf {
		digits[i], _ = strconv.Atoi(string(r))
	}

	sum := 0
	for i := 0; i < 9; i++ {
		sum += digits[i] * (10 - i)
	}
	digit1 := (sum * 10) % 11
	if digit1 == 10 {
		digit1 = 0
	}
	if digit1 != digits[9] {
		return false
	}

	sum = 0
	for i := 0; i < 10; i++ {
		sum += digits[i] * (11 - i)
	}
	digit2 := (sum * 10) % 11
	if digit2 == 10 {
		digit2 = 0
	}
	return digit2 == digits[10]
}

func isValidCNPJ(cnpj string) bool {
	if len(cnpj) != 14 {
		return false
	}

	if allDigitsEqual(cnpj) {
		return false
	}

	digits := make([]int, 14)
	for i, r := range cnpj {
		digits[i], _ = strconv.Atoi(string(r))
	}

	weights1 := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	sum := 0
	for i := 0; i < 12; i++ {
		sum += digits[i] * weights1[i]
	}
	digit1 := sum % 11
	if digit1 < 2 {
		digit1 = 0
	} else {
		digit1 = 11 - digit1
	}
	if digit1 != digits[12] {
		return false
	}

	weights2 := []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	sum = 0
	for i := 0; i < 13; i++ {
		sum += digits[i] * weights2[i]
	}
	digit2 := sum % 11
	if digit2 < 2 {
		digit2 = 0
	} else {
		digit2 = 11 - digit2
	}
	return digit2 == digits[13]
}

func allDigitsEqual(s string) bool {
	first := s[0]
	for _, c := range s {
		if byte(c) != first {
			return false
		}
	}
	return true
}
