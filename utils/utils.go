package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func ErrorCheck(e error) {
	if e != nil {
		panic(e)
	}
}

func GetInputData(day int) string {
	client := http.Client{}
	url := fmt.Sprintf("https://adventofcode.com/2023/day/%d/input", day)
	session := os.Getenv("AOC_SESSION")

	req, err := http.NewRequest("GET", url, nil)
	ErrorCheck(err)

	req.Header = http.Header{
		"Content-Type": {"application/json"},
		"Cookie":       {fmt.Sprintf("session=%s", session)},
	}

	res, err := client.Do(req)
	ErrorCheck(err)

	body, err := io.ReadAll(res.Body)
	ErrorCheck(err)

	return strings.TrimSuffix(string(body), "\n")
}
