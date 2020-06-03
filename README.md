# 数据库表/字段变更检测

1. build

   ```bash
   go build
   ```

2. 初始化

   ```bash
   db_check -type init -username root -password 123456 -host localhost -port 3306 -dbname test
   ```

3. 检测
    > 需要将 db_check.json 文件放到当前路径下

   ```bash
   db_check -type check -username root -password 123456 -host localhost -port 3306 -dbname test
   ```

4. 参数说明

    > 默认参数即示例参数, type默认为init

   | 参数     | 说明                                              |
   | -------- | ------------------------------------------------- |
   | type     | 要执行的操作, 初始化使用 `init`  检测使用 `check` |
   | username | 数据库用户名                                      |
   | password | 数据库密码                                        |
   | host     | 数据库服务地址                                    |
   | port     | 端口号                                            |
   | dbname   | 数据库名                                          |

