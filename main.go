package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
)

const (
	baseURL       = "https://api.hibob.com/v1"
	serviceUserID = "SERVICE-USER_ID"
	apiToken      = "API_TOKEN"
)

// Create a reusable function to add headers
func addHeaders(req *http.Request) {
	// Encode credentials in Base64
	credentials := serviceUserID + ":" + apiToken
	encodedCredentials := base64.StdEncoding.EncodeToString([]byte(credentials))

	// Adding request headers including auth using the  encoded credentials
	req.Header.Add("Authorization", "Basic "+encodedCredentials)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
}

func main() {

	//Make the API call to obtain the employee profile from the HiBob API
	profiles, err := getEmployeeProfiles()
	if err != nil {
		fmt.Println("Error getting profiles:", err)
		return
	}

	//Here you are simply looping through the returned records and printing the details to logs.
	for _, employee := range profiles {
		fmt.Printf("ID: %s, Name: %s %s, Email: %s\n", employee.ID, employee.FirstName, employee.Surname, employee.Email)
		fmt.Printf("Display Name: %s\n", employee.DisplayName)
		fmt.Printf("Start Date: %s\n", employee.Work.StartDate)
		fmt.Printf("Manager: %s\n", employee.Work.Manager)
		fmt.Println("---")
		// Process other employee data
	}

	//Make the API call to obtain details all employees that are out of office from the HiBob API
	fromDate := "2024-09-16"
	toDate := "2024-09-25"
	includePending := true
	outOfOfficeData, err := getWhoIsOutOfOffice(fromDate, toDate, includePending)
	if err != nil {
		fmt.Println("Error getting out of office data:", err)
		return
	}

	for _, out := range outOfOfficeData.Outs {
		fmt.Printf("Employee: %s\n", out.EmployeeDisplayName)
		fmt.Printf("Start Date: %s\n", out.StartDate)
		fmt.Printf("End Date: %s\n", out.EndDate)
		fmt.Printf("Policy Type: %s\n", out.PolicyTypeDisplayName)
		fmt.Printf("Status: %s\n", out.Status)
		if out.PercentageOfDay != 0 {
			fmt.Printf("Percentage of day: %f\n", out.PercentageOfDay)
		}
		fmt.Println("---")
	}

	//creating and sending time-off request
	timeOffRequest := TimeOffRequest{
		StartDatePortion: "all_day",
		EndDatePortion:   "morning",
		RequestRangeType: "days",
		PolicyType:       "Holiday",
		StartDate:        "2024-05-29",
		EndDate:          "2024-06-03",
		Description:      "Worn out from work",
	}

	//replace employee id with real one
	responseBody, err := submitTimeOffRequest("2983224227961766275", timeOffRequest)
	if err != nil {
		fmt.Println("Error submitting time off request:", err)
		return
	}
	fmt.Println(string(responseBody))

}
