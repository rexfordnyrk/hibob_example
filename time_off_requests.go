package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type OutOfOffice struct {
	PolicyTypeDisplayName string  `json:"policyTypeDisplayName"`
	Visibility            string  `json:"visibility"`
	EndDate               string  `json:"endDate"`
	RequestId             int     `json:"requestId"`
	PolicyType            string  `json:"policyType"`
	StartPortion          string  `json:"startPortion,omitempty"`
	EmployeeId            string  `json:"employeeId"`
	EmployeeDisplayName   string  `json:"employeeDisplayName"`
	EndPortion            string  `json:"endPortion,omitempty"`
	Type                  string  `json:"type"`
	StartDate             string  `json:"startDate"`
	Status                string  `json:"status"`
	PercentageOfDay       float64 `json:"percentageOfDay,omitempty"`
}

type OutOfOfficeResponse struct {
	Outs []OutOfOffice `json:"outs"`
}

func getWhoIsOutOfOffice(fromDate, toDate string, includePending bool) (*OutOfOfficeResponse, error) {

	urlObj, err := url.Parse(baseURL + "/timeoff/whosout")
	if err != nil {
		return nil, err
	}

	//setting query params
	query := urlObj.Query()
	query.Set("from", fromDate)
	query.Set("to", toDate)
	if includePending {
		query.Set("includePending", "true")
	}
	urlObj.RawQuery = query.Encode()

	req, err := http.NewRequest("GET", urlObj.String(), nil)
	if err != nil {
		return nil, err
	}
	//Attaching the auth headers
	addHeaders(req)

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

	var outOfOfficeResponse *OutOfOfficeResponse
	err = json.Unmarshal(body, &outOfOfficeResponse)
	if err != nil {
		return nil, err
	}

	return outOfOfficeResponse, nil
}

type TimeOffRequest struct {
	StartDatePortion string `json:"startDatePortion"`
	EndDatePortion   string `json:"endDatePortion"`
	RequestRangeType string `json:"requestRangeType"`
	PolicyType       string `json:"policyType"`
	StartDate        string `json:"startDate"`
	EndDate          string `json:"endDate"`
	Description      string `json:"description"`
}

func submitTimeOffRequest(eID string, request TimeOffRequest) ([]byte, error) {

	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", baseURL+"/timeoff/employees/"+eID+"/requests", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	addHeaders(req)

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

	return body, nil
}
