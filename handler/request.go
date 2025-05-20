package handler

import "fmt"

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	IsRemote *bool  `json:"isRemote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func errRequiredParam(name, typ string) error {
	return fmt.Errorf("parameter %s of type %s is required", name, typ)
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" && r.Company == "" && r.Location == "" && r.Link == "" && r.IsRemote == nil && r.Salary <= 0 {
		return fmt.Errorf("request body is empty or invalid")
	}

	if r.Role == "" {
		return errRequiredParam("role", "string")
	}

	if r.Company == "" {
		return errRequiredParam("company", "string")
	}

	if r.Location == "" {
		return errRequiredParam("location", "string")
	}

	if r.Link == "" {
		return errRequiredParam("link", "string")
	}

	if r.IsRemote == nil {
		return errRequiredParam("isRemote", "bool")
	}

	if r.Salary <= 0 {
		return errRequiredParam("salary", "int64")
	}

	return nil
}
