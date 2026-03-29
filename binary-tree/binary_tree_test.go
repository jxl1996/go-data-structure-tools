package binary_tree

import (
	"fmt"
	"testing"
)

func TestPrint(t *testing.T) {
	// input := []string{"3", "9", "20", "null", "null", "15", "7"}
	input := []string{"1", "2", "2", "null", "3", "null", "3"}
	root := Generate(input)
	root.Print()
	// root.Draw()

	// str := "[1,2,2,null,3,null,3]"
	str := "[] "
	root2 := GenerateByArrayString(str)
	root2.Print()
	// root2.Draw()
	root2.PreOrder()
	root2.InOrder()
	root2.PostOrder()
	root2.LevelOrder()
	root2.MorrisInOrder()
	root2.MorrisPreOrder()
	root2.MorrisPostOrder()
	fmt.Println("高度", root2.GetHeight())
}

func TestPrint2(t *testing.T) {
	root := GenerateRandom(10)
	// root := GenerateCompleteTree(10, 10, 10000)
	// root := GenerateBST(30, 10, 10000)
	// root := GenerateFullTree(3)
	// root.Draw("root.png", func(n *TreeNode) NodeStyle {
	// 	if n.Val > 50 {
	// 		return StyleSpecial
	// 	}
	// 	return StyleDefault
	// })

	root.Draw("root.png", func(n *TreeNode) NodeStyle {
		// 假设你有一个计算平衡因子的函数
		if n.GetBalanceFactor() >= 2 {
			return StyleWarning // 失衡节点标红
		}
		return StyleDefault
	})

	fmt.Println("是否为满二叉树：", root.IsFullTree())
	fmt.Println("是否为完全二叉树：", root.IsCompleteTree())
	fmt.Println("是否为二叉搜索树：", root.IsBST())
	fmt.Println("是否平衡：", root.IsBalanced())

	root.PreOrder()
	root.InOrder()
	root.PostOrder()
	root.LevelOrder()
	root.MorrisInOrder()
	root.MorrisPreOrder()
	root.MorrisPostOrder()
	root.Print()
}
