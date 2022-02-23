package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"com.ocbc.smb/constants"
	"com.ocbc.smb/dto"
	"com.ocbc.smb/model"
	"com.ocbc.smb/services"
	"com.ocbc.smb/utils"
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

	id := dto.CurrUserID
	utils.LogInfo("get by User ID : " + strconv.Itoa(id))
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
		utils.LogInfo("Failed get body request : " + err.Error())
		res.ErrDesc = constants.ERR_DESC_50_BODY_REQUEST
		res.ErrCode = constants.ERR_CODE_50
		res.Contents = err.Error()
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	}

	bodyReq, _ := json.Marshal(&req)
	utils.LogInfo("Register : " + string(bodyReq))
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
		utils.LogInfo("Failed get body request : " + err.Error())
		res.ErrDesc = constants.ERR_DESC_50_BODY_REQUEST
		res.ErrCode = constants.ERR_CODE_50
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	}

	bodyReq, _ := json.Marshal(&req)
	utils.LogInfo("Login : " + string(bodyReq))
	c.JSON(http.StatusOK, AccountService.Login(&req))

	return
}
