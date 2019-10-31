package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// https://pokeapi.co/api/v2/pokemon/25

func main() {
	resp, err := http.Get("https://pokeapi.co/api/v2/pokemon/25")

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, bodyerr := ioutil.ReadAll(resp.Body)

	if bodyerr != nil {
		log.Fatalln(bodyerr)
	}

	fmt.Println(string(body))
}
