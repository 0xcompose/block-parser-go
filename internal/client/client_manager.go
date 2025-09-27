package client

import (
	"fmt"
	"sync"

	"github.com/ethereum/go-ethereum/rpc"
)

type ClientManager struct {
	clients map[ChainID]*rpc.Client
	mu      sync.RWMutex
}

var (
	manager     *ClientManager
	managerOnce sync.Once
)

func GetClientManager() *ClientManager {
	managerOnce.Do(func() {
		manager = &ClientManager{
			clients: make(map[ChainID]*rpc.Client),
		}
	})
	return manager
}

func (cm *ClientManager) GetClient(chainID ChainID) (*rpc.Client, error) {
	cm.mu.RLock()
	if client, exists := cm.clients[chainID]; exists {
		cm.mu.RUnlock()
		return client, nil
	}
	cm.mu.RUnlock()

	// Need to create client
	cm.mu.Lock()
	defer cm.mu.Unlock()

	// Double check after acquiring write lock
	if client, exists := cm.clients[chainID]; exists {
		return client, nil
	}

	client, err := NewClient(chainID)

	if err != nil {
		return nil, fmt.Errorf("failed to create client for chain %d: %w", chainID, err)
	}

	cm.clients[chainID] = client
	return client, nil
}

func (cm *ClientManager) Close() {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	for _, client := range cm.clients {
		client.Close()
	}
	cm.clients = make(map[ChainID]*rpc.Client)
}
