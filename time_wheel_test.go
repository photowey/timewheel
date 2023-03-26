package timewheel

import (
	"container/list"
	"reflect"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	type args struct {
		interval time.Duration
		size     int
		job      Job
	}
	tests := []struct {
		name string
		args args
		want *TimeWheel
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.interval, tt.args.size, tt.args.job); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeWheel_Add(t *testing.T) {
	type fields struct {
		interval   time.Duration
		ticker     *time.Ticker
		size       int
		slots      []*list.List
		timer      map[any]int
		position   int
		job        Job
		addChan    chan Task
		removeChan chan any
		stopChan   chan bool
	}
	type args struct {
		delay time.Duration
		key   any
		data  any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tw := &TimeWheel{
				interval:   tt.fields.interval,
				ticker:     tt.fields.ticker,
				size:       tt.fields.size,
				wheel:      tt.fields.slots,
				timer:      tt.fields.timer,
				position:   tt.fields.position,
				job:        tt.fields.job,
				addChan:    tt.fields.addChan,
				removeChan: tt.fields.removeChan,
				stopChan:   tt.fields.stopChan,
			}
			if got := tw.Add(tt.args.delay, tt.args.key, tt.args.data); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeWheel_AddTask(t *testing.T) {
	type fields struct {
		interval   time.Duration
		ticker     *time.Ticker
		size       int
		slots      []*list.List
		timer      map[any]int
		position   int
		job        Job
		addChan    chan Task
		removeChan chan any
		stopChan   chan bool
	}
	type args struct {
		task Task
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tw := &TimeWheel{
				interval:   tt.fields.interval,
				ticker:     tt.fields.ticker,
				size:       tt.fields.size,
				wheel:      tt.fields.slots,
				timer:      tt.fields.timer,
				position:   tt.fields.position,
				job:        tt.fields.job,
				addChan:    tt.fields.addChan,
				removeChan: tt.fields.removeChan,
				stopChan:   tt.fields.stopChan,
			}
			if got := tw.AddTask(tt.args.task); got != tt.want {
				t.Errorf("AddTask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeWheel_Remove(t *testing.T) {
	type fields struct {
		interval   time.Duration
		ticker     *time.Ticker
		size       int
		slots      []*list.List
		timer      map[any]int
		position   int
		job        Job
		addChan    chan Task
		removeChan chan any
		stopChan   chan bool
	}
	type args struct {
		taskKey any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tw := &TimeWheel{
				interval:   tt.fields.interval,
				ticker:     tt.fields.ticker,
				size:       tt.fields.size,
				wheel:      tt.fields.slots,
				timer:      tt.fields.timer,
				position:   tt.fields.position,
				job:        tt.fields.job,
				addChan:    tt.fields.addChan,
				removeChan: tt.fields.removeChan,
				stopChan:   tt.fields.stopChan,
			}
			if got := tw.Remove(tt.args.taskKey); got != tt.want {
				t.Errorf("Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeWheel_Start(t *testing.T) {
	type fields struct {
		interval   time.Duration
		ticker     *time.Ticker
		size       int
		slots      []*list.List
		timer      map[any]int
		position   int
		job        Job
		addChan    chan Task
		removeChan chan any
		stopChan   chan bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tw := &TimeWheel{
				interval:   tt.fields.interval,
				ticker:     tt.fields.ticker,
				size:       tt.fields.size,
				wheel:      tt.fields.slots,
				timer:      tt.fields.timer,
				position:   tt.fields.position,
				job:        tt.fields.job,
				addChan:    tt.fields.addChan,
				removeChan: tt.fields.removeChan,
				stopChan:   tt.fields.stopChan,
			}
			tw.Start()
		})
	}
}

func TestTimeWheel_Stop(t *testing.T) {
	type fields struct {
		interval   time.Duration
		ticker     *time.Ticker
		size       int
		slots      []*list.List
		timer      map[any]int
		position   int
		job        Job
		addChan    chan Task
		removeChan chan any
		stopChan   chan bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tw := &TimeWheel{
				interval:   tt.fields.interval,
				ticker:     tt.fields.ticker,
				size:       tt.fields.size,
				wheel:      tt.fields.slots,
				timer:      tt.fields.timer,
				position:   tt.fields.position,
				job:        tt.fields.job,
				addChan:    tt.fields.addChan,
				removeChan: tt.fields.removeChan,
				stopChan:   tt.fields.stopChan,
			}
			tw.Stop()
		})
	}
}

func TestTimeWheel_add(t *testing.T) {
	type fields struct {
		interval   time.Duration
		ticker     *time.Ticker
		size       int
		slots      []*list.List
		timer      map[any]int
		position   int
		job        Job
		addChan    chan Task
		removeChan chan any
		stopChan   chan bool
	}
	type args struct {
		task *Task
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tw := &TimeWheel{
				interval:   tt.fields.interval,
				ticker:     tt.fields.ticker,
				size:       tt.fields.size,
				wheel:      tt.fields.slots,
				timer:      tt.fields.timer,
				position:   tt.fields.position,
				job:        tt.fields.job,
				addChan:    tt.fields.addChan,
				removeChan: tt.fields.removeChan,
				stopChan:   tt.fields.stopChan,
			}
			tw.add(tt.args.task)
		})
	}
}

func TestTimeWheel_remove(t *testing.T) {
	type fields struct {
		interval   time.Duration
		ticker     *time.Ticker
		size       int
		slots      []*list.List
		timer      map[any]int
		position   int
		job        Job
		addChan    chan Task
		removeChan chan any
		stopChan   chan bool
	}
	type args struct {
		taskKey any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tw := &TimeWheel{
				interval:   tt.fields.interval,
				ticker:     tt.fields.ticker,
				size:       tt.fields.size,
				wheel:      tt.fields.slots,
				timer:      tt.fields.timer,
				position:   tt.fields.position,
				job:        tt.fields.job,
				addChan:    tt.fields.addChan,
				removeChan: tt.fields.removeChan,
				stopChan:   tt.fields.stopChan,
			}
			tw.remove(tt.args.taskKey)
		})
	}
}

func TestTimeWheel_scan(t *testing.T) {
	type fields struct {
		interval   time.Duration
		ticker     *time.Ticker
		size       int
		slots      []*list.List
		timer      map[any]int
		position   int
		job        Job
		addChan    chan Task
		removeChan chan any
		stopChan   chan bool
	}
	type args struct {
		slot *list.List
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tw := &TimeWheel{
				interval:   tt.fields.interval,
				ticker:     tt.fields.ticker,
				size:       tt.fields.size,
				wheel:      tt.fields.slots,
				timer:      tt.fields.timer,
				position:   tt.fields.position,
				job:        tt.fields.job,
				addChan:    tt.fields.addChan,
				removeChan: tt.fields.removeChan,
				stopChan:   tt.fields.stopChan,
			}
			tw.scan(tt.args.slot)
		})
	}
}

