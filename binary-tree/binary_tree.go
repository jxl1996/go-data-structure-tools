package binary_tree

import (
	"fmt"
	"math"
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

	// 1. 计算每个子树所需的最小宽度
	widths := make(map[*TreeNode]int)
	var getWidth func(n *TreeNode) int
	getWidth = func(n *TreeNode) int {
		if n == nil {
			return 0
		}
		valStr := fmt.Sprintf("(%d)", n.Val)
		lw, rw := getWidth(n.Left), getWidth(n.Right)
		// 核心：宽度至少要能放下当前节点文字，且预留左右间距
		w := lw + rw + 2
		if len(valStr) > w {
			w = len(valStr)
		}
		widths[n] = w
		return w
	}
	getWidth(root)

	// 2. 第一次递归：初步计算坐标
	canvasNodes := []DrawNode{}
	var buildPos func(n *TreeNode, x, y int)
	buildPos = func(n *TreeNode, x, y int) {
		if n == nil {
			return
		}
		valStr := fmt.Sprintf("(%d)", n.Val)
		lw := 0
		if n.Left != nil {
			lw = widths[n.Left]
		}

		// 确保 mid 至少在左边之后
		mid := x + lw
		canvasNodes = append(canvasNodes, DrawNode{
			ValStr: valStr,
			X:      mid - len(valStr)/2,
			Mid:    mid,
			Y:      y,
			Node:   n,
		})

		if n.Left != nil {
			buildPos(n.Left, x, y+2)
		}
		if n.Right != nil {
			buildPos(n.Right, mid+1, y+2)
		}
	}
	buildPos(root, 0, 0)

	// 3. 关键修正：防止负数坐标并计算确切边界
	minX, maxX, maxY := 0, 0, 0
	for _, d := range canvasNodes {
		if d.X < minX {
			minX = d.X
		}
		if d.X+len(d.ValStr) > maxX {
			maxX = d.X + len(d.ValStr)
		}
		if d.Y > maxY {
			maxY = d.Y
		}
	}

	// 将所有坐标平移，确保从 0 开始且不越界
	offsetX := -minX
	totalWidth := maxX + offsetX + 5 // 额外给点余量

	canvas := make([][]rune, maxY+2)
	for i := range canvas {
		canvas[i] = []rune(strings.Repeat(" ", totalWidth))
	}

	// 4. 填充文字（带偏移量）
	for i := range canvasNodes {
		d := &canvasNodes[i]
		d.X += offsetX
		d.Mid += offsetX
		for j, r := range d.ValStr {
			canvas[d.Y][d.X+j] = r
		}
	}

	// 5. 绘制连线
	for _, d := range canvasNodes {
		if d.Node.Left == nil && d.Node.Right == nil {
			continue
		}

		// 画支点
		canvas[d.Y+1][d.Mid] = '┴'

		if d.Node.Left != nil {
			for _, cn := range canvasNodes {
				if cn.Node == d.Node.Left {
					canvas[d.Y+1][cn.Mid] = '╭'
					for i := cn.Mid + 1; i < d.Mid; i++ {
						canvas[d.Y+1][i] = '─'
					}
				}
			}
		}
		if d.Node.Right != nil {
			for _, cn := range canvasNodes {
				if cn.Node == d.Node.Right {
					canvas[d.Y+1][cn.Mid] = '╮'
					for i := d.Mid + 1; i < cn.Mid; i++ {
						canvas[d.Y+1][i] = '─'
					}
				}
			}
		}
	}

	// 6. 输出结果
	for _, row := range canvas {
		fmt.Println(strings.TrimRight(string(row), " "))
	}
}

