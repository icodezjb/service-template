package definition

// 定义一些公共常用字段
const (
	MetadataRequestId         = "request_id"
	MetadataDeviceId          = ""
	MetadataPackageId         = ""
	MetadataVersion           = ""
	MetadataTimeReciveRequest = ""
)

// Environment 定义当前服务环境
type Environment string

const (
	EnvPro   Environment = "pro"
	EnvPre               = "pre"
	EnvTest              = "test"
	EnvDebug             = "debug"
)
