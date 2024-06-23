package rest

func (a *API) userRoutes() {
	a.router.GET("/users", a.getUsers)
	userGroup := a.router.Group("user")
	{
		userGroup.POST("", a.createUser)
		userGroup.GET(":id", a.getUserByID)
		userGroup.PUT(":id", a.updateUser)
		userGroup.DELETE(":id", a.deleteUser)
	}
}
