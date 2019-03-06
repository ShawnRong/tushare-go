package TuShare

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDaily(t *testing.T) {
	ast := assert.New(t)
	params := make(map[string]string)
	var fields []string
	// Check params
	_, err := client.Daily(params, fields)
	if err != nil {
		ast.Equal(err.Error(), "Need one argument ts_code or trade_date")
	}

	params["trade_date"] = "20181101"
	resp, err := client.Daily(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestDailyInvalidDateArgs(t *testing.T) {
	ast := assert.New(t)
	params := make(map[string]string)
	params["trade_date"] = "2018-11-01"
	var fields []string
	_, err := client.Daily(params, fields)

	if err != nil {
		ast.Equal(err.Error(), "Please input right date format YYYYMMDD!")
	}
}

func TestWeekly(t *testing.T) {
	ast := assert.New(t)
	params := make(map[string]string)
	var fields []string
	// Check params
	_, err := client.Weekly(params, fields)
	if err != nil {
		ast.Equal(err.Error(), "Need one argument ts_code or trade_date")
	}

	params["trade_date"] = "20181101"
	resp, err := client.Weekly(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestWeeklyInvalidDateArgs(t *testing.T) {
	ast := assert.New(t)
	params := make(map[string]string)
	params["trade_date"] = "2018-11-01"
	var fields []string
	_, err := client.Weekly(params, fields)

	if err != nil {
		ast.Equal(err.Error(), "Please input right date format YYYYMMDD!")
	}
}

func TestMonthly(t *testing.T) {
	ast := assert.New(t)
	var fields []string
	params := make(map[string]string)
	// Check params
	_, err := client.Monthly(params, fields)
	if err != nil {
		ast.Equal(err.Error(), "Need one argument ts_code or trade_date")
	}

	params["trade_date"] = "20181101"
	resp, err := client.Monthly(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestMonthlyInvalidDateArgs(t *testing.T) {
	ast := assert.New(t)
	params := make(map[string]string)
	params["trade_date"] = "2018-11-01"
	var fields []string
	_, err := client.Monthly(params, fields)

	if err != nil {
		ast.Equal(err.Error(), "Please input right date format YYYYMMDD!")
	}
}

func TestDailyBasic(t *testing.T) {
	ast := assert.New(t)
	params := make(map[string]string)
	var fields []string
	// Check params
	_, err := client.DailyBasic(params, fields)
	if err != nil {
		ast.Equal(err.Error(), "Need one argument ts_code or trade_date")
	}

	params["trade_date"] = "20181101"
	resp, err := client.DailyBasic(params, fields)

	if err != nil {
		if resp.Code == -2002 {
			ast.Equal(err.Error(), "Your point is not enough to use this api")
		}
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestDailyBasicInvalidDateArgs(t *testing.T) {
	ast := assert.New(t)
	params := make(map[string]string)
	params["trade_date"] = "2018-11-01"
	var fields []string
	_, err := client.DailyBasic(params, fields)

	if err != nil {
		ast.Equal(err.Error(), "Please input right date format YYYYMMDD!")
	}
}

func TestAdjFactor(t *testing.T) {
	ast := assert.New(t)
	params := make(map[string]string)
	var fields []string
	// Check params
	_, err := client.AdjFactor(params, fields)
	if err != nil {
		ast.Equal(err.Error(), "Need one argument ts_code or trade_date")
	}

	params["trade_date"] = "20181101"
	resp, err := client.AdjFactor(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestAdjFactorInvalidDateArgs(t *testing.T) {
	ast := assert.New(t)
	params := make(map[string]string)
	params["trade_date"] = "2018-11-01"
	var fields []string
	_, err := client.AdjFactor(params, fields)

	if err != nil {
		ast.Equal(err.Error(), "Please input right date format YYYYMMDD!")
	}
}

func TestSuspend(t *testing.T) {
	ast := assert.New(t)
	params := make(map[string]string)
	paramsTest := make(map[string]string)
	var fields []string
	// Check params
	paramsTest["ts_code"] = "000001.SZ"
	paramsTest["resume_date"] = "20181102"
	_, err := client.Suspend(paramsTest, fields)
	if err != nil {
		ast.Equal(err.Error(), "Need one argument among ts_code, suspend_date, resume_date")
	}

	params["suspend_date"] = "20181101"
	resp, err := client.Suspend(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestSuspendInvalidDateArgs(t *testing.T) {
	ast := assert.New(t)
	params := make(map[string]string)
	params["suspend_date"] = "2018-11-01"
	var fields []string
	_, err := client.Suspend(params, fields)

	if err != nil {
		ast.Equal(err.Error(), "Please input right date format YYYYMMDD!")
	}
}
