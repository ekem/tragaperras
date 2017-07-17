package main

import (
	"./minecraft"
	"log"
	"os"
)

func main() {
	installDir := "./tmp"

	if err := os.Mkdir(installDir, os.ModeDir); !os.IsExist(err) {
		log.Print("Creating ", installDir)
	}

	l := minecraft.LatestVersion()
	minecraft.Get_a_version(l.Snapshot, 1)
}
