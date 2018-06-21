package api

import (
	"database/sql"
)

type Server struct {
	Data *DataManager
}

func NewServer(dbConnection *sql.DB) *Server {
	return &Server{
		Data: &DataManager{DB: dbConnection},
	}
}
