package redis

type Config struct {
	Name     string  `mapstructure:"name" json:"name" yaml:"name"`             // 代表当前实例的名字
	Addr     string  `mapstructure:"addr" json:"addr" yaml:"addr"`             // 服务器地址:端口
	Password string  `mapstructure:"password" json:"password" yaml:"password"` // 密码
	DB       int     `mapstructure:"db" json:"db" yaml:"db"`                   // 单实例模式下redis的哪个数据库
	Cluster  Cluster `mapstructure:"cluster" json:"cluster" yaml:"cluster"`    //集群配置
}

type Cluster struct {
	Enable   bool     `mapstructure:"enable" json:"enable" yaml:"enable"`         // 是否使用集群模式
	AddrList []string `mapstructure:"addr-list" json:"addrList" yaml:"addr-list"` // 集群模式下的节点地址列表
}
