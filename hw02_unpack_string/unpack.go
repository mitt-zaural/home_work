package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode/utf8"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var countNext int
	var err error
	strlen := len(s)

	result := ""
	for i := 0; i < strlen; {
		c, rsize := utf8.DecodeRuneInString(s[i:]) // выделяем руну из строки
		if string(c) != "\\" {
			_, err = strconv.Atoi(string(c))
		} else {
			i += rsize
			err = ErrInvalidString
		}
		if err == nil { // удачно сконвертировали - значит первый символ цифра, что является ошибкой
			return "", ErrInvalidString // возвращаем пустую строку и строку описания ошибки
		}
		if i < strlen-1 { // индекс ещё в диапазоне длины строки?
			c1, rsize2 := utf8.DecodeRuneInString(s[i+rsize:]) // читаем следующую руну
			countNext, err = strconv.Atoi(string(c1))          // следующая руна - это число?
			if err == nil && countNext != 0 {                  // да, это ненулевое число
				result += strings.Repeat(string(c), countNext-1)
				i += rsize2 // увеличим индекс, перескочим число
			}
			if err == nil && countNext == 0 { // это нулевое число
				i += rsize2 // увеличим индекс, перескочим число
				c = 0
			}
		}
		if c != 0 {
			result += string(c)
		}
		i += rsize // увеличиваем индекс с учетом размера обработанной руны
	}
	return result, nil
}
