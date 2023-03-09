package main

import (
	"os"
	"fmt"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
// EnvValue помогает отличить пустые файлы от файлов с первой пустой строкой
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
// ReadDir читает указанный каталог и возвращает карту переменных env
// Переменные представлены в виде файлов, где имя файла — имя переменной, первая строка файла — значение
func ReadDir(dir string) (Environment, error) {
	// Place your code here
	var ev EnvValue
	var env Environment = make(map[string]EnvValue)
	var buffer []byte = make([]byte, 100*1024)

// чтение списка файлов в указанном пути dir
	files, err := os.ReadDir(dir)
// есть доступ?
	if err != nil {
		return nil, err
	}
//fmt.Println("files = ", files)
	// доступ есть - чтение всех файлов, в. т.ч. являющихся подкаталогами
	for _, val := range files {
		if !val.IsDir() { // если это не подкаталог, то обработать файл
			finfo, _ := val.Info()
fmt.Printf("File %s size: %d\n", finfo.Name(), finfo.Size())
			if finfo.Size() > 0 {
				f, ferr := os.Open(dir + finfo.Name()) // открыть файл на чтение
				if ferr != nil { // файл открылся?
fmt.Printf("%s access fault\n", finfo.Name())
					return nil, ferr
				}
				flen, fread_err := f.Read(buffer) // читаем его содержимое
fmt.Println("length = ", flen)
				if fread_err != nil {
fmt.Printf("reading fault. %s\n", read_err)
					return nil, read_err
				}
				if flen == 0 {
					ev.NeedRemove = true
				}
				f.Close()
				ev.Value = string(buffer[:flen])
			}			
fmt.Println("ev.Value = ", ev.Value)
			env[finfo.Name()] = ev // наполняем карту структурой типа EnvValue
		}
	}
	return env, nil	
}
