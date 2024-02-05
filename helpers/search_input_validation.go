package helpers

import (
	"event-hunters/dto"
	"fmt"
	"strconv"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/gofiber/fiber/v2"
)

func CategoriesFormHandler(c *fiber.Ctx) []int {
	var selectedCategoriesId []int
	checkboxes := []string{
		"Entertainment & Performance",
		"Art & Culture",
		"Charity",
		"Sports",
		"Competition",
		"Education & Career",
		"Expo",
	}
	for _, checkbox := range checkboxes {
		if value := c.FormValue(checkbox); value != "" {
			intValue, err := strconv.Atoi(value)
			if err != nil {
				// Handle the error if the conversion fails
				fmt.Println("Error converting string to int:", err)
				continue
			}
			selectedCategoriesId = append(selectedCategoriesId, intValue)
		}
	}
	fmt.Println("THIS IS SELECTED CATEGORIES FORM HANDLER")
	fmt.Println(selectedCategoriesId)
	return selectedCategoriesId
}

func ParseSearchDate(minRegDateStr string, maxRegDateStr string, minEventStartDateStr string, maxEventStartDateStr string) (dto.ParsedEventSearchDate, error) {
	// Parse date strings into time.Time objects
	ParsedEventSearchDate := dto.ParsedEventSearchDate{}
	var minRegDate time.Time
	var maxRegDate time.Time
	var minEventStartDate time.Time
	var maxEventStartDate time.Time
	var err error

	// Parse date strings into time.Time objects
	if minRegDateStr != "" {
		minRegDate, err = time.Parse("01/02/2006", minRegDateStr)
		if err != nil {
			return ParsedEventSearchDate, err
		}
	}
	if maxRegDateStr != "" {
		maxRegDate, err = time.Parse("01/02/2006", minRegDateStr)
		if err != nil {
			return ParsedEventSearchDate, err
		}
	}

	if minEventStartDateStr != "" {
		minEventStartDate, err = time.Parse("01/02/2006", minEventStartDateStr)
		if err != nil {
			return ParsedEventSearchDate, err
		}

	}
	if maxEventStartDateStr != "" {
		maxEventStartDate, err = time.Parse("01/02/2006", maxEventStartDateStr)
		if err != nil {
			return ParsedEventSearchDate, err
		}
	}

	ParsedEventSearchDate.MinRegDate = minRegDate
	ParsedEventSearchDate.MaxRegDate = maxRegDate
	ParsedEventSearchDate.MinEventStartDate = minEventStartDate
	ParsedEventSearchDate.MaxEventStartDate = maxEventStartDate

	return ParsedEventSearchDate, nil
}

func DateValidation(ParsedDateResponse dto.ParsedEventSearchDate) error {
	fmt.Println(ParsedDateResponse)
	if !ParsedDateResponse.MinRegDate.IsZero() && !ParsedDateResponse.MaxRegDate.IsZero() {
		if ParsedDateResponse.MinRegDate.After(ParsedDateResponse.MaxRegDate) {
			return errors.New("minDate cannot be more than maxDate")
		}
	}

	if !ParsedDateResponse.MinEventStartDate.IsZero() && !ParsedDateResponse.MaxEventStartDate.IsZero() {
		if ParsedDateResponse.MinEventStartDate.After(ParsedDateResponse.MaxEventStartDate) {
			return errors.New("minDate cannot be more than maxDate")
		}
	}
	return nil
}
