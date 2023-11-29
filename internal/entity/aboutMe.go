package entity

import "fmt"

type AboutMe struct {
	Id   int
	Text string
}

const (
	AboutMeTextError = "'text' should be between 1-1000 characters."
)

func NewAboutMe(id int, text string) (*AboutMe, error) {
	err := validateText(text)
	if err != nil {
		return nil, err
	}

	return &AboutMe{
		Id:   id,
		Text: text,
	}, nil
}

func validateText(text string) error {
	if len(text) < 1 || len(text) > 1000 {
		return fmt.Errorf(AboutMeTextError)
	}
	return nil
}
