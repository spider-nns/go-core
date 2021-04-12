package main

import (
	"fmt"
	"sync"
)

//工作任务
type task struct {
	begin  int
	end    int
	result chan<- int
}

//任务执行，计算begin 到 end 的和
//执行结果写入结果 chan result
func (t *task) do() {
	sum := 0
	for i := t.begin; i <= t.end; i++ {
		sum += i
	}
	t.result <- sum
}
func main() {
	//创建任务管道
	taskChain := make(chan task, 10)
	//创建结果通道
	resultChain := make(chan int, 10)
	//wait 用于同步等待任务的执行
	wait := &sync.WaitGroup{}
	go InitTask(taskChain, resultChain, 100)
	//初始化task的goroutine进行处理
	go DistributeTask(taskChain, wait, resultChain)
	//通过结果通道获取结果并汇总
	sum := ProcessResult(resultChain)
	fmt.Println("sum=", sum)

}

// InitTask 构建task 并写入task通道
func InitTask(taskChan chan<- task, r chan int, p int) {
	qu := p / 10
	mod := p % 10
	high := qu * 10
	for j := 0; j < qu; j++ {
		b := 10*j + 1
		e := 10 * (j + 1)
		tsk := task{
			begin:  b,
			end:    e,
			result: r,
		}
		taskChan <- tsk
	}
	if mod != 0 {
		tsk := task{
			begin:  high + 1,
			end:    p,
			result: r,
		}
		taskChan <- tsk
	}
	close(taskChan)
}

// DistributeTask 读取task chan,每个task启动一个worker goroutine 进行处理，等待每个task运行完，关闭结果通道
func DistributeTask(taskChan <-chan task, wait *sync.WaitGroup, result chan int) {
	for v := range taskChan {
		wait.Add(1)
		go ProcessTask(v, wait)
	}
	wait.Wait()
	close(result)
}

// ProcessTask 读取结果通道，汇总结果
func ProcessTask(t task, wait *sync.WaitGroup) {
	t.do()
	wait.Done()
}

// ProcessResult 处理具体工作，将处理结果发送到结果通道
func ProcessResult(resultChan chan int) int {
	sum := 0
	for r := range resultChan {
		sum += r
	}
	return sum
}
