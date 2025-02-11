package handler

import (
	"fmt"
)

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)

}

// Create Opening
type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" && r.Company == "" && r.Location == "" && r.Link == "" && r.Salary == 0 && r.Remote == nil {
		return fmt.Errorf("malformed request body or empty params")
	}

	if r.Role == "" {
		return errParamIsRequired("role", "string")
	}
	if r.Company == "" {
		return errParamIsRequired("company", "string")
	}
	if r.Location == "" {
		return errParamIsRequired("location", "string")
	}
	if r.Link == "" {
		return errParamIsRequired("link", "string")
	}

	// **

	if r.Remote == nil {
		return errParamIsRequired("remote", "boll")
	}

	if r.Salary == 0 {
		return errParamIsRequired("salary", "int64")
	}
	return nil
}

// Update Opening
type UpdateOpeningRequest struct {
	ID       uint   `json:"id"`
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *UpdateOpeningRequest) Validate() error {
	if r.Role != "" && r.Company != "" && r.Location != "" && r.Link != "" && r.Salary >= 0 && r.Remote != nil {
		return nil
	}
	return fmt.Errorf("at least one valid field must be provided")

}
