package types

type Page struct {
	HasNextPage     bool `json:"hasNextPage"`
	HasPreviousPage bool `json:"hasPreviousPage"`
	Quantity        int  `json:"quantity"`
}
