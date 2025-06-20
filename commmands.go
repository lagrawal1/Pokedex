package main

import (
	pokecache "bootdev/Pokedex/internal"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

var Cache *pokecache.Cache

type config struct {
	Loc_Next_Off     int
	Loc_Previous_Off int
	Parameters       []string
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

type loc_area_pok struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
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
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	fmt.Println()

	for _, value := range Commands {
		fmt.Printf("%s: %s\n", value.name, value.description)
	}

	return nil
}

func commandMap(conf *config) error {

	url := fmt.Sprint("https://pokeapi.co/api/v2/location-area?limit=20&offset=", conf.Loc_Next_Off)
	var data location_area
	var unmar_data []byte

	if val, ok := Cache.Get(url); ok {
		unmar_data = val
	} else {
		res, err := http.Get(url)

		if err != nil {
			return err
		}

		unmar_data, err = io.ReadAll(res.Body)

		if err != nil {
			return err
		}

		defer res.Body.Close()

		Cache.Add(url, unmar_data)

	}

	err := json.Unmarshal(unmar_data, &data)

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

	conf.Loc_Next_Off -= 20
	conf.Loc_Previous_Off -= 20

	url := fmt.Sprint("https://pokeapi.co/api/v2/location-area?limit=20&offset=", conf.Loc_Previous_Off)

	var data location_area

	if val, ok := Cache.Get(url); ok {
		err := json.Unmarshal(val, &data)

		if err != nil {
			return err
		}

	} else {
		res, err := http.Get(url)

		if err != nil {
			return err
		}

		cache_data, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		defer res.Body.Close()

		err = json.Unmarshal(cache_data, &data)

		if err != nil {
			return err
		}

		Cache.Add(url, cache_data)

	}

	for _, value := range data.Results {
		fmt.Println(value.Name)
	}

	return nil
}

func commandExplore(conf *config) error {
	if len(conf.Parameters) == 0 {
		fmt.Print("Location Parameter not given. Give input in 'explore <location_area>' ")
		return nil
	}

	loc_exp := conf.Parameters[0]

	url := fmt.Sprint("https://pokeapi.co/api/v2/location-area/", loc_exp)
	var data loc_area_pok
	var unmar_data []byte

	if val, ok := Cache.Get(url); ok {
		unmar_data = val
	} else {
		res, err := http.Get(url)

		if err != nil {
			return err
		}

		unmar_data, err = io.ReadAll(res.Body)

		if err != nil {
			return err
		}

		defer res.Body.Close()

		Cache.Add(url, unmar_data)

	}

	err := json.Unmarshal(unmar_data, &data)

	if err != nil {
		return err
	}
	fmt.Println("Exploring ", loc_exp, "...")
	fmt.Println("Found Pokemon:")
	for _, value := range data.PokemonEncounters {
		fmt.Println("- ", value.Pokemon.Name)
	}

	return nil

}