func TestTimeWheel_schedule(t *testing.T) {
	type fields struct {
		interval   time.Duration
		ticker     *time.Ticker
		size       int
		slots      []*list.List
		timer      map[any]int
		position   int
		job        Job
		addChan    chan Task
		removeChan chan any
		stopChan   chan bool
	}
	type args struct {
		delay time.Duration
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantPosition int
		wantRound    int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tw := &TimeWheel{
				interval:   tt.fields.interval,
				ticker:     tt.fields.ticker,
				size:       tt.fields.size,
				wheel:      tt.fields.slots,
				timer:      tt.fields.timer,
				position:   tt.fields.position,
				job:        tt.fields.job,
				addChan:    tt.fields.addChan,
				removeChan: tt.fields.removeChan,
				stopChan:   tt.fields.stopChan,
			}
			gotPosition, gotRound := tw.schedule(tt.args.delay)
			if gotPosition != tt.wantPosition {
				t.Errorf("schedule() gotPosition = %v, want %v", gotPosition, tt.wantPosition)
			}
			if gotRound != tt.wantRound {
				t.Errorf("schedule() gotRound = %v, want %v", gotRound, tt.wantRound)
			}
		})
	}
}

func TestTimeWheel_start(t *testing.T) {
	type fields struct {
		interval   time.Duration
		ticker     *time.Ticker
		size       int
		slots      []*list.List
		timer      map[any]int
		position   int
		job        Job
		addChan    chan Task
		removeChan chan any
		stopChan   chan bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tw := &TimeWheel{
				interval:   tt.fields.interval,
				ticker:     tt.fields.ticker,
				size:       tt.fields.size,
				wheel:      tt.fields.slots,
				timer:      tt.fields.timer,
				position:   tt.fields.position,
				job:        tt.fields.job,
				addChan:    tt.fields.addChan,
				removeChan: tt.fields.removeChan,
				stopChan:   tt.fields.stopChan,
			}
			tw.start()
		})
	}
}

func TestTimeWheel_tick(t *testing.T) {
	type fields struct {
		interval   time.Duration
		ticker     *time.Ticker
		size       int
		slots      []*list.List
		timer      map[any]int
		position   int
		job        Job
		addChan    chan Task
		removeChan chan any
		stopChan   chan bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tw := &TimeWheel{
				interval:   tt.fields.interval,
				ticker:     tt.fields.ticker,
				size:       tt.fields.size,
				wheel:      tt.fields.slots,
				timer:      tt.fields.timer,
				position:   tt.fields.position,
				job:        tt.fields.job,
				addChan:    tt.fields.addChan,
				removeChan: tt.fields.removeChan,
				stopChan:   tt.fields.stopChan,
			}
			tw.tick()
		})
	}
}
