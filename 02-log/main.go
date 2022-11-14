package main

import (
	"fmt"
	"log"
	"os"
)

//
//// Formatando o log
//func formatLog() {
//	var warning *log.Logger
//	warning = log.New(os.Stdout,
//		"WARNING: ",
//		log.Ldate|log.Ltime|log.Lshortfile) // data hora e o arquivo
//
//	warning.Println("aviso numero 1")
//	warning.Println("aviso numero 2")
//	//ex WARNING: 2022/11/13 20:19:22 main.go:23: aviso numero 1
//}

func fileLogging() {
	fmt.Println("--------------- file log ---------------")

	file, err := os.OpenFile("02-log/logs.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var logFile *log.Logger
	logFile = log.New(file, "App: ", log.Ldate|log.Ltime|log.Lshortfile)

	logFile.Println("message de error  1")
	logFile.Println("message de error  2")
	logFile.Println("message de error  3")
	fmt.Println("Fim")
}

func main() {
	//formatLog()
	fileLogging()
}
