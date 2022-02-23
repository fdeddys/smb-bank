package constants

const (
	API_VERSION1 = "1.0-1"
	API_VERSION2 = "2.0-1"
)

const (
	TokenSecretKey        = "s#creT_99-key$"
	TokenExpiredInMinutes = 8 * 60 * 60

	MAX_LENGTH_PASSWORD = 8
)

// ERROR CODE

// GENERAL
const (
	ERR_CODE_00         = "00"
	ERR_DESC_00_SUCCESS = "Success"
)

// INFRASTRUCTURE
const (
	ERR_CODE_50              = "50"
	ERR_DESC_50_BODY_REQUEST = "Failed to read body request"

	ERR_CODE_51               = "51"
	ERR_DESC_51_AUTHORIZATION = "Invalid Authorization "

	ERR_CODE_52               = "52"
	ERR_CODE_52_TOKEN_EXPIRED = "Token expired !"

	ERR_CODE_53       = "53"
	ERR_DESC_53_LOGIN = "Invalid Username or Password"

	ERR_CODE_54                = "54"
	ERR_DESC_54_GENERATE_TOKEN = "Failed generate token"
)

// DATABASE
const (
	ERR_CODE_01               = "01"
	ERR_DESC_01_SAVE_DATABASE = "Failed save data "

	ERR_CODE_02          = "02"
	ERR_DESC_02_GET_DATA = "Failed get data"
)

// USER
const (
	ERR_CODE_10                = "10"
	ERR_DESC_10_USERNAME_EXIST = "Username already exist "
)
