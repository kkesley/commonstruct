package stdreq

//ListStdResponse holds the standard response for list
type ListStdResponse struct {
	Page      *int    `json:"page"`
	Limit     *int    `json:"limit"`
	Count     int64   `json:"count"`
	NextToken *string `json:"next_token"`
}
