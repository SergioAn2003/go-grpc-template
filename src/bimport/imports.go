package bimport

import "grpc-test/internal/bridge"

type BridgeImports struct {
	Bridge Bridge
}

func (b *BridgeImports) InitBridge(info bridge.Info) {
	b.Bridge = Bridge{
		Info: info,
	}
}

func NewEmptyBridge() *BridgeImports {
	return &BridgeImports{}
}
