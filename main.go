package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/c-bata/go-prompt"
	"github.com/manifoldco/promptui"
)

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "get", Description: "Command: Gets the query"},
		{Text: "nameStartsWith", Description: "Query: Gets characters with name"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func main() {

	// data := get_data("characters?nameStartsWith=iron")
	// m, ok := gjson.Parse(data).Value().(map[string]interface{})

	// if ok {
	// 	fmt.Println(m["copyright"])
	// }
	validate := func(input string) error {
		_, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return errors.New("Invalid number")
		}
		return nil
	}

	myprompt := promptui.Prompt{
		Label:    "Number",
		Validate: validate,
	}

	fmt.Println("Please select table.")
	t := prompt.Input("marvel> ", completer)
	fmt.Println("You selected " + t)

	result, err := myprompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose %q\n", result)

}
