package handler

import (
	"errors"
	"strings"
)

type StudentRequest struct {
	Name     string `json:"name"`
	Address  string `json:"address"`
	City     string `json:"city"`
	Country  string `json:"country"`
	Pincode  string `json:"pincode"`
	SatScore int64  `json:"sat_score"`
	Passed   string `json:"passed"`
}

// If "PASS" -> 1, if "FAIL" -> 0. Returns an error if the input is invalid.
func (sr *StudentRequest) ConvertPassed() (bool, error) {
	// Normalize the input to uppercase to accept "pass", "Pass", "FAIL", etc.
	switch strings.ToUpper(sr.Passed) {
	case "PASS":
		return true, nil
	case "FAIL":
		return false, nil
	default:
		return false, errors.New("invalid value for 'passed': must be 'PASS' or 'FAIL'")
	}
}

type UpdateScoreRequest struct {
	Name     string `json:"name"`
	SatScore int    `json:"sat_score"`
	Passed   bool   `json:"passed"`
}
