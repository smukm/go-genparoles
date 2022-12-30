package main

import (
	"flag"
	"fmt"
	"genparole/pkg/gen"
	"log"
	"strconv"
)

/**
* структура для хранения зависимостей всего приложения
 */
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	
	//infoLog := log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime)
	//errorLog := log.New(os.Stderr, "ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile)

	// initialize structure with application dependency
	/*app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}*/

	// get command line parameters
	// -lines=5 number of words
	lines := flag.String("lines", "5", "Число слов")
	flag.Parse()
	//fmt.Println("Заданное число слов:", *lines)
	iLines, err := strconv.Atoi(*lines)
	check(err)
	iLines *=2

	words := gen.GenWords(iLines)

	// output result words
	for i := 0; i<iLines; i=i+2 {
		fmt.Println(words[i] + words[i+1])
	}
}
