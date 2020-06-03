package model

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/jinzhu/gorm"
)

type Table struct {
	Name   string   // 表名
	Fields []*Field // 字段信息
	Status string   // 删除标记
}

// GetTableList 获取数据库中所有的表
func (item *Table) GetTableList(db *gorm.DB) (tableList []Table) {
	tableNameList := make([]string, 0)
	rows, err := db.Raw("show tables").Rows()

	if err != nil {
		fmt.Printf("err: %#v\n", err)
		panic("获取数据库表名失败")
	}
	defer rows.Close()
	for rows.Next() {
		var tmp string
		err := rows.Scan(&tmp)
		if err != nil {
			fmt.Printf("err: %#v\n", err)
			panic("表名映射失败")
		}
		tableNameList = append(tableNameList, tmp)
	}
	tableList = make([]Table, 0)
	for _, name := range tableNameList {
		table := &Table{Name: name}
		// 字段信息
		table.solveFields(db)
		tableList = append(tableList, *table)
	}
	return tableList
}

// solveFields 设置表中所有的字段
func (item *Table) solveFields(db *gorm.DB) {
	field := new(Field)
	fields := field.GetFieldsByTableName(db, item.Name)
	item.Fields = fields
}

// SetStatus 修改表和表字段的状态
func (item *Table) SetStatus(status string) {
	item.Status = status
	for i := range item.Fields {
		item.Fields[i].Status = status
	}
}

// InitTables 检索数据库中所有表的信息
func InitTables() (tableList []Table) {
	// 获取一个数据库中所有表的数据
	table := new(Table)
	db := GetDb()
	tableList = table.GetTableList(db)
	return tableList
}

// GetMd5 获取表的md5 目前的逻辑是计算每个字段的md5合是否相同
func (item *Table) GetMd5() string {
	h := md5.New()
	str := ""
	for _, field := range item.Fields {
		str += field.Name
		str += field.Type
		str += field.Key
		str += field.Default
		str += field.Extra
	}
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
