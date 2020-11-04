# 数据库表/字段变更检测

1. build

   ```bash
   go build
   ```

2. 使用
  ```bash
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
  ```
