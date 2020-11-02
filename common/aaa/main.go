// scrape website
// The Go Programming ebook

package main

import "fmt"
import "log"
import "os"
import "encoding/json"

type Configuration struct {
	Address      string
	ReadTimeout  int64
	WriteTimeout int64
	Static       string
}

var config Configuration
var logger *log.Logger

func loadConfig() {
	file, err := os.Open("common/aaa/config.json")
	if err != nil {
		log.Fatalln("Cannot open config file", err)
	}
	decoder := json.NewDecoder(file)
	config = Configuration{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalln("Cannot get configuration from file", err)
	}
}

func main() {
	loadConfig()
	fmt.Println(config.Address)
}
