package web

import (
	"log"
	"net/http"

	"github.com/jackc/pgx/v4"
	"github.com/zecodein/thecarwash.kz/domain"
)

func errorHandler(err error) int {
	log.Println(err)
	switch err {
	case domain.ErrInvalidData:
		return http.StatusBadRequest
	case domain.ErrUniqueData:
		return http.StatusConflict
	case pgx.ErrNoRows:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
