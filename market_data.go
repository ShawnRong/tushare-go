package tushare

import "fmt"

// Daily 获取股票行情数据, 日线行情
func (api *TuShare) Daily(params map[string]string, fields []string) (*APIResponse, error) {
	// Check params
	_, hasTsCode := params["ts_code"]
	_, hasTradeDate := params["trade_date"]

	// ts_code & trade_date required
	if (!hasTsCode && !hasTradeDate) || (hasTsCode && hasTradeDate) {
		return nil, fmt.Errorf("Need one argument ts_code or trade_date")
	}

	if dateFormat := IsDateFormat(params["trade_date"], params["start_date"], params["end_date"]); !dateFormat {
		return nil, fmt.Errorf("please input right date format YYYYMMDD")
	}

	body := map[string]interface{}{
		"api_name": "daily",
		"token":    api.token,
		"fields":   fields,
		"params":   params,
	}

	return api.postData(body)
}

// Weekly 获取股票行情数据, 周线行情
func (api *TuShare) Weekly(params map[string]string, fields []string) (*APIResponse, error) {
	// Check params
	_, hasTsCode := params["ts_code"]
	_, hasTradeDate := params["trade_date"]

	// ts_code & trade_date required
	if (!hasTsCode && !hasTradeDate) || (hasTsCode && hasTradeDate) {
		return nil, fmt.Errorf("Need one argument ts_code or trade_date")
	}

	if dateFormat := IsDateFormat(params["trade_date"], params["start_date"], params["end_date"]); !dateFormat {
		return nil, fmt.Errorf("please input right date format YYYYMMDD")
	}

	body := map[string]interface{}{
		"api_name": "weekly",
		"token":    api.token,
		"fields":   fields,
		"params":   params,
	}

	return api.postData(body)
}

// Monthly 获取A股月线数据
func (api *TuShare) Monthly(params map[string]string, fields []string) (*APIResponse, error) {
	// Check params
	_, hasTsCode := params["ts_code"]
	_, hasTradeDate := params["trade_date"]

	// ts_code & trade_date required
	if (!hasTsCode && !hasTradeDate) || (hasTsCode && hasTradeDate) {
		return nil, fmt.Errorf("Need one argument ts_code or trade_date")
	}

	if dateFormat := IsDateFormat(params["trade_date"], params["start_date"], params["end_date"]); !dateFormat {
		return nil, fmt.Errorf("please input right date format YYYYMMDD")
	}

	body := map[string]interface{}{
		"api_name": "monthly",
		"token":    api.token,
		"fields":   fields,
		"params":   params,
	}

	return api.postData(body)
}

// DailyBasic 获取全部股票每日重要的基本面指标，可用于选股分析、报表展示等
func (api *TuShare) DailyBasic(params map[string]string, fields []string) (*APIResponse, error) {
	// Check params
	_, hasTsCode := params["ts_code"]
	_, hasTradeDate := params["trade_date"]

	// ts_code & trade_date required
	if (!hasTsCode && !hasTradeDate) || (hasTsCode && hasTradeDate) {
		return nil, fmt.Errorf("Need one argument ts_code or trade_date")
	}

	if dateFormat := IsDateFormat(params["trade_date"], params["start_date"], params["end_date"]); !dateFormat {
		return nil, fmt.Errorf("please input right date format YYYYMMDD")
	}

	body := map[string]interface{}{
		"api_name": "daily_basic",
		"token":    api.token,
		"fields":   fields,
		"params":   params,
	}

	return api.postData(body)
}

// AdjFactor 获取股票复权因子，可提取单只股票全部历史复权因子，也可以提取单日全部股票的复权因子
func (api *TuShare) AdjFactor(params map[string]string, fields []string) (*APIResponse, error) {
	// Check params
	_, hasTsCode := params["ts_code"]
	_, hasTradeDate := params["trade_date"]

	// ts_code & trade_date required
	if (!hasTsCode && !hasTradeDate) || (hasTsCode && hasTradeDate) {
		return nil, fmt.Errorf("Need one argument ts_code or trade_date")
	}

	if dateFormat := IsDateFormat(params["trade_date"], params["start_date"], params["end_date"]); !dateFormat {
		return nil, fmt.Errorf("please input right date format YYYYMMDD")
	}

	body := map[string]interface{}{
		"api_name": "adj_factor",
		"token":    api.token,
		"fields":   fields,
		"params":   params,
	}

	return api.postData(body)
}

// Suspend 获取股票每日停复牌信息
func (api *TuShare) Suspend(params map[string]string, fields []string) (*APIResponse, error) {
	// Check params
	_, hasTsCode := params["ts_code"]
	_, hasTradeDate := params["suspend_date"]
	_, hasResumeDate := params["resume_date"]

	// ts_code & trade_date & resume_date only required
	argsCount := 0
	if hasTsCode {
		argsCount++
	}
	if hasTradeDate {
		argsCount++
	}
	if hasResumeDate {
		argsCount++
	}

	if argsCount != 1 {
		return nil, fmt.Errorf("Need one argument among ts_code, suspend_date, resume_date")
	}

	if dateFormat := IsDateFormat(params["suspend_date"], params["resume_date"]); !dateFormat {
		return nil, fmt.Errorf("please input right date format YYYYMMDD")
	}

	body := map[string]interface{}{
		"api_name": "suspend",
		"token":    api.token,
		"fields":   fields,
		"params":   params,
	}

	return api.postData(body)
}
