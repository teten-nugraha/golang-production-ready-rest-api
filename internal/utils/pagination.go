package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type Pagination struct {
	Page     int
	PageSize int
}

func GetPagination(c *gin.Context) Pagination {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	if page < 1 {
		page = 1
	}

	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	return Pagination{
		Page:     page,
		PageSize: pageSize,
	}
}
