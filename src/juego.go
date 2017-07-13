package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	// Set a path to a single executable/script or binary.
	var mc_path = flag.String("path", "C:/Program Files (x86)/Minecraft", "Path to executable")

	// The name of the target to be executed.
	var mc_target = "MinecraftLauncher.exe"

	// Produce a path in the filesystem to exec.
	exec_line := fmt.Sprintf("%s/%s", *mc_path, mc_target)

	// Print the line to stdout.
	fmt.Printf("%s\n", exec_line)

	// Execute the command.
	cmd := exec.Command(exec_line)

	// Display stdout and stderr as combined output.
	stdoutStderr, err := cmd.CombinedOutput()

	// If go errors before the executed code returns 0, then log and exit as a
	// fatal error.
	if err != nil {
		log.Fatal(err)
	}

	// Print the combined output of the exec to the current shell's stdout.
	fmt.Printf("%s\n", stdoutStderr)
}
