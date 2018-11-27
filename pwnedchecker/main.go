package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Usage: pwnedchecker example@example.com")
		os.Exit(1)
	}

	email := args[0]
	breaches, err := fetchBreaches(email)
	if err != nil {
		panic(err)
	}

	for _, breach := range breaches {
		fmt.Printf("PWNED: %s at %s\n", breach.Domain, breach.BreachDate)
	}

	if len(breaches) != 0 {
		fmt.Printf("You have been pwned a total of %d time(s).\n", len(breaches))
	} else {
		fmt.Println("You have not been pwned!")
	}
}

type breaches []struct {
	Domain     string `json:"Domain"`
	BreachDate string `json:"BreachDate"`
}

func fetchBreaches(email string) (breaches, error) {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	request, err := http.NewRequest(
		"GET",
		fmt.Sprintf("https://haveibeenpwned.com/api/v2/breachedaccount/%s", email),
		nil,
	)
	if err != nil {
		return breaches{}, err
	}
	request.Header.Set("User-Agent", "pwnedChecker v1.0")

	response, err := client.Do(request)
	if err != nil {
		return breaches{}, err
	}
	defer response.Body.Close()

	body := breaches{}
	json.NewDecoder(response.Body).Decode(&body)

	return body, nil
}
