package types

// Interest type
type Interest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// CreateInterest type
type CreateInterest struct {
	Name string
}
