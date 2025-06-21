package main

import (
	"encoding/json"
	"errors"
	"os"
)

func Exists() {
	filename := "internal/profile.json"
	_, err := os.Stat(filename)

	if errors.Is(err, os.ErrNotExist) {
		os.Create("internal/profile.json")
	}

}

func Save() error {
	filename := "internal/profile.json"

	data, err := json.Marshal(Pokedex)

	if err != nil {
		return err
	}

	err = os.WriteFile(filename, data, 0644)

	if err != nil {
		return err
	}

	return nil
}
