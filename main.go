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
	n2 := rand.Intn(806) + 1
	n3 := rand.Intn(806) + 1
	n4 := rand.Intn(806) + 1
	n5 := rand.Intn(806) + 1
	n6 := rand.Intn(806) + 1

	dexNums := []int{n1, n2, n3, n4, n5, n6}

	f, fileErr := os.Create("pokemon.txt")
	errorCheck(fileErr)

	for _, num := range dexNums {
		resp, apiErr := http.Get(fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%d", num))
		errorCheck(apiErr)

		body, bodyErr := ioutil.ReadAll(resp.Body)
		errorCheck(bodyErr)
		defer resp.Body.Close()

		p := pokemon{}
		jsonErr := json.Unmarshal(body, &p)
		errorCheck(jsonErr)
		fmt.Println(p.Name)

		_, writeErr := f.WriteString(fmt.Sprintf("%s\n", p.Name))
		errorCheck(writeErr)
	}

	closeErr := f.Close()
	errorCheck(closeErr)

}
