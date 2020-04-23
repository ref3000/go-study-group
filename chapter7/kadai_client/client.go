package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/gommon/log"
)

func main() {
	err := post(123, "ほげ")
	if err != nil {
		log.Error(err)
	}
}

type Request struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
}

func post(userId int, name string) error {
	jsonBytes, err := json.Marshal(&Request{
		UserID: userId,
		Name:   name,
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		"POST",
		"http://localhost:8080/user_fortune",
		bytes.NewBuffer(jsonBytes),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}
