package air

import (
	"github.com/name1e5s/acdc/config"
	"sync"
)

type Air struct {
	sync.RWMutex
	Power     bool
	scheduler Scheduler
	config    config.AirConfig
}

var (
	air  *Air
	once sync.Once
)

func GetAir() *Air {
	once.Do(func() {
		air = &Air{
			scheduler: Scheduler{
				queue: make(TaskQueue, 0),
			},
			config: config.GetConfig().AirConfig,
			Power:  false,
		}
	})
	return air
}

func (a *Air) SetConfig(config config.AirConfig) {
	a.Lock()
	a.config = config
	a.Unlock()
}

func (a *Air) Init() {
	a.Lock()
	a.scheduler.Init()
	a.Unlock()
}

func (a *Air) AddTask(task *Task) {
	a.Lock()
	a.scheduler.AddTask(task)
	a.Unlock()
}

func (a *Air) GetTaskByRoomID(roomID uint) Task {
	a.RLock()
	defer a.RUnlock()
	return a.scheduler.GetTaskByRoomID(roomID)
}

func (a *Air) RemoveTaskByRoomID(roomID uint) {
	a.Lock()
	a.scheduler.RemoveTaskByRoomID(roomID)
	a.Unlock()
}

func (a *Air) UpdateTaskByRoomID(roomID uint, currentTemperature float32, targetTemperature float32, fanSpeed uint) {
	a.Lock()
	a.scheduler.UpdateTaskByRoomID(roomID, currentTemperature, targetTemperature, fanSpeed)
	a.Unlock()
}

func (a *Air) UpdatePriorityByRoomID(roomID uint, priority int) {
	a.Lock()
	a.scheduler.UpdatePriorityByRoomID(roomID, priority)
	a.Unlock()
}

func (a *Air) GetServeOptionByRoomID(roomID uint) bool {
	a.RLock()
	defer a.RUnlock()
	index, ok := a.scheduler.getTaskIndexByRoomID(roomID)
	if ok == false {
		return ok
	}
	if index < a.config.MaxServeSize {
		a.scheduler.queue[index].CurrentServiceCount += 1
		return true
	}
	return false
}
