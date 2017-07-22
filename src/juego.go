package main

import (
	"flag"
	"fmt"
	"log"
	"minecraft"
	"os"
	. "trasto"
)

var (
	auth      = flag.Bool("auth", false, "Authenticates a username and password.")
	latest    = flag.Bool("latest", false, "Get the latest version number.")
	debug     = flag.Bool("debug", false, "Prints debug output to stdout.")
	client    = flag.Bool("client", false, "Downloads the client jar.")
	arguments = flag.Bool("arguments", false, "Get the arguments for a version.")
	server    = flag.Bool("server", false, "Downloads the server jar.")
)

/**
 * Authenticate a Minecraft user via the public mojang authentication system.
 */
func authUser() {
	accessToken, err := minecraft.Authenticate(
		os.Getenv("MC_USERNAME"),
		os.Getenv("MC_PASSWORD"),
	)

	Check(err)

	fmt.Print(accessToken)
}

/**
 * Get the arguments for a Minecraft client jar.
 */
func getArguments(version string) (arguments string) {
	j := minecraft.GetJugar(getLatest())

	fmt.Printf(
		"Main class:\n%s\n\nArguments:\n%s\n",
		j.MainClass,
		j.MinecraftArguments,
	)

	return j.MinecraftArguments
}

func getLatest() (version string) {
	return minecraft.GetLatestVersion().Snapshot
}

// Initialize the directory to which items are downloaded.
func initFilesystem() {
	// Define a working directory.
	installDir := "./tmp"

	// Create the working directory if it does not exist.
	if err := os.Mkdir(installDir, os.ModeDir); !os.IsExist(err) {
		log.Print("Creating ", installDir)
	}
}

// Initialize an environment for a Jugar.
func initEnv(space string) (location string) {
	location = fmt.Sprintf("./tmp/%s", space)
	if err := os.MkdirAll(location, os.ModePerm); err != nil {
		log.Print(err)
	}
	return
}

func installEnv(p int) {
	location := initEnv(minecraft.GetPackage(p))
	minecraft.DownloadVersion(getLatest(), location, p)
}

func main() {
	flag.Parse()

	switch {
	case *latest:
		log.Print(getLatest())
	case *auth:
		authUser()
	case *arguments:
		getArguments(getLatest())
	case *client:
		installEnv(minecraft.CLIENT)
	case *server:
		installEnv(minecraft.SERVER)
	}
}
