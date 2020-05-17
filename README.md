# acdc

## 如何运行？

1. 参照 [这里](https://blog.csdn.net/AdolphKevin/article/details/105480530) 安装并设置 Go 的环境。
2. 在该目录内执行 `go run main.go`
3. 开始测试

如果安装了 GoLand 的话，在设置好 `$GOPROXY` 之后直接使用 GoLand 打开该项目即可。

## API 文档

在成功运行该项目后，打开 http://localhost:8080/swagger/index.html 即可看到全部的 API 文档，并可以直接在网页端进行测试。

### apiKey

挂锁的 API 都需要预先配置 APIKey 才可以运行，配置方法：

- 点击右上角的 Authorize
- 输入 `Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNTkwNzY0MjY4LCJvcmlnX2lhdCI6MTU4OTY4NDI2OCwicGF5bG9hZCI6IntcInVzZXJfaWRcIjoxLFwidXNlcm5hbWVcIjpcInJvb3RcIixcIlJvbGVcIjozMX0ifQ.NBD7YA-DZSFQB9YcLWixjxxo_aL_ECBu_cDQkVC2lYg`
- 点击 Authorize