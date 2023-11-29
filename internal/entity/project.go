package entity

import (
	"fmt"
	"regexp"
)

type Project struct {
	Id          int
	Name        string
	Description string
	StartDate   string
	EndDate     string
	DateOrder   string
	Skills      []*Skill
	Link        string
}

const (
	ProjectNameError        = "'name' should be between 3-30 characters."
	ProjectStartDateError   = "'startDate' should be in format 'MM/YYYY'."
	ProjectEndDateError     = "'endDate' should be in format 'MM/YYYY'."
	ProjectLinkError        = "'link' should be between 4-250 characters."
	ProjectStartDatePattern = "^(0[1-9]|1[0-2])/20[0-9]{2}$"
	ProjectEndDatePattern   = "^(0[1-9]|1[0-2])/20[0-9]{2}$"
)

func NewProject(id int, name, description, startDate, endDate, link string) (*Project, error) {
	err := validateProjectName(name)
	if err != nil {
		return nil, err
	}
	err = validateProjectStartDate(startDate)
	if err != nil {
		return nil, err
	}

	err = validateProjectEndDate(endDate)
	if err != nil {
		return nil, err
	}

	err = validateProjectLink(link)
	if err != nil {
		return nil, err
	}
	project := &Project{
		Id:          id,
		Name:        name,
		Description: description,
		StartDate:   startDate,
		EndDate:     endDate,
		Link:        link,
	}
	project.CreateOrder()
	return project, nil
}

func validateProjectName(name string) error {
	if len(name) < 3 || len(name) > 30 {
		return fmt.Errorf(ProjectNameError)
	}
	return nil
}

func validateProjectStartDate(startDate string) error {
	if matched, _ := regexp.MatchString(ProjectStartDatePattern, startDate); !matched {
		return fmt.Errorf(ProjectStartDateError)
	}
	return nil
}

func validateProjectEndDate(endDate string) error {
	if matched, _ := regexp.MatchString(ProjectEndDatePattern, endDate); !matched {
		return fmt.Errorf(ProjectEndDateError)
	}
	return nil
}

func validateProjectLink(link string) error {
	if len(link) < 4 || len(link) > 250 {
		return fmt.Errorf(ProjectLinkError)
	}
	return nil
}

func (p *Project) CreateOrder() {
	p.DateOrder = p.StartDate[3:] + p.StartDate[:2] + p.EndDate[3:] + p.EndDate[:2]
}

func (p *Project) AddSkill(skill *Skill) {
	skill.Project = append(skill.Project, p)
	p.Skills = append(p.Skills, skill)
}
