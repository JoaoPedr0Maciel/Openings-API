package handler

import "fmt"

func ErrParamerterIsRequired(name, typ string) error {
	return fmt.Errorf("param (%s) is required (type '%s')", name, typ)
}

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   string `json:"salary"`
}

func (r *CreateOpeningRequest) Validate() error {

	if r.Role == "" || r.Company == "" || r.Location == "" || r.Salary == "" || r.Link == "" || r.Remote == nil {
		return fmt.Errorf("the request body are empty or malformed")
	}

	if r.Role == "" {
		return ErrParamerterIsRequired("role", "string")
	}

	if r.Company == "" {
		return ErrParamerterIsRequired("company", "string")
	}

	if r.Location == "" {
		return ErrParamerterIsRequired("location", "string")
	}

	if r.Salary == "" {
		return ErrParamerterIsRequired("salary", "string")
	}

	if r.Remote == nil {
		return ErrParamerterIsRequired("remote", "bool")
	}

	if r.Link == "" {
		return ErrParamerterIsRequired("link", "string")
	}

	return nil
}


type UpdateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   string `json:"salary"`
}

func (r *UpdateOpeningRequest) Validate() error {

	// if any field is provided, return no error
	if r.Role != "" || r.Company != "" || r.Location != "" || r.Salary != "" || r.Link != "" || r.Remote != nil {
		return nil
	}

	return fmt.Errorf("at least one field must be provided")
}