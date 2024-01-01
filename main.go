package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"example.com/diary-cli/entry"
)

func main() {
	today := time.Now().Format("2006/01/02")
	datePrompt := fmt.Sprintf("Date (%s):", today)
	date := getUserInput(datePrompt)
	if date == "" {
		date = today
	}

	body := getUserInput("How was your day?")

	var dayEntry entry.Entry

	dayEntry, err := entry.New(date, body)

	if err != nil {
		fmt.Println(err)
		return
	}

	dayEntry.Display()
	e := dayEntry.Save()
	if e != nil {
		fmt.Println("saving note failed")
		return
	}
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
