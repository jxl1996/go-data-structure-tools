package binary_tree

import "fmt"

// PreOrder 前序遍历: 根 -> 左 -> 右
func (root *TreeNode) PreOrder() {
	fmt.Print("前序遍历：")
	if root == nil {
		fmt.Println()
		return
	}

	stack := []*TreeNode{root}
	for len(stack) > 0 {
		// Pop
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		fmt.Printf("%d ", node.Val)

		// 栈是后进先出，所以先压右孩子，再压左孩子
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	fmt.Println()
}

// InOrder 中序遍历 左 -> 根 -> 右
func (root *TreeNode) InOrder() {
	fmt.Print("中序遍历：")
	if root == nil {
		fmt.Println()
		return
	}

	var stack []*TreeNode
	curr := root

	for curr != nil || len(stack) > 0 {
		// 1. 尽可能向左走，并把路径上的节点压入栈
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		// 2. 此时左边已经到底了，弹出栈顶节点（最近的根）
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// 3. 访问节点
		fmt.Printf("%d ", curr.Val)

		// 4. 转向右子树，重复上述过程
		curr = curr.Right
	}
	fmt.Println()
}

// PostOrder 后序遍历: 左 -> 右 -> 根
func (root *TreeNode) PostOrder() {
	fmt.Print("后序遍历：")
	if root == nil {
		fmt.Println()
		return
	}

	var res []int
	stack := []*TreeNode{root}

	// 1. 按照 "根 -> 右 -> 左" 的顺序访问并存入数组
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)

		// 注意：这里先压左再压右，弹出时就是先右后左
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}

	// 2. 反转结果数组，得到 "左 -> 右 -> 根"
	for i := len(res) - 1; i >= 0; i-- {
		fmt.Printf("%d ", res[i])
	}
	fmt.Println()
}

// LevelOrder 层序遍历
func (root *TreeNode) LevelOrder() {
	fmt.Print("层序遍历：")
	if root == nil {
		fmt.Println()
		return
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		// 出队
		node := queue[0]
		queue = queue[1:]

		fmt.Printf("%d ", node.Val)

		// 左孩子入队
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		// 右孩子入队
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	fmt.Println()
}

// MorrisInOrder Morris中序遍历
func (root *TreeNode) MorrisInOrder() {
	fmt.Print("Morris中序遍历：")
	curr := root
	var pre *TreeNode

	for curr != nil {
		if curr.Left == nil {
			// 1. 如果左子树为空，访问当前节点，然后进入右子树
			fmt.Printf("%d ", curr.Val)
			curr = curr.Right
		} else {
			// 2. 找到当前节点在左子树上的“前驱节点”（左子树里最右边的节点）
			pre = curr.Left
			for pre.Right != nil && pre.Right != curr {
				pre = pre.Right
			}

			if pre.Right == nil {
				// 3a. 建立“线索”：将前驱节点的右指针指向当前节点，然后进入左子树
				pre.Right = curr
				curr = curr.Left
			} else {
				// 3b. 断开“线索”：说明左子树已访问完，恢复原树结构，访问当前节点，进入右子树
				pre.Right = nil
				fmt.Printf("%d ", curr.Val)
				curr = curr.Right
			}
		}
	}
	fmt.Println()
}

// MorrisPreOrder Morris前序遍历
func (root *TreeNode) MorrisPreOrder() {
	fmt.Print("Morris前序遍历：")
	curr := root
	for curr != nil {
		if curr.Left == nil {
			// 情况1：左子树为空，打印并进入右子树
			fmt.Printf("%d ", curr.Val)
			curr = curr.Right
		} else {
			// 情况2：找到左子树的最右节点（前驱节点）
			pre := curr.Left
			for pre.Right != nil && pre.Right != curr {
				pre = pre.Right
			}

			if pre.Right == nil {
				// 第一次访问：打印根节点，建立线索，进入左子树
				fmt.Printf("%d ", curr.Val)
				pre.Right = curr
				curr = curr.Left
			} else {
				// 第二次访问：说明左子树已处理完，断开线索，进入右子树
				pre.Right = nil
				curr = curr.Right
			}
		}
	}
	fmt.Println()
}

// 辅助函数：反转右孩子路径
func reversePath(from *TreeNode) *TreeNode {
	var prev *TreeNode
	curr := from
	for curr != nil {
		next := curr.Right
		curr.Right = prev
		prev = curr
		curr = next
	}
	return prev
}

// 辅助函数：逆序打印路径
func printReversePath(from *TreeNode) {
	tail := reversePath(from)
	curr := tail
	for curr != nil {
		fmt.Printf("%d ", curr.Val)
		curr = curr.Right
	}
	reversePath(tail) // 恢复原状
}

// MorrisPostOrder Morris后序遍历
func (root *TreeNode) MorrisPostOrder() {
	fmt.Print("Morris后序遍历：")

	// 创建一个哑节点（Dummy Node），让它的左孩子指向根节点
	dummy := &TreeNode{Left: root}
	curr := dummy

	for curr != nil {
		if curr.Left == nil {
			curr = curr.Right
		} else {
			pre := curr.Left
			for pre.Right != nil && pre.Right != curr {
				pre = pre.Right
			}

			if pre.Right == nil {
				pre.Right = curr
				curr = curr.Left
			} else {
				// 关键点：在断开连接前，逆序打印左子树的右边界
				pre.Right = nil
				printReversePath(curr.Left)
				curr = curr.Right
			}
		}
	}
	fmt.Println()
}
