package coincheck

// Pagination represents the pagination of coincheck API.
// It is possible to get by dividing the data.
type Pagination struct {
	// Limit is the number of data to get.
	Limit int `json:"limit"`
	// PaginationOrder is the order of the data. You can specify "desc" or "asc".
	PaginationOrder PaginationOrder `json:"order"`
	// StartingAfter is the ID of the data to start getting.
	// Greater than the specified ID. For example, if you specify 3, you will get data from ID 4.
	StartingAfter int `json:"starting_after,omitempty"`
	// EndingBefore is the ID of the data to end getting.
	// Less than the specified ID. For example, if you specify 3, you will get data up to ID 2.
	EndingBefore int `json:"ending_before,omitempty"`
}

// PaginationOrder represents the order of the pagination.
type PaginationOrder string

const (
	// PaginationOrderDesc is the order of the pagination in descending order.
	PaginationOrderDesc PaginationOrder = "desc"
	// PaginationOrderAsc is the order of the pagination in ascending order.
	PaginationOrderAsc PaginationOrder = "asc"
)
