package minecraft

import (
	. "../trasto"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
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

func LatestVersion() (l Latest) {
	versions := Find_versions()

	var r Record
	err := json.Unmarshal(versions, &r)
	Check(err)

	l = r.Latest
	return
}

func Find_versions() (versions []byte) {
	filename := "./tmp/version_manifest.json"
	url := "https://launchermeta.mojang.com/mc/game/version_manifest.json"

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		log.Print("Minecraft version manifest not found.")

		Get_a_file(filename, url)
	}

	versions, err := ioutil.ReadFile(filename)

	Check(err)

	return versions
}

func Get_a_version(id string, opt int) {
	versions := Find_versions()

	var r Record
	err := json.Unmarshal(versions, &r)
	Check(err)

	var jugar Jugar
	response := Make_a_request(r.Versions[0].URL)
	buf, err := ioutil.ReadAll(response.Body)
	log.Printf("%s", buf)
	err = json.Unmarshal(buf, &jugar)
	Check(err)

	log.Print(jugar.MinecraftArguments)

	switch opt {
	case 1:
		Get_a_file("./tmp/server.jar", jugar.Downloads.Server.URL)
		fallthrough
	default:
		Get_a_file("./tmp/client.jar", jugar.Downloads.Client.URL)
	}
}
