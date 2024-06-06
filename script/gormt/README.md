### gorm文档

https://gorm.io/zh_CN/docs/index.html

### 数据库文件生成
* 如果要新加表，需要修改config
* cd script/gormt
* 执行后将bbj_db.go 移动到 model文件夹下
*  ./gormt && cp _model/bbj_db.go ../../model
* 修改 package  name  _model
