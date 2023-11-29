package entity

import "fmt"

type Skill struct {
	Id                  int
	Name                string
	Type                string
	SimpleIconsIconSlug string
	Project             []*Project
}

const (
	SkillNameError                = "'name' should be between 1-15 characters."
	SkillSimpleIconsIconSlugError = "'simpleIconsIconSlug' should be between 0-50 characters."
	SkillTypeError                = "Invalid 'type'."
)

var validSkillTypes = map[string]bool{
	"Language":          true,
	"Framework/Library": true,
	"Database":          true,
	"Software":          true,
}

func NewSkill(id int, name, skillType, simpleIconsIconSlug string) (*Skill, error) {
	err := validateSkillName(name)
	if err != nil {
		return nil, err
	}

	err = validateSkillType(skillType)
	if err != nil {
		return nil, err
	}

	err = validateSimpleIconsIconSlug(simpleIconsIconSlug)
	if err != nil {
		return nil, err
	}

	return &Skill{
		Id:                  id,
		Name:                name,
		Type:                skillType,
		SimpleIconsIconSlug: simpleIconsIconSlug,
	}, nil
}

func validateSkillName(name string) error {
	if len(name) < 1 || len(name) > 15 {
		return fmt.Errorf(SkillNameError)
	}
	return nil
}

func validateSkillType(skillType string) error {
	if !validSkillTypes[skillType] {
		return fmt.Errorf(SkillTypeError)
	}
	return nil
}

func validateSimpleIconsIconSlug(simpleIconsIconSlug string) error {
	if len(simpleIconsIconSlug) > 50 {
		return fmt.Errorf(SkillSimpleIconsIconSlugError)
	}
	return nil
}
