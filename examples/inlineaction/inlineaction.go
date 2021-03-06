package main

import (
	"github.com/uuosio/chain"
)

//packer
type Transfer struct {
	From     chain.Name
	To       chain.Name
	Quantity chain.Asset
	Memo     string
}

func main() {
	_, _, action := chain.GetApplyArgs()
	if action == chain.NewName("sayhello") {
		a := chain.NewAction(chain.PermissionLevel{chain.NewName("helloworld11"), chain.ActiveName}, chain.NewName("eosio.token"), chain.NewName("transfer"))

		t := Transfer{}
		t.From = chain.NewName("helloworld11")
		t.To = chain.NewName("eosio")

		// Send 1.0 EOS
		t.Quantity.Amount = 10000
		t.Quantity.Symbol = chain.NewSymbol("EOS", 4)

		a.Data = t.Pack()
		a.Send()
	}

}
