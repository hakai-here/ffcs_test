package models

import (
	"ffcs/api/constants"
	"ffcs/api/utils"
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/exp/slices"
	"gorm.io/gorm"
)

type Users struct {
	BaseModel
	Rollno    string            `json:"rollno" gorm:"unique"`
	Firstname string            `json:"firstname"`
	Lastname  string            `json:"lastname"`
	Email     string            `json:"email"`
	Branch    string            `json:"branch"`
	IsAdmin   bool              `json:"-"`
	Password  string            `json:"-"`
	Courses   []SelectedCourses `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (user *Users) BeforeCreate(_ *gorm.DB) error {
	user.ID = uuid.New()
	if !slices.Contains(constants.Branches, user.Branch) {
		return fmt.Errorf("[NOS] : the students branch is not supported")
	}
	user.Password = utils.Hashes(user.Password)
	return nil
}
