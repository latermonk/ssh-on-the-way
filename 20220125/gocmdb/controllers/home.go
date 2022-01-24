package controllers

import (
	"github.com/astaxie/beego"
	"github.com/imsilence/gocmdb/controllers/auth"
	"net/http"
)

type HomeController struct {
	auth.LoginRequiredController
}

func (c *HomeController) Index() {
	c.Redirect(beego.URLFor("UserPageController.Index"), http.StatusFound)
}
