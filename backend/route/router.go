package route

import (
	"ECHO-GORM/api"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORS())
	//user
	e.GET("/", api.Home)

	e.GET("/users", api.GetUsers)

	e.GET("/users/:id", api.GetUser)

	e.POST("/users", api.PostUser)

	e.PUT("/users/:id", api.UpdateUser)

	e.DELETE("/users/:id", api.DeleteUser)

	//notes
	e.GET("/notes", api.GetNotes)

	e.GET("/notes/:id", api.GetNote)

	e.POST("/notes", api.PostNote)

	e.PUT("/notes/:id", api.UpdateNote)

	e.DELETE("/notes/:id", api.DeleteNote)
	return e

}

// func post() *echo.Echo {
// 	p := echo.New()

// 	p.POST("/", api.Home)

// 	// p.POST("/users", api.PostUsers)
// 	return p
// }

// func put() *echo.Echo {
// 	i := echo.New()

// 	i.PUT("/", api.Home)

// 	// i.POST("/users", api.PutUsers)
// 	return i
// }

// func delete() *echo.Echo {
// 	d := echo.New()

// 	d.DELETE("/", api.Home)

// 	// d.DELETE("/users", api.DeleteUsers)
// 	return d
// }
