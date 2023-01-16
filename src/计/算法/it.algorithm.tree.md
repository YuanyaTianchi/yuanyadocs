

+++

title = "it.algorithm.tree"
description = "it.algorithm.tree"
tags = ["it", "algorithm"]

# it.algorithm.tree

## 技巧总结

三种遍历：二叉树题目思路大都离不开其三种遍历方式（都属于 dfs）。

- **前序遍历**：即父节点在前，顺序为 root-左子-右子。从 root 往下顺序构造结果，或子树依赖于一个已经正确的 root
- **中序遍历**：即父节点在中，顺序为 左子-root-右子。
- **后序遍历**：即父节点在后，顺序为 左子-右子-root

```go
func doTraverse(root *TreeNode) {
    // todo something // 前序遍历
	doTraverse(root.Left)
	// todo something // 中序遍历
	doTraverse(root.Right)
	// todo something  // 后序遍历
}
```

> 虽然 root 一般指整个树的根节点，但在递归中当前节点一般也以 root 命名，因为我们只关注当前节点，即当前(子)树的 root，而不会在乎递归的前后关系，默认**直接认为左右子树已经成为或有了预想结果**，因为跳入递归本就是一大忌讳



几个核心点

- **something**：思考当前节点下， root 的本职，以及其左右子节点具体要做的事情，它们一个都不能被忽略。
- **somewhere**：判断使用哪种遍历方式，这决定了 something 在代码内的前中后位置。主要取决于结果依赖关系，如果子树依赖于已出结果的 root，则表示应前序；如果 root 依赖于子树，则需后序......一个题目也可能有两种甚至三种遍历方式，都可以解决。一般带入三种遍历方式的思维特点即可得出答案，但切忌跳入递归
- **递归边界**：如果 something 中包含一些边界判断，不妨通过构筑辅助函数，使传入的参数可以直接构成边界判断是最方便的



## 数据结构

```go
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
```



## 遍历

### 二叉树的中序遍历

题目：[94. 二叉树的中序遍历](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/)

思路

- something：因为要求返回一个节点值的数组，因此需要初始化一个数组，并在中序位置添加 root 节点值即可
- somewhere：已由题目给出

```go
var a []int

func inorderTraversal(root *TreeNode) []int {
	a = make([]int, 0, 0) // leetcode 上如果使用全局变量需要重新初始化，因为调用多个用例将一直使用同一个变量
	doTraverse(root)
	// doTraverse(root, a)
	return a
}

func doTraverse(root *TreeNode) {
	if root == nil {
		return
	}

	//a = append(a, root.Val) // 前序遍历
	doTraverse(root.Left)
	a = append(a, root.Val) // 中序遍历
	doTraverse(root.Right)
	//a = append(a, root.Val) // 后序遍历
}
```



### 验证二叉搜索树

