package wallet_test

import (
	"testing"

	"github.com/realjf/blockchain_demo/wallet"
)

func TestWallet(t *testing.T) {
	cases := map[string]struct {
		f func()
	}{
		"makewallet": {
			f: func() {
				w := wallet.MakeWallet()
				w.Address()
			},
		},
	}

	for name, ts := range cases {
		t.Run(name, func(t *testing.T) {
			ts.f()
		})
	}
}
