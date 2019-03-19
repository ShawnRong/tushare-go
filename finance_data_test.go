package tushare

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInCome(t *testing.T) {
	ast := assert.New(t)
	params := make(map[string]string)
	var fields []string
	// Check params
	_, err := client.Income(params, fields)
	if err != nil {
		ast.Equal(err.Error(), "ts_code is a required argument")
	}
	params["ts_code"] = "000001.SZ"
	resp, err := client.Income(params, fields)

	if err != nil {
		if resp.Code == -2002 {
			ast.Equal(err.Error(), "Your point is not enough to use this api")
		}
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestBalanceSheet(t *testing.T) {
	ast := assert.New(t)
	params := make(map[string]string)
	var fields []string
	// Check params
	_, err := client.BalanceSheet(params, fields)
	if err != nil {
		ast.Equal(err.Error(), "ts_code is a required argument")
	}
	params["ts_code"] = "000001.SZ"
	resp, err := client.BalanceSheet(params, fields)

	if err != nil {
		if resp.Code == -2002 {
			ast.Equal(err.Error(), "Your point is not enough to use this api")
		}
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestCashFlow(t *testing.T) {
	ast := assert.New(t)
	params := make(map[string]string)
	var fields []string
	// Check params
	_, err := client.CashFlow(params, fields)
	if err != nil {
		ast.Equal(err.Error(), "ts_code is a required argument")
	}
	params["ts_code"] = "000001.SZ"
	resp, err := client.CashFlow(params, fields)

	if err != nil {
		if resp.Code == -2002 {
			ast.Equal(err.Error(), "Your point is not enough to use this api")
		}
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestForecast(t *testing.T) {
	ast := assert.New(t)
	params := make(map[string]string)
	var fields []string
	// Check params
	_, err := client.Forecast(params, fields)
	if err != nil {
		ast.Equal(err.Error(), "Need one argument ts_code or ann_date")
	}
	params["ts_code"] = "000001.SZ"
	resp, err := client.Forecast(params, fields)

	if err != nil {
		if resp.Code == -2002 {
			ast.Equal(err.Error(), "Your point is not enough to use this api")
		}
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestDividend(t *testing.T) {
	ast := assert.New(t)
	params := make(map[string]string)
	var fields []string
	resp, err := client.Dividend(params, fields)
	if err != nil {
		if resp.Code == -2002 {
			ast.Equal(err.Error(), "Your point is not enough to use this api")
		}
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestExpress(t *testing.T) {
	ast := assert.New(t)
	params := make(map[string]string)
	var fields []string
	// Check params
	_, err := client.Express(params, fields)
	if err != nil {
		ast.Equal(err.Error(), "ts_code is a required argument")
	}
	params["ts_code"] = "000001.SZ"
	resp, err := client.Express(params, fields)

	if err != nil {
		if resp.Code == -2002 {
			ast.Equal(err.Error(), "Your point is not enough to use this api")
		}
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestFinaIndicator(t *testing.T) {
	ast := assert.New(t)
	params := make(map[string]string)
	var fields []string
	// Check params
	_, err := client.FinaIndicator(params, fields)
	if err != nil {
		ast.Equal(err.Error(), "ts_code is a required argument")
	}
	params["ts_code"] = "000001.SZ"
	resp, err := client.FinaIndicator(params, fields)

	if err != nil {
		if resp.Code == -2002 {
			ast.Equal(err.Error(), "Your point is not enough to use this api")
		}
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestFinaAudit(t *testing.T) {
	ast := assert.New(t)
	params := make(map[string]string)
	var fields []string
	// Check params
	_, err := client.FinaAudit(params, fields)
	if err != nil {
		ast.Equal(err.Error(), "ts_code is a required argument")
	}
	params["ts_code"] = "000001.SZ"
	resp, err := client.FinaAudit(params, fields)

	if err != nil {
		if resp.Code == -2002 {
			ast.Equal(err.Error(), "Your point is not enough to use this api")
		}
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestFinaMainbz(t *testing.T) {
	ast := assert.New(t)
	params := make(map[string]string)
	var fields []string
	// Check params
	_, err := client.FinaMainbz(params, fields)
	if err != nil {
		ast.Equal(err.Error(), "ts_code is a required argument")
	}
	params["ts_code"] = "000001.SZ"
	resp, err := client.FinaMainbz(params, fields)

	if err != nil {
		if resp.Code == -2002 {
			ast.Equal(err.Error(), "Your point is not enough to use this api")
		}
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestDisclosureDate(t *testing.T) {
	ast := assert.New(t)
	params := make(map[string]string)
	var fields []string
	resp, err := client.DisclosureDate(params, fields)
	if err != nil {
		if resp.Code == -2002 {
			ast.Equal(err.Error(), "Your point is not enough to use this api")
		}
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}
