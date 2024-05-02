package model

type PaginatedBook struct {
	Books            []*Book
	PageNo           int64
	PageSize         int64
	TotalPageNo      int64
	TotalRecordCount int64
	MinPrice         string
	MaxPrice         string
	IsLoggedIn       bool
	Username         string
}

func (p *PaginatedBook) HasPreviousPage() bool {
	return p.PageNo > 1
}

func (p *PaginatedBook) HasNextPage() bool {
	return p.PageNo < p.TotalPageNo
}

func (p *PaginatedBook) GetPreviousPageNo() int64 {
	if p.PageNo > 1 {
		return p.PageNo - 1
	}
	return 1
}

func (p *PaginatedBook) GetNextPageNo() int64 {
	if p.PageNo < p.TotalPageNo {
		return p.PageNo + 1
	}
	return p.TotalPageNo
}
