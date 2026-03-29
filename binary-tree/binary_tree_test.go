package binary_tree

import "testing"

func TestPrint(t *testing.T) {
	input := []string{"3", "9", "20", "null", "null", "15", "7"}
	// input := []string{"1", "2", "2", "null", "3", "null", "3"}
	root := Generate(input)
	root.Print()
	root.Draw()

	// str := "[1,2,2,null,3,null,3]"
	str := "[3,9,20,null,null,15,7] "
	root2 := GenerateByArrayString(str)
	// root2.Print()
	// root2.Draw()
	root2.PreOrder()
	root2.InOrder()
	root2.PostOrder()
	root2.LevelOrder()
	root2.MorrisInOrder()
	root2.MorrisPreOrder()
	root2.MorrisPostOrder()
}
