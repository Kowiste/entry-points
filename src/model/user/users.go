package user

import "context"



type UsersGetFilter struct {
	//filters for get user list
}

func (u Users) Get(c context.Context, filter *UsersGetFilter) (err error) {
	return
}
