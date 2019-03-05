# TuShare golang version
基于[tushare api](https://tushare.pro/document/2)的golang版本。
## 使用方法：
```go
	c := TuShare.New("你的token")
	// 参数
	params := make(map[string]string)
	// 字段
	var fields []string
	// 根据api 请求对应的接口
	data, _ := c.StockBasic(params, fields)	
```

具体参数接口请查看[TuShare API](https://tushare.pro/document/2)文档