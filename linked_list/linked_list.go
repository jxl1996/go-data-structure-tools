package linked_list

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// ListNode 是 LeetCode 标准单链表节点定义
type ListNode struct {
	Val  int
	Next *ListNode
}

// Generate 根据整数切片快速构建链表，并返回头节点
func Generate(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	// 使用哨兵节点（Dummy Node）简化逻辑
	dummy := &ListNode{}
	current := dummy

	for _, val := range nums {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}

	return dummy.Next
}

// GenerateByArrayString 根据一个类数组的字符串生成链表(方便leetcode中复制直接生成)
// 形参：类似于 [1,4,3,2,5,2]
func GenerateByArrayString(arrayString string) *ListNode {
	arrayString = strings.Trim(arrayString, " ")
	arrayString = strings.TrimLeft(arrayString, "[")
	arrayString = strings.TrimRight(arrayString, "]")
	if len(arrayString) == 0 {
		return nil
	}
	strArr := strings.Split(arrayString, ",")
	intArr := make([]int, 0)
	for _, str := range strArr {
		n, _ := strconv.Atoi(str)
		intArr = append(intArr, n)
	}
	return Generate(intArr)
}

// Print 可视化打印链表
func (head *ListNode) Print(msg ...string) {
	if len(msg) > 0 {
		fmt.Print(msg[0], " ")
	}
	var nodes []string
	curr := head

	for curr != nil {
		nodes = append(nodes, strconv.Itoa(curr.Val))
		curr = curr.Next
	}

	if len(nodes) == 0 {
		fmt.Println("Empty List: nil")
		return
	}

	// 模拟 LeetCode 常见的展示风格: 1 -> 2 -> 3 -> NULL
	fmt.Printf("%s\n", strings.Join(nodes, " -> "))
}

// GenerateRandom 生成指定长度的链表
// size: 节点个数
// bounds: 可选参数。
//
//	如果不传，默认随机范围 [0, 100)
//	如果传 1 个，范围为 [0, bounds[0])
//	如果传 2 个，范围为 [bounds[0], bounds[1])
func GenerateRandom(size int, bounds ...int) *ListNode {
	if size <= 0 {
		return nil
	}

	// 默认值
	minVal, maxVal := 0, 100

	// 灵活处理参数
	if len(bounds) == 1 {
		// 如果只传一个，通常定义为 [0, maxVal)
		maxVal = bounds[0]
	} else if len(bounds) >= 2 {
		minVal = bounds[0]
		maxVal = bounds[1]
	}

	// 关键：如果范围非法，直接交换，确保逻辑稳健
	if minVal > maxVal {
		minVal, maxVal = maxVal, minVal
	}

	// 如果相等，Intn 会 panic，强制给一个差值
	if minVal == maxVal {
		maxVal = minVal + 1
	}

	// Go 1.20+ 建议使用这种方式，或者直接用 rand.Intn (现在已经自动加盐了)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	dummy := &ListNode{}
	curr := dummy
	for i := 0; i < size; i++ {
		// r.Intn(n) 返回 [0, n) 的数
		val := r.Intn(maxVal-minVal) + minVal
		curr.Next = &ListNode{Val: val}
		curr = curr.Next
	}

	return dummy.Next
}