// GetHeight 获取树高度
func (root *TreeNode) GetHeight() int {
	if root == nil {
		return 0
	}
	l := root.Left.GetHeight()
	r := root.Right.GetHeight()
	return max(l, r) + 1
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

// GenerateBST 生成一个包含 n 个节点的平衡二叉搜索树
// size: 节点个数
// bounds: 可选参数。
//
//	如果不传，默认随机范围 [0, 100)
//	如果传 1 个，范围为 [0, bounds[0])
//	如果传 2 个，范围为 [bounds[0], bounds[1])
func GenerateBST(n int, bounds ...int) *TreeNode {
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

// GenerateFullTree 根据指定高度生成满二叉树
// height: 树的高度（根节点为第 1 层）
// bounds: 可选参数。
//
//	如果不传，默认随机范围 [0, 100)
//	如果传 1 个，范围为 [0, bounds[0])
//	如果传 2 个，范围为 [bounds[0], bounds[1])
func GenerateFullTree(height int, bounds ...int) *TreeNode {
	if height <= 0 {
		return nil
	}

	// 1. 统一处理随机数范围逻辑
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

	// 2. 调用递归构建函数
	return buildFullRecursive(height, minVal, maxVal, r)
}

func buildFullRecursive(h int, min, max int, r *rand.Rand) *TreeNode {
	if h <= 0 {
		return nil
	}

	// 创建当前节点
	node := &TreeNode{
		Val: r.Intn(max-min) + min,
	}

	// 如果还没达到叶子层，则必须同时生成左右孩子
	if h > 1 {
		node.Left = buildFullRecursive(h-1, min, max, r)
		node.Right = buildFullRecursive(h-1, min, max, r)
	}

	return node
}

// IsFullTree 判断二叉树是否为满二叉树
func (root *TreeNode) IsFullTree() bool {
	if root == nil {
		return true
	}
	// 如果是叶子节点，返回 true
	if root.Left == nil && root.Right == nil {
		return true
	}
	// 如果左右孩子都存在，递归检查左右子树
	if root.Left != nil && root.Right != nil {
		return root.Left.IsFullTree() && root.Right.IsFullTree()
	}
	// 只有一个孩子，肯定不是满二叉树
	return false
}

// IsCompleteTree 判断当前树是否为完全二叉树
func (root *TreeNode) IsCompleteTree() bool {
	if root == nil {
		return true
	}

	// 使用切片模拟队列进行层序遍历
	queue := []*TreeNode{root}
	// 标记位：一旦发现某个节点为 nil，后续所有节点都不能有值
	foundNil := false

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr == nil {
			// 第一次遇到空节点，标记后续不能再有节点
			foundNil = true
		} else {
			// 如果之前已经遇到过空节点，但现在又遇到了非空节点
			// 说明结构不连续，不是完全二叉树
			if foundNil {
				return false
			}
			// 将左右孩子入队（包括 nil，用于占位判断）
			queue = append(queue, curr.Left)
			queue = append(queue, curr.Right)
		}
	}

	return true
}

// IsBST 判断当前树是否为二叉搜索树
func (root *TreeNode) IsBST() bool {
	return validateBST(root, math.MinInt64, math.MaxInt64)
}

func validateBST(node *TreeNode, min, max int64) bool {
	if node == nil {
		return true
	}

	// 当前节点值必须在 (min, max) 范围内
	val := int64(node.Val)
	if val <= min || val >= max {
		return false
	}

	// 递归检查左子树：更新上限为当前节点值
	// 递归检查右子树：更新下限为当前节点值
	return validateBST(node.Left, min, val) && validateBST(node.Right, val, max)
}

// IsBalanced 判断当前二叉树是否为平衡二叉树
func (root *TreeNode) IsBalanced() bool {
	return checkHeight(root) != -1
}

// checkHeight 递归函数：
// 如果平衡，返回该子树的高度；
// 如果不平衡，返回 -1。
func checkHeight(node *TreeNode) float64 {
	if node == nil {
		return 0
	}

	// 1. 递归检查左子树
	leftHeight := checkHeight(node.Left)
	if leftHeight == -1 {
		return -1
	}

	// 2. 递归检查右子树
	rightHeight := checkHeight(node.Right)
	if rightHeight == -1 {
		return -1
	}

	// 3. 核心逻辑：判断当前节点是否平衡
	// 如果左右子树高度差 > 1，则不平衡
	if math.Abs(leftHeight-rightHeight) > 1 {
		return -1
	}

	// 4. 如果平衡，返回当前节点的高度 (左右子树最大高度 + 1)
	return math.Max(leftHeight, rightHeight) + 1
}

// GetBalanceFactor 获取平衡因子
func (root *TreeNode) GetBalanceFactor() int {
	if root == nil {
		return 0
	}
	// 左子树高度 - 右子树高度
	return root.Left.GetHeight() - root.Right.GetHeight()
}

// GetImbalancedNodes 找出所有平衡因子绝对值 >= 2 的节点
func (root *TreeNode) GetImbalancedNodes() map[*TreeNode]bool {
	imbalanced := make(map[*TreeNode]bool)

	var check func(*TreeNode) int
	check = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		lh := check(node.Left)
		rh := check(node.Right)

		// 计算并检查平衡因子
		if math.Abs(float64(lh-rh)) >= 2 {
			imbalanced[node] = true
		}

		// 返回当前节点高度
		return int(math.Max(float64(lh), float64(rh))) + 1
	}

	check(root)
	return imbalanced
}
