/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
    if root == nil {
        return []string{}
    }

    rootValStr := strconv.Itoa(root.Val)

    if root.Left == nil && root.Right == nil {
        return []string{rootValStr}
    }

    l, r := binaryTreePaths(root.Left), binaryTreePaths(root.Right)

    for i := range l {
        l[i] = rootValStr + "->" + l[i]
    }

    for i := range r {
        r[i] = rootValStr + "->" + r[i]
    }

    return append(l, r...)
}