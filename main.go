package main

import (
	pokecache "bootdev/Pokedex/internal"
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func main() {
	Commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "<add map description>",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "add mapb description>",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore a Location Area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a Pokemon",
			callback:    commmandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a Pokemon that you have caught",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "View the names of the pokemon you caught",
			callback:    commandPokedex,
		},
		"login": {
			name:        "login",
			description: "Login. Enter login <username> <password>",
			callback:    commandLogin,
		},
		"newuser": {
			name:        "newuser",
			description: "Creates a new profile for a user",
			callback:    commandNewuser,
		},
	}

	Exists()

	profile_data, err := os.ReadFile("internal/profile.json")

	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(profile_data, &Profiles)

	scanner := bufio.NewScanner(os.Stdin)
	conf := config{Loc_Next_Off: 0, Loc_Previous_Off: -20, Catch_Chance: 25}
	Cache = pokecache.NewCache(20 * time.Second)

	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		input := scanner.Text()

		cleaned_input := cleanInput(input)

		curr_command := cleaned_input[0]

		val, ok := Commands[curr_command]

		conf.Parameters = make([]string, len(cleaned_input[1:]))
		copy(conf.Parameters, cleaned_input[1:])

		if ok {
			err := val.callback(&conf)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}

		err = Save(conf.User)

		if err != nil {
			fmt.Println(err)
		}

	}

}
