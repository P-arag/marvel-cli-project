package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/gookit/color"

	"github.com/manifoldco/promptui"
)

func setupCheck() bool {
	file, _ := ioutil.ReadFile("config")
	lines := strings.Split(string(file), "\n")

	if len(lines) == 2 {
		return true
	}
	return false
}

func setup(override bool) {
	privatePrompt := promptui.Prompt{
		Label: "Enter Private Key",
		Mask:  '♠', 
	}
	publicPrompt := promptui.Prompt{
		Label: "Enter Private Key",
	}

	if !setupCheck() || override {
		color.Info.Prompt("This is a one time configuration for the CLI, please fill out the form below..")
		privateKey, err := privatePrompt.Run()
		publicKey, err := publicPrompt.Run()
		if err != nil {
			color.Error.Println("Prompt failed %v\n", err)
			return
		}
		fmt.Println("Your Credentials...")
		fmt.Println("Private Key <: " + privateKey)
		fmt.Println("Public Key <: " + publicKey) 
		s := []byte("private: " + privateKey + "\npublic: " + publicKey)
		ioutil.WriteFile("config", s, 0644)
	}

}
