package tushare

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var token = ""
var client = New(token)

func init() {
	envToken := os.Getenv("TUSHARE_TOKEN")
	if envToken != "" {
		token = envToken
	}
}

func TestMain(m *testing.M) {
	client = New(token)
	os.Exit(m.Run())
}

func TestStockBasic(t *testing.T) {
	params := make(map[string]string)
	params["is_hs"] = "N"
	params["list_status"] = "L"
	params["exchange"] = "SSE"
	var fields []string
	resp, err := client.StockBasic(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestInvalidField(t *testing.T) {
	ast := assert.New(t)
	params := make(map[string]string)
	var fields []string
	fields = append(fields, "invalid_field")
	resp, err := client.StockBasic(params, fields)

	if err != nil {
		if resp.Code == -2001 {
			ast.Equal(err.Error(), fmt.Sprintf("Argument error: %s", resp.Msg))
		}
	}
}

func TestTradeCal(t *testing.T) {
	params := make(map[string]string)
	params["exchange"] = "SSE"
	params["start_date"] = "2017-01-01"
	params["end_date"] = "2019-01-01"
	params["is_open"] = "1"
	var fields []string
	resp, err := client.TradeCal(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}

	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestHSConst(t *testing.T) {
	params := make(map[string]string)
	params["hs_type"] = "SH"
	params["is_new"] = "1"
	var fields []string
	resp, err := client.HSConst(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestHSConstParamsRequired(t *testing.T) {
	ast := assert.New(t)
	params := make(map[string]string)
	var fields []string
	_, err := client.HSConst(params, fields)
	if err != nil {
		ast.Equal(err.Error(), "hs_type required")
	}
}

func TestNameChange(t *testing.T) {
	params := make(map[string]string)
	params["ts_code"] = "000001.SZ"
	params["start_date"] = "2017-01-01"
	params["end_date"] = "2019-01-01"
	var fields []string
	resp, err := client.NameChange(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestStockCompany(t *testing.T) {
	ast := assert.New(t)
	params := make(map[string]string)
	params["exchange"] = "SSE"
	var fields []string
	resp, err := client.StockCompany(params, fields)

	if err != nil {
		if resp.Code == -2002 {
			ast.Equal(err.Error(), "Your point is not enough to use this api")
		}
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestNewShare(t *testing.T) {
	ast := assert.New(t)
	params := make(map[string]string)
	params["start_date"] = "2017-01-01"
	params["end_date"] = "2019-01-01"
	var fields []string
	resp, err := client.NewShare(params, fields)

	if err != nil {
		if resp.Code == -2002 {
			ast.Equal(err.Error(), "Your point is not enough to use this api")
		}
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}
