package conf

import (
	"github.com/go-ini/ini"
)

var (
	Cfg *ini.File

	// 数据库相关
	Username string // 用户名
	Password string // 密码
	Host     string // 主机地址
	Port     string // 端口号
	Dbname   string // 数据库名
)

func LoadDb() {
	section := LoadSection("db")
	Username = section.Key("username").String()
	Password = section.Key("password").String()
	Host = section.Key("host").String()
	Port = section.Key("port").String()
	Dbname = section.Key("dbname").String()
}

// GetSection的封装, 给出节点名称返回节点信息
func LoadSection(sectionName string) *ini.Section {
	var section, _ = Cfg.GetSection(sectionName)
	return section
}

func LoadIni(iniFilePath string) {
	var err error
	// 加载并解析ini文件
	Cfg, err = ini.Load(iniFilePath)
	if err != nil {
		panic("无法加载解析配置文件: " + iniFilePath + ": " + err.Error())
	}

	LoadDb()
}
