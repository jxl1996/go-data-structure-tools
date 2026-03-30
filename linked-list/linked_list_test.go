package linked_list

import "testing"

func TestPrint(t *testing.T) {
	head := Generate([]int{})
	head.Print()

	head2 := GenerateByArrayString("[1,4,3,2,5,2]")
	head2.Print("链表展示：")

	head3 := GenerateRandom(50, 50, 500)
	head3.Print()
}
