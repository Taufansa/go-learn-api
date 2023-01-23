package model

import (
	"time"
)

type People struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Peoples []People

func newPeople() *People {
	return &People{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
