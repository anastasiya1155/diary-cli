package entry

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Entry struct {
	Date      string    `json:"date"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

func New(date, body string) (Entry, error) {
	if date == "" || body == "" {
		return Entry{}, errors.New("date and body are required")
	}

	return Entry{
		Date:      date,
		Body:      body,
		CreatedAt: time.Now(),
	}, nil
}

func (entry Entry) Display() {
	fmt.Println(entry.Date)
	fmt.Println(entry.Body)
}

func (entry Entry) Save() error {
	fileName := strings.ReplaceAll(entry.Date, "/", "-") + ".json"

	json, err := json.Marshal(entry)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}
