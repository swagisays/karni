package model

import (
	"github.com/swagisays/karni/karni"
)

var userSchema = karni.Schema(map[string]karni.Field{
	"email": {Type: karni.String, Required: true, Trim: true, Lowercase: true, Unique: true},
	"password": {
		Type: karni.String,
	},
})

func InitUserModel() *karni.ModelStruct {
	User := karni.Model("users", userSchema)
	return User
}
