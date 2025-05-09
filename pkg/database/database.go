package database

type SetupOption struct {
	Host string
	Port int

	Username     string
	Password     string
	DatabaseName string

	PoolSize int
}

// Setup 初始化数据库
func Setup(option *SetupOption) error {
	panic("todo")
}
