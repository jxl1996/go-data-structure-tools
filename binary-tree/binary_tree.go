package binary_tree

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Generate  生成一棵二叉树
// 形参：类似于 []string{"3", "9", "20", "null", "null", "15", "7"}
func Generate(nodes []string) *TreeNode {
	if len(nodes) == 0 || nodes[0] == "null" {
		return nil
	}
	v, _ := strconv.Atoi(nodes[0])
	root := &TreeNode{Val: v}
	q := []*TreeNode{root}
	i := 1
	for len(q) > 0 && i < len(nodes) {
		curr := q[0]
		q = q[1:]
		if i < len(nodes) && nodes[i] != "null" {
			val, _ := strconv.Atoi(nodes[i])
			curr.Left = &TreeNode{Val: val}
			q = append(q, curr.Left)
		}
		i++
		if i < len(nodes) && nodes[i] != "null" {
			val, _ := strconv.Atoi(nodes[i])
			curr.Right = &TreeNode{Val: val}
			q = append(q, curr.Right)
		}
		i++
	}
	return root
}

// GenerateByArrayString 根据一个类数组的字符串生成二叉树(方便leetcode中复制直接生成)
// 形参：类似于 [1,2,2,null,3,null,3]
func GenerateByArrayString(arrayString string) *TreeNode {
	arrayString = strings.Trim(arrayString, " ")
	arrayString = strings.TrimLeft(arrayString, "[")
	arrayString = strings.TrimRight(arrayString, "]")
	arrayString = strings.ToLower(arrayString)
	if len(arrayString) == 0 {
		return nil
	}
	strArr := strings.Split(arrayString, ",")
	return Generate(strArr)
}

// DrawNode 增强版 DrawNode，记录原始指针以便判断叶子状态
type DrawNode struct {
	ValStr    string
	X, Mid, Y int
	Node      *TreeNode // 增加原始指针引用
}

func (root *TreeNode) Print() {
	if root == nil {
		fmt.Println("Empty Tree")
		return
	}

	widths := make(map[*TreeNode]int)
	var getWidth func(n *TreeNode) int
	getWidth = func(n *TreeNode) int {
		if n == nil {
			return 0
		}
		vLen := len(fmt.Sprintf("(%d)", n.Val))
		lw, rw := getWidth(n.Left), getWidth(n.Right)
		w := lw + rw + 4
		if vLen > w {
			w = vLen
		}
		widths[n] = w
		return w
	}
	getWidth(root)

	canvasNodes := []DrawNode{}
	var buildPos func(n *TreeNode, x, y int)
	buildPos = func(n *TreeNode, x, y int) {
		if n == nil {
			return
		}
		valStr := fmt.Sprintf("(%d)", n.Val)
		vLen := len(valStr)
		lw := 0
		if n.Left != nil {
			lw = widths[n.Left]
		}

		mid := x + lw + 2
		if n.Left == nil {
			mid = x + 2
		}

		canvasNodes = append(canvasNodes, DrawNode{
			ValStr: valStr,
			X:      mid - vLen/2,
			Mid:    mid,
			Y:      y,
			Node:   n, // 存入引用
		})

		if n.Left != nil {
			buildPos(n.Left, x, y+2)
		}
		if n.Right != nil {
			buildPos(n.Right, mid+2, y+2)
		}
	}
	buildPos(root, 0, 0)

	maxY, maxX := 0, 0
	for _, n := range canvasNodes {
		if n.Y > maxY {
			maxY = n.Y
		}
		if n.X+len(n.ValStr) > maxX {
			maxX = n.X + len(n.ValStr)
		}
	}

	canvas := make([][]rune, maxY+2)
	for i := range canvas {
		canvas[i] = []rune(strings.Repeat(" ", maxX+10))
	}

	// 绘制文字
	for _, d := range canvasNodes {
		for i, r := range d.ValStr {
			canvas[d.Y][d.X+i] = r
		}
	}

	// 核心修正：绘制支点和连线
	var linkLines func(n *TreeNode, d DrawNode)
	linkLines = func(n *TreeNode, d DrawNode) {
		if n == nil {
			return
		}

		// --- 关键判断：只有不是叶子节点才画支点 ---
		if n.Left != nil || n.Right != nil {
			canvas[d.Y+1][d.Mid] = '┴'
		}

		if n.Left != nil {
			for _, cn := range canvasNodes {
				if cn.Y == d.Y+2 && cn.Mid < d.Mid && cn.Node == n.Left {
					canvas[d.Y+1][cn.Mid] = '╭'
					for i := cn.Mid + 1; i < d.Mid; i++ {
						canvas[d.Y+1][i] = '─'
					}
					linkLines(n.Left, cn)
					break
				}
			}
		}
		if n.Right != nil {
			for _, cn := range canvasNodes {
				if cn.Y == d.Y+2 && cn.Mid > d.Mid && cn.Node == n.Right {
					canvas[d.Y+1][cn.Mid] = '╮'
					for i := d.Mid + 1; i < cn.Mid; i++ {
						canvas[d.Y+1][i] = '─'
					}
					linkLines(n.Right, cn)
					break
				}
			}
		}
	}

	for _, d := range canvasNodes {
		if d.Y == 0 {
			linkLines(root, d)
			break
		}
	}

	for _, row := range canvas {
		line := strings.TrimRight(string(row), " ")
		if line != "" {
			fmt.Println(line)
		}
	}
	fmt.Println()
}

