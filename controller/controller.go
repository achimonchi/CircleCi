package controller

import (
	"github.com/achimonchi/CircleCi/entity"
)

func GetData() entity.User {
	var user1 entity.User
	user1.Email = "reyhan.jovie@dana.id"
	user1.Name = "Reyhan Jovie"
	return user1
}
