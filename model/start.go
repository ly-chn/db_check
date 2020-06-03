package model

import (
	"db_check/conf"
	"fmt"
	"os"
)

const TypeCheck = "check" // 检测变更
const TypeStart = "start" // 初始化

var (
	OperationType = ""
	IniFilePath   = ""
)

func init() {
	// 判断参数格式是否合法
	args := os.Args
	OperationType = args[1]
	if OperationType != TypeCheck && OperationType != TypeStart {
		fmt.Print(`
            Usage: 
                db_check [command] [json/ini] [ini]
            The commands are:
                start  将当前数据库结构保存为 dbname@host.json 文件
                check  检测文件变更
			Examples:
				db_check start localhost.ini
				db_check check a.json b.json
            `)
		os.Exit(0)
	}

	if OperationType == TypeStart {
		IniFilePath = args[2]
		conf.LoadIni(IniFilePath)
	}

	// 初始化 mysql
	InitDb()
}
