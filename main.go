package main

import (
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/fatih/color"

	"ya/utils"
)

func main() {
	shortcuts, err := utils.LoadShortcuts()
	if err != nil {
		color.Red("Error loading shortcuts:", err)
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		color.Red("Usage: ya <shortcut> \n for shortcuts use: ya help")
		os.Exit(1)
	}

	shortcut := os.Args[1]

	switch shortcut {
	case "version", "-v":
		version := utils.GetAppVersion()
		color.Green("Ya version: %s", version)
		return
	case "help":
		color.Green("Available shortcuts:")
		for key := range shortcuts {
			color.Yellow(" - %s :", key)
			color.Green("   %s", shortcuts[key])
		}
		color.Green("\nTo add a new shortcut use: ya add <shortcut> '<command>'")
		color.Green("To remove a shortcut use: ya remove <shortcut>")
		return
	case "add":
		if len(os.Args) < 4 {
			color.Red("Usage: ya add <shortcut> '<command>'")
			os.Exit(1)
		}
		shortcutName := os.Args[2]
		command := os.Args[3]
		if utils.IsInvalidString(shortcutName) || utils.IsInvalidString(command) {
			color.Red("Usage: ya add <shortcut> '<command>'")
			os.Exit(1)
		}
		utils.AddShortcut(shortcutName, command)
		return
	case "remove":
		if len(os.Args) < 3 {
			color.Red("Usage: ya remove <shortcut>")
			os.Exit(1)
		}
		utils.RemoveShortcut(os.Args[2])
		return
	}

	command, exists := shortcuts[shortcut]

	// i added this because i was wondering how i would have been using this
	// if i cannot pass more arguments to the shortcut
	// so hear it is
	if len(os.Args) > 2 {
		args := os.Args[2:]
		command += " " + strings.Join(args, " ")
	}

	if !exists {
		color.Red("Unknown shortcut: %s\n to add a new shortcut use: ya add <shortcut> '<command>'", shortcut)
		os.Exit(1)
	}

	// if we are here, detect if it is windows, you get?
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("powershell", "-NoProfile", "-NonInteractive", "-Command", command)
	} else {
		// for linux and macOS typeshit
		cmd = exec.Command("bash", "-c", command)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	cmdError := cmd.Run()
	if cmdError != nil {
		color.Red("Command failed: %v", cmdError)
		os.Exit(1)
	}
}
