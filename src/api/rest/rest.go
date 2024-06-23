package rest

import (
	"github.com/gin-gonic/gin"
)

type API struct {
	router *gin.Engine
}

func New() (api *API) {
	return &API{}
}
func (a *API) Init() (err error) {

	a.router = gin.Default()

	a.userRoutes()
	go func() {
		err = a.router.Run(":8080") // listen and serve on port 8080
		if err != nil {
			panic(err)
		}
	}()
	return
}
