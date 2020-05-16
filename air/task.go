package air

import (
	"math"
	"time"
)

const (
	LowPriority = iota
	MediumPriority
	HighPriority
)

type Task struct {
	RoomID              uint
	LastModifiedTime    time.Time
	Priority            int
	DefaultPriority     int
	CurrentTemperature  float32
	TargetTemperature   float32
	FanSpeed            uint
	CurrentServiceCount uint
}

func (t *Task) clear() {
	t.LastModifiedTime = time.Now()
	t.CurrentServiceCount = 0
}

func (t *Task) OnFanSpeedChanged() {
	if t.CurrentServiceCount != 0 {
		GenerateFanSpeedChangedTicket(t)
	}
	t.clear()
}

func (t *Task) OnPriorityChanged() {
	if t.CurrentServiceCount != 0 {
		GeneratePriorityChangedTicket(t)
	}
	t.clear()
}

type TaskQueue []*Task

func (q TaskQueue) Len() int { return len(q) }

func (q TaskQueue) Less(i, j int) bool {
	this := q[i]
	other := q[j]
	thisDiffTemperature := math.Abs(float64(this.TargetTemperature - this.CurrentTemperature))
	otherDiffTemperature := math.Abs(float64(other.TargetTemperature - other.CurrentTemperature))
	if this.Priority > other.Priority {
		return true
	} else if this.Priority == other.Priority {
		if thisDiffTemperature > otherDiffTemperature {
			return true
		} else if math.Abs(thisDiffTemperature-otherDiffTemperature) < 0.01 {
			return this.FanSpeed > other.FanSpeed
		} else {
			return false
		}
	} else {
		return false
	}
}

func (q TaskQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *TaskQueue) Push(x interface{}) {
	*q = append(*q, x.(*Task))
}

func (q *TaskQueue) Pop() interface{} {
	old := *q
	n := len(old)
	task := old[n-1]
	*q = old[0 : n-1]
	return task
}
