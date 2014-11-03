package controllers

import (
	"github.com/astaxie/beego"

	"api/models"
)

// Operations about object
type TestController struct {
	beego.Controller
}

// @Title Get
// @Description find area test info
// @Param	string	area
// @router /getdata [get]
func (this *TestController) GetData() {
	area := this.GetString("area")
	if "" == area {
		this.Data["json"] = "Param area is empty"
		this.ServeJson()
	}

	common_info := beego.AppConfig.String("common_info")

	if "" == common_info {
		this.Data["json"] = "Config collection is not exists"
		this.ServeJson()
	}

	info, err := models.TGetSaveData(common_info, area)
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = info
	}
	this.ServeJson(true)

}
