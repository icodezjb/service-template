package definition

// 定义一些公共常用字段
const (
	MetadataRequestId         = "request_id"
	MetadataDeviceId          = ""
	MetadataPackageId         = ""
	MetadataRequestSource     = ""
	MetadataRequestProtocol   = ""
	MetadataClientVersion     = ""
	MetadataTimeReciveRequest = ""
)

// 定义请求来源
const (
	RequestSourceApp       = "app"
	RequestSourceThirdPart = "thirdpart"
)

// 定义请求协议
const (
	RequestProtocolHTTP = "HTTP"
	RequestProtocolGRPC = "GRPC"
)

// Environment 定义当前服务环境
type Environment string

const (
	EnvPro   Environment = "pro"
	EnvPre               = "pre"
	EnvTest              = "test"
	EnvDebug             = "debug"
)
