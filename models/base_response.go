package models

import (
	"math"
	"net/http"
	"strconv"
)

type Meta struct {
	Status  bool   `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}
type PaginationMeta struct {
	Status      bool        `json:"status"`
	Code        int         `json:"code"`
	Message     string      `json:"message"`
	CurrentPage int         `json:"current_page"`
	NextPage    interface{} `json:"next_page"`
	PrevPage    interface{} `json:"prev_page"`
	PerPage     int         `json:"per_page"`
	PageCount   int         `json:"page_count"`
	TotalCount  int         `json:"total_count"`
}
type ModelResponse struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}
type ModelPaginationResponse struct {
	Meta interface{} `json:"meta"`
	Data interface{} `json:"data"`
}
type EmptyStruct struct{}
type EmptyResponse struct {
	Content interface{} `json:"content"`
}

func Response(code int, msg string, data interface{}) *ModelResponse {
	res := new(ModelResponse)
	meta := new(Meta)
	meta.Code = code
	switch code {
	case http.StatusOK:
		meta.Status = true
	case http.StatusInternalServerError:
		meta.Status = false
	case http.StatusBadRequest:
		meta.Status = false
	case http.StatusCreated:
		meta.Status = true
	case http.StatusNotFound:
		meta.Status = false
	case http.StatusUnauthorized:
		meta.Status = false
	}
	meta.Message = msg
	res.Meta = *meta
	res.Data = data

	return res
}

func PaginationResponse(code, total int, page, perPage string, data interface{}) *ModelPaginationResponse {
	res := new(ModelPaginationResponse)
	convPage, _ := strconv.Atoi(page)
	convPerPage, _ := strconv.Atoi(perPage)
	page_count := int(math.Ceil(float64(total) / float64(convPerPage)))
	next_page := convPage + 1
	prev_page := convPage - 1
	if math.Ceil(float64(total)/float64(convPerPage)) == float64(convPage) {
		next_page = convPage
	}
	if page == "1" {
		prev_page = convPage
	}
	meta := PaginationMeta{
		Message:     "success",
		Code:        code,
		Status:      true,
		CurrentPage: convPage,
		NextPage:    next_page,
		PrevPage:    prev_page,
		PerPage:     convPerPage,
		PageCount:   page_count,
		TotalCount:  total,
	}
	res.Meta = meta
	res.Data = data

	return res
}
