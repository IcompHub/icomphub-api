package handlers

import "icomphub-api/dtos"

func Paginate[T any](items []T, count uint64, pageNumber uint64, pageSize uint64) dtos.PaginationDTO[T] {
	return dtos.PaginationDTO[T]{
		TotalItems: count,
		TotalPages: (count + pageSize) / pageSize,
		PageNumber: pageNumber,
		PageSize:   pageSize,
		Items:      items,
	}
}
