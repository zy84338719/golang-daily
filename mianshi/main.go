package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverselist(head *ListNode) *ListNode {
	var temp *ListNode
	pre := head
	for pre != nil {
		t := pre.Next
		pre.Next = temp
		temp = pre
		pre = t
	}
	return temp
}

type NodeTree struct {
	Val   int
	Left  *NodeTree
	Right *NodeTree
}

func levelOrder(node *NodeTree) (res [][]int) {
	res = [][]int{}
	if node == nil {
		return
	}
	queue := []*NodeTree{node}
	for len(queue) > 0 {
		//创建临时队列,保存层节点
		tmpQ := []*NodeTree{}
		tmpRes := []int{}
		for j := 0; j < len(queue); j++ {
			node := queue[j]
			tmpRes = append(tmpRes, node.Val)
			if node.Right != nil {
				tmpQ = append(tmpQ, node.Right)
			}
			if node.Left != nil {
				tmpQ = append(tmpQ, node.Left)
			}
		}
		//处理下一层
		queue = tmpQ
		res = append(res, tmpRes)
	}
	return
}
func main() {

}
