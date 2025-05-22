## fastgo 项目

# gorm
```sh
# 根据数据库表生成  Model 文件
go install gorm.io/gen/tools/gentool@latest
gentool -db mysql -dsn 'root:123456@tcp(127.0.0.1:3306)/fastgo' -onlyModel -modelPkgName internal/apiserver/model
```