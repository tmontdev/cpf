package cpf

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

var numberRX = regexp.MustCompile("[^0-9]")
var defErr = errors.New("not a valid CPF")

const cpfLength = 11

type CPF string

func isOnlyNumbers(s string) bool {
	return !numberRX.MatchString(s)
}

func onlyNumbers(s string) string {
	return numberRX.ReplaceAllString(s, "")
}

func isValidLength(s string) bool {
	return len(s) == cpfLength
}

func ensureValidPattern(cpf string) string {
	if !isOnlyNumbers(cpf) {
		cpf = onlyNumbers(cpf)
	}
	if isValidLength(cpf) {
		cpf = fmt.Sprintf("%011s", cpf)
	}
	return cpf
}

func digitValidation(sequence string) int {
	sum := 0
	initial := len(sequence) + 1
	for i, n := range sequence {
		factor := initial - i
		sum += factor * int(n-'0')
	}
	remainder := sum % cpfLength
	digit := cpfLength - remainder
	if digit >= 10 {
		digit = 0
	}
	return digit
}

func FromString(source string) (CPF, error) {
	raw := ensureValidPattern(source)
	cpf := CPF(raw)
	number, _ := strconv.Atoi(raw)
	if len(raw) != cpfLength || number < 100000000 {
		return cpf, defErr
	}
	firstDigit := digitValidation(raw[:9])
	if strconv.Itoa(firstDigit) != raw[9:10] {
		return cpf, defErr
	}
	secondDigit := digitValidation(raw[:10])
	if strconv.Itoa(secondDigit) != raw[10:] {
		return cpf, defErr
	}
	return cpf, nil
}
func FromInt(i int) (CPF, error) {
	return FromString(strconv.Itoa(i))
}
func FromInt64(i int64) (CPF, error) {
	return FromString(strconv.FormatInt(i, 10))
}

func FromBytes(bytes []byte) (CPF, error) {
	return FromString(string(bytes))
}

func IsValid(cpf string) bool {
	_, err := FromString(cpf)
	return err == nil
}

func Must(cpf string) CPF {
	compiled, err := FromString(cpf)
	if err != nil {
		panic(err)
	}
	return compiled
}

func (c CPF) String() string {
	return string(c)
}

func (c CPF) IsValid() bool {
	return IsValid(c.String())
}
func (c CPF) Format() string {
	s := c.String()
	return fmt.Sprintf("%s.%s.%s-%s", s[:3], s[3:6], s[6:9], s[9:])
}
