package main

import (
	"encoding/json"
	"fmt"
	"log"
	"redistest/cache/redis"
)

type simple struct {
	Foo  string `json:"foo"`
	Blah int    `json:"blah"`
}

func main() {

	handle, err := redis.Connect()
	if err != nil {
		log.Fatalln(err)
	}

	defer handle.Close()

	// Ping Check
	reply, err := handle.Ping()
	fmt.Println(reply, err)

	// Set data
	handle.SetKey("name", "some name")
	handle.SetKey("number", 100)
	handle.SetKey("json", "{\"foo\":\"bar\",\"blah\":10}")

	name, err := handle.GetKey("name")
	number, err := handle.GetKey("number")
	jsonStr, err := handle.GetKey("json")

	fmt.Println(name.(string), err)
	fmt.Println(number.(int64), err)

	var t simple
	json.Unmarshal(jsonStr.([]byte), &t)
	fmt.Println(t)

	keys := []string{"name", "number"}
	rec, err := handle.GetRecord(keys)
	fmt.Println(rec, err)

}
