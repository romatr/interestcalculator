package model

type InterestResponse struct {
	ID string `json:"-"`
	Interest float64 `json:"interest"`
}

func (r InterestResponse) GetName() string {
	return "interest_calculations"
}

func (r InterestResponse) GetID() string {
	return r.ID
}