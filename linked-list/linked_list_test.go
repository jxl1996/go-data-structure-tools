package linked_list

import "testing"

func TestPrint(t *testing.T) {
	head := Generate([]int{2, 5, 6, 8, 9, 0})
	head.Print()

	head2 := GenerateByArrayString("[1,4,3,2,5,2]")
	head2.Print()
}
