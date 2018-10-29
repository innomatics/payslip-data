package main 

import (
	"encoding/json"
	"io/ioutil"
	"log"
    "os"
)

func openFile(fileName string) (file *os.File) {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func createFile(fileName string) (file *os.File) {
	f, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func loadJSONFile(f *os.File, o interface{})  {
	defer f.Close()
	bytes, _ := ioutil.ReadAll(f)
	err := json.Unmarshal(bytes, o)
	if err != nil {
		log.Fatal(err)
	}
}
