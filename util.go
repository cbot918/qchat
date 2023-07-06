package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var log = fmt.Println

type Config struct {
	Port string
	Web  string
}

func GetConfig() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		return Config{}, err
	}
	cfg := Config{
		Port: os.Getenv("PORT"),
		Web:  os.Getenv("WEB"),
	}
	return cfg, nil
}

func IsFirstMsg(msg []byte) bool {
	flag := "_init_user_name"
	return strings.Contains(string(msg), "\""+flag+"\":")
}
