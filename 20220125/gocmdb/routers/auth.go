package routers

import (
	"github.com/astaxie/beego"
	"github.com/imsilence/gocmdb/controllers/auth"
)

func init() {
	beego.AutoRouter(&auth.AuthController{})
}
