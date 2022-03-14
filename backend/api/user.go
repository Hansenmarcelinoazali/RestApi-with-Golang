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

func GetUsers(c echo.Context) error {
	db := db.DbManager()
	users := []model.User{}
	// db.Find(&users)
	// spew.Dump(json.Marshal(users))
	// return c.JSON(http.StatusOK, users)
	// response := helpers.Response{
	// 	StatusCode: http.StatusOK,
	// 	Message:    "ok",
	// 	Data:       users,
	var kirim error
	// }
	err := db.Find(&users).Error
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
			Data:       users,
		}
		kirim = c.JSON(http.StatusOK, response)
	}

	return kirim
}

func GetUser(c echo.Context) error {
	db := db.DbManager()
	user := model.User{}
	idUser := c.Param("id")

	response := helpers.Response{}

	var kirim error

	err := db.First(&user, idUser).Error
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
			Data:       user,
		}
		kirim = c.JSON(http.StatusOK, response)
	}

	return kirim

}

func PostUser(c echo.Context) error {
	db := db.DbManager()
	// user := model.User{}
	u := new(model.User)
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

func UpdateUser(c echo.Context) error {
	db := db.DbManager()
	// user := model.User{}
	u := new(model.User)
	idUser := c.Param("id")
	err := c.Bind(u)
	if err != nil {
		panic("Failed to Parsing Json")
	}
	u.ID, err = strconv.Atoi(idUser)
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

func DeleteUser(c echo.Context) error {
	db := db.DbManager()
	u := new(model.User)
	idUser := c.Param("id")
	var kirim error
	response := helpers.Response{}
	// err = db.Create(&u).Error
	// fmt.Println(u)
	err := db.Where("id = ?", idUser).First(&u).Delete(&u).Error
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


