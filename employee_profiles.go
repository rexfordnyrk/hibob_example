package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Employee A struct tyoe for an employee profile matching the json structure
type Employee struct {
	ID          string `json:"id"`
	FirstName   string `json:"firstName"`
	Surname     string `json:"surname"`
	Email       string `json:"email"`
	DisplayName string `json:"displayName"`
	Personal    struct {
		Honorific      string   `json:"honorific"`
		ShortBirthDate string   `json:"shortBirthDate"`
		Nationality    []string `json:"nationality"`
		Pronouns       string   `json:"pronouns"`
	} `json:"personal"`
	About struct {
		Avatar          string   `json:"avatar"`
		Hobbies         []string `json:"hobbies"`
		FoodPreferences []string `json:"foodPreferences"`
		SocialData      struct {
			Linkedin string `json:"linkedin"`
			Twitter  string `json:"twitter"`
			Facebook string `json:"facebook"`
		} `json:"socialData"`
		Superpowers []string `json:"superpowers"`
	} `json:"about"`
	Work struct {
		ShortStartDate string `json:"shortStartDate"`
		StartDate      string `json:"startDate"`
		Manager        string `json:"manager"`
		TenureDuration struct {
			PeriodISO  string `json:"periodISO"`
			SortFactor int    `json:"sortFactor"`
			Humanize   string `json:"humanize"`
		} `json:"tenureDuration"`
		Custom               []string `json:"custom"`
		DurationOfEmployment struct {
			PeriodISO  string `json:"periodISO"`
			SortFactor int    `json:"sortFactor"`
			Humanize   string `json:"humanize"`
		} `json:"durationOfEmployment"`
		ReportsToIdInCompany int `json:"reportsToIdInCompany"`
		EmployeeIdInCompany  int `json:"employeeIdInCompany"`
		ReportsTo            struct {
			DisplayName string `json:"displayName"`
			Email       string `json:"email"`
			Surname     string `json:"surname"`
			FirstName   string `json:"firstName"`
			ID          string `json:"id"`
		} `json:"reportsTo"`
		WorkMobile            string  `json:"workMobile"`
		WorkPhone             string  `json:"workPhone"`
		IndirectReports       int     `json:"indirectReports"`
		SiteID                int     `json:"siteID"`
		TenureDurationYears   float64 `json:"tenureDurationYears"`
		Department            string  `json:"department"`
		TenureYears           int     `json:"tenureYears"`
		IsManager             bool    `json:"isManager"`
		Title                 string  `json:"title"`
		Site                  string  `json:"site"`
		OriginalStartDate     string  `json:"originalStartDate"`
		ActiveEffectiveDate   string  `json:"activeEffectiveDate"`
		DirectReports         int     `json:"directReports"`
		SecondLevelManager    string  `json:"secondLevelManager"`
		DaysOfPreviousService int     `json:"daysOfPreviousService"`
		YearsOfService        float64 `json:"yearsOfService"`
	} `json:"work"`
}

type ProfilesResponse struct {
	Employees []Employee `json:"employees"`
}

func getEmployeeProfiles() ([]Employee, error) {

	//creating request
	req, err := http.NewRequest("GET", baseURL+"/profiles", nil)
	if err != nil {
		//do propper
		return nil, err
	}
	//Attaching the auth headers
	addHeaders(req)

	//Sending request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	//create a struct to hold the response.
	var employeesResponse ProfilesResponse
	err = json.Unmarshal(body, &employeesResponse)
	if err != nil {
		return nil, err
	}

	return employeesResponse.Employees, nil
}
