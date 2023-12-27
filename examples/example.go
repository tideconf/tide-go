package main

import (
	"fmt"
	"log"

	tide "github.com/tideconf/tide-go/pkg/parser"
)

func main() {
	cfg, err := tide.NewTIDE("./examples/example.tide")
	if err != nil {
		log.Fatal(err)
	}

	// Retrieve a string value
	password, err := cfg.GetString("database.credentials.password")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database Password:", password)

	port, err := cfg.GetInt("database.port")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database Port:", port)

	// Retrieve an array of strings
	features, err := cfg.GetArray("myApp.features")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Features:")
	for _, feature := range features {
		fmt.Println("-", feature)
	}

	numbers, err := cfg.GetArray("myApp.numbers")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Numbers:")
	for _, number := range numbers {
		fmt.Println("-", number)
	}

	logLevel, err := cfg.GetString("logging.level")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Logging Level:", logLevel)
}
