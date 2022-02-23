package router

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"com.ocbc.smb/constants"
	"com.ocbc.smb/controller"
	"com.ocbc.smb/dto"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {

	r := gin.New()

	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "DELETE", "PUT"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		MaxAge:           86400,
	}))

	AccountController := new(controller.AccountController)

	router_version := ""

	router_version = "/api/" + constants.API_VERSION1 + "/"
	api := r.Group(router_version + "account")
	api.GET("", AccountController.GetUser)
	api.POST("", AccountController.Register)

	return r
}

func tokenChecking(c *gin.Context) {

	res := dto.ContentResponse{}
	tokenString := c.Request.Header.Get("Authorization")

	// Prefix Bearer tidak ada
	if strings.HasPrefix(tokenString, "Bearer ") == false {
		res.ErrCode = constants.ERR_CODE_51
		res.ErrDesc = constants.ERR_DESC_51_AUTHORIZATION
		c.JSON(http.StatusUnauthorized, res)
		c.Abort()
		return
	}

	// Isi Bearer gagal signing
	tokenString = strings.Replace(tokenString, "Bearer ", "", -1)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			res.ErrCode = constants.ERR_CODE_51
			res.ErrDesc = constants.ERR_DESC_51_AUTHORIZATION
			c.JSON(http.StatusUnauthorized, res)
			c.Abort()
		}
		return []byte(constants.TokenSecretKey), nil
	})

	if token != nil && err == nil {
		claims := token.Claims.(jwt.MapClaims)

		unixNano := time.Now().UnixNano()
		timeNowInInt := unixNano / 1000000

		tokenCreated := (claims["tokenCreated"])
		dto.CurrUser = (claims["user"]).(string)

		fmt.Println("now : ", timeNowInInt)
		fmt.Println("token created time : ", tokenCreated)
		fmt.Println("user by token : ", dto.CurrUser)

		// Cek DATE TOKEN CREATED is valid ?
		tokenCreatedInString := tokenCreated.(string)
		tokenCreatedInInt, errTokenExpired := strconv.ParseInt(tokenCreatedInString, 10, 64)
		if errTokenExpired != nil {
			res.ErrDesc = constants.ERR_CODE_52_TOKEN_EXPIRED
			res.ErrCode = constants.ERR_CODE_52
			c.JSON(http.StatusUnauthorized, res)
			c.Abort()
			return
		}

		// Cek Token masih valid
		if ((timeNowInInt - tokenCreatedInInt) / 1000) > constants.TokenExpiredInMinutes {
			res.ErrDesc = constants.ERR_CODE_52_TOKEN_EXPIRED
			res.ErrCode = constants.ERR_CODE_52
			c.JSON(http.StatusUnauthorized, res)
			c.Abort()
			return
		}

	} else {
		res.ErrCode = constants.ERR_CODE_51
		res.ErrDesc = constants.ERR_DESC_51_AUTHORIZATION
		c.JSON(http.StatusUnauthorized, res)
		c.Abort()
		return
	}
}
