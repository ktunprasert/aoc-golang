package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

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
	sessionCookie := http.Cookie{
		Name:     "session",
		Value:    cookie,
		Path:     "/",
		Domain:   ".adventofcode.com",
		Expires:  time.Now().Add(time.Duration(24 * 365 * time.Hour)),
		Secure:   true,
		HttpOnly: true,
	}
	req.AddCookie(&sessionCookie)

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to download input: %s", res.Status)
	}

	file, err := os.Create(fmt.Sprintf("input/%s-%s.txt", year, day))
	if err != nil {
		return err
	}

	defer file.Close()

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	fmt.Printf("Truncated input:\n%s...\n", string(bytes[:100]))

	written, err := file.Write(bytes)
	if err != nil {
		return err
	}

	fmt.Printf("Wrote %d bytes to input/%s-%s.txt\n", written, year, day)

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
