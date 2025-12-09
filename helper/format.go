package helper

import (
	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/dto"
)

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

func FormatListResult[T any](
	data []T,
) []T {
	if len(data) == 0 {
		return []T{}
	} else {
		return data
	}
}
