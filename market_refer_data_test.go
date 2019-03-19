package tushare

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoneyflowHsgt(t *testing.T) {
	ast := assert.New(t)
	params := make(map[string]string)
	var fields []string
	// Check params
	_, err := client.MoneyflowHsgt(params, fields)
	if err != nil {
		ast.Equal(err.Error(), "Need one argument trade_date or start_date")
	}
	params["trade_date"] = "20181101"
	resp, err := client.MoneyflowHsgt(params, fields)

	if err != nil {
		if resp.Code == -2002 {
			ast.Equal(err.Error(), "Your point is not enough to use this api")
		}
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestHsgtTop10(t *testing.T) {
	ast := assert.New(t)
	params := make(map[string]string)
	var fields []string
	// Check params
	_, err := client.HsgtTop10(params, fields)
	if err != nil {
		ast.Equal(err.Error(), "Need one argument ts_code or trade_date")
	}
	params["ts_code"] = "000001.SZ"
	resp, err := client.HsgtTop10(params, fields)

	if err != nil {
		if resp.Code == -2002 {
			ast.Equal(err.Error(), "Your point is not enough to use this api")
		}
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestGgtTop10(t *testing.T) {
	ast := assert.New(t)
	params := make(map[string]string)
	var fields []string
	// Check params
	_, err := client.GgtTop10(params, fields)
	if err != nil {
		ast.Equal(err.Error(), "Need one argument ts_code or trade_date")
	}
	params["ts_code"] = "000001.SZ"
	resp, err := client.GgtTop10(params, fields)

	if err != nil {
		if resp.Code == -2002 {
			ast.Equal(err.Error(), "Your point is not enough to use this api")
		}
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestMargin(t *testing.T) {
	ast := assert.New(t)
	params := make(map[string]string)
	var fields []string
	// Check params
	_, err := client.Margin(params, fields)
	if err != nil {
		ast.Equal(err.Error(), "trade_date is a required argument")
	}
	params["trade_date"] = "20181101"
	resp, err := client.Margin(params, fields)

	if err != nil {
		if resp.Code == -2002 {
			ast.Equal(err.Error(), "Your point is not enough to use this api")
		}
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestMarginDetail(t *testing.T) {
	ast := assert.New(t)
	params := make(map[string]string)
	var fields []string
	// Check params
	_, err := client.MarginDetail(params, fields)
	if err != nil {
		ast.Equal(err.Error(), "trade_date is a required argument")
	}
	params["trade_date"] = "20181101"
	resp, err := client.MarginDetail(params, fields)

	if err != nil {
		if resp.Code == -2002 {
			ast.Equal(err.Error(), "Your point is not enough to use this api")
		}
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestTop10Holders(t *testing.T) {
	ast := assert.New(t)
	params := make(map[string]string)
	var fields []string
	// Check params
	_, err := client.Top10Holders(params, fields)
	if err != nil {
		ast.Equal(err.Error(), "ts_code is a required argument")
	}
	params["ts_code"] = "000001.SZ"
	resp, err := client.Top10Holders(params, fields)

	if err != nil {
		if resp.Code == -2002 {
			ast.Equal(err.Error(), "Your point is not enough to use this api")
		}
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestTop10FloatHolders(t *testing.T) {
	ast := assert.New(t)
	params := make(map[string]string)
	var fields []string
	// Check params
	_, err := client.Top10FloatHolders(params, fields)
	if err != nil {
		ast.Equal(err.Error(), "ts_code is a required argument")
	}
	params["ts_code"] = "000001.SZ"
	resp, err := client.Top10FloatHolders(params, fields)

	if err != nil {
		if resp.Code == -2002 {
			ast.Equal(err.Error(), "Your point is not enough to use this api")
		}
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestTopList(t *testing.T) {
	ast := assert.New(t)
	params := make(map[string]string)
	var fields []string
	// Check params
	_, err := client.TopList(params, fields)
	if err != nil {
		ast.Equal(err.Error(), "trade_date is a required argument")
	}
	params["trade_date"] = "20181101"
	resp, err := client.TopList(params, fields)

	if err != nil {
		if resp.Code == -2002 {
			ast.Equal(err.Error(), "Your point is not enough to use this api")
		}
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestTopInst(t *testing.T) {
	ast := assert.New(t)
	params := make(map[string]string)
	var fields []string
	// Check params
	_, err := client.TopInst(params, fields)
	if err != nil {
		ast.Equal(err.Error(), "trade_date is a required argument")
	}
	params["trade_date"] = "20181101"
	resp, err := client.TopInst(params, fields)

	if err != nil {
		if resp.Code == -2002 {
			ast.Equal(err.Error(), "Your point is not enough to use this api")
		}
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestPledgeStat(t *testing.T) {
	ast := assert.New(t)
	params := make(map[string]string)
	var fields []string
	// Check params
	_, err := client.PledgeStat(params, fields)
	if err != nil {
		ast.Equal(err.Error(), "ts_code is a required argument")
	}
	params["ts_code"] = "000001.SZ"
	resp, err := client.PledgeStat(params, fields)

	if err != nil {
		if resp.Code == -2002 {
			ast.Equal(err.Error(), "Your point is not enough to use this api")
		}
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestPledgeDetail(t *testing.T) {
	ast := assert.New(t)
	params := make(map[string]string)
	var fields []string
	// Check params
	_, err := client.PledgeDetail(params, fields)
	if err != nil {
		ast.Equal(err.Error(), "ts_code is a required argument")
	}
	params["ts_code"] = "000001.SZ"
	resp, err := client.PledgeDetail(params, fields)

	if err != nil {
		if resp.Code == -2002 {
			ast.Equal(err.Error(), "Your point is not enough to use this api")
		}
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestRepurchase(t *testing.T) {
	ast := assert.New(t)
	params := make(map[string]string)
	var fields []string
	resp, err := client.Repurchase(params, fields)
	if err != nil {
		if resp.Code == -2002 {
			ast.Equal(err.Error(), "Your point is not enough to use this api")
		}
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestConcept(t *testing.T) {
	ast := assert.New(t)
	params := make(map[string]string)
	var fields []string
	resp, err := client.Concept(params, fields)
	if err != nil {
		if resp.Code == -2002 {
			ast.Equal(err.Error(), "Your point is not enough to use this api")
		}
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestConceptDetail(t *testing.T) {
	ast := assert.New(t)
	params := make(map[string]string)
	var fields []string
	// Check params
	_, err := client.ConceptDetail(params, fields)
	if err != nil {
		ast.Equal(err.Error(), "id is a required argument")
	}
	params["id"] = "000001"
	resp, err := client.ConceptDetail(params, fields)

	if err != nil {
		if resp.Code == -2002 {
			ast.Equal(err.Error(), "Your point is not enough to use this api")
		}
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestShareFloat(t *testing.T) {
	ast := assert.New(t)
	params := make(map[string]string)
	var fields []string
	resp, err := client.ShareFloat(params, fields)
	if err != nil {
		if resp.Code == -2002 {
			ast.Equal(err.Error(), "Your point is not enough to use this api")
		}
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestBlockTrade(t *testing.T) {
	ast := assert.New(t)
	params := make(map[string]string)
	var fields []string
	resp, err := client.BlockTrade(params, fields)
	if err != nil {
		if resp.Code == -2002 {
			ast.Equal(err.Error(), "Your point is not enough to use this api")
		}
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestStkAccount(t *testing.T) {
	ast := assert.New(t)
	params := make(map[string]string)
	var fields []string
	resp, err := client.StkAccount(params, fields)
	if err != nil {
		if resp.Code == -2002 {
			ast.Equal(err.Error(), "Your point is not enough to use this api")
		}
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestStkHolderNumber(t *testing.T) {
	ast := assert.New(t)
	params := make(map[string]string)
	var fields []string
	resp, err := client.StkHolderNumber(params, fields)
	if err != nil {
		if resp.Code == -2002 {
			ast.Equal(err.Error(), "Your point is not enough to use this api")
		}
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}
