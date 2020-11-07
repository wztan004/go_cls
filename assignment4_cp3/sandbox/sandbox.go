package main

import (
	// "assignment4_cp3/datastruct"
	"assignment4_cp3/utils"
	// "encoding/csv"
	// "encoding/hex"
	// "errors"
	"fmt"
	// "log"
	// "net/http"
	// "os"
	// "sync"
	// "time"
)

func main() {
	// var wg sync.WaitGroup
	// wg.Add(2)
	// firstchan := make(chan [][]string)
	// secondchan := make(chan [][]string)
	// go utils.ReadFileConcurrently(`assignment4_cp3/confidential/venues_202009.csv`, firstchan, &wg)
	// go utils.ReadFileConcurrently(`assignment4_cp3/confidential/venues_202010.csv`, secondchan, &wg)
	// x1 := <- firstchan
	// x2 := <- secondchan
	// wg.Wait()
	// x1 = append(x1, x2...)
	// fmt.Println(x1)
	updateVenueCSV := func() {
		for _, j := range utils.ReadMultipleFilesConcurrently() {

				fmt.Println(j)


		}
	}

	updateVenueCSV()
}

// func ReadFileConcurrently(path string, ch chan <-[][]string) {
// 	defer wg.Done()
// 	file, err := os.Open(path)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()

// 	reader := csv.NewReader(file)
// 	record, err := reader.ReadAll()
// 	if err != nil {
// 		panic(err)
// 	}
// 	ch <- record
// }
