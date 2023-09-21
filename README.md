# go-template
go service template

## 生成命令

### 依赖注入

```bash
wire ./internal
```


### 生成 pb服务文件

```bash
protoc -I=pb --go_out=pb/v1 --go_opt=paths=source_relative --go-grpc_out=pb/v1 --go-grpc_opt=paths=source_relative pb/log.proto
```

### run

```bash
docker run -it --rm -p 7007:7007 -v $PWD/config:/app/config gp-template:0.0.1
```

## feature
