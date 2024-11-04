# 配置数据库相关

```
当前Mysql版本: 8.0.38
当前Redis版本: 6.0.16
```

## 数据库配置

### 修改db.go文件
```go
// 把yourname:password@/yourdb改成你的数据库用户名,密码和数据库名称
db, err := gorm.Open("mysql", "yourname:password@/yourdb?charset=utf8&parseTime=True&loc=Local")
```

### 修改redis.go文件
```go
// 如果你的redis没有设置密码,则Password: ""
// 如果redis设置了密码,则Password: "yourpassword"
// Addr: ""改成你的redis地址和端口
rdb = redis.NewClient(&redis.Options{
    Addr:     "localhost:6379",
    Password: "",
    DB:       0,
})
```
