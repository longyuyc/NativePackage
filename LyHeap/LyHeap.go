//heap包提供了对任意类型（实现了heap.Interface接口）的堆操作。
//（最小）堆是具有“每个节点都是以其为根的子树中最小值”属性的树。
//树的最小元素为其根元素，索引0的位置。
//heap是常用的实现优先队列的方法。要创建一个优先队列，
//实现一个具有使用（负的）优先级作为比较的依据的Less方法的Heap接口，
//如此一来可用Push添加项目而用Pop取出队列最高优先级的项目。

package LyHeap

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}
func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h IntHeap) Swap(i, j int) {
	h[i],
		h[j] = h[j],
		h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func Example_intHeap() {
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Printf("minimun:%d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Print("%d ", heap.Pop(h))
	}
}
