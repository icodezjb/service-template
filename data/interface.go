package data

// RedisData 定义redis服务需要实现的方法
type RedisData interface {
	UserAccountExist(account string) (bool, error)
}

// MysqlData 定义Mysql服务需要实现的方法
type MysqlData interface {
}

// ExternalData 定义外部第三方数据需要实现的方法
type ExternalData interface {
}

// MicroServiceData 定义内部微服务需要实现的方法
type MicroServiceData interface {
}
