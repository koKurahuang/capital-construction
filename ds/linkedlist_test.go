// some tests make rings and dead loops
package ds

import (
	"fmt"
	"testing"
)

func init() {
	lp1.Value = 1
	lp2.Value = 2
	lp3.Value = 3
	lp4.Value = 4

	reset()
}

var lp1, lp2, lp3, lp4 ListPoint
var ll LinkedList

func print() {
	one := ll.Head
	for one!= nil {
		fmt.Println(one.Value)
		one = one.Next
	}
	fmt.Println("###################")
}
func reset() {
	lp1.Next = &lp2
	lp2.Next = &lp3
	lp3.Next = &lp4
	lp4.Next = nil

	ll.Head = &lp1
}

func TestLinkedList_DeletePosition(t *testing.T) {
	ll.DeletePosition(0)
	print()
	reset()
	ll.DeletePosition(2)
	print()
	reset()
	ll.DeletePosition(3)
	print()
	reset()
	ll.DeletePosition(5)
	print()
	reset()
	ll.DeletePosition(0)
	ll.DeletePosition(1)
	print()
}

func TestLinkedList_Delete(t *testing.T) {
	ll.Delete(&lp1)
	print()
	reset()
	ll.Delete(&lp2)
	print()
	reset()
	ll.Delete(&lp3)
	print()
	reset()
	ll.Delete(&lp4)
	print()
	reset()
	ll.Delete(&lp1)
	ll.Delete(&lp4)
	print()
	reset()

	var testlp ListPoint
	testlp.Value =1
	testlp.Next = &lp2
	ll.Delete(&testlp)
	print()
}

func TestLinkedList_Contains(t *testing.T) {
	fmt.Println(ll.Contains(&lp1))
	fmt.Println(ll.Contains(&lp1))
	fmt.Println(ll.Contains(&lp1))
	fmt.Println(ll.Contains(&lp1))
	ll.Delete(&lp3)
	fmt.Println(ll.Contains(&lp3))
	reset()
	var testlp ListPoint
	testlp.Value =1
	testlp.Next = &lp2
	fmt.Println(ll.Contains(&testlp))
}

func TestLinkedList_GetHead(t *testing.T) {
	fmt.Println(ll.GetHead().Value)
	ll.Delete(&lp1)
	fmt.Println(ll.GetHead().Value)
}

func TestLinkedList_GetTail(t *testing.T) {
	fmt.Println(ll.GetTail().Value)
	ll.Delete(&lp4)
	fmt.Println(ll.GetTail().Value)
}

func TestLinkedList_GetPosition(t *testing.T) {
	fmt.Println(ll.GetPosition(0).Value)
	ll.Delete(&lp1)
	fmt.Println(ll.GetPosition(0).Value)
	reset()
	fmt.Println(ll.GetPosition(6).Value)
}

// dead loop
func TestLinkedList_SetAtHead(t *testing.T) {
	var testlp ListPoint
	testlp.Value =0
	ll.SetAtHead(&testlp)
	print()
	reset()
	ll.SetAtHead(&lp3)
	print()
	reset()
	ll.SetAtHead(&lp4)
	print()
	reset()
	ll.SetAtHead(&lp1)
	print()
	reset()
}

func TestLinkedList_SetAtTail(t *testing.T) {
	var testlp ListPoint
	testlp.Value =0
	ll.SetAtTail(&testlp)
	print()
	reset()
	ll.SetAtTail(&lp3)
	print()
	reset()
	ll.SetAtTail(&lp4)
	print()
	reset()
	ll.SetAtTail(&lp1)
	print()
	reset()
}

func TestLinkedList_SetPosition(t *testing.T) {
	var testlp ListPoint
	testlp.Value =0
	ll.SetPosition(0, &testlp)
	print()
	reset()
	ll.SetPosition(1, &testlp)
	print()
	reset()
	ll.SetPosition(2, &testlp)
	print()
	reset()
	ll.SetPosition(3, &testlp)
	print()
	reset()
	ll.SetPosition(4, &testlp)
	print()
	reset()
	ll.SetPosition(6, &testlp)
	print()
	reset()
}