package definition

// 定义版本号类型
type Version string

// 定义服务端版本号
const (
	VersionLatest Version = ""

	Version_1     Version = "v1"
	Version_1_1_2 Version = "v1.1.2"
)
