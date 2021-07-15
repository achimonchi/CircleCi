package controller

import (
	"testing"

	"github.com/achimonchi/CircleCi/entity"
)

func TestGetData(t *testing.T) {
	var user entity.User
	user.Email = "reyhan.jovie@dana.id"
	user.Name = "Reyhan Jovie"

	data := GetData()
	// assert.Equal(t, data, user, "Harus sama")
	if data != user {
		t.Errorf("SALAH!, seharusnya %v bukan %v", user, data)
	}
}
