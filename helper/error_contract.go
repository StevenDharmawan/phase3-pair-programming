package helper

import (
	"net/http"
	"pair-programming/model/web"
)

func ErrBadRequest(detail any) web.ErrorResponse {
	return web.ErrorResponse{
		Code:    http.StatusBadRequest,
		Message: "Bad Request",
		Detail:  detail,
	}
}

func ErrInternalServer(detail any) web.ErrorResponse {
	return web.ErrorResponse{
		Code:    http.StatusInternalServerError,
		Message: "Internal Server Error",
		Detail:  detail,
	}
}
