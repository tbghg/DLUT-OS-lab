package main

import (
	"fmt"
)

// Resource 资源列表
type Resource struct {
	A int // 资源A的数量
	B int // 资源B的数量
	C int // 资源C的数量
}

// Process 进程的资源统计
type Process struct {
	Max    Resource // 进程最大需要的资源
	Alloc  Resource // 进程已经分配到的资源
	Need   Resource // 进程还需要的资源
	Finish bool     // 进程是否已经完成
}

var (
	available Resource  = Resource{17, 5, 20} // 系统可用资源
	processes []Process = []Process{          // 进程列表
		{Resource{5, 5, 9}, Resource{2, 1, 2}, Resource{3, 4, 7}, false},
		{Resource{5, 3, 6}, Resource{4, 0, 2}, Resource{1, 3, 4}, false},
		{Resource{4, 0, 11}, Resource{4, 0, 5}, Resource{0, 0, 6}, false},
		{Resource{4, 2, 5}, Resource{2, 0, 4}, Resource{2, 2, 1}, false},
		{Resource{4, 2, 4}, Resource{3, 1, 4}, Resource{1, 1, 0}, false},
	}
)

func main() {
	if isSafe() {
		fmt.Println("T0时刻是安全状态")
		fmt.Println("安全序列为：", getSafeSeq())
	} else {
		fmt.Println("T0时刻不是安全状态")
	}

	fmt.Println("请输入进程P2请求的资源数：")
	var m, n, p int
	_, _ = fmt.Scan(&m, &n, &p)
	req := Resource{m, n, p}

	if canAllocate(1, req) {
		allocate(1, req)
		fmt.Println("分配成功")
	} else {
		fmt.Println("无法分配")
	}
}

// 判断当前状态是否安全
func isSafe() bool {
	work := available                      // 可用资源
	finish := make([]bool, len(processes)) // 每个进程是否已经完成

	for i := range finish {
		finish[i] = processes[i].Finish
	}

	for {
		found := false
		for i := range processes {
			if !finish[i] && processes[i].Need.A <= work.A && processes[i].Need.B <= work.B && processes[i].Need.C <= work.C {
				work.A += processes[i].Alloc.A // 分配资源
				work.B += processes[i].Alloc.B
				work.C += processes[i].Alloc.C
				finish[i] = true // 标记进程已完成
				found = true
			}
		}
		if !found {
			break
		}
	}

	for _, f := range finish { // 如果有进程未完成，则不安全
		if !f {
			return false
		}
	}
	return true
}

// 获取安全序列
func getSafeSeq() []int {
	work := available                      // 可用资源
	finish := make([]bool, len(processes)) // 每个进程是否已经完成
	seq := make([]int, len(processes))     // 安全序列
	count := 0

	for count < len(processes) {
		found := false
		for i := range processes {
			if !finish[i] && processes[i].Need.A <= work.A && processes[i].Need.B <= work.B && processes[i].Need.C <= work.C {
				work.A += processes[i].Alloc.A // 分配资源
				work.B += processes[i].Alloc.B
				work.C += processes[i].Alloc.C
				finish[i] = true   // 标记进程已完成
				seq[count] = i + 1 // 记录安全序列
				count++
				found = true
			}
		}
		if !found {
			break
		}
	}

	return seq
}

// 判断是否可以分配资源给某个进程
func canAllocate(pid int, req Resource) bool {
	return req.A <= processes[pid].Need.A && req.B <= processes[pid].Need.B && req.C <= processes[pid].Need.C &&
		req.A <= available.A && req.B <= available.B && req.C <= available.C
}

func allocate(pid int, req Resource) {
	processes[pid].Alloc.A += req.A
	processes[pid].Alloc.B += req.B
	processes[pid].Alloc.C += req.C
	processes[pid].Need.A -= req.A
	processes[pid].Need.B -= req.B
	processes[pid].Need.C -= req.C
	available.A -= req.A
	available.B -= req.B
	available.C -= req.C
}
