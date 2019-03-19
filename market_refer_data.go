package tushare

import (
	"fmt"
)

// MoneyflowHsgt 获取沪股通、深股通、港股通每日资金流向数据
func (api *TuShare) MoneyflowHsgt(params map[string]string, fields []string) (*APIResponse, error) {
	// Check param
	_, hasTradeDate := params["trade_date"]
	_, hasStartDate := params["start_date"]

	// trade_date & start_date required
	if (!hasTradeDate && !hasStartDate) || (hasTradeDate && hasStartDate) {
		return nil, fmt.Errorf("Need one argument trade_date or start_date")
	}

	body := map[string]interface{}{
		"api_name": "moneyflow_hsgt",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}

	return api.postData(body)
}

// HsgtTop10 获取沪股通、深股通每日前十大成交详细数据
func (api *TuShare) HsgtTop10(params map[string]string, fields []string) (*APIResponse, error) {
	// Check param
	_, hasTsCode := params["ts_code"]
	_, hasTradeDate := params["trade_date"]

	// ts_code & trade_date required
	if (!hasTsCode && !hasTradeDate) || (hasTsCode && hasTradeDate) {
		return nil, fmt.Errorf("Need one argument ts_code or trade_date")
	}

	body := map[string]interface{}{
		"api_name": "hsgt_top10",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}

	return api.postData(body)
}

// GgtTop10 获取港股通每日成交数据，其中包括沪市、深市详细数据
func (api *TuShare) GgtTop10(params map[string]string, fields []string) (*APIResponse, error) {
	// Check param
	_, hasTsCode := params["ts_code"]
	_, hasTradeDate := params["trade_date"]

	// ts_code & trade_date required
	if (!hasTsCode && !hasTradeDate) || (hasTsCode && hasTradeDate) {
		return nil, fmt.Errorf("Need one argument ts_code or trade_date")
	}

	body := map[string]interface{}{
		"api_name": "ggt_top10",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}

	return api.postData(body)
}

// Margin 获取融资融券每日交易汇总数据
func (api *TuShare) Margin(params map[string]string, fields []string) (*APIResponse, error) {
	// Check params
	_, hasTsCode := params["trade_date"]
	if !hasTsCode {
		return nil, fmt.Errorf("trade_date is a required argument")
	}

	body := map[string]interface{}{
		"api_name": "margin",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}

	return api.postData(body)
}

// MarginDetail 获取沪深两市每日融资融券明细
func (api *TuShare) MarginDetail(params map[string]string, fields []string) (*APIResponse, error) {
	// Check params
	_, hasTsCode := params["trade_date"]
	if !hasTsCode {
		return nil, fmt.Errorf("trade_date is a required argument")
	}

	body := map[string]interface{}{
		"api_name": "margin_detail",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}

	return api.postData(body)
}

// Top10Holders 获取上市公司前十大股东数据，包括持有数量和比例等信息
func (api *TuShare) Top10Holders(params map[string]string, fields []string) (*APIResponse, error) {
	// Check params
	_, hasTsCode := params["ts_code"]
	if !hasTsCode {
		return nil, fmt.Errorf("ts_code is a required argument")
	}

	body := map[string]interface{}{
		"api_name": "top10_holders",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}

	return api.postData(body)
}

// Top10FloatHolders 获取上市公司前十大流通股东数据
func (api *TuShare) Top10FloatHolders(params map[string]string, fields []string) (*APIResponse, error) {
	// Check params
	_, hasTsCode := params["ts_code"]
	if !hasTsCode {
		return nil, fmt.Errorf("ts_code is a required argument")
	}

	body := map[string]interface{}{
		"api_name": "top10_floatholders",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}

	return api.postData(body)
}

// TopList 龙虎榜每日交易明细
func (api *TuShare) TopList(params map[string]string, fields []string) (*APIResponse, error) {
	// Check params
	_, hasTsCode := params["trade_date"]
	if !hasTsCode {
		return nil, fmt.Errorf("trade_date is a required argument")
	}

	body := map[string]interface{}{
		"api_name": "top_list",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}

	return api.postData(body)
}

// TopInst 龙虎榜机构成交明细
func (api *TuShare) TopInst(params map[string]string, fields []string) (*APIResponse, error) {
	// Check params
	_, hasTsCode := params["trade_date"]
	if !hasTsCode {
		return nil, fmt.Errorf("trade_date is a required argument")
	}

	body := map[string]interface{}{
		"api_name": "top_inst",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}

	return api.postData(body)
}

// PledgeStat 获取股权质押统计数据
func (api *TuShare) PledgeStat(params map[string]string, fields []string) (*APIResponse, error) {
	// Check params
	_, hasTsCode := params["ts_code"]
	if !hasTsCode {
		return nil, fmt.Errorf("ts_code is a required argument")
	}

	body := map[string]interface{}{
		"api_name": "pledge_stat",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}

	return api.postData(body)
}

// PledgeDetail 获取股权质押明细数据
func (api *TuShare) PledgeDetail(params map[string]string, fields []string) (*APIResponse, error) {
	// Check params
	_, hasTsCode := params["ts_code"]
	if !hasTsCode {
		return nil, fmt.Errorf("ts_code is a required argument")
	}

	body := map[string]interface{}{
		"api_name": "pledge_detail",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}

	return api.postData(body)
}

// Repurchase 获取上市公司回购股票数据
func (api *TuShare) Repurchase(params map[string]string, fields []string) (*APIResponse, error) {

	body := map[string]interface{}{
		"api_name": "repurchase",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}

	return api.postData(body)
}

// Concept 获取概念股分类，目前只有ts一个来源，未来将逐步增加来源
func (api *TuShare) Concept(params map[string]string, fields []string) (*APIResponse, error) {

	body := map[string]interface{}{
		"api_name": "concept",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}

	return api.postData(body)
}

// ConceptDetail 获取概念股分类明细数据
func (api *TuShare) ConceptDetail(params map[string]string, fields []string) (*APIResponse, error) {
	// Check params
	_, hasTsCode := params["id"]
	if !hasTsCode {
		return nil, fmt.Errorf("id is a required argument")
	}

	body := map[string]interface{}{
		"api_name": "concept_detail",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}

	return api.postData(body)
}

// ShareFloat 获取限售股解禁
func (api *TuShare) ShareFloat(params map[string]string, fields []string) (*APIResponse, error) {

	body := map[string]interface{}{
		"api_name": "share_float",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}

	return api.postData(body)
}

// BlockTrade 大宗交易
func (api *TuShare) BlockTrade(params map[string]string, fields []string) (*APIResponse, error) {

	body := map[string]interface{}{
		"api_name": "block_trade",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}

	return api.postData(body)
}

// StkAccount 获取股票账户开户数据，统计周期为一周
func (api *TuShare) StkAccount(params map[string]string, fields []string) (*APIResponse, error) {

	body := map[string]interface{}{
		"api_name": "stk_account",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}

	return api.postData(body)
}

// StkHolderNumber 获取上市公司股东户数数据，数据不定期公布
func (api *TuShare) StkHolderNumber(params map[string]string, fields []string) (*APIResponse, error) {

	body := map[string]interface{}{
		"api_name": "stk_holdernumber",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}

	return api.postData(body)
}
