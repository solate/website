

db.createCollection("stencil")


http://www.jb51.net/article/52498.htm

mongodump -h IP --port 端口 -u 用户名 -p 密码 -d 数据库 -o 文件存在路径



mongorestore -h IP --port 端口 -u 用户名 -p 密码 -d 数据库 --drop 文件存在路径





## 编译


GOOS=linux go build -ldflags "-w" -o home home.go




GOOS=linux go build -ldflags "-w" -o management management.go



scp home zhangjin@59.151.42.5:~/

