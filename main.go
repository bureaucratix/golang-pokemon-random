package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type pokemon struct {
	Name string `json:"name"`
}

// https://pokeapi.co/api/v2/pokemon/25

func errorCheck(err error) {
	if err != nil {
		log.Fatalln(err)
		return
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	n1 := rand.Intn(806) + 1
	resp, apiErr := http.Get(fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%d", n1))
	errorCheck(apiErr)

	defer resp.Body.Close()

	body, bodyErr := ioutil.ReadAll(resp.Body)
	errorCheck(bodyErr)

	p := pokemon{}
	jsonErr := json.Unmarshal(body, &p)
	errorCheck(jsonErr)

	f, fileErr := os.Create("pokemon.txt")
	errorCheck(fileErr)

	_, writeErr := f.WriteString(p.Name)
	errorCheck(writeErr)

	closeErr := f.Close()
	errorCheck(closeErr)

	fmt.Println(p.Name)
}
