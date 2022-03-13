package dto

import "movies-api/database"

type JsonResponse struct {
	Type    string  `json:"type"`
	Data    []database.Movie `json:"data"`
	Message string  `json:"message"`
}