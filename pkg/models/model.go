package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// Basemodel for the all gorm proceedings 'gorm.Model'
type BaseModel struct {
	ID        uuid.UUID     `gorm:"primaryKey;type:uuid"`
	CreatedAt time.Time     `json:"-"`
	UpdatedAt time.Time     `json:"-"`
	DeletedAt *sql.NullTime `gorm:"index" json:"-"`
}
