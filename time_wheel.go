package timewheel

import (
	"container/list"
	"time"
)

// Job 任务回调函数
type Job func(any)

// Handler Task 自定义事件回调函数
type Handler func(any)

type TimeWheel struct {
	interval   time.Duration // 时间间隔
	ticker     *time.Ticker
	size       int          // 槽数量
	slots      []*list.List // 轮槽 [ [bucket], [bucket],..., [bucket] ]
	timer      map[any]int  // 任务池, key:唯一标识 value: 槽 position
	position   int
	job        Job // Job 任务回调函数
	addChan    chan Task
	removeChan chan any
	stopChan   chan bool
}

// New 时间轮构造函数
func New(interval time.Duration, size int, job Job) *TimeWheel {
	// 理论上 task 增加了 自定义函数 HandleFun 可以不用指定 job 通用回调函数,
	// 但为了避免开发者又通过添加普通任务的函数(tw#Add)来添加任务,
	// 故, 这儿还是强制判断一下 job 非空
	// 换一种思路 ->
	// -> 如果 允许 job 为空 -> 那么,在添加任务的时候, job 为 nil, task.HandleFun 就必须指定
	if interval <= 0 || size <= 0 || job == nil {
		return nil
	}
	tw := &TimeWheel{
		interval:   interval,
		slots:      make([]*list.List, size),
		timer:      make(map[any]int),
		position:   0,
		job:        job,
		size:       size,
		addChan:    make(chan Task),
		removeChan: make(chan any),
		stopChan:   make(chan bool),
	}

	tw.init()

	return tw
}

// Start 启动时间轮
func (tw *TimeWheel) Start() {
	tw.ticker = time.NewTicker(tw.interval)

	go func() {
		tw.start()
	}()
}

// Stop 停止时间轮
func (tw *TimeWheel) Stop() {
	tw.stopChan <- true
}

// Add 添加普通任务
func (tw *TimeWheel) Add(delay time.Duration, key any, data any) bool {
	return tw.AddTask(Task{
		Delay: delay,
		Key:   key,
		Args:  data,
	})
}

// AddTask 添加指定任务 -> 开发者可以自定义 task 的信息, 最重要的是可以自定义回调函数
func (tw *TimeWheel) AddTask(task Task) bool {
	if task.Delay <= 0 {
		return false
	}
	if tw.job == nil && task.HandleFun == nil {
		return false
	}

	tw.addChan <- task

	return true
}

// Remove 移除任务
func (tw *TimeWheel) Remove(taskKey any) bool {
	if taskKey == nil {
		return false
	}

	tw.removeChan <- taskKey

	return true
}

func (tw *TimeWheel) init() {
	for i := 0; i < tw.size; i++ {
		tw.slots[i] = list.New()
	}
}

func (tw *TimeWheel) start() {
	for {
		select {
		case <-tw.ticker.C:
			tw.tick()
		case taskAdd := <-tw.addChan:
			tw.add(&taskAdd)
		case taskKey := <-tw.removeChan:
			tw.remove(taskKey)
		case <-tw.stopChan:
			tw.ticker.Stop()
			return
		}
	}
}

func (tw *TimeWheel) tick() {
	bucket := tw.slots[tw.position]
	tw.scan(bucket)
	if tw.position == tw.size-1 {
		tw.position = 0
	} else {
		tw.position++
	}
}

func (tw *TimeWheel) add(task *Task) {
	position, round := tw.schedule(task.Delay)
	task.round = round

	tw.slots[position].PushBack(task)
	if task.Key != nil {
		tw.timer[task.Key] = position
	}
}

func (tw *TimeWheel) remove(taskKey any) {
	if position, ok := tw.timer[taskKey]; ok {
		bucket := tw.slots[position]
		for element := bucket.Front(); element != nil; {
			task := element.Value.(*Task)
			if task.Key == taskKey {
				delete(tw.timer, task.Key)
				bucket.Remove(element)
			}

			element = element.Next()
		}
	}
}

func (tw *TimeWheel) scan(bucket *list.List) {
	for element := bucket.Front(); element != nil; {
		task := element.Value.(*Task)
		if task.round > 0 {
			task.round--
			element = element.Next()
			continue
		}

		// TODO 采用 协程池 来处理
		go func() {
			// TODO 如果 task 指定了回调函数, 就执行 task.HandleFun
			//  否则, 执行通用回调函数
			if task.HandleFun != nil {
				task.HandleFun(task.Args)
			} else {
				tw.job(task.Args)
			}
		}()

		next := element.Next()
		bucket.Remove(element)
		if task.Key != nil {
			delete(tw.timer, task.Key)
		}
		element = next
	}
}

func (tw *TimeWheel) schedule(delay time.Duration) (position int, round int) {
	delaySeconds := int(delay.Seconds())
	intervalSeconds := int(tw.interval.Seconds())
	round = delaySeconds / intervalSeconds / tw.size
	position = (tw.position + delaySeconds/intervalSeconds) % tw.size

	return
}
