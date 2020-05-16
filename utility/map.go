package utility

import "sync"

type Map struct {
	sync.RWMutex
	internal map[uint]uint
}

func NewMap() *Map {
	return &Map{
		internal: make(map[uint]uint),
	}
}

func (m *Map) Load(key uint) (value uint, ok bool) {
	m.RLock()
	result, ok := m.internal[key]
	m.RUnlock()
	return result, ok
}

func (m *Map) Delete(key uint) {
	m.Lock()
	delete(m.internal, key)
	m.Unlock()
}

func (m *Map) Store(key, value uint) {
	m.Lock()
	m.internal[key] = value
	m.Unlock()
}

func (m *Map) FindKeyByValue(value uint) (uint, bool) {
	m.RLock()
	defer m.RUnlock()
	for k, v := range m.internal {
		if v == value {
			return k, true
		}
	}
	return 0, false
}
