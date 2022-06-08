# golang

### 一、Golang类库包使用教程

##### 1、引用或更新该类库包
```go get -u github.com/daimayun/go```
##### 2、引用指定版本的类库包
```go get github.com/daimayun/go/qrcode@v0.0.8```

```go install github.com/daimayun/go/qrcode@v0.0.8```

### 二、跨平台打包

##### 1、Mac下编译Linux平台的64位可执行程序
```CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go```

##### 2、Mac下编译Windows平台的64位可执行程序
```CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go```

##### 3、Linux下编译Mac平台的64位可执行程序
```CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go```

##### 4、Linux下编译Windows平台的64位可执行程序
```CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go```

##### 5、Windows下编译Mac平台的64位可执行程序
```SET CGO_ENABLED=0 SET GOOS=darwin3 SET GOARCH=amd64 go build main.go```

##### 6、Windows下编译Linux平台的64位可执行程序
```SET CGO_ENABLED=0 SET GOOS=linux SET GOARCH=amd64 go build main.go```
