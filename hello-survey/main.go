package main

import (
	"fmt"
	"log"

	"github.com/AlecAivazis/survey/v2"
)

// the questions to ask
var qs = []*survey.Question{
	{
		Name: "username",
		Prompt: &survey.Input{
			Message: "Please type your username:",
		},
		Validate: survey.Required,
	},
	{
		Name: "password",
		Prompt: &survey.Password{
			Message: "Please type your password:",
		},
		Validate: survey.Required,
	},
	{
		Name: "exampleDropDown",
		Prompt: &survey.Select{
			Message: "Example dropdown:",
			Options: []string{"one", "two", "three"},
			Default: "red",
		},
	},
}

func main() {
	// the answers will be written to this struct
	answers := struct {
		Username        string // survey will match the question and field names
		ExampleDropDown string `survey:"exampleDropDown"` // or you can tag fields to match a specific name
		Password        string `survey:"password"`        // or you can tag fields to match a specific name
	}{}

	// perform the questions
	err := survey.Ask(qs, &answers)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("Welcome %s!\n", answers.Username)
}
