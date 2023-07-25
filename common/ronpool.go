package common

import (
	"github.com/ethereum/go-ethereum/log"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

const (
	MaxRonTasks             = 24
	MaxRonTasksQueue        = 8192 // say we have 64 arbTxs to send for each peer, and 5 or more 'trade' servers, 24*64*2=3072.
	MaxRonHuntingTasks      = 24
	MaxRonHuntingTasksOnP2p = 8
	MaxRonHuntingTasksQueue = 64                                   // this is one single queue for all hunting tasks
	MaxArbTxAllowedInQueue  = 512                                  // at maximum, we allow 512 arbTx queued to be broadcast
	MaxRonTaskQueueSize     = MaxArbTxAllowedInQueue * MaxRonTasks // say we have 500 peers, we allow maximum 512 arbTxs to queue, then: 500*512/(500/24) = 512*24 = 12288
)

type SafeChan struct {
	Channel chan struct{}
	closed  bool
	lock    sync.Mutex
}

func NewSafeChan() *SafeChan {
	safeChan := SafeChan{}
	safeChan.Channel = make(chan struct{})
	return &safeChan
}
func (c *SafeChan) Close() {
	c.lock.Lock()
	if !c.closed {
		c.closed = true
		close(c.Channel)
	}
	c.lock.Unlock()
}

var (
	MutexMaxProc sync.Mutex

	RonTaskChan         = [MaxRonTasks]chan func(){}
	RonTaskChanNextIdle = int64(0)
	RonPoolClose        = NewSafeChan()

	RonHuntingTaskChan         = [MaxRonHuntingTasks]chan func(){}
	RonHuntingTaskChanNextIdle = int64(0)
	RonHuntingPoolClose        = NewSafeChan()

	runningRonTask     = int64(0)
	runningHuntingTask = int64(0)

	//ronTaskQueue        = NewRonQueue(MaxRonTaskQueueSize)
	//ronHuntingTaskQueue = NewRonQueue(MaxRonHuntingTasksQueue)

	maxRonHuntingTasks = int64(MaxRonHuntingTasks)
)

func IncGoMaxProcs(number int) int {
	MutexMaxProc.Lock()
	current := runtime.GOMAXPROCS(0)
	current += number
	runtime.GOMAXPROCS(current)
	MutexMaxProc.Unlock()
	return current
}

func RonTaskDispatch(task func()) bool {
	next := atomic.AddInt64(&RonTaskChanNextIdle, 1)
	idleIndex := (next - 1) % MaxRonTasks
	timeoutTimer := time.NewTimer(10 * time.Second)
	select {
	case RonTaskChan[idleIndex] <- task:
	case <-timeoutTimer.C:
		timeoutTimer.Stop()
		log.Warn("RonTaskDispatch channel write timeout")
		return false
	}

	return true
}

func RonTaskPool() {
	current := IncGoMaxProcs(MaxRonTasks)
	log.Info("RonFi arb LockOSThread for RonTaskPool", "procs", current)

	for i := 0; i < MaxRonTasks; i++ {
		RonTaskChan[i] = make(chan func(), 3) // i.e. only 3 arbTxs parallel sending is allowed
	}
	logIdle := false

	for i := 0; i < MaxRonTasks; i++ {
		channelIndex := i
		go func() {
			runtime.LockOSThread() // 1st priority routine
			for {
				select {
				case task := <-RonTaskChan[channelIndex]:
					if running := atomic.AddInt64(&runningRonTask, 1); !logIdle && running >= MaxRonTasks {
						log.Info("RonFi RonTaskPool busy", "running", running)
						logIdle = true
					}
					task()
					if running := atomic.AddInt64(&runningRonTask, -1); running <= 0 && logIdle {
						logIdle = false
						log.Info("RonFi RonTaskPool idle", "running", running)
					}
				case <-RonPoolClose.Channel:
					runtime.UnlockOSThread()
					log.Info("RonFi RonTaskPool exit")
					return
				}
			}
		}()
	}
}

func RonHuntingTaskDispatch(task func()) bool {
	next := atomic.AddInt64(&RonHuntingTaskChanNextIdle, 1)
	idleIndex := (next - 1) % maxRonHuntingTasks
	select {
	case RonHuntingTaskChan[idleIndex] <- task:
	default:
		log.Warn("RonHuntingTaskDispatch channel full")
		return false
	}
	return true
}

func RonHuntingTaskPool(maxTasks int) {
	RonHuntingPoolClose = NewSafeChan()
	maxRonHuntingTasks = int64(maxTasks) // save this into static variable
	current := IncGoMaxProcs(maxTasks)
	log.Info("RonFi arb LockOSThread for RonHuntingTaskPool", "procs", current)

	for i := 0; i < maxTasks; i++ {
		RonHuntingTaskChan[i] = make(chan func(), MaxRonHuntingTasksQueue)
	}
	logIdle := false
	for i := 0; i < maxTasks; i++ {
		channelIndex := i
		go func() {
			runtime.LockOSThread() // 1st priority routine
			for {
				select {
				case task := <-RonHuntingTaskChan[channelIndex]:
					if running := atomic.AddInt64(&runningHuntingTask, 1); !logIdle && running >= MaxRonHuntingTasks {
						logIdle = true
						log.Info("RonFi huntingTaskPool busy", "running", running)
					}
					task()
					if running := atomic.AddInt64(&runningHuntingTask, -1); running <= 0 && logIdle {
						logIdle = false
						log.Info("RonFi huntingTaskPool idle", "running", running)
					}
				case <-RonHuntingPoolClose.Channel:
					runtime.UnlockOSThread()
					remains := IncGoMaxProcs(-1)
					log.Info("RonFi RonHuntingTaskPool exit", "procs", remains)
					return
				}
			}
		}()
	}
}
