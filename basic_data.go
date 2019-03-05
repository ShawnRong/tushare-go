package TuShare

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
	return resp, nil
}

