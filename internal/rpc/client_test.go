package rpc

import "testing"

func TestGetClientMainnet(t *testing.T) {
	client := GetClient(Mainnet)

	if client == nil {
		t.Errorf("Expected client to be not nil")
	}
}
