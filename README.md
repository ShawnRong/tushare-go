# TuShare restful api golang version
![](https://travis-ci.org/ShawnRong/tushare-go.svg?branch=master)
[![codecov](https://codecov.io/gh/ShawnRong/tushare-go/branch/master/graph/badge.svg)](https://codecov.io/gh/ShawnRong/tushare-go)
[![Documentation](https://godoc.org/github.com/ShawnRong/tushare-go?status.svg)](https://godoc.org/github.com/ShawnRong/tushare-go)


基于[tushare api](https://tushare.pro/document/2)的golang版本。
## 使用方法：
```go
	c := tushare.New("你的token")
	// 参数
	params := make(map[string]string)
	// 字段
	var fields []string
	// 根据api 请求对应的接口
	data, _ := c.StockBasic(params, fields)	
```

具体参数接口请查看[tushare API](https://tushare.pro/document/2)文档