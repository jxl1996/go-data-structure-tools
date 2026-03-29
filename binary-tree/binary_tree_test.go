package binary_tree

import (
	"fmt"
	"testing"
)

func TestPrint(t *testing.T) {
	input := []string{"3", "9", "20", "null", "null", "15", "7"}
	// input := []string{"1", "2", "2", "null", "3", "null", "3"}
	root := Generate(input)
	root.Print()
	root.Draw()

	// str := "[1,2,2,null,3,null,3]"
	str := "[] "
	root2 := GenerateByArrayString(str)
	root2.Print()
	root2.Draw()
	root2.PreOrder()
	root2.InOrder()
	root2.PostOrder()
	root2.LevelOrder()
	root2.MorrisInOrder()
	root2.MorrisPreOrder()
	root2.MorrisPostOrder()
}

func TestPrint2(t *testing.T) {
	root := GenerateRandom(20)
	// root := GenerateCompleteTree(31)
	// root := GenerateBST(30)
	// root := GenerateFullTree(5)
	root.Draw()

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
