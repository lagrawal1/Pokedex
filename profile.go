package main

import (
	"encoding/json"
	"errors"
	"os"

	"golang.org/x/crypto/bcrypt"
)

type Profile_t struct {
	Password []byte
	Pokedex  map[string]Pokemon
}

var Profiles map[string]Profile_t

func Exists() {
	filename := "internal/profile.json"
	_, err := os.Stat(filename)
	Profiles = make(map[string]Profile_t)

	if errors.Is(err, os.ErrNotExist) {
		os.Create("internal/profile.json")
	}
}

func Save(username string) error {
	filename := "internal/profile.json"

	data, err := json.Marshal(Profiles)

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
