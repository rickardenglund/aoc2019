package main

import (
	"testing"
)

//nolint
func TestPriorityQueue_Swap(t *testing.T) {
	//items := []*state{
	//	{cost: 3},
	//	{cost: 5},
	//	{cost: 1},
	//	{cost: 8},
	//	{cost: 4},
	//}
	//
	//// Create a priority queue, put the items in it, and
	//// establish the priority queue (heap) invariants.
	//pq := make(PriorityQueue, 0) //len(items))
	////i := 0
	////for value, priority := range items {
	////	pq[i] = &Item{
	////		value:    value,
	////		priority: priority,
	////		index:    i,
	////	}
	////	i++
	////}
	//heap.Init(&pq)
	//
	//for i := range items {
	//	heap.Push(&pq, &Item{
	//		value:    items[i],
	//		priority: 100 - items[i].cost,
	//	})
	//}
	//fmt.Printf("%v\n", len(items))
	//// Insert a new item and then modify its priority.
	////item := &Item{
	////	value:    "orange",
	////	priority: 1,
	////}
	////heap.Push(&pq, item)
	////pq.update(item, item.value, 5)
	//
	//// Take the items out; they arrive in decreasing priority order.
	//for pq.Len() > 0 {
	//	item := heap.Pop(&pq).(*Item)
	//	fmt.Printf("%v - %v\n", item.priority, item.value.cost)
	//}
}
