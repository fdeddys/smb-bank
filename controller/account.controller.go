package controller

import (
	"net/http"
	"strconv"

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

// GetByUser ...
func (a *AccountController) GetByUser(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	res := AccountService.GetAccount(id)

	c.JSON(http.StatusOK, res)
	c.Abort()
	return
}
