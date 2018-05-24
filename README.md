# mygo
study go

godoc -http :6060

// 查看测试代码覆盖率
// go test -coverprofile=c.out
// go tool cover -html=c.out

// go test -bench .
// go test -bench . -cpuprofile cpu.out
// go test -memprofile mem.out
// go tool pprof mem.out
// go tool pprof cpu.out   然后输入web  或是quit 保证下载了svg
// https://graphviz.gitlab.io/_pages/Download/Download_source.html