package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"com.ocbc.smb/constants"
	"com.ocbc.smb/dto"
	"com.ocbc.smb/model"
	"com.ocbc.smb/services"
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
		fmt.Println("Error, body Request ")
		res.ErrCode = constants.ERR_CODE_01
		res.ErrDesc = constants.ERR_CODE_01
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, AccountService.RegisterAccount(&req))

	return
}
