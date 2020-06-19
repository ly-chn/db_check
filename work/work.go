package work

import (
	"db_check/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Start 初始化
func Start() {
	// 获取一个数据库中所有表的数据
	tableList := model.InitTables()
	// 保存folder对象到本地json
	marshal, _ := json.Marshal(tableList)
	err := ioutil.WriteFile(model.FileListPath, marshal, os.ModePerm)
	if err != nil {
		panic("文件保存失败: " + err.Error())
	}
	fmt.Println("操作完成, 请携带 " + model.FileListPath + " 比对信息")
}

// Check 检测变更
func Check() {
	// 第一个数据库信息
	oldFileList := getFileListFromFile(model.OldFileListPath)
	// 第二个数据库信息
	newFileList := getFileListFromFile(model.NewFileListPath)
	compare := model.Compare(oldFileList, newFileList)
	model.PrintCompare(compare)
}

// 本地json文件解析
func getFileListFromFile(path string) (tableList []model.Table) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(`读取文件发生错误: ` + err.Error())
	}

	// 本地文件
	tableList = make([]model.Table, 0)
	err = json.Unmarshal(bytes, &tableList)
	if err != nil {
		panic("文件解析错误: " + err.Error())
	}
	return tableList
}
