package main

/*
	1. 模拟两种磁盘移臂调度算法：sstf 算法和 scan 算法
	2. 能对两种算法给定任意序列不同的磁盘请求序列，显示响应磁盘请求的过程。
	3. 能统计和报告不同算法情况下响应请求的顺序、移臂的总量。
*/

import (
	"fmt"
	"math"
	"sort"
)

// DiskRequest 磁盘请求表示访问某磁盘扇区的请求
type DiskRequest struct {
	sector int // 要访问的扇区
}

// sstf 磁盘调度算法
func sstf(requests []DiskRequest, currentSector int) {
	var (
		totalMove int // 总寻道时间
		rs        []DiskRequest
	)
	rs = append(rs, requests...)
	// 处理所有请求
	for len(rs) > 0 {
		// 找到距离当前位置最近的请求
		closest := -1
		closestDistance := int(1e9)
		for i, r := range rs {
			distance := int(math.Abs(float64(r.sector - currentSector)))
			if distance < closestDistance {
				closest = i
				closestDistance = distance
			}
		}

		// 服务最接近的请求
		fmt.Println("  处理扇区", rs[closest].sector, "的请求")
		currentSector = rs[closest].sector
		totalMove += closestDistance
		rs = append(rs[:closest], rs[closest+1:]...)
	}

	fmt.Println("sstf移臂总量:", totalMove)
}

// scan 磁盘调度算法
func scan(requests []DiskRequest, currentSector, maxRound int) {
	var (
		totalMove int // 总寻道时间
		rs        []DiskRequest
	)
	rs = append(rs, requests...)
	sort.Slice(rs, func(i, j int) bool {
		return rs[i].sector < rs[j].sector
	})

	// 根据方向查找最近的请求
	var closest int
	closestDistance := int(1e9)
	for i, r := range rs {
		distance := int(math.Abs(float64(r.sector - currentSector)))
		if distance < closestDistance {
			closest = i
			closestDistance = distance
		}
	}

	if rs[closest].sector-currentSector > 0 {
		// 向右移动，扫描所有大于currentSector的requests，走到右侧顶端
		for i := closest; i < len(rs); i++ {
			fmt.Println("  处理扇区", rs[i].sector, "的请求")
		}
		totalMove += rs[len(rs)-1].sector - currentSector
		if closest > 0 {
			totalMove += maxRound - rs[len(rs)-1].sector
			for i := closest - 1; i >= 0; i-- {
				fmt.Println("  处理扇区", rs[i].sector, "的请求")
			}
		}
		fmt.Println("scan移臂总量:", totalMove)
		return
	}
	for i := closest; i >= 0; i-- {
		fmt.Println("  处理扇区", rs[i].sector, "的请求")
	}
	totalMove += currentSector - rs[0].sector
	if closest < len(rs)-1 {
		totalMove += rs[0].sector
		for i := closest + 1; i < len(rs); i++ {
			fmt.Println("  处理扇区", rs[i].sector, "的请求")
		}
	}
	fmt.Println("scan移臂总量:", totalMove)
}

/*
	requests := []DiskRequest{
		{10}, {7}, {30}, {20},
		{15}, {25}, {8}, {31},
	}
	currentSector := 23 // 目前指针位置
	maxRound := 49      // 磁盘最大磁道
*/

func main() {
	// 读取用户输入的磁盘请求列表、当前指针位置和最大磁道数
	var n, currentSector, maxRound int
	fmt.Println("请输入磁盘请求列表长度：")
	fmt.Scan(&n)
	requests := make([]DiskRequest, n)
	fmt.Println("请输入磁盘请求列表：")
	for i := 0; i < n; i++ {
		var sector int
		fmt.Printf("请输入第 %d 个请求的磁道号：", i+1)
		fmt.Scan(&sector)
		requests[i] = DiskRequest{sector}
	}
	fmt.Println("请输入当前指针位置：")
	fmt.Scan(&currentSector)
	fmt.Println("请输入最大磁道数：")
	fmt.Scan(&maxRound)

	fmt.Println("sstf:")
	sstf(requests, currentSector)

	fmt.Println("\nscan:")
	scan(requests, currentSector, maxRound)
}
