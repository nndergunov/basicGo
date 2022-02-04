package pool

type workerPool interface {
	Start()
	AddTask(task func())
}

type Pool struct {
	workerNumber int
	taskChan     chan func()
}

func (p *Pool) Start() {
	for i := 0; i < p.workerNumber; i++ {
		go func() {
			for task := range p.taskChan {
				task()
			}
		}()
	}
}

func (p *Pool) AddTask(task func()) {
	p.taskChan <- task
}

func NewPool(workerNumber int) *Pool {
	taskChan := make(chan func(), 24)

	return &Pool{
		workerNumber: workerNumber,
		taskChan:     taskChan,
	}
}