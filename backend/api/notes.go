package api

import (
	"ECHO-GORM/api/helpers"
	"ECHO-GORM/db"
	"ECHO-GORM/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetNotes(c echo.Context) error {
	db := db.DbManager()
	notes := []model.Note{}
	// db.Find(&users)
	// spew.Dump(json.Marshal(users))
	// return c.JSON(http.StatusOK, users)
	// response := helpers.Response{
	// 	StatusCode: http.StatusOK,
	// 	Message:    "ok",
	// 	Data:       users,
	var kirim error
	// }
	err := db.Find(&notes).Error
	response := helpers.Response{}
	// err := errors.New("itu error bangsat")

	if err != nil {
		response = helpers.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed",
			Data:       err.Error(),
		}
		kirim = c.JSON(http.StatusOK, response)
	} else {
		response = helpers.Response{
			StatusCode: http.StatusOK,
			Message:    "ok",
			Data:       notes,
		}
		kirim = c.JSON(http.StatusOK, response)
	}

	return kirim
}

func GetNote(c echo.Context) error {
	db := db.DbManager()
	note := model.Note{}
	idNote := c.Param("id")

	response := helpers.Response{}

	var kirim error

	err := db.First(&note, idNote).Error
	if err != nil {
		response = helpers.Response{
			StatusCode: http.StatusNotFound,
			Message:    "failed",
			Data:       nil,
		}
		kirim = c.JSON(http.StatusOK, response)
	} else {
		response = helpers.Response{
			StatusCode: http.StatusOK,
			Message:    "ok",
			Data:       note,
		}
		kirim = c.JSON(http.StatusOK, response)
	}

	return kirim

}

func PostNote(c echo.Context) error {
	db := db.DbManager()
	// user := model.User{}
	u := new(model.Note)
	err := c.Bind(u)
	if err != nil {
		panic("Failed to Parsing Json")
	}

	response := helpers.Response{}

	var kirim error

	err = db.Create(&u).Error
	if err != nil {
		response = helpers.Response{
			StatusCode: http.StatusNotFound,
			Message:    "failed",
			Data:       nil,
		}
		kirim = c.JSON(http.StatusOK, response)
	} else {
		response = helpers.Response{
			StatusCode: http.StatusCreated,
			Message:    "ok",
			Data:       u,
		}
		kirim = c.JSON(http.StatusCreated, response)
	}

	return kirim
}

func UpdateNote(c echo.Context) error {
	db := db.DbManager()
	// user := model.User{}
	u := new(model.Note)
	idNote := c.Param("id")
	err := c.Bind(u)
	if err != nil {
		panic("Failed to Parsing Json")
	}
	u.ID, err = strconv.Atoi(idNote)
	if err != nil {
		panic("Failed Convert ID")
	}
	response := helpers.Response{}

	var kirim error

	// err = db.Create(&u).Error
	fmt.Println(u)
	err = db.Model(&u).Updates(&u).First(&u).Error

	if err != nil {
		response = helpers.Response{
			StatusCode: http.StatusNotFound,
			Message:    "failed",
			Data:       nil,
		}
		kirim = c.JSON(http.StatusNotFound, response)
	} else {
		response = helpers.Response{
			StatusCode: http.StatusOK,
			Message:    "ok",
			Data:       u,
		}
		kirim = c.JSON(http.StatusOK, response)
	}

	return kirim
}

func DeleteNote(c echo.Context) error {
	db := db.DbManager()
	u := new(model.Note)
	idNote := c.Param("id")
	var kirim error
	response := helpers.Response{}
	// err = db.Create(&u).Error
	// fmt.Println(u)
	err := db.Where("id = ?", idNote).First(&u).Delete(&u).Error
	if err != nil {
		response = helpers.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed",
			Data:       nil,
		}
		kirim = c.JSON(http.StatusInternalServerError, response)

	} else {
		response = helpers.Response{
			StatusCode: http.StatusOK,
			Message:    "data terhapus",
			Data:       nil,
		}
		kirim = c.JSON(http.StatusOK, response)

	}

	return kirim
}
