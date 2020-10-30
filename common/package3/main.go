// cross-package logging
// https://stackoverflow.com/questions/18361750/correct-approach-to-global-logging-in-golang

package main

import (
	"os"
	"log"
	"sync"
	"io"
)

type logger struct {
    filename string
    *log.Logger
}

var logge *logger
var once sync.Once

// start loggeando
func GetInstance() *logger {
    once.Do(func() {
        logge = createLogger("common/logging3/log.txt")
    })
    return logge
}

func createLogger(fname string) *logger {
    file, _ := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

    return &logger{
        filename: fname,
        Logger:   log.New(io.MultiWriter(file, os.Stderr), "My app Name ", log.Ldate|log.Ltime|log.Lshortfile),
    }
}


func main() {
	l := GetInstance()

    l.Println("Starting")
}

