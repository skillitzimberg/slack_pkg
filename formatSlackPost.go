package main

import (
	"github.com/apptreesoftware/go-workflow/pkg/step"
)

type FormatSlackPost struct {

}

func (f FormatSlackPost) Name() string {
	return "formatSlackPost"
}

func (f FormatSlackPost) Version() string {
	return "1.0"
}

type FormattedPost struct {
	Post string
}

func (f FormatSlackPost) Execute(in step.Context) (interface{}, error) {

	input := FormatSlackPost{}
	err := in.BindInputs(&input)
	if err != nil {
		log.Fatal(err, "BindInputs: ")
		return nil, err
	}
		formattedPost := `{
			"text": "Vacasa"
		}`
	output := &FormatSlackPost{
		Post: formattedPost,
	}
	return output, nil
}

//func (f FormatSlackPost) execute(input FormatSlackPost) ( *FormatSlackPost,
//	error) {
//	formattedPost := `{
//		"text": "Vacasa"
//	}`
//
//	return formattedPost, nil
//}

// STEP_NAME=formatSlackPost STEP_VERSION=1.0 go run slack_pkg/