// GenerateRandom 生成一个包含 n 个节点的随机二叉树
// size: 节点个数
// bounds: 可选参数。
//
//	如果不传，默认随机范围 [0, 100)
//	如果传 1 个，范围为 [0, bounds[0])
//	如果传 2 个，范围为 [bounds[0], bounds[1])
func GenerateRandom(n int, bounds ...int) *TreeNode {
	if n <= 0 {
		return nil
	}

	// 1. 处理随机范围
	minVal, maxVal := 0, 100
	if len(bounds) == 1 {
		maxVal = bounds[0]
	} else if len(bounds) >= 2 {
		minVal, maxVal = bounds[0], bounds[1]
	}
	if minVal >= maxVal {
		minVal, maxVal = maxVal, minVal+1
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// 2. 调用内部递归函数
	return generateRecursive(n, minVal, maxVal, r)
}

// GenerateCompleteTree 生成一个包含 n 个节点的完全二叉树
// size: 节点个数
// bounds: 可选参数。
//
//	如果不传，默认随机范围 [0, 100)
//	如果传 1 个，范围为 [0, bounds[0])
//	如果传 2 个，范围为 [bounds[0], bounds[1])
func GenerateCompleteTree(n int, bounds ...int) *TreeNode {
	if n <= 0 {
		return nil
	}

	// 1. 处理随机数范围逻辑 (保持与你链表方法一致的健壮性)
	minVal, maxVal := 0, 100
	if len(bounds) == 1 {
		maxVal = bounds[0]
	} else if len(bounds) >= 2 {
		minVal, maxVal = bounds[0], bounds[1]
	}

	// 确保范围合法
	if minVal > maxVal {
		minVal, maxVal = maxVal, minVal
	}
	if minVal == maxVal {
		maxVal = minVal + 1
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// 2. 预先创建所有节点并存入切片
	// 这样我们可以通过索引直接访问任何一个节点
	nodes := make([]*TreeNode, n)
	for i := 0; i < n; i++ {
		nodes[i] = &TreeNode{
			Val: r.Intn(maxVal-minVal) + minVal,
		}
	}

	// 3. 建立父子连接关系
	// 根据完全二叉树性质：
	// 对于索引为 i 的节点：
	// - 左孩子索引 = 2*i + 1
	// - 右孩子索引 = 2*i + 2
	for i := 0; i < n/2; i++ {
		leftIndex := 2*i + 1
		rightIndex := 2*i + 2

		if leftIndex < n {
			nodes[i].Left = nodes[leftIndex]
		}
		if rightIndex < n {
			nodes[i].Right = nodes[rightIndex]
		}
	}

	// 返回根节点
	return nodes[0]
}

func generateRecursive(n int, min, max int, r *rand.Rand) *TreeNode {
	if n <= 0 {
		return nil
	}

	// 创建当前节点
	root := &TreeNode{
		Val: r.Intn(max-min) + min,
	}

	// 将剩余的 n-1 个节点随机分配给左右子树
	leftCount := r.Intn(n) // 随机生成 0 到 n-1 之间的数
	rightCount := n - 1 - leftCount

	root.Left = generateRecursive(leftCount, min, max, r)
	root.Right = generateRecursive(rightCount, min, max, r)

	return root
}

// GenerateRandomBST 生成一个包含 n 个节点的平衡二叉搜索树
// size: 节点个数
// bounds: 可选参数。
//
//	如果不传，默认随机范围 [0, 100)
//	如果传 1 个，范围为 [0, bounds[0])
//	如果传 2 个，范围为 [bounds[0], bounds[1])
func GenerateRandomBST(n int, bounds ...int) *TreeNode {
	if n <= 0 {
		return nil
	}

	// 1. 处理随机范围
	minVal, maxVal := 0, 100
	if len(bounds) == 1 {
		maxVal = bounds[0]
	} else if len(bounds) >= 2 {
		minVal, maxVal = bounds[0], bounds[1]
	}

	// 确保范围能够容纳 n 个不重复的数
	if maxVal-minVal < n {
		maxVal = minVal + n
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// 2. 生成 n 个不重复的随机数
	set := make(map[int]struct{})
	for len(set) < n {
		set[r.Intn(maxVal-minVal)+minVal] = struct{}{}
	}

	// 3. 转化为有序切片
	nums := make([]int, 0, n)
	for k := range set {
		nums = append(nums, k)
	}
	sort.Ints(nums)

	// 4. 递归构建平衡 BST
	return buildBST(nums, 0, len(nums)-1)
}

// buildBST 核心算法：每次取中间值作为根，确保左右子树高度差不超过 1
func buildBST(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}

	mid := left + (right-left)/2
	root := &TreeNode{Val: nums[mid]}

	root.Left = buildBST(nums, left, mid-1)
	root.Right = buildBST(nums, mid+1, right)

	return root
}
