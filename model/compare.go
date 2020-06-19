package model

import (
	"fmt"
	"os"
)

const (
	StatusDeleted = "删除"
	StatusCreate  = "新增"
	StatusUpdate  = "变更"
	StatusNothing = "未变"
)

// Compare 比对
func Compare(tableListLocal []Table, tableList []Table) []Table {
	// 变更记录
	record := make([]Table, 0)
	// 判断已删除的表
	for _, tableLocal := range tableListLocal {
		flag := true
		for _, table := range tableList {
			if tableLocal.Name == table.Name {
				flag = false
				break
			}
		}
		// 已删除
		if flag {
			tableLocal.SetStatus(StatusDeleted)
			record = append(record, tableLocal)
		}
	}
	// 判断新增的表
	for _, table := range tableList {
		flag := true
		for _, tableLocal := range tableListLocal {
			if tableLocal.Name == table.Name {
				flag = false
			}
		}
		// 记录新增
		if flag {
			table.SetStatus(StatusCreate)
			record = append(record, table)
		}
	}
	// 判断修改的表
	for _, tableLocal := range tableListLocal {
		for _, table := range tableList {
			if tableLocal.Name == table.Name {
				// 无变化的表
				if table.GetMd5() == tableLocal.GetMd5() {
					tableLocal.SetStatus(StatusNothing)
					break
				}
				tableLocal.Status = StatusUpdate
				// 删除的字段
				for _, fieldLocal := range tableLocal.Fields {
					flag := true
					for _, field := range table.Fields {
						if fieldLocal.Name == field.Name {
							flag = false
							break
						}
					}
					if flag {
						fieldLocal.Status = StatusDeleted
					}
				}
				// 未变/修改的字段
				for _, fieldLocal := range tableLocal.Fields {
					for _, field := range table.Fields {
						message := ""
						if fieldLocal.Name == field.Name {
							fieldLocal.Status = StatusNothing

							if fieldLocal.Type != field.Type {
								message += fmt.Sprintf(" 类型变更: %v --> %v", fieldLocal.Type, field.Type)
							}
							if fieldLocal.Key != field.Key {
								message += fmt.Sprintf(" 键变更: %v --> %v", fieldLocal.Key, field.Key)
							}
							if fieldLocal.Default != field.Default {
								message += fmt.Sprintf(" 默认值变更: %v --> %v", fieldLocal.Default, field.Default)
							}
							if fieldLocal.Extra != field.Extra {
								message += fmt.Sprintf(" 其它变更: %v --> %v", fieldLocal.Extra, field.Extra)
							}

							if message != "" {
								fieldLocal.Status = StatusUpdate
								fieldLocal.Message = message
							}
						}
					}
				}
				// 新增的字段
				for _, field := range table.Fields {
					flag := true
					var fieldLocal *Field
					for _, fieldLocal = range tableLocal.Fields {
						if fieldLocal.Name == field.Name {
							flag = false
							break
						}
					}
					if flag {
						field.Status = StatusCreate
						tableLocal.Fields = append(tableLocal.Fields, field)
						break
					}
				}
				record = append(record, tableLocal)
			}
		}
	}
	return record
}

// PrintCompare 打印比对结果
func PrintCompare(compare []Table) {
	text := ""
	// 打印新增的表
	text += "新增的表: \n\n"
	for _, table := range compare {
		if table.Status == StatusCreate {
			text += "\t" + table.Name + "\n"
			for _, field := range table.Fields {
				text += fmt.Sprintf("\t\t%v\n", field.Name)
			}
		}
	}
	// 打印删除的表
	text += "\n删除的表: \n\n"
	for _, table := range compare {
		if table.Status == StatusDeleted {
			text += "\t" + table.Name + "\n"
			for _, field := range table.Fields {
				text += fmt.Sprintf("\t\t%v\n", field.Name)
			}
		}
	}
	// 变更的表
	text += "\n变更的表: \n\n"
	for _, table := range compare {
		if table.Status == StatusUpdate {
			text += "\t" + table.Name + "\n"
			for _, field := range table.Fields {
				if field.Status != StatusNothing {
					text += fmt.Sprintf("\t\t%v: %v", field.Status, field.Name)
					if field.Message != "" {
						text += " ----" + field.Message
					}
					text += "\n"
				}
			}
		}
	}
	// 保存到本地
	SaveStringToFile(text, CompareResultPath)
}

// SaveStringToFile 字符串写入到本地
func SaveStringToFile(str string, path string) {
	dstFile, err := os.Create(path)
	if err != nil {
		fmt.Println("文件保存失败: " + err.Error())
		os.Exit(0)
	}
	defer dstFile.Close()
	_, err = dstFile.WriteString(str)
	if err != nil {
		fmt.Println("文件保存失败: " + err.Error())
		os.Exit(0)
	}

	fmt.Println("比对完成请在 " + path + " 文件中查看结果")

}
