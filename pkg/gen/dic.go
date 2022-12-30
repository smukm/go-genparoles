package gen

import (
	"bufio"
	"bytes"
	"io"
	"math/rand"
	"os"
	"time"
)

// возвращает количество строк из файла
func CountLinesInFile(fileName string) (int, error) {
	file, err := os.Open(fileName)

	if err != nil {
		return 0, err
	}

	buf := make([]byte, 1024)
	lines := 0

	for {
		readBytes, err := file.Read(buf)

		if err != nil {
			if readBytes == 0 && err == io.EOF {
				err = nil
			}
			return lines, err
		}

		lines += bytes.Count(buf[:readBytes], []byte{'\n'})
	}
}

/*
*	возвращает заданное количество случайных строк из файла словаря
*   объявляем метод как функцию структуры application
 */
func GetRandomLines(filename string, lines int) ([]string, error) {
	// число строк в словаре
	maxLines, err := CountLinesInFile(filename)
	if err != nil {
		return nil, err
	}

	// срезы для индексов строк и результата
	rndmLines := make([]int, lines)
	ret := make([]string, lines)

	// формируем массив со случайными индексами строк
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < lines; i++ {
		rndmLines[i] = rand.Intn(maxLines)
	}

	// открытие файла словаря
	dic, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	// создание сканера
	scaner := bufio.NewScanner(dic)
	idx := 0
	retIdx := 0
	for scaner.Scan() {
		for _, rndmLn := range rndmLines {
			if idx == rndmLn {
				ret[retIdx] = scaner.Text()
				retIdx++
			}
		}
		// если результирующий массив уже заполнен, выходим
		if retIdx == lines {
			break
		}

		idx++
	}

	return ret, nil
}
