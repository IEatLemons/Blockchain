package ganache

import (
	"fmt"
	"testing"
)

var Host = "HTTP://127.0.0.1:7545"
var Gan *Ganache

func InitG() {
	err := Init(Host)
	if err != nil {

	}
	Gan = New()
}

func TestNewCli(t *testing.T) {
	InitG()
	balance, err := Gan.Balance("0x2eE0523b0Aa301A85D897cD2cbC226064dCA4A45", nil)
	if err != nil {
		fmt.Println("err", err.Error())
	}

	fmt.Println("balance", balance)
}
