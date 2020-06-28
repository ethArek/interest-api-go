package types

// Interest type
type Interest struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Language int    `json:"language"`
}

// CreateInterest type
type CreateInterest struct {
	Name string
}