题目：[98. 验证二叉搜索树](https://leetcode-cn.com/problems/validate-binary-search-tree/)

思路：二叉搜索树的中序遍历是一个升序序列，因此中序遍历过程中，只要出现前一个元素大于当前元素则可以返回 false

```go
var preItem int

func isValidBST(root *TreeNode) bool {
	preItem = math.MinInt64
	return tr(root)
}

func tr(root *TreeNode) bool {
	rootRes := true
	if root == nil {
		return rootRes
	}

	leftRes := tr(root.Left)
	
	// 中序位置，判断前一个元素是否大于当前元素
	if preItem >= root.Val {
		return false
	}
	// 替换前一个元素
	preItem = root.Val
	
	rightRes := tr(root.Right)

	return rootRes && leftRes && rightRes
}
```



## 职责

题目：[226. 翻转二叉树](https://leetcode-cn.com/problems/invert-binary-tree/)

```shell
输入:
     4
   /   \
  2     7
 / \   / \
1   3 6   9
输入:
     4
   /   \
  7     2
 / \   / \
9   6 3   1
```

思路：

- something：从上至下，每个节点的左右子节点进行了位置对换
- somewhere：很容易观察到，前序、后序均可，但中序不可，因为先递归左节点，然后交换左右子节点，使左节点变成了右节点，再递归右节点（之前的递归过的左节点），而之前的右子节点（现在的左子节点）将完全不被递归到

```go
func invertTree(root *TreeNode) *TreeNode {
	doInvert(root)
	return root
}

func doInvert(root *TreeNode) {
	if root == nil {
		return
	}
	root.Left, root.Right = root.Right, root.Left
	doInvert(root.Left)
	doInvert(root.Right)
}
```



## 遍历选择

题目：[116. 填充每个节点的下一个右侧节点指针](https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node/)

<img src="img/it.algorithm.tree.填充每个节点的下一个右侧节点指针.png" style="zoom:50%;" />

```go
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node // 新 field
}
```

思路

- something
  - 左子节点：左子节点的 Next 指向右子节点，但这明显只完成了左子节点该做的事情，而右子节点的 Next 还没有着落
  - 右子节点：右子节点的 Next 指向是跨子树的，但结合 root 的 Next 看，其实就是右子节点的 Next 指向了 root 节点的 Next 节点的左子节点
  - root：这里不要把 root 的 Next 指向当作 root 的本职，因为这是由 root 作为其父节点的左子节点时完成的
- somewhere：因为右子节点的指向建立在 root 已有 next 的基础之上，左子节点的右子节点也是建立在 root 对左子节点进行过操作的基础之上的，所以中序也不可，前序可

```go
func connect(root *Node) *Node {
	doConnect(root)
	return root
}

func doConnect(root *Node) {
	if root == nil {
		return
	}

	if root.Left != nil {
		root.Left.Next = root.Right
	}
	if root.Next != nil && root.Right != nil {
		root.Right.Next = root.Next.Left
	}
	doConnect(root.Left)
	doConnect(root.Right)
}
```

```java
// 主函数
Node connect(Node root) {
    if (root == null) return null;
    connectTwoNode(root.left, root.right);
    return root;
}

// 定义：输入两个节点，将它俩连接起来
void connectTwoNode(Node node1, Node node2) {
    if (node1 == null || node2 == null) {
        return;
    }
    /**** 前序遍历位置 ****/
    // 将传入的两个节点连接
    node1.next = node2;

    // 连接相同父节点的两个子节点
    connectTwoNode(node1.left, node1.right);
    connectTwoNode(node2.left, node2.right);
    // 连接跨越父节点的两个子节点
    connectTwoNode(node1.right, node2.left);
}
```



## 递归思想

题目：[114. 二叉树展开为链表](https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/)

<img src="img/it.algorithm.tree.二叉树展开为链表.jpg" style="zoom:50%;" />

思路：

- something：右子节点嫁接到了左子节点的右子树最右侧的叶子节点上，左子节点成为 root 的右子节点
- somewhere：我服了，他猫在下面，直接上去捆啊我服了，他猫在下面，直接上去捆啊
  - 前序：先做 something，会发现找不到左子节点的右子树还并不是一个链表，嫁接到其最右侧的叶子节点上是明显错误的
  - 后序：**直接认为**左右子树都已经在递归方法下展开为链表了，做 something 即可

```go
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)

	// 左子节点移为右子节点，并用左子节点临时保存右子节点
	// 同一行中，即使 root.Left 被赋值为 nil，root.Right 也能被正常赋值为之前的 root.Left
	root.Left, root.Right = root.Right, root.Left

	// 找到最右叶子节点，并将左子节点（之前的右子节点）嫁接到最右叶子节点下
	rightLeafNodeOfLeft := root
	for rightLeafNodeOfLeft.Right != nil {
		rightLeafNodeOfLeft = rightLeafNodeOfLeft.Right
	}
	rightLeafNodeOfLeft.Right, root.Left = root.Left, nil
}
```



## 递归边界

题目：[654. 最大二叉树](https://leetcode-cn.com/problems/maximum-binary-tree/)

<img src="img/it.algorithm.tree.最大二叉树.jpg" style="zoom:50%;" />

```shell
输入：nums = [3,2,1,6,0,5]
输出：[6,3,5,null,2,0,null,null,1]
```

思路

```go
func constructMaximumBinaryTree(nums []int) *TreeNode {
	maxIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[maxIndex] < nums[i] {
			maxIndex = i
		}
	}

	return &TreeNode{
		Val: nums[maxIndex],
		Left: constructMaximumBinaryTree(nums[:maxIndex]),
		Right: constructMaximumBinaryTree(nums[maxIndex+1:]),
	}
}
```

可以发现会有下标越界的问题，一般这种边界问题，可以通过边界比较来处理，但是直接比较会比较麻烦，通过一个辅助函数将其作为递归结束条件则方便的多

```go
func constructMaximumBinaryTree(nums []int) *TreeNode {
	return doConstruct(nums, 0, len(nums))
}

func doConstruct(nums []int, left, right int) *TreeNode {
	if left == right {
		return nil
	}

	maxIndex := left
	for i := left + 1; i < right; i++ {
		if nums[maxIndex] < nums[i] {
			maxIndex = i
		}
	}

	return &TreeNode{
		Val:   nums[maxIndex],
		Left:  doConstruct(nums, left, maxIndex),
		Right: doConstruct(nums, maxIndex+1, right),
	}
}
```





题目：[105. 从前序与中序遍历序列构造二叉树](https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/)

```shell
Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
Output: [3,9,20,null,null,15,7]

Input: preorder = [-1], inorder = [-1]
Output: [-1]
```

