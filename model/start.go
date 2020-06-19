package model

import (
	"db_check/conf"
	"fmt"
	"os"
)

const TypeCheck = "compare" // 检测变更
const TypeStart = "start"   // 初始化

var (
	OperationType     = "" // 操作类型
	IniFilePath       = "" // 数据库配置文件路径
	FileListPath      = "" // 数据库信息保存位置
	OldFileListPath   = "" // 第一个数据库文件的储存位置
	NewFileListPath   = "" // 第二个数据库文件的储存位置
	CompareResultPath = "" // 比对结果保存位置
)

func init() {
	// 判断参数格式是否合法
	args := os.Args
	if len(args) == 1 || (args[1] != TypeCheck && args[1] != TypeStart) {
		errParam()
	}

	OperationType = args[1]
	if OperationType == TypeStart {
		if len(args) != 4 {
			errParam()
		}
		IniFilePath = args[2]
		FileListPath = args[3]
		conf.LoadIni(IniFilePath)
		// 初始化 mysql
		InitDb()
	} else if OperationType == TypeCheck {
		if len(args) != 5 {
			errParam()
		}
		// 文件路径
		OldFileListPath = args[2]
		NewFileListPath = args[3]
		CompareResultPath = args[4]
	}
}

// 参数错误, 程序无法运行
func errParam() {
	fmt.Println(`
Usage: 
    db_check [command] [json/ini] [json]
The commands are:
    start      将当前数据库结构保存为文件
    compare    检测文件变更
Examples:
	// 将 root@localhost.ini 中配置的数据库信息保存到 a.json 件中
	db_check start root@localhost.ini a.json 
	// 比对 b.json 对于 a.json 的变化, 并将结果保存到 c.txt 中
	db_check compare a.json b.json c.txt 
            `)
	os.Exit(0)
}
