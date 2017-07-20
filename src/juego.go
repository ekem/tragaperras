package main

import (
	"log"
	"minecraft"
	"os"
)

func auth_user() {
	minecraft.Authenticate(
		os.Getenv("MC_USERNAME"),
		os.Getenv("MC_SERVERID"),
		os.Getenv("MC_PASSWORD"),
		[]byte(""),
	)
}

func init() {
	// Define a working directory.
	installDir := "./tmp"

	// Create the working directory.
	if err := os.Mkdir(installDir, os.ModeDir); !os.IsExist(err) {
		log.Print("Creating ", installDir)
	}
}

func install_server_and_client() {
	// Get the latest Minecraft version.
	l := minecraft.LatestVersion()

	// Download the latest version
	minecraft.Get_a_version(l.Snapshot, 1)
}

func main() {
	auth_user()
}
