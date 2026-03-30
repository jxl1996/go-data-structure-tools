package binary_tree

import (
	"fmt"
	"github.com/fogleman/gg"
	"image/color"
	"math"
	"strconv"
)

// NodeStyle 定义节点视觉类型
type NodeStyle int

const (
	StyleDefault NodeStyle = iota // 默认：经典蓝
	StyleWarning                  // 警告：橙红色（如 AVL 失衡）
	StyleSpecial                  // 特殊：清新绿（如新插入或搜索路径）
)

// Draw 绘制二叉树
// filename: 保存的文件名
// styleFunc: 可选。传入一个函数，根据节点属性返回对应的 Style
func (root *TreeNode) Draw(filename string, styleFunc func(*TreeNode) NodeStyle) {
	if root == nil {
		fmt.Println("空树，跳过绘制")
		return
	}

	if filename == "" {
		filename = "tree.png"
	}

	depth := root.GetHeight()
	// 动态计算尺寸，确保大树不挤压
	imgWidth := int(math.Pow(2, float64(depth-1))*50.0 + 200)
	imgHeight := int(float64(depth)*90.0 + 120)

	dc := gg.NewContext(imgWidth, imgHeight)
	dc.SetRGB(1, 1, 1) // 纯白背景
	dc.Clear()

	// 初始位置参数
	initialX := float64(imgWidth) / 2
	initialY := 60.0
	initialOffsetX := float64(imgWidth) / 4

	// 1. 第一次遍历：绘制所有连线（确保线条在节点下方）
	drawLines(dc, root, initialX, initialY, initialOffsetX)

	// 2. 第二次遍历：根据 styleFunc 绘制所有节点
	drawNodes(dc, root, initialX, initialY, initialOffsetX, styleFunc)

	if err := dc.SavePNG(filename); err != nil {
		fmt.Println("❌ 保存失败:", err)
	} else {
		fmt.Println("✨ 图片已生成:", filename)
	}
}

// 绘制连线
func drawLines(dc *gg.Context, node *TreeNode, x, y, offsetX float64) {
	if node == nil {
		return
	}
	dc.SetRGBA(0.3, 0.3, 0.3, 0.8)
	dc.SetLineWidth(2.0)

	if node.Left != nil {
		lx, ly := x-offsetX, y+90.0
		dc.DrawLine(x, y, lx, ly)
		dc.Stroke()
		drawLines(dc, node.Left, lx, ly, offsetX/2)
	}
	if node.Right != nil {
		rx, ry := x+offsetX, y+90.0
		dc.DrawLine(x, y, rx, ry)
		dc.Stroke()
		drawLines(dc, node.Right, rx, ry, offsetX/2)
	}
}

// 绘制节点（包含逻辑变色）
func drawNodes(dc *gg.Context, node *TreeNode, x, y, offsetX float64, styleFunc func(*TreeNode) NodeStyle) {
	if node == nil {
		return
	}

	// --- A. 确定配色 ---
	var startCol, endCol, strokeCol color.RGBA
	style := StyleDefault
	if styleFunc != nil {
		style = styleFunc(node)
	}

	switch style {
	case StyleWarning: // 橙红
		startCol = color.RGBA{255, 120, 100, 255}
		endCol = color.RGBA{200, 50, 40, 255}
		strokeCol = color.RGBA{180, 40, 30, 255}
	case StyleSpecial: // 翠绿
		startCol = color.RGBA{100, 230, 150, 255}
		endCol = color.RGBA{40, 160, 90, 255}
		strokeCol = color.RGBA{30, 130, 70, 255}
	default: // 经典蓝
		startCol = color.RGBA{80, 160, 240, 255}
		endCol = color.RGBA{40, 100, 180, 255}
		strokeCol = color.RGBA{30, 80, 150, 255}
	}

	// --- B. 绘图细节 ---
	// 阴影
	dc.SetRGBA(0, 0, 0, 0.1)
	dc.DrawCircle(x+1.5, y+1.5, 25.0)
	dc.Fill()

	// 渐变主体
	grad := gg.NewLinearGradient(x, y-25.0, x, y+25.0)
	grad.AddColorStop(0, startCol)
	grad.AddColorStop(1, endCol)
	dc.SetFillStyle(grad)
	dc.DrawCircle(x, y, 25.0)
	dc.Fill()

	// 边框
	dc.SetColor(strokeCol)
	dc.SetLineWidth(2)
	dc.DrawCircle(x, y, 25.0)
	dc.Stroke()

	// 文字 (白色并微调居中)
	dc.SetRGB(1, 1, 1)
	valStr := strconv.Itoa(node.Val)
	dc.DrawStringAnchored(valStr, x, y, 0.5, 0.5)

	// 递归
	if node.Left != nil {
		drawNodes(dc, node.Left, x-offsetX, y+90.0, offsetX/2, styleFunc)
	}
	if node.Right != nil {
		drawNodes(dc, node.Right, x+offsetX, y+90.0, offsetX/2, styleFunc)
	}
}
