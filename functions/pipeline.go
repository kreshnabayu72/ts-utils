package functions

import "github.com/betsv2/tsv2-go-utils/dto"

func FormatPaginationResult[T any](
	data []dto.PaginationResult[T],
) dto.PaginationResult[T] {
	if len(data) == 0 {
		return dto.PaginationResult[T]{
			TotalRecords: 0,
			Data:         []T{},
		}
	} else {
		return data[0]
	}
}
