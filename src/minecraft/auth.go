package minecraft

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	. "trasto"
)

type Agent struct {
	Name    string `json:"name"`
	Version int    `json:"version"`
}

type Payload struct {
	Agent       Agent  `json:"agent"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	ClientToken string `json:"clientToken"`
	RequestUser bool   `json:"requestUser"`
}

type AuthResponse struct {
	Error             string    `json:"error"`
	ErrorMessage      string    `json:"errorMessage"`
	Cause             string    `json:"cause"`
	User              User      `json:"user"`
	AvailableProfiles []Profile `json:"availableProfile"`
	SelectedProfile   Profile   `json:"selectedProfile"`
	AccessToken       string    `json:"accessToken"`
	ClientToken       string    `json:"clientToken"`
}

type Profile struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Legacy bool   `json:"legacy"`
}

type Property struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type User struct {
	Id         string     `json:"id"`
	Properties []Property `json:"properties"`
}

var ErrorAuthFailed = errors.New("Authentication failed.")

type jsonResponse struct {
	ID string `json:"id"`
}

func Authenticate(
	username,
	password string) (string, error) {

	url := "https://authserver.mojang.com"

	payload := Payload{
		Agent: Agent{
			Name:    "Minecraft",
			Version: 1,
		},
		Username:    username,
		Password:    password,
		ClientToken: "100",
		RequestUser: true,
	}

	b, err := json.Marshal(payload)

	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

	// Create a new POST request.
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/authenticate", url), bytes.NewBuffer(b))

	// Set custem header for agent.
	req.Header.Set("X-Custom-Header", "lapis2")
	// Set request type to 'application/json'.
	req.Header.Set("Content-Type", "application/json")

	// Create an instance of an http client.
	client := http.Client{}
	// Capture the response and error for the http making the POST request.
	response, err := client.Do(req)
	Check(err)

	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	defer response.Body.Close()

	if Debug {
		log.Print("Printing Status: ", response.Status)
		log.Print("Printing Header: ", response.Header)
		log.Print("Printing Body: ", string(body))
	}

	var a AuthResponse
	err = json.Unmarshal(body, &a)
	Check(err)

	if a.Error != "" {
		log.Fatal(a.ErrorMessage)
	}

	if Debug {
		log.Print(a.AccessToken)
	}

	return a.AccessToken, nil
}
