package schedule

import (
	"fmt"
	"time"

	sel "github.com/chazari-x/hmtpk_get_groups/selenium"
	log "github.com/sirupsen/logrus"
	"github.com/tebeka/selenium"
)

type Schedule struct {
	selenium *sel.Selenium
}

func NewSchedule(selenium *sel.Selenium) *Schedule {
	return &Schedule{selenium}
}

func (s *Schedule) GetGroups() []string {
	var groups []string

	if err := s.selenium.OpenURL("https://hmtpk.ru/ru/students/schedule/"); err != nil {
		log.Error(err)
		return nil
	}

	if err := s.selenium.ClickToElement(selenium.ByCSSSelector, "#select2-group-container"); err != nil {
		log.Error(err)
		return nil
	}

	for i := 2; ; i++ {
		id, err := s.selenium.GetElementID(selenium.ByCSSSelector, fmt.Sprintf("#select2-group-results > li:nth-child(%d)", i))
		if err != nil {
			log.Error(err)
			break
		}

		name, err := s.selenium.GetElementText(selenium.ByCSSSelector, fmt.Sprintf("#select2-group-results > li:nth-child(%d)", i))
		if err != nil {
			log.Error(err)
			break
		}

		groups = append(groups, fmt.Sprintf("{label: '%s', value: '%s'},", name, id))
	}

	return groups
}

func (s *Schedule) GetTeachers() []string {
	var teachers []string

	if err := s.selenium.OpenURL("https://hmtpk.ru/ru/teachers/schedule/"); err != nil {
		log.Error(err)
		return nil
	}

	time.Sleep(time.Second * 5)

	if err := s.selenium.ClickToElement(selenium.ByCSSSelector, "#zstfiltr > div > div:nth-child(1) > span > span.selection > span > span.select2-selection__rendered"); err != nil {
		log.Error(err)
		return nil
	}

	for i := 3; ; i++ {
		id, err := s.selenium.GetElementID(selenium.ByCSSSelector, fmt.Sprintf("body > span > span > span.select2-results > ul.select2-results__options > li:nth-child(%d)", i))
		if err != nil {
			log.Error(err)
			break
		}

		name, err := s.selenium.GetElementText(selenium.ByCSSSelector, fmt.Sprintf("body > span > span > span.select2-results > ul.select2-results__options > li:nth-child(%d)", i))
		if err != nil {
			log.Error(err)
			break
		}

		teachers = append(teachers, fmt.Sprintf("{label: '%s', value: '%s'},", name, id))
	}

	return teachers
}
