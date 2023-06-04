package main

import (
	"fmt"
)

func main() {
	// 读取用户输入的页面引用序列和页帧数
	var n, frames int
	fmt.Println("请输入页面引用序列长度：")
	// 例如 refs (20) 7, 0, 1, 2, 0, 3, 0, 4, 2, 3, 0, 3, 2, 1, 2, 0, 1, 7, 0, 1
	fmt.Scan(&n)
	refs := make([]int, n)
	fmt.Println("请输入页面引用序列：")
	for i := 0; i < n; i++ {
		fmt.Scan(&refs[i])
	}
	fmt.Println("请输入分配给进程的物理页面数目：")
	// 例如 frames 3
	fmt.Scan(&frames)

	// 根据选择的算法，定义相应的缺页中断替换函数
	mp := map[string]func([]int, []int, int) int{
		"OPTIMAL": optimalReplacement,
		"LRU":     lruReplacement,
		"FIFO":    fifoReplacement,
	}
	// 定义算法名称列表
	names := []string{"OPTIMAL", "LRU", "FIFO"}
	// 遍历算法名称列表，分别使用不同算法进行页面替换操作，并输出替换结果
	for _, name := range names {
		// 调用 pageReplace 函数进行页面替换操作，并返回缺页中断次数和被淘汰的页号列表
		faults, eliminated := pageReplace(refs, frames, mp[name])
		// 打印输出替换结果
		fmt.Printf("%s:\n", name)
		fmt.Printf("  淘汰的页号: %v\n", eliminated)
		fmt.Printf("  缺页次数: %d\n", faults)
		fmt.Printf("  缺页率: %.2f%%\n", float64(faults)/float64(len(refs))*100)
	}
}

// pageReplace 函数，使用指定算法进行页面替换操作，并返回缺页中断次数和被淘汰的页号列表
func pageReplace(refs []int, frames int, algorithm func([]int, []int, int) int) (int, []int) {
	var (
		faults     int                      // 缺页中断次数
		eliminated []int                    // 被淘汰的页号列表
		memory     = make([]int, 0, frames) // 内存页帧列表
	)
	// 遍历页面引用序列
	for i, ref := range refs {
		// 如果内存页帧列表中已经包含该页，则跳过本次遍历
		if contains(memory, ref) {
			continue
		}
		// 如果内存页帧列表未满，则添加该页
		if len(memory) < frames {
			memory = append(memory, ref)
		} else {
			// 如果内存页帧列表已满，则使用指定算法进行页面替换，并记录被淘汰的页号
			replaceIdx := algorithm(memory, refs, i)
			eliminated = append(eliminated, memory[replaceIdx])
			memory[replaceIdx] = ref
		}
		faults++
	}
	return faults, eliminated
}

// contains 函数，判断列表中是否包含指定值
func contains(arr []int, value int) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}

// optimalReplacement 函数，使用 OPTIMAL 算法进行页面替换
func optimalReplacement(memory []int, refs []int, idx int) int {
	// 只保留当前位置及以后的页面引用序列
	refs = refs[idx:]
	maxIdx := 0       // 最远以后不使用的页号在内存页帧列表中的位置
	maxDistance := -1 // 最远以后不使用的页号距离当前位置的距离
	// 遍历内存页帧列表中的每个页号
	for i, page := range memory {
		// 查找该页号在当前位置及以后的页面引用序列中最近的距离
		distance := findBackDistance(refs, page)
		if distance == -1 {
			// 如果在当前位置及以后的页面引用序列中不存在该页号，则直接返回该页号在内存页帧列表中的位置
			return i
		}
		if distance > maxDistance {
			// 如果找到了一个距离更远的页号，则更新最远以后不使用的页号在内存页帧列表中的位置和距离
			maxIdx = i
			maxDistance = distance
		}
	}
	return maxIdx
}

// findBackDistance 函数，查找指定页号在当前位置及以后的页面引用序列中最近的距离
func findBackDistance(refs []int, value int) int {
	for i, v := range refs {
		if v == value {
			return i
		}
	}
	return -1
}

// findFrontDistance 函数，查找指定页号在当前位置及以前的页面引用序列中最近的距离
func findFrontDistance(refs []int, value int) int {
	for i := len(refs) - 1; i > 0; i-- {
		if refs[i] == value {
			return i
		}
	}
	return -1
}

// lruReplacement 函数，使用 LRU 算法进行页面替换
func lruReplacement(memory []int, refs []int, idx int) int {
	// 只保留当前位置及以前的页面引用序列
	refs = refs[:idx]
	maxIdx := 0       // 最久未使用的页号在内存页帧列表中的位置
	maxDistance := -1 // 最久未使用的页号距离当前位置的距离
	// 遍历内存页帧列表中的每个页号
	for i, page := range memory {
		// 查找该页号在当前位置及以前的页面引用序列中最久未使用的距离
		distance := findFrontDistance(refs, page)
		if distance == -1 {
			// 如果在当前位置及以前的页面引用序列中不存在该页号，则直接返回该页号在内存页帧列表中的位置
			return i
		}
		if distance > maxDistance {
			// 如果找到了一个距离更久远的页号，则更新最久未使用的页号在内存页帧列表中的位置和距离
			maxIdx = i
			maxDistance = distance
		}
	}
	return maxIdx
}

// fifoReplacement 函数，使用 FIFO 算法进行页面替换
func fifoReplacement(memory []int, refs []int, idx int) int {
	return 0 // 返回内存页帧列表中最早添加的页号的位置
}
