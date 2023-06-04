package main

import (
	"fmt"
	"sort"
)

type Process struct {
	ID          string // 进程ID
	ArrivalTime int    // 到达时间
	ServiceTime int    // 要求服务时间
}

/*
	processes := []Process{
		{"A", 0, 3},
		{"B", 2, 6},
		{"C", 4, 4},
		{"D", 6, 5},
		{"E", 8, 2},
	}
*/

func main() {
	// 读取用户输入的进程列表
	var n int
	fmt.Println("请输入进程列表长度：")
	fmt.Scan(&n)
	processes := make([]Process, n)
	fmt.Println("请输入进程信息：")
	for i := 0; i < n; i++ {
		var id string
		var arrivalTime, serviceTime int
		fmt.Printf("请输入第 %d 个进程的 ID：", i+1)
		fmt.Scan(&id)
		fmt.Printf("请输入第 %d 个进程的到达时间：", i+1)
		fmt.Scan(&arrivalTime)
		fmt.Printf("请输入第 %d 个进程的要求服务时间：", i+1)
		fmt.Scan(&serviceTime)
		processes[i] = Process{id, arrivalTime, serviceTime}
	}

	fmt.Println("FCFS:")
	fcfs(processes)

	fmt.Println("\nRR (q=1):")
	rr(processes, 1)

	fmt.Println("\nSJF:")
	sjf(processes)

	fmt.Println("\nHRN:")
	hrn(processes)

}

func fcfs(processes []Process) {
	var (
		turnaroundTime         float64 // 周转时间
		weightedTurnaroundTime float64 // 带权周转时间
		curTime                int     // 当前时间
	)

	for _, p := range processes {
		waitTime := curTime - p.ArrivalTime // 等待时间 = 当前时间 - 到达时间
		curTime += p.ServiceTime
		turnaroundTime = float64(waitTime + p.ServiceTime)               // 周转时间 = 等待时间 + 服务时间
		weightedTurnaroundTime = turnaroundTime / float64(p.ServiceTime) // 带权周转时间 = 周转时间 / 服务时间

		fmt.Printf("Process %s: Completion Time: %d, Turnaround Time: %.2f, Weighted Turnaround Time: %.2f\n",
			p.ID, curTime, turnaroundTime, weightedTurnaroundTime)
	}
}

func rr(processes []Process, quantum int) {
	var (
		turnaroundTime         float64                // 周转时间
		weightedTurnaroundTime float64                // 带权周转时间
		curTime                int                    // 当前时间
		queue                  []Process              // 正在轮转的进程队列
		processQueue           []Process              // 待加入的进程队列
		serviceTime            = make(map[string]int) // 进程原始服务时间
	)

	for _, p := range processes {
		serviceTime[p.ID] = p.ServiceTime
	}
	processQueue = append(processQueue, processes...)
	// 对进程按到达时间进行排序
	sort.Slice(processQueue, func(i, j int) bool {
		return processQueue[i].ArrivalTime < processQueue[j].ArrivalTime
	})
	// 插入第一个进程
	queue = append(queue, processQueue[0])
	processQueue = processQueue[1:]

	// 开始模拟调度
	for len(queue) > 0 {
		// 记录当前进程是否完成
		var done bool
		// 取出队首进程
		p := queue[0]
		queue = queue[1:]
		// 当前时间小于进程到达时间，直接跳过
		if curTime < p.ArrivalTime {
			curTime = p.ArrivalTime
		}
		// 执行当前进程，直到完成或者时间片用完
		if p.ServiceTime <= quantum {
			// 进程完成
			curTime += p.ServiceTime
			turnaroundTime = float64(curTime - p.ArrivalTime)                    // 计算周转时间
			weightedTurnaroundTime = turnaroundTime / float64(serviceTime[p.ID]) // 计算带权周转时间
			fmt.Printf("Process %s: Completion Time: %d, Turnaround Time: %.2f, Weighted Turnaround Time: %.2f\n",
				p.ID, curTime, turnaroundTime, weightedTurnaroundTime)
			done = true
		} else {
			// 时间片用完
			curTime += quantum
			p.ServiceTime -= quantum // 减去已经用掉的时间
		}
		// 将到达时间在当前时间之前的所有未完成的进程加入队列
		var count int
		for _, p := range processQueue {
			if p.ArrivalTime <= curTime {
				queue = append(queue, p)
				count++
			}
		}
		processQueue = processQueue[count:]
		if !done {
			queue = append(queue, p) // 把进程重新放回队尾
		}
	}
}

