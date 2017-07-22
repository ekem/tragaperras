package minecraft

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	. "trasto"
)

const (
	SERVER = iota
	CLIENT
)

type Latest struct {
	Snapshot string `json:"snapshot"`
	Release  string `json:"release"`
}

type Version struct {
	Id          string `json:"id"`
	Type        string `json:"type"`
	ReleaseTime string `json:"time"`
	Time        string `json:"releaseTime"`
	URL         string `json:"url"`
}

type Record struct {
	Latest   Latest    `json:"latest"`
	Versions []Version `json:"versions"`
}

type Jugar struct {
	Downloads          Downloads `json:"downloads"`
	MainClass          string    `json:"mainClass"`
	MinecraftArguments string    `json:"minecraftArguments"`
}

type Downloads struct {
	Client Download `json:"client"`
	Server Download `json:"server"`
}

type Download struct {
	Sha1 string `json:"sha1"`
	Size int    `json:"size"`
	URL  string `json:"url"`
}

func GetLatestVersion() (l Latest) {
	versions := GetVersionFile()

	var r Record
	err := json.Unmarshal(versions, &r)
	Check(err)

	l = r.Latest
	return
}

func GetPackage(p int) string {
	switch p {
	case SERVER:
		return "server"
	case CLIENT:
		return "client"
	}

	return ""
}

func GetVersionFile() []byte {
	return findVersionFile()
}

func findVersionFile() (versions []byte) {
	filename := "./tmp/version_manifest.json"
	url := "https://launchermeta.mojang.com/mc/game/version_manifest.json"

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		log.Print("Minecraft version manifest not found.")

		DownloadFile(filename, url)
	}

	versions, err := ioutil.ReadFile(filename)

	Check(err)

	return versions
}

func GetArguments(version string) (arguments string) {
	versions := GetVersionFile()

	var r Record
	err := json.Unmarshal(versions, &r)
	Check(err)

	response := MakeRequest(r.Versions[0].URL)
	buf, err := ioutil.ReadAll(response.Body)

	var j Jugar
	err = json.Unmarshal(buf, &j)
	Check(err)

	arguments = j.MinecraftArguments

	return
}

func GetJugar(id string) (j Jugar) {
	versions := GetVersionFile()

	var r Record
	err := json.Unmarshal(versions, &r)
	Check(err)

	response := MakeRequest(r.Versions[0].URL)
	buf, err := ioutil.ReadAll(response.Body)

	if Debug {
		log.Printf(string(buf))
	}

	err = json.Unmarshal(buf, &j)
	Check(err)

	return
}

func DownloadVersion(id string, location string, opt int) {
	var jugar Jugar

	jugar = GetJugar(id)

	var packageType string

	switch opt {
	case CLIENT:
		packageType = "client"
	case SERVER:
		packageType = "server"
	}

	DownloadFile(
		fmt.Sprintf("%s/%s.jar", location, packageType),
		jugar.Downloads.Server.URL,
	)
}
