package linked_list

import (
	"fmt"
	"strconv"
	"strings"
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
	strArr := strings.Split(arrayString, ",")
	intArr := make([]int, 0)
	for _, str := range strArr {
		n, _ := strconv.Atoi(str)
		intArr = append(intArr, n)
	}
	return Generate(intArr)
}

// Print 可视化打印链表
func (head *ListNode) Print() {
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
