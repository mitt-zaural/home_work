package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type TWord struct {
	word string
	freq int
}

var (
	words  []TWord // промежуточный срез структур для сортировки, подсчёта количества встреченных слов
	result []string
)

func Top10(sl string) []string {
	// Place your code here.
	var cnt int = 0 // счётчик дубликатов слов, встречающихся в исходной строке
	if sl == "" {   // исходная строка - пустая?
		return nil // тогда вернём пустой указатель - нечего обрабатывать
	}
	f := strings.Fields(sl) // разложим исходную строку в срез строк
	// сортируем срез строк в алфавитном порядке для упрощения подсчёта дубликатов
	sort.Slice(f, func(i, j int) bool { return f[i] < f[j] })
	var tmpstr string // временная буферная строка
	for idx, val := range f {
		if idx == 0 { // первый элемент среза?
			tmpstr = val // запоминаем строку - первый элемент
			cnt = 0      // обнулим счётчик
		}
		// подсчёт числа дубликатов
		if val == tmpstr {
			cnt++
		} else {
			words = append(words, TWord{tmpstr, cnt})
			tmpstr = val
			cnt = 1
		}
		if idx == len(f)-1 { // последний элемент среза?
			words = append(words, TWord{val, cnt})
		}
	}
	// сортировка по алфавиту
	sort.Slice(words, func(i, j int) bool { return words[i].word < words[j].word })
	// сортировка по частоте
	sort.Slice(words, func(i, j int) bool { return (words[i].freq > words[j].freq) })
	for idx, val := range words {
		if idx < 10 {
			result = append(result, val.word)
		} else {
			break
		}
	}
	return result
}
