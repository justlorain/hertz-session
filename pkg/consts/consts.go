package consts

// constants
const (
	TCP              = "tcp"
	UserTableName    = "users"
	RedisAddr        = "127.0.0.1:6379"
	MaxIdleNum       = 10
	RedisPasswd      = ""
	SessionSecretKey = "session-secret"
	CSRFSecretKey    = "csrf-secret"
	CSRFKeyLookUp    = "form:csrf"
	Username         = "username"
	HertzSession     = "HERTZ-SESSION"
	MySQLDefaultDSN  = "root:114514@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local"
)

// error msg
const (
	Success     = "success"
	RegisterErr = "user already exists"
	LoginErr    = "wrong username or password"
	PageErr     = "please login first"
	CSRFErr     = "csrf exception"
)
