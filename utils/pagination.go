package utils
 
import (
    _ "github.com/go-sql-driver/mysql"
)

type Pagination struct {
	Page int
	PageSize int
	Count int
	TotalPages int
	PrePage int
	NextPage int
	Ranges []int
}

func Set(page, pageSize, count int) Pagination {
	totalPages := count / pageSize
    if count % pageSize > 0 {
        totalPages += 1
	}
	
	var prePage int
	if page > 1 {
		prePage = page - 1
	} else {
		prePage = page
	}

	var nextPage int
	if page < totalPages {
		nextPage = page + 1
	} else {
		nextPage = page
	}

	var ranges []int
	if totalPages <= 5 {
		for i := page; i <= totalPages; i ++ {
			ranges = append(ranges, i)
		}
	} else if totalPages - page < 5 {
		for i := totalPages -5; i <= totalPages; i++ {
			ranges = append(ranges, i)
		}
	} else {
		for i := page; i < page+5 ; i++ {
			ranges = append(ranges, i)
		}
	}

	var pg Pagination
	pg.Page = page
	pg.PageSize = pageSize
	pg.PrePage = prePage
	pg.TotalPages = totalPages
	pg.Count = count
	pg.NextPage = nextPage
	pg.Ranges = ranges
	return pg
}
