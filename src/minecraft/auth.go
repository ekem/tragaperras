package minecraft

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

//var Instance = Authenticator{}

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

type Response struct {
	AccessToken string `json:"accessToken"`
	ClientToken string `json:"clientToken"`
}

type AuthResponse struct {
	Error        string `json:"error"`
	ErrorMessage string `json:"errorMessage"`
	Cause        string `json:"cause"`
	AccessToken  string `json:"accessToken"`
	ClientToken  string `json:"clientToken"`
}

type Profile struct {
	ID   string `json:"ID"`
	Name string `json:"name"`
}

var ErrorAuthFailed = errors.New("Authentication failed.")

type jsonResponse struct {
	ID string `json:"id"`
}

func Authenticate(
	username,
	serverID,
	sharedSecret string, publicKey []byte) (string, error) {
	sha := sha1.New()
	sha.Write([]byte(serverID))
	sha.Write([]byte(sharedSecret))
	sha.Write(publicKey)
	hash := sha.Sum(nil)

	negative := (hash[0] & 0x80) == 0x80

	if negative {
		twosCompliment(hash)
	}

	buf := hex.EncodeToString(hash)

	if negative {
		buf = "-" + buf
	}

	hashString := strings.TrimLeft(buf, "0")

	url := "https://authserver.mojang.com"

	log.Print(string(sharedSecret))

	payload := Payload{
		Agent: Agent{
			Name:    "Minecraft",
			Version: 1,
		},
		Username:    username,
		Password:    string(sharedSecret),
		ClientToken: hashString,
		RequestUser: true,
	}

	b, err := json.Marshal(payload)

	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/authenticate", url), bytes.NewBuffer(b))

	req.Header.Set("X-Custom-Header", "lapis")
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	response, err := client.Do(req)

	if err != nil {
		panic(err)
		os.Exit(-1)
	}

	defer response.Body.Close()

	log.Print("Printing Status: ", response.Status)

	log.Print("Printing Header: ", response.Header)

	body, _ := ioutil.ReadAll(response.Body)

	log.Print("Printing Body: ", string(body))

	defer response.Body.Close()

	var a AuthResponse
	err = json.Unmarshal(body, &a)
	Check(err)
	/*
		dec := json.NewDecoder(response.Body)
		res := &jsonResponse{}
		err = dec.Decode(res)

		if err != nil {
			return "", ErrorAuthFailed
		}

		if len(res.ID) != 32 {
			return "", ErrorAuthFailed
		}

		return res.ID, nil
	*/
}

func twosCompliment(p []byte) {
	carry := true
	for i := len(p) - 1; i >= 0; i-- {
		p[i] = ^p[i]
		if carry {
			carry = p[i] == 0xFF
			p[i]++
		}
	}
}
