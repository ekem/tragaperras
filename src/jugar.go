package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

type Juga struct {
	Name, Path, Target string
}

func main() {
	// Create an object to hold game information.
	juga := Juga{}

	// Set a path to a single executable/script or binary.
	juga.Path = *flag.String(
		"path", "C:/Program Files (x86)/Minecraft", "Path to executable.")

	// The name of the target to be executed.
	juga.Target = *flag.String(
		"target", "MinecraftLauncher.exe", "Target name.")

	// Determine if help has been called.
	if flag.ErrHelp != nil {
		flag.PrintDefaults()
		os.Exit(0)
	}

	// Produce a path in the filesystem to exec.
	exec_line := fmt.Sprintf("%s/%s", juga.Path, juga.Target)

	// Print the line to stdout.
	fmt.Printf("%s\n", exec_line)

	// Execute the command.
	cmd := exec.Command("date")

	// Display stdout and stderr as combined output.
	output, err := cmd.CombinedOutput()

	// If go errors before the executed code returns 0, then log and exit as a
	// fatal error.
	if err != nil {
		log.Fatal(err)
	}

	// Print the combined output of the exec to the current shell's stdout.
	fmt.Printf("%s\n", output)
}
