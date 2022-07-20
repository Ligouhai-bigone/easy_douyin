package constants

const (
	UserTableName              = "user_douyin"
	SecretKey                  = "secret key"
	ApiServiceName             = "api_service"
	UserServiceName            = "user_service"
	CommentServiceName         = "comment_service"
	VideoServiceName           = "video_service"
	MySQLDefaultDSN            = "douyin:douyin@tcp(localhost:9910)/douyin?charset=utf8&parseTime=True&loc=Local"
	EtcdAddress                = "127.0.0.1:2379"
	CPURateLimit       float64 = 80.0
	DefaultLimit               = 10
)
