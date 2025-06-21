package main

import (
	"encoding/json"
	"errors"
	"os"

	"golang.org/x/crypto/bcrypt"
)

type Profile_t struct {
	Username string
	Password string
	Pokedex  map[string]Pokemon
}

func Exists() {
	filename := "internal/profile.json"
	_, err := os.Stat(filename)

	if errors.Is(err, os.ErrNotExist) {
		os.Create("internal/profile.json")
		Pokedex = make(map[string]Pokemon)
		Save()
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

func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		return "", err
	}

	return string(hashed), nil
}
