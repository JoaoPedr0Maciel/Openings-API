package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Opening struct {
	gorm.Model
	Role string
	Company string
	Location string
	Remote bool
	Link string
	Salary string
}

type OpeningResponse struct {
	ID          uint   `json:"id"`
	Role        string `json:"role"`
	Company     string `json:"company"`
	Location     string `json:"location"`
	Remote      bool   `json:"remote"`
	Link         string `json:"link"`
	Salary      string `json:"salary"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
}