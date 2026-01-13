package main

import (
	"fmt"
	"os"
	"os/exec"

	"ya/utils"
)

func main() {
	shortcuts, err := utils.LoadShortcuts()
	if err != nil {
		fmt.Println("Error loading shortcuts:", err)
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage: ya <shortcut> \n for shortcuts use: ya help")
		os.Exit(1)
	}

	shortcut := os.Args[1]

	switch shortcut {
	case "help":
		fmt.Println("Available shortcuts:")
		for key := range shortcuts {
			fmt.Println(" -", key, ":", shortcuts[key])
		}
		return
	case "add":
		if len(os.Args) < 4 {
			fmt.Println("Usage: ya add <shortcut> '<command>'")
			os.Exit(1)
		}
		shortcutName := os.Args[2]
		command := os.Args[3]
		if utils.IsInvalidString(shortcutName) || utils.IsInvalidString(command) {
			fmt.Println("Usage: ya add <shortcut> '<command>'")
			os.Exit(1)
		}
		utils.AddShortcut(shortcutName, command)
		return
	case "remove":
		if len(os.Args) < 3 {
			fmt.Println("Usage: ya remove <shortcut>")
			os.Exit(1)
		}
		utils.RemoveShortcut(os.Args[2])
		return
	}

	command, exists := shortcuts[shortcut]

	if !exists {
		fmt.Printf("Unknown shortcut: %s\n", shortcut+"\n to add a new shortcut use: ya add <shortcut> '<command>'")
		os.Exit(1)
	}

	cmd := exec.Command("powershell", "-NoProfile", "-NonInteractive", "-Command", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	cmdError := cmd.Run()
	if cmdError != nil {
		fmt.Println("Command failed:", cmdError)
		os.Exit(1)
	}
}
