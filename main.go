package main

import (
	"db_check/model"
	"db_check/work"
)

func main() {
	// 初始化数据库信息
	if model.OperationType == model.TypeStart {
		work.Start()
	} else if model.OperationType == model.TypeCheck {
		work.Check()
	}
}
