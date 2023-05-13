package controllers

import (
	"encoding/json"
	"ffcs/api/utils"
	"ffcs/pkg/models"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (r Repository) AdminCreateSlot(c *fiber.Ctx) error { // creating the sloat
	var slot models.Slots

	if err := c.BodyParser(&slot); err != nil {
		log.Println(err)
		return c.Status(http.StatusUnprocessableEntity).JSON(message.Error("unable to procress the requested resource"))
	}

	if err := r.Rdb.DeleteCache("sloatsadmin"); err != nil {
		log.Fatal(err.Error())
	}
	slot.ID = utils.CustomUUID()
	slot.AvaliableSeat = slot.TotalSeat
	if err := r.DB.Create(&slot).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(message.Error("Internal server error , unable to create the slot"))
	}
	return c.Status(http.StatusCreated).JSON(message.Success("successfully creatred the slot", fiber.Map{
		"course_code": slot.CourseCode,
		"faculty":     slot.CourseFaculty,
	}))
}

func (r Repository) AdminReadallSloat(c *fiber.Ctx) error { // reading all the sloats
	var sloats []models.Slots

	data, err := r.Rdb.GetCache("sloatsadmin")
	if err == nil {
		if err := json.Unmarshal([]byte(data), &sloats); err != nil {
			log.Println("Unable to parse redis data")
		}
		return c.JSON(message.Success("Found in cache database", sloats))
	}

	if err := r.DB.Find(&sloats).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(message.Error("Unable to procress the request at this moment "))
	}
	if len(sloats) == 0 {
		return c.Status(http.StatusNotFound).JSON(message.Error("No data found"))
	}
	jsonData, err := json.Marshal(sloats)
	if err != nil {
		log.Println("unable to marshel the data")
	} else {
		if err := r.Rdb.SetCache("sloatsadmin", jsonData); err != nil {
			log.Println("unable to set a cache")
		}
	}
	return c.JSON(message.Success("Fetched all the data", sloats))
}

// read one slot
func (r Repository) AdminReadoneSlot(c *fiber.Ctx) error {
	var sloat models.Slots
	id := c.Params("id")
	data, err := r.Rdb.GetCache(id)
	if err == nil {
		if err := json.Unmarshal([]byte(data), &sloat); err != nil {
			log.Println(err)
		}
		return c.JSON(message.Success("Found in cache database", sloat))
	}
	if err := r.DB.First(&sloat, "id = ?", id).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(message.Error("Unable to process the request"))
	}

	jsonData, err := json.Marshal(sloat)
	if err == nil {
		if err := r.Rdb.SetCache(id, jsonData); err != nil {
			log.Println(err)
		}
	}
	return c.JSON(message.Success("Fetched", sloat))
}
