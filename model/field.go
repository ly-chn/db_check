package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Field struct {
	Name    string // 字段名
	Type    string // 字段类型
	Key     string // 键
	Default string // 默认值
	Extra   string // 其它
	Status  string `json:"-"` // 删除标记
	Message string `json:"-"` // 变更详情
}

// GetFieldsByTableName 通过表格获取字段信息
func (item *Field) GetFieldsByTableName(db *gorm.DB, tableName string) (fieldList []*Field) {
	// 初始化
	fieldList = make([]*Field, 0)

	// 获取字段
	rows, err := db.Raw("desc " + tableName).Rows()
	if err != nil {
		fmt.Printf("err: %#v\n", err)
		panic("根据表明获取字段信息失败")
	}
	defer rows.Close()

	// 解析sql执行结果
	for rows.Next() {
		var field = new(Field)
		var Tmp interface{}
		var Default interface{}
		err := rows.Scan(&field.Name,
			&field.Type,
			&Tmp,
			&field.Key,
			&Default,
			&field.Extra)
		if err != nil {
			fmt.Printf("err: %#v\n", err)
			panic("字段名映射失败")
		}
		// default信息需要间接转换一下
		if Default != nil {
			field.Default = string(Default.([]uint8))
		}
		fieldList = append(fieldList, field)
	}

	return fieldList
}
