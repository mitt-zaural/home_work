package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var count int
	var err error
	var strlen = len(s)

	result := ""
	for i := 0; i < strlen; i++ {
		c := s[i : i+1] // выделяем символ из строки
		if c != "\\" {
			count, err = strconv.Atoi(c)
		} else {
			i++
			err = ErrInvalidString
		}
		if err == nil { // удачно сконвертировали - значит первый символ цифра, что является ошибкой
			return "", ErrInvalidString // возвращаем пустую строку и строку описания ошибки
		}
		if i < strlen-1 { // индекс ещё в диапазоне строки?
			var c1 string = s[i+1 : i+2]  // читаем следующий символ
			count, err = strconv.Atoi(c1) // следующий символ - это число?
			if err == nil && count != 0 { // да, это ненулевое число
				result += strings.Repeat(c, count-1)
				i++ // увеличим индекс, перескочим число
			}
			if err == nil && count == 0 { // это нулевое число
				i++ // увеличим индекс, перескочим число
				c = ""
			}
		}
		result += c
	}
	return result, nil
}