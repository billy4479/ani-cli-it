package main

import (
	"os"
	"os/exec"

	"github.com/fatih/color"
)

const (
	PLAYER_CMD         = "mpv"
	PLAYER_HEADER_FLAG = "--http-header-fields='Referer: https://www.animesaturn.it/'"
)

func PlayURL(URL string) error {
	cmd := exec.Command(PLAYER_CMD, PLAYER_HEADER_FLAG, URL)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	color.New(color.FgGreen).Printf("[+] Starting mpv...\n")

	return cmd.Run()
}
