package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type config struct {
	Loc_Next_Off     int
	Loc_Previous_Off int
}

type location_area struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

var Commands map[string]cliCommand

func commandExit(conf *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(conf *config) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:\n")

	for _, value := range Commands {
		fmt.Printf("%s: %s\n", value.name, value.description)
	}

	return nil
}

func commandMap(conf *config) error {

	url := fmt.Sprint("https://pokeapi.co/api/v2/location-area?limit=20&offset=", conf.Loc_Next_Off)

	res, err := http.Get(url)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	var data location_area

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&data)

	if err != nil {
		return err
	}

	for _, value := range data.Results {
		fmt.Println(value.Name)
	}

	conf.Loc_Next_Off += 20
	conf.Loc_Previous_Off += 20
	return nil
}

func commandMapb(conf *config) error {

	if conf.Loc_Previous_Off == -20 {
		fmt.Println("You're on the first page")
		return nil
	}

	url := fmt.Sprint("https://pokeapi.co/api/v2/location-area?limit=20&offset=", conf.Loc_Previous_Off)

	res, err := http.Get(url)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	var data location_area
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&data)

	if err != nil {
		return err
	}

	for _, value := range data.Results {
		fmt.Println(value.Name)
	}

	conf.Loc_Next_Off -= 20
	conf.Loc_Previous_Off -= 20
	return nil
}
