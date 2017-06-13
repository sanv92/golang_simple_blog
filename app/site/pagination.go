package site

// Pagination ////////////////////////
//
type Pagination struct {
	PerPage     int
	TotalAmount int
	CurrentPage int
	TotalPage   int
}

func NewPagination(totalAmount, perPage, currentPage int) *Pagination {
	if currentPage == 0 {
		currentPage = 1
	}

	n := (totalAmount + perPage - 1) / perPage
	if currentPage > n {
		currentPage = n
	}

	return &Pagination{
		PerPage:     perPage,
		TotalAmount: totalAmount,
		CurrentPage: currentPage,
		TotalPage:   n,
	}
}
