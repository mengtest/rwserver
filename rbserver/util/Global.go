package util

import (
	"../../rbwork/network"
	"sync"
)

//定义全局存储在线用户，键值为（IP+SN）或用户ID（建立连接时为IP,登录认证后为userId）
var Clients = NewClientMap()

//map 并发存取
type ClientMap struct {
	lock *sync.RWMutex
	bm   map[string]*network.TcpClient
}

func NewClientMap() *ClientMap {
	return &ClientMap{
		lock: new(sync.RWMutex),
		bm:   make(map[string]*network.TcpClient),
	}
}

func (m *ClientMap) GetMap() map[string]*network.TcpClient {
	return m.bm
}

//Get from maps return the k's value
func (m *ClientMap) Get(k string)*network.TcpClient {
	m.lock.RLock()
	defer m.lock.RUnlock()
	if val, ok := m.bm[k]; ok {
		return val
	}
	return nil
}

// Maps the given key and value. Returns false
// if the key is already in the map and changes nothing.
func (m *ClientMap) Set(k string,v *network.TcpClient) bool {
	m.lock.Lock()
	defer m.lock.Unlock()
	if val, ok := m.bm[k]; !ok {
		m.bm[k] = v
	} else if val != v {
		m.bm[k] = v
	} else {
		return false
	}
	return true
}

// Returns true if k is exist in the map.
func (m *ClientMap) Check(k string) bool {
	m.lock.RLock()
	defer m.lock.RUnlock()
	if _, ok := m.bm[k]; !ok {
		return false
	}
	return true
}

func (m *ClientMap) Delete(k string) {
	m.lock.Lock()
	defer m.lock.Unlock()
	delete(m.bm, k)
}
