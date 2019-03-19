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

// BalanceSheet 获取上市公司资产负债表
func (api *TuShare) BalanceSheet(params map[string]string, fields []string) (*APIResponse, error) {
	// Check params
	_, hasTsCode := params["ts_code"]
	if !hasTsCode {
		return nil, fmt.Errorf("ts_code is a required argument")
	}

	body := map[string]interface{}{
		"api_name": "balancesheet",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}

	return api.postData(body)
}

// CashFlow 获取上市公司现金流量表
func (api *TuShare) CashFlow(params map[string]string, fields []string) (*APIResponse, error) {
	// Check params
	_, hasTsCode := params["ts_code"]
	if !hasTsCode {
		return nil, fmt.Errorf("ts_code is a required argument")
	}

	body := map[string]interface{}{
		"api_name": "cashflow",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}

	return api.postData(body)
}

// Forecast 获取业绩预告数据
func (api *TuShare) Forecast(params map[string]string, fields []string) (*APIResponse, error) {
	// Check param
	_, hasTsCode := params["ts_code"]
	_, hasAnnDate := params["ann_date"]

	// ts_code & ann_date required
	if (!hasTsCode && !hasAnnDate) || (hasTsCode && hasAnnDate) {
		return nil, fmt.Errorf("Need one argument ts_code or ann_date")
	}

	body := map[string]interface{}{
		"api_name": "forecast",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}

	return api.postData(body)
}

// Dividend 分红送股数据
func (api *TuShare) Dividend(params map[string]string, fields []string) (*APIResponse, error) {

	body := map[string]interface{}{
		"api_name": "dividend",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}

	return api.postData(body)
}

// Express 获取上市公司业绩快报
func (api *TuShare) Express(params map[string]string, fields []string) (*APIResponse, error) {
	// Check params
	_, hasTsCode := params["ts_code"]
	if !hasTsCode {
		return nil, fmt.Errorf("ts_code is a required argument")
	}

	body := map[string]interface{}{
		"api_name": "express",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}

	return api.postData(body)
}

// FinaIndicator 获取上市公司财务指标数据
func (api *TuShare) FinaIndicator(params map[string]string, fields []string) (*APIResponse, error) {
	// Check params
	_, hasTsCode := params["ts_code"]
	if !hasTsCode {
		return nil, fmt.Errorf("ts_code is a required argument")
	}

	body := map[string]interface{}{
		"api_name": "fina_indicator",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}

	return api.postData(body)
}

// FinaAudit 获取上市公司定期财务审计意见数据
func (api *TuShare) FinaAudit(params map[string]string, fields []string) (*APIResponse, error) {
	// Check params
	_, hasTsCode := params["ts_code"]
	if !hasTsCode {
		return nil, fmt.Errorf("ts_code is a required argument")
	}

	body := map[string]interface{}{
		"api_name": "fina_audit",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}

	return api.postData(body)
}

// FinaMainbz 获得上市公司主营业务构成，分地区和产品两种方式
func (api *TuShare) FinaMainbz(params map[string]string, fields []string) (*APIResponse, error) {
	// Check params
	_, hasTsCode := params["ts_code"]
	if !hasTsCode {
		return nil, fmt.Errorf("ts_code is a required argument")
	}

	body := map[string]interface{}{
		"api_name": "fina_mainbz",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}

	return api.postData(body)
}

// DisclosureDate 获取财报披露计划日期
func (api *TuShare) DisclosureDate(params map[string]string, fields []string) (*APIResponse, error) {

	body := map[string]interface{}{
		"api_name": "disclosure_date",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}

	return api.postData(body)
}
