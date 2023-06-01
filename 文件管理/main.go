package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
)

/*
给出一个磁盘块序列:1、2、3、......、500，初始状态所有块为 空的，每块的大小为 2k。
选择使用 空闲表、空闲盘区链、位示图 三种算法之一来管理空闲块。对于基于块的索引分配执行以下步骤:
	1. 随机生成 2k-10k 的文件 50 个，文件名为 1.txt、2.txt、......、50.txt，按照上述算法存储到模拟磁盘中。
	2. 删除奇数.txt(1.txt、3.txt、......、49.txt)文件
	3. 新创建 5 个文件(A.txt、B.txt、C.txt、D.txt、E.txt)，大小为:7k、5k、2k、9k、3.5k，按照与(1)相同的算法存储到模拟磁盘中。
	4. 给出文件 A.txt、B.txt、C.txt、D.txt、E.txt 的盘块存储状态和所有空闲区块的状态。
*/

const (
	diskSize  = 500
	blockSize = 2 * 1024
)

type File struct {
	name   string
	size   int
	blocks []int
}

func main() {
	// 初始化位示图和文件列表
	bitMap := make([]bool, diskSize)
	files := make([]*File, 50)
	// 随机生成 50 个文件并写入模拟磁盘
	for i := 0; i < 50; i++ {
		fileSize := rand.Intn(8*1024) + 2*1024
		files[i] = createFile(fmt.Sprintf("%d.txt", i+1), fileSize, &bitMap)
	}
	// 删除奇数.txt文件
	for i := 0; i < len(files); i += 2 {
		deleteFile(files[i], &bitMap)
		files[i] = nil
	}
	// 新创建 5 个文件并写入模拟磁盘
	newFiles := []struct {
		name string
		size int
	}{
		{"A.txt", 7 * 1024},
		{"B.txt", 5 * 1024},
		{"C.txt", 2 * 1024},
		{"D.txt", 9 * 1024},
		{"E.txt", 3.5 * 1024},
	}
	for _, f := range newFiles {
		files = append(files, createFile(f.name, f.size, &bitMap))
	}
	// 输出文件盘块存储状态和空闲区块状态
	for _, file := range files {
		if file == nil {
			continue
		}
		fmt.Printf("文件名：%s，盘块：%v\n", file.name, file.blocks)
	}
	freeBlocks := getFreeBlocks(&bitMap)
	fmt.Printf("空闲区块：%v\n", freeBlocks)
}

// 创建文件
func createFile(name string, size int, bitMap *[]bool) *File {
	numBlocks := (size + blockSize - 1) / blockSize
	blocks := allocateBlocks(numBlocks, bitMap)
	return &File{name: name, size: size, blocks: blocks}
}

// 删除文件
func deleteFile(file *File, bitMap *[]bool) {
	for _, block := range file.blocks {
		(*bitMap)[block] = false
	}
}

// 分配盘块
func allocateBlocks(numBlocks int, bitMap *[]bool) []int {
	freeBlocks := getFreeBlocks(bitMap)
	if numBlocks > len(freeBlocks) {
		_, _ = fmt.Fprintf(os.Stderr, "没有足够的空闲区块，need(%d) freeBlocks(%d)", numBlocks, freeBlocks)
		os.Exit(1)
	}
	allocated := freeBlocks[:numBlocks]
	for _, block := range allocated {
		(*bitMap)[block] = true
	}
	return allocated
}

// 获取空闲区块
func getFreeBlocks(bitMap *[]bool) []int {
	var freeBlocks []int
	for i, used := range *bitMap {
		if !used {
			freeBlocks = append(freeBlocks, i)
		}
	}
	sort.Ints(freeBlocks)
	return freeBlocks
}
