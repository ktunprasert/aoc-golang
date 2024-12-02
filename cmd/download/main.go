package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/profclems/go-dotenv"
)

var (
	cookie    string
	AOC_INPUT = "https://adventofcode.com/%s/day/%s/input"
)

func downloadInput(year, day string) error {
	req, err := http.NewRequest("GET", fmt.Sprintf(AOC_INPUT, year, day), nil)
	if err != nil {
		return err
	}

	req.Header.Set("Cookie", cookie)

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	file, err := os.Create(fmt.Sprintf("input/%s-%s.txt", year, day))
	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}

func main() {
	err := dotenv.Load()
	if err != nil {
		panic(err)
	}

	cookie = dotenv.GetString("COOKIE")

	argv := os.Args[1:]
	if len(argv) != 2 {
		fmt.Println("Usage: download <year> <day>")
		os.Exit(1)
	}
	year := argv[0]
	day := argv[1]
	fmt.Printf("Downloading input for %s %s\n", year, day)
	err = downloadInput(year, day)
	if err != nil {
		panic(err)
	}
}
