package binary_tree

import (
	"fmt"
	"github.com/fogleman/gg"
	"math"
	"strconv"
)

const (
	nodeRadius = 25.0 // 节点半径
	levelSpace = 80.0 // 层间距
	minSpace   = 40.0 // 同层节点最小间距
)

// Draw 保存为图片
// 形参： 可传入图片名称， 可不传，不传默认为binary_tree.png
func (root *TreeNode) Draw(opts ...string) {
	var filename = "binary_tree.png"
	if len(opts) > 0 {
		filename = opts[0]
	}

	// 动态计算图片尺寸
	depth := getDepth(root)
	totalWidth := getSubtreeWidth(root) * 2
	imgWidth := int(totalWidth + 100)
	imgHeight := int(float64(depth)*levelSpace + 100)

	dc := gg.NewContext(imgWidth, imgHeight)
	dc.SetRGB(1, 1, 1) // 背景涂白
	dc.Clear()

	// 载入内置字体（或指定路径字体）
	if err := dc.LoadFontFace("/Library/Fonts/Arial.ttf", 18); err != nil {
		// 如果没字体，会使用默认简单字体
	}

	drawNode(dc, root, float64(imgWidth)/2, 50, totalWidth/2)

	err := dc.SavePNG(filename)
	if err != nil {
		fmt.Println("保存失败:", err)
	} else {
		fmt.Println("图片已生成:", filename)
	}
}

// 计算子树需要的总宽度
func getSubtreeWidth(node *TreeNode) float64 {
	if node == nil {
		return 0
	}
	lw := getSubtreeWidth(node.Left)
	rw := getSubtreeWidth(node.Right)
	// 宽度由左右子树加上间距决定，叶子节点给一个基础宽度
	width := lw + rw + minSpace
	return math.Max(width, nodeRadius*2.5)
}

// 递归绘制
func drawNode(dc *gg.Context, node *TreeNode, x, y, width float64) {
	if node == nil {
		return
	}

	// 1. 绘制到子节点的连线
	dc.SetRGB(0.5, 0.5, 0.5) // 灰色线条
	dc.SetLineWidth(2)

	if node.Left != nil {
		lx := x - width/4
		ly := y + levelSpace
		dc.DrawLine(x, y, lx, ly)
		dc.Stroke()
		drawNode(dc, node.Left, lx, ly, width/2)
	}
	if node.Right != nil {
		rx := x + width/4
		ry := y + levelSpace
		dc.DrawLine(x, y, rx, ry)
		dc.Stroke()
		drawNode(dc, node.Right, rx, ry, width/2)
	}

	// 2. 绘制节点圆形
	dc.SetRGB(1, 1, 1) // 白色背景
	dc.DrawCircle(x, y, nodeRadius)
	dc.FillPreserve()
	dc.SetRGB(0.2, 0.6, 0.9) // 蓝色边框
	dc.SetLineWidth(3)
	dc.Stroke()

	// 3. 绘制文字
	dc.SetRGB(0, 0, 0) // 黑色文字
	valStr := strconv.Itoa(node.Val)
	dc.DrawStringAnchored(valStr, x, y, 0.5, 0.5)
}

func getDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return 1 + int(math.Max(float64(getDepth(node.Left)), float64(getDepth(node.Right))))
}
