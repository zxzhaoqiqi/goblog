# goblog


## 常用命令

### 环境变量

```
# 修改 Go 相关的环境变量
go env -w  GOPROXY=https://goproxy.cn,direct
```

### mod 命令

```
# 初始化 mod， 生成 go.mod
go mod init xxx

# 整理依赖使用，执行时会把未使用的 module 移除掉
go mod tidy

```

### 依赖命令

```
# 下载依赖
go get xxx

# 清空 Go Modules 缓存
go clean -modcache

# 下载依赖
go mod download

# 查看依赖
go mod graph

# 编辑 go.mod 文件
go mod edit
```

## 目录结构
- go.mod 类似 composer.json
- go.sum 类似 comooser.lock