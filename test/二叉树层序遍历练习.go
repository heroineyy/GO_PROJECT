package main

import "container/list"

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil{
		return res
	}

	queue:=list.New()//先new一个队列
	queue.PushBack(root)//将root入队列
	var tmpArr []int
	for queue.Len()>0{//que不为空
		length :=queue.Len()//求队列的长度
		for i:=0;i<length;i++{
			node :=queue.Remove(queue.Front()).(*TreeNode)//出队列
			if node.Left!=nil{
				queue.PushBack(node.Left)
			}
			if node.Right!=nil{
				queue.PushBack(node.Right)
			}
			tmpArr =append(tmpArr,node.Val)
		}
		res = append(res,tmpArr)
		tmpArr=[]int{}//清空
	}
	return res
}


