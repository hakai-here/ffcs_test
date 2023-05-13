package controllers

import (
	"encoding/json"
	"ffcs/pkg/models"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func getIdandBranch(c *fiber.Ctx) (uuid.UUID, string) {
	uid := c.Locals("userid").(uuid.UUID)
	branch := c.Locals("branch").(string)
	return uid, branch
}

func (r Repository) ApiReadbyCourse(c *fiber.Ctx) error {
	var courses []models.Slots
	uid, branch := getIdandBranch(c)
	data, err := r.Rdb.GetCache(uid.String())
	if err == nil {
		if err := json.Unmarshal([]byte(data), &courses); err != nil {
			log.Printf("[Parse error] %s", err.Error())
		}
		return c.JSON(message.Success("fetched from cache", courses))
	}

	if err := r.DB.Where("course_branch = ?", branch).Find(&courses).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(message.Error("Unable to get the data"))
	}
	if len(courses) == 0 {
		return c.Status(http.StatusNotFound).JSON(message.Error("No data found for your branch"))
	}
	jsonData, err := json.Marshal(courses)
	if err == nil {
		if err := r.Rdb.SetCache(uid.String(), jsonData); err != nil {
			log.Println("[Redis] unable to set course cache")
		}
	}

	return c.JSON(message.Success("fetched", courses))
}

func (r Repository) RegisterCourse(c *fiber.Ctx) error {
	uid, _ := getIdandBranch(c)
	cid := c.Params("id")

	var course models.Slots
	if err := r.DB.Where("id = ?", cid).Find(&course).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(message.Error("Course not found"))
	}
	course.AvaliableSeat--
	if err := r.DB.Save(&course).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(message.Error("Unable to update"))
	}
	s := &models.SelectedCourses{
		CourseName: course.CourseName,
		CourseCode: course.CourseCode,
		CourseId:   course.ID,
		UsersId:    &uid,
	}
	if err := r.DB.Create(&s).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(message.Error(err.Error()))
	}
	return c.JSON(message.Success("booked your sloat", uid))

}
