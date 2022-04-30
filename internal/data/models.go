package data

import (
	"database/sql"
)

type Healthcheck struct {
	Status string `json:"status"`
	Environment string `json:"environment"`
	Version string `json:"version"`
}

type SimplePayload struct {
	Payload string `json:"payload"`
}

type Models struct {
	Users UserModel
	Files FileModel
	Timetables TimetableModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		UserModel{db},
		FileModel{db},
		TimetableModel{db},
	}
}
