package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"time"
)

type Headers struct {
	Name    string `json: "name"`
	File    string `json: "file"`
	Project string `json: "project"`
	Company string `json: "company"`
}

const YYYYMMDD = "2006/01/02"
const YYYY = "2006"

func main() {

	var filename string
	flag.StringVar(&filename, "f", "headers.json", "File to load header info from.")
	flag.Parse() // after declaring flags we need to call it

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := Headers{}
	_ = json.Unmarshal([]byte(file), &data)

	fmt.Println(
		`// 
//  ` + data.File + `
//  ` + data.Project + `
//
//  Created by ` + data.Name + ` on ` + time.Now().Format(YYYYMMDD) + `
//  Copyright © ` + time.Now().Format(YYYY) + ` ` + data.Company + `. All rights reserved.
//`)
}
