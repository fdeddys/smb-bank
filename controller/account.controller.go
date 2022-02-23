package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"com.ocbc.smb/constants"
	"com.ocbc.smb/dto"
	"com.ocbc.smb/model"
	"com.ocbc.smb/services"
	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// AccountController ...
type AccountController struct {
	DB *gorm.DB
}

// AccountService ...
var AccountService = new(services.AccountService)

// GetUser ...
func (a *AccountController) GetUser(c *gin.Context) {

	id := 1
	c.JSON(http.StatusOK, AccountService.GetAccount(id))
	return
}

// Register Account ...
func (a *AccountController) Register(c *gin.Context) {

	req := model.Account{}
	res := dto.ContentResponse{}

	body := c.Request.Body
	dataBodyReq, _ := ioutil.ReadAll(body)

	if err := json.Unmarshal(dataBodyReq, &req); err != nil {
		logs.Info("Failed get body request")
		res.ErrDesc = constants.ERR_DESC_50_BODY_REQUEST
		res.ErrCode = constants.ERR_CODE_50
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, AccountService.RegisterAccount(&req))

	return
}

// Login Account ...
func (a *AccountController) Login(c *gin.Context) {

	req := dto.LoginDto{}
	res := dto.ContentResponse{}

	body := c.Request.Body
	dataBodyReq, _ := ioutil.ReadAll(body)

	if err := json.Unmarshal(dataBodyReq, &req); err != nil {
		logs.Info("Failed get body request")
		res.ErrDesc = constants.ERR_DESC_50_BODY_REQUEST
		res.ErrCode = constants.ERR_CODE_50
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, AccountService.Login(&req))

	return
}
