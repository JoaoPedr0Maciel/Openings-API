package handler

import "fmt"


func errParamerterIsRequired(name, typ string) error {
	return fmt.Errorf("param (%s) is required (type '%s')", name, typ)
}

type CreateOpeningRequest struct {
	Role        string `json:"role"`
	Company     string `json:"company"`
	Location    string `json:"location"`
	Remote      *bool  `json:"remote"`
	Link        string `json:"link"`
	Salary      string `json:"salary"`
}

func (r *CreateOpeningRequest) Validate() error {

	if r.Role == "" || r.Company == "" || r.Location == "" || r.Salary == "" || r.Link == "" || r.Remote == nil {
    return fmt.Errorf("the request body are empty or malformed")
  }

	if r.Role == "" {
		return errParamerterIsRequired("role", "string")
	}

	if r.Company == "" {
    return errParamerterIsRequired("company", "string")
  }

	if r.Location == "" {
    return errParamerterIsRequired("location", "string")
  }

	if r.Salary == "" {
    return errParamerterIsRequired("salary", "string")
  }

	if r.Remote == nil {
    return errParamerterIsRequired("remote", "bool")
  }

	if r.Link == "" {
    return errParamerterIsRequired("link", "string")
  }


	return nil
} 