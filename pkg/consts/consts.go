package consts

// constants
const (
	TCP           = "tcp"
	UserTableName = "users"
	RedisAddr     = "127.0.0.1:6379"
	MaxIdleNum    = 10
	RedisPasswd   = ""
	SecretKey     = "secret"
	Username      = "username"
	HertzSession  = "HERTZ-SESSION"
	// root:114514@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local
	// gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local
	MySQLDefaultDSN = "root:114514@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
)

// error msg
const (
	Success     = "success"
	RegisterErr = "user already exists"
	LoginErr    = "wrong username or password"
	PageErr     = "please login first"
)
