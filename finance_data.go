package tushare

import (
	"fmt"
)

// Income 获取上市公司财务利润表数据
func (api *TuShare) Income(params map[string]string, fields []string) (*APIResponse, error) {
	// Check params
	_, hasTsCode := params["ts_code"]
	if !hasTsCode {
		return nil, fmt.Errorf("ts_code is a required argument")
	}

	body := map[string]interface{}{
		"api_name": "income",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}

	return api.postData(body)
}
