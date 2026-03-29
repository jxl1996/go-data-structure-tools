package binary_tree

import (
	"fmt"
	"strconv"
	"strings"
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
