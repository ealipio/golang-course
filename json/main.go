package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	// generateJSONWithNoTags()
	// generateJSONWithTags()
	// decodeJSON()
	defer fmt.Println("I will run after the function completes")

	fmt.Println("Hello Peru!")
}
func decodeJSON() {
	type Person struct {
		Name    string   `json:"name"`
		Age     int      `json:"age"`
		Hobbies []string `json:"hobbies"`
	}

	jsonStringData := `{"name":"George","age":40,"hobbies":["Cycling", "Cheese","Techno"]}`
	jsonByteData := []byte(jsonStringData)

	p := Person{}
	err := json.Unmarshal(jsonByteData, &p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", p)
}
func generateJSONWithTags() {
	type Person struct {
		Name    string   `json:"name"`
		Age     int      `json:"age"`
		Hobbies []string `json:"hobbies"`
	}

	hobbies := []string{"Chatting", "Hacking", "Coding"}
	p := Person{
		Name:    "Cachimbo",
		Age:     2,
		Hobbies: hobbies,
	}
	jsonByteData, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}
	jsonStringData := string(jsonByteData)
	fmt.Printf("%+v\n", jsonStringData)
}

func generateJSONWithNoTags() {
	type Person struct {
		Name    string
		Age     int
		Hobbies []string
	}

	hobbies := []string{"Chatting", "Hacking", "Coding"}
	p := Person{
		Name:    "Cachimbo",
		Age:     2,
		Hobbies: hobbies,
	}
	// fmt.Printf("%+v\n", p)
	jsonByteData, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%+v\n", jsonByteData)
	jsonStringData := string(jsonByteData)
	fmt.Printf("%+v\n", jsonStringData)
}
