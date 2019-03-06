package TuShare

import "fmt"

// 获取基础信息数据，包括股票代码、名称、上市日期、退市日期等
func (api *TuShare)StockBasic(params map[string]string, fields []string) (*ApiResponse, error) {
	body := map[string]interface{}{
		"api_name": "stock_basic",
		"token": api.token,
	}
	bodyParam := make(map[string] string)
	// Add params
	if value, ok := params["is_hs"]; ok {
		bodyParam["is_hs"] = value
	}
	if value, ok := params["list_status"]; ok {
		bodyParam["list_status"] = value
	}
	if value, ok := params["exchange"]; ok {
		bodyParam["exchange"] = value
	}

	body["params"] = bodyParam
	body["fields"] = fields

	req, err := api.request("POST", API_URL, body)
	if err != nil {
		return nil, err
	}
	resp, err := api.doRequest(req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// 获取各大交易所交易日历数据,默认提取的是上交所
func (api *TuShare)TradeCal(params map[string]string, fields []string) (*ApiResponse, error) {
	body := map[string]interface{}{
		"api_name": "trade_cal",
		"token": api.token,
	}
	bodyParam := make(map[string] string)
	// Add params
	if value, ok := params["exchange"]; ok {
		bodyParam["exchange"] = value
	}
	if value, ok := params["start_date"]; ok {
		bodyParam["start_date"] = value
	}
	if value, ok := params["end_date"]; ok {
		bodyParam["end_date"] = value
	}
	if value, ok := params["is_open"]; ok {
		bodyParam["is_open"] = value
	}

	body["params"] = bodyParam
	body["fields"] = fields

	req, err := api.request("POST", API_URL, body)
	if err != nil {
		return nil, err
	}
	resp, err := api.doRequest(req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// 获取沪股通、深股通成分数据
func (api *TuShare)HSConst(params map[string]string, fields []string) (*ApiResponse, error) {
	body := map[string]interface{}{
		"api_name": "hs_const",
		"token": api.token,
	}
	bodyParam := make(map[string] string)
	// Add params
	if value, ok := params["hs_type"]; !ok {
		return nil, fmt.Errorf("hs_type required!")
	} else {
		bodyParam["hs_type"] = value
	}

	if value, ok := params["is_new"]; ok {
		bodyParam["is_new"] = value
	}

	body["params"] = bodyParam
	body["fields"] = fields

	req, err := api.request("POST", API_URL, body)
	if err != nil {
		return nil, err
	}
	resp, err := api.doRequest(req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// 历史名称变更记录
func (api *TuShare)NameChange(params map[string]string, fields []string) (*ApiResponse, error) {
	body := map[string]interface{}{
		"api_name": "namechange",
		"token": api.token,
	}
	bodyParam := make(map[string] string)
	// Add params
	if value, ok := params["ts_code"]; ok {
		bodyParam["ts_code"] = value
	}

	if value, ok := params["start_date"]; ok {
		bodyParam["start_date"] = value
	}

	if value, ok := params["end_date"]; ok {
		bodyParam["end_date"] = value
	}

	body["params"] = bodyParam
	body["fields"] = fields

	req, err := api.request("POST", API_URL, body)
	if err != nil {
		return nil, err
	}
	resp, err := api.doRequest(req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// 获取上市公司基础信息
func (api *TuShare)StockCompany(params map[string]string, fields []string) (*ApiResponse, error) {
	body := map[string]interface{}{
		"api_name": "stock_company",
		"token": api.token,
	}
	bodyParam := make(map[string] string)
	// Add params
	if value, ok := params["exchange"]; ok {
		bodyParam["exchange"] = value
	}

	body["params"] = bodyParam
	body["fields"] = fields

	req, err := api.request("POST", API_URL, body)
	if err != nil {
		return nil, err
	}
	resp, err := api.doRequest(req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// 获取新股上市列表数据
func (api *TuShare)NewShare(params map[string]string, fields []string) (*ApiResponse, error) {
	body := map[string]interface{}{
		"api_name": "new_share",
		"token": api.token,
	}
	bodyParam := make(map[string] string)
	// Add params
	if value, ok := params["start_date"]; ok {
		bodyParam["start_date"] = value
	}
	if value, ok := params["end_date"]; ok {
		bodyParam["end_date"] = value
	}

	body["params"] = bodyParam
	body["fields"] = fields

	req, err := api.request("POST", API_URL, body)
	if err != nil {
		return nil, err
	}
	resp, err := api.doRequest(req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
