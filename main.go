package main

import (
	"db_check/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// 初始化数据库信息
	if model.OperationType == model.TypeStart {
		// 获取一个数据库中所有表的数据
		tableList := model.InitTables()
		// 保存folder对象到本地json
		marshal, _ := json.Marshal(tableList)

		err := ioutil.WriteFile("./db_check.json", marshal, os.ModePerm)
		if err != nil {
			panic("文件保存失败: " + err.Error())
		}
		fmt.Println("操作完成, 请携带 db_check.json 比对信息")
	} else if model.OperationType == model.TypeCheck {
		bytes, err := ioutil.ReadFile("./db_check.json")
		if err != nil {
			panic(`读取文件发生错误: ` + err.Error() + `
            如果没有初始化,请使用以下参数:
            -type init
        		`)
		}

		// 本地文件
		tableListLocal := make([]model.Table, 0)
		err = json.Unmarshal(bytes, &tableListLocal)
		if err != nil {
			panic("请检查 db_check.json 文件是否发生变更: " + err.Error())
		}

		// 获取一个数据库中所有表的数据
		tableList := model.InitTables()
		compare := model.Compare(tableListLocal, tableList)
		model.PrintCompare(compare)
	}
}
