

db.createCollection("stencil")


http://www.jb51.net/article/52498.htm

mongodump -h IP --port 端口 -u 用户名 -p 密码 -d 数据库 -o 文件存在路径



mongorestore -h IP --port 端口 -u 用户名 -p 密码 -d 数据库 --drop 文件存在路径





## 编译


GOOS=linux go build -ldflags "-w" -o home home.go




GOOS=linux go build -ldflags "-w" -o management management.go



scp home zhangjin@59.151.42.5:~/



## 设计


### 用户中心

#### 登录


#### 注册


#### 登出

#### 用户管理

#### 角色管理