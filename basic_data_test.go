package TuShare

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
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
	ast := assert.New(t)

	params := make(map[string]string)
	var fields []string
	resp, err := client.StockBasic(params, fields)

	//Check api without token
	if(token == "")	{
		ast.Equal(resp.Code, -2001, "response %v", resp)
	}

	if err != nil {
		t.Fatal(err)
	}
}
