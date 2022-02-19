package main

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

var ErrNotEnoughoptions = errors.New("At least one option is required")
var inputReader = bufio.NewReader(os.Stdin)

type Askable struct {
	Name string
	URL  string
}

func readLine() (string, error) {
	input, err := inputReader.ReadString('\n')
	if err != nil {
		return "", err
	}
	input = strings.ReplaceAll(input, "\r", "")
	input = strings.ReplaceAll(input, "\n", "")
	return input, nil
}

func MakeMenu(noDefault bool, options []*Askable) (int, error) {
	if len(options) == 0 {
		return -1, ErrNotEnoughoptions
	}

	run := true
	for run {
		for i, o := range options {
			if noDefault {
				color.Cyan("- [%d] %s", i+1, o.Name)
			} else {
				color.Cyan("- [%d] %s", i, o.Name)
			}
		}
		if noDefault {
			color.New(color.FgBlue).Printf("[?] Your option [1-%d]: ", len(options))
		} else {
			color.New(color.FgBlue).Printf("[?] Your option [0-%d] (default: 0): ", len(options)-1)
		}
		input, err := readLine()
		if err != nil {
			return -1, err
		}

		if inputN, err := strconv.ParseInt(input, 10, 32); err == nil {
			n := int(inputN)
			if noDefault {
				n -= 1
			}

			if n >= len(options) || n < 0 {
				if noDefault {
					color.New(color.FgYellow).Printf("[!] Option %d was not found.\n", inputN)
					continue
				}
				color.New(color.FgYellow).Printf("[!] Option %d was not found, falling back on default.\n", inputN)
				return 0, nil
			}

			color.New(color.FgGreen).Printf("[+] %s selected.\n", options[n].Name)
			return n, nil
		} else if noDefault {
			color.New(color.FgYellow).Printf("[!] Invalid option.")
		}
		run = noDefault
	}

	color.New(color.FgGreen).Printf("[+] Default option selected.\n")
	return 0, nil
}
