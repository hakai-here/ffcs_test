package models

import (
	"ffcs/api/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Models associated with  Courses

type SelectedCourses struct { // to be connected with the user's profile to reflect the selected course
	CourseName string     `json:"name"`
	CourseCode string     `json:"code"`
	CourseId   uuid.UUID  `json:"course-id"`
	UsersId    *uuid.UUID `json:"-"` // foregin key
}

type Slots struct { // for all the avaliable courses
	BaseModel
	CourseName    string `json:"name"`
	CourseCode    string `json:"code"`
	CourseSlot    string `json:"slot"`
	CourseBranch  string `json:"branch"`
	CourseFaculty string `json:"faculty"`
	TotalSeat     int    `json:"total_seat"`
	AvaliableSeat int    `json:"avaliable_seat"`
}

func (s Slots) BeforeCreate(_ *gorm.DB) error {
	s.ID = utils.CustomUUID() // creating a custom uuids starting with ffc500
	return nil
}
