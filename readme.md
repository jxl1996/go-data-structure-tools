```bash
go get github.com/jxl1996/go-data-structure-tools@master
```



```go
package main

import (
	binary_tree "github.com/jxl1996/go-data-structure-tools/binary-tree"
	linked_list "github.com/jxl1996/go-data-structure-tools/linked-list"
)

type ListNode = linked_list.ListNode
type TreeNode = binary_tree.TreeNode

func main() {
	head := linked_list.GenerateByArrayString("[1,4,3,2,5,2] ")
	head.Print()

	root := binary_tree.GenerateByArrayString("[3,9,20,null,null,15,7] ")
	root.Print()
	root.Draw()
}
```

