package air

import (
	"container/heap"
	"math"
)

type Scheduler struct {
	queue TaskQueue
}

func (s *Scheduler) Init() {
	heap.Init(&s.queue)
}

func (s *Scheduler) AddTask(task *Task) {
	heap.Push(&s.queue, task)
}

func (s *Scheduler) getTaskIndexByRoomID(roomID uint) (int, bool) {
	index := 0
	flag := false
	for i, val := range s.queue {
		if val.RoomID == roomID {
			flag = true
			index = i
		}
	}
	return index, flag
}

func (s *Scheduler) GetTaskByRoomID(roomID uint) Task {
	i, flag := s.getTaskIndexByRoomID(roomID)
	if flag {
		return *s.queue[i]
	}
	return Task{
		RoomID: 0,
	}
}

func (s *Scheduler) RemoveTaskByRoomID(roomID uint) {
	i, ok := s.getTaskIndexByRoomID(roomID)
	if ok {
		if s.queue[i].CurrentServiceCount != 0 {
			GenerateShutdownTicket(s.queue[i])
		}
		heap.Remove(&s.queue, i)
	}
}

func (s *Scheduler) UpdateTaskByRoomID(roomID uint, currentTemperature float32, targetTemperature float32, fanSpeed uint) {
	i, ok := s.getTaskIndexByRoomID(roomID)
	if ok {
		if s.queue[i].FanSpeed != fanSpeed {
			s.queue[i].OnFanSpeedChanged()
		}
		diffTemperature := math.Abs(float64(targetTemperature - currentTemperature))
		s.queue[i].FanSpeed = fanSpeed
		s.queue[i].CurrentTemperature = currentTemperature
		s.queue[i].TargetTemperature = targetTemperature
		if diffTemperature < 1.0 && s.queue[i].DefaultPriority > MediumPriority {
			s.queue[i].Priority = MediumPriority
		} else if diffTemperature < 0.5 && s.queue[i].DefaultPriority > LowPriority {
			s.queue[i].Priority = LowPriority
		} else {
			s.queue[i].Priority = s.queue[i].DefaultPriority
		}
		heap.Fix(&s.queue, i)
	}
}

func (s *Scheduler) UpdatePriorityByRoomID(roomID uint, priority int) {
	i, ok := s.getTaskIndexByRoomID(roomID)
	if ok {
		if s.queue[i].DefaultPriority != priority {
			s.queue[i].OnPriorityChanged()
		}
		s.queue[i].DefaultPriority = priority
		heap.Fix(&s.queue, i)
	}
}