func sjf(processes []Process) {
	var (
		turnaroundTime         float64   // 周转时间
		weightedTurnaroundTime float64   // 带权周转时间
		curTime                int       // 当前时间
		processQueue           []Process // 进程队列，存储所有未完成的进程
	)

	processQueue = append(processQueue, processes...) // 将所有进程添加到进程队列中
	sort.Slice(processQueue, func(i, j int) bool {
		return processQueue[i].ArrivalTime < processQueue[j].ArrivalTime
	})

	for len(processQueue) > 0 {
		// 选出到达时间小于等于当前时间的进程中，服务时间最短的进程
		minIndex := 0 // 记录服务时间最短进程在进程队列中的下标
		for i, p := range processQueue {
			if p.ArrivalTime <= curTime {
				if p.ServiceTime < processQueue[minIndex].ServiceTime {
					minIndex = i
				}
			} else {
				break // 如果有进程的到达时间比当前时间还晚，就跳出循环，等待下一个进程到达
			}
		}
		p := processQueue[minIndex]
		processQueue = append(processQueue[:minIndex], processQueue[minIndex+1:]...) // 从进程队列中删除服务时间最短的进程

		waitTime := curTime - p.ArrivalTime                              // 等待时间 = 当前时间 - 到达时间
		curTime += p.ServiceTime                                         // 更新当前时间
		turnaroundTime = float64(waitTime + p.ServiceTime)               // 周转时间 = 等待时间 + 服务时间
		weightedTurnaroundTime = turnaroundTime / float64(p.ServiceTime) // 带权周转时间 = 周转时间 / 服务时间

		fmt.Printf("Process %s: Completion Time: %d, Turnaround Time: %.2f, Weighted Turnaround Time: %.2f\n",
			p.ID, curTime, turnaroundTime, weightedTurnaroundTime)
	}
}

func hrn(processes []Process) {
	var (
		turnaroundTimes         []float64 // 周转时间
		weightedTurnaroundTimes []float64 // 带权周转时间
		curTime                 int       // 当前时间
		processQueue            []Process // 进程队列
	)
	processQueue = append(processQueue, processes...) // 将所有进程加入队列

	for len(processQueue) > 0 {
		// 计算每个进程的响应比
		responseRatios := make([]float64, len(processQueue))
		for i, p := range processQueue {
			waitTime := curTime - p.ArrivalTime                                          // 等待时间
			responseRatios[i] = float64(waitTime+p.ServiceTime) / float64(p.ServiceTime) // 计算响应比
		}

		// 找出响应比最大的进程
		maxIndex := 0
		for i, r := range responseRatios {
			if r > responseRatios[maxIndex] {
				maxIndex = i
			}
		}

		p := processQueue[maxIndex]                                                  // 找到响应比最大的进程
		processQueue = append(processQueue[:maxIndex], processQueue[maxIndex+1:]...) // 从队列中删除该进程

		waitTime := curTime - p.ArrivalTime                               // 该进程等待时间
		curTime += p.ServiceTime                                          // 当前时间加上服务时间，即该进程完成时间
		turnaroundTime := float64(waitTime + p.ServiceTime)               // 周转时间 = 等待时间 + 服务时间
		weightedTurnaroundTime := turnaroundTime / float64(p.ServiceTime) // 带权周转时间 = 周转时间 / 服务时间

		turnaroundTimes = append(turnaroundTimes, turnaroundTime)                         // 将该进程的周转时间加入周转时间列表
		weightedTurnaroundTimes = append(weightedTurnaroundTimes, weightedTurnaroundTime) // 将该进程的带权周转时间加入带权周转时间列表

		fmt.Printf("Process %s: Completion Time: %d, Turnaround Time: %.2f, Weighted Turnaround Time: %.2f\n",
			p.ID, curTime, turnaroundTime, weightedTurnaroundTime)
	}
}
