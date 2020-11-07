package log

// import (
// 	co "assignment4_cp1/constants"
// 	"os"
// 	"log"
// 	"io"
// )

// var (
// 	Warning *log.Logger // Be concerned
// 	Error *log.Logger // Critical problem
// )

// func init() {
// 	file, err := os.OpenFile(co.LOG_FILE, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
// 	if err != nil {
// 		log.Fatalln("Failed to open error log file:", err)
//     }
//     defer file.Close()

// 	Warning = log.New(io.MultiWriter(file, os.Stderr), "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
// 	Error = log.New(io.MultiWriter(file, os.Stderr), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
// }

// func main() {
// 	Warning.Println("There is something you need to know about")
//     Error.Println("Something has failed")
// }



// // type logger struct {
// //     filename string
// //     *log.Logger
// // }

// // var logge *logger
// // var once sync.Once

// // // start loggeando
// // func GetInstance() *logger {
// //     once.Do(func() {
// //         logge = createLogger(co.LOG_FILE)
// //     })
// //     return logge
// // }

// // func createLogger(fname string) *logger {
// //     file, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
// //     if err != nil {
// //         log.Fatalln("Failed to open error log file:", err)
// //     }

// //     return &logger{
// //         filename: fname,
// //         Logger: log.New(io.MultiWriter(file, os.Stderr), "Warning: ", log.Ldate|log.Ltime|log.Lshortfile), //returns pointer to logger
// //     }
// // }

// // func main() {
// //     l := GetInstance()
// //     c := GetInstance()

// //     l.Println("Starting")
// //     c.Println("Starting")
// // }