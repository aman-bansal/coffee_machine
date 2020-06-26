package model

import (
	"errors"
	"sync"
)

//this is used to maintain machine state i.e. which outlet is available which isn't
//this state will always be managed at the run time. Even if the machine restarts all the outlets will available
type MachineOutletPool struct {
	numOfOutlets     int
	isMachineReady   bool
	availableOutlets map[int]struct{}
	inuseOutlets     map[int]struct{}
	mutex            sync.Mutex
}

func NewMachineOutletPool(numOfOutlets int) *MachineOutletPool {
	idx := 1
	availableOutlets := make(map[int]struct{})
	for {
		if idx > numOfOutlets { break }
		availableOutlets[idx] = struct{}{}
		idx = idx + 1
	}
	return &MachineOutletPool{
		numOfOutlets:     numOfOutlets,
		availableOutlets: availableOutlets,
		inuseOutlets:     make(map[int]struct{}),
	}
}

func (m *MachineOutletPool) MarkAllOutletsAvailable() {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if len(m.inuseOutlets) != 0 {
		for key, val := range m.inuseOutlets {
			m.availableOutlets[key] = val
		}
		m.inuseOutlets = make(map[int]struct{})
	}
}

func (m *MachineOutletPool) GetAvailableOutlet() (int, error) {
	m.mutex.Lock()
	if len(m.availableOutlets) == 0 {
		m.mutex.Unlock()
		return 0, errors.New("no outletId available. Please wait sometime")
	}

	var outletId int
	for key := range m.availableOutlets {
		outletId = key
		break
	}

	if outletId == 0 {
		m.mutex.Unlock()
		return 0, errors.New("no outletId available. Please wait sometime")
	}

	delete(m.availableOutlets, outletId)
	m.inuseOutlets[outletId] = struct{}{}
	m.mutex.Unlock()
	return outletId, nil
}

func (m *MachineOutletPool) MarkOutletFree(outletId int) error {
	m.mutex.Lock()
	if _, ok := m.inuseOutlets[outletId]; ok {
		delete(m.inuseOutlets, outletId)
		m.availableOutlets[outletId] = struct{}{}
		m.mutex.Unlock()
		return nil
	}

	m.mutex.Unlock()
	return errors.New("outlet is already free")
}