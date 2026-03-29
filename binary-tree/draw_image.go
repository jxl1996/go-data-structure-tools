package binary_tree

import (
	"fmt"
	"github.com/fogleman/gg"
	"image/color"
	"math"
	"strconv"
)

const (
	nodeRadius = 25.0 // 节点半径
	levelSpace = 90.0 // 层间距
	minSpace   = 50.0 // 同层节点最小间距
)

// Draw 完美直线美化版：保存为图片（双遍历修复）
func (root *TreeNode) Draw(opts ...string) {
	if root == nil {
		fmt.Println("空树，跳过绘制")
		return
	}

	var filename = "tree.png"
	if len(opts) > 0 {
		filename = opts[0]
	}

	// 1. 动态计算画布尺寸
	depth := getDepth(root)
	// 宽度随深度指数增长，防止深层挤压
	imgWidth := int(math.Pow(2, float64(depth-1))*minSpace + 200)
	imgHeight := int(float64(depth)*levelSpace + 120)

	dc := gg.NewContext(imgWidth, imgHeight)

	// 背景：纯白
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	// 字体设置（请根据系统环境修改路径，或注销使用默认字体）
	if err := dc.LoadFontFace("/Library/Fonts/Arial.ttf", 18); err != nil {
		// 找不到字体时会回退到系统默认
	}

	// 初始位置
	initialX := float64(imgWidth) / 2
	initialY := 60.0
	initialOffsetX := float64(imgWidth) / 4

	// 2. 第一次遍历：绘制所有连线
	drawAllLinesStraight(dc, root, initialX, initialY, initialOffsetX)

	// 3. 第二次遍历：绘制所有节点
	drawAllNodesStraight(dc, root, initialX, initialY, initialOffsetX)

	// 4. 保存结果
	if err := dc.SavePNG(filename); err != nil {
		fmt.Println("❌ 保存失败:", err)
	} else {
		fmt.Println("✨ 完美修复版直线图片已生成:", filename)
	}
}

// 第一次遍历：绘制所有连线 (直线)
func drawAllLinesStraight(dc *gg.Context, node *TreeNode, x, y, offsetX float64) {
	if node == nil {
		return
	}

	// 连线样式
	dc.SetRGBA(0.3, 0.3, 0.3, 0.8) // 深灰色实线
	dc.SetLineWidth(2.0)

	// 绘制到左子节点的连线
	if node.Left != nil {
		lx, ly := x-offsetX, y+levelSpace
		dc.DrawLine(x, y, lx, ly)
		dc.Stroke()
		// 递归绘制左子树连线
		drawAllLinesStraight(dc, node.Left, lx, ly, offsetX/2)
	}

	// 绘制到右子节点的连线
	if node.Right != nil {
		rx, ry := x+offsetX, y+levelSpace
		dc.DrawLine(x, y, rx, ry)
		dc.Stroke()
		// 递归绘制右子树连线
		drawAllLinesStraight(dc, node.Right, rx, ry, offsetX/2)
	}
}

// 第二次遍历：绘制所有节点
func drawAllNodesStraight(dc *gg.Context, node *TreeNode, x, y, offsetX float64) {
	if node == nil {
		return
	}

	// A. 先绘制节点阴影 (增加立体感)
	dc.SetRGBA(0, 0, 0, 0.1)
	dc.DrawCircle(x+1, y+1, nodeRadius)
	dc.Fill()

	// B. 绘制节点圆柱色渐变
	grad := gg.NewLinearGradient(x, y-nodeRadius, x, y+nodeRadius)
	grad.AddColorStop(0, color.RGBA{R: 80, G: 160, B: 240, A: 255}) // 亮蓝
	grad.AddColorStop(1, color.RGBA{R: 40, G: 100, B: 180, A: 255}) // 深蓝

	dc.SetFillStyle(grad)
	dc.DrawCircle(x, y, nodeRadius)
	dc.Fill()

	// C. 节点外圈
	dc.SetRGB(0.2, 0.5, 0.8)
	dc.SetLineWidth(2)
	dc.DrawCircle(x, y, nodeRadius)
	dc.Stroke()

	// D. 文字 (白色)
	dc.SetRGB(1, 1, 1)
	valStr := strconv.Itoa(node.Val)
	dc.DrawStringAnchored(valStr, x, y, 0.5, 0.5)

	// 递归绘制左子树节点
	if node.Left != nil {
		drawAllNodesStraight(dc, node.Left, x-offsetX, y+levelSpace, offsetX/2)
	}
	// 递归绘制右子树节点
	if node.Right != nil {
		drawAllNodesStraight(dc, node.Right, x+offsetX, y+levelSpace, offsetX/2)
	}
}

// getDepth 辅助函数：获取树深度
func getDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	l := getDepth(node.Left)
	r := getDepth(node.Right)
	if l > r {
		return l + 1
	}
	return r + 1
}
