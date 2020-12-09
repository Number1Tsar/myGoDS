package tree

import (
	"github.com/Number1Tsar/myGoDS/pkg/comparator"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func newTree(comp comparator.Comparator) *AVLTree {
	return NewAVLTree(comp)
}


func TestAVLTree_InsertFirst(t *testing.T) {
	tree := newTree(comparator.IntComparator)

	tree.Insert(5)

	assert.Equal(t, true, tree.Find(5))
}


func TestAVLTree_Insert(t *testing.T) {
	tree := newTree(comparator.IntComparator)
	value := []int{7,3,9,8,10,4}
	for _, val := range value{
		tree.Insert(val)
	}

	for _, val := range value{
		assert.Equal(t, true, tree.Find(val))
	}
}


func TestAVLTree_ErasePresent(t *testing.T) {
	tree := newTree(comparator.IntComparator)
	value := []int{7,3,9,8,10,4}
	for _, val := range value{
		tree.Insert(val)
	}

	tree.Erase(9)

	assert.Equal(t, false, tree.Find(9))
}


func TestAVLTree_Size1(t *testing.T) {
	tree := newTree(comparator.IntComparator)
	value := []int{7,3,9,8,10,4}
	for _, val := range value{
		tree.Insert(val)
	}

	tree.Erase(1)

	assert.Equal(t, len(value), tree.Size())
}

func TestAVLTree_Size2(t *testing.T) {
	tree := newTree(comparator.IntComparator)
	value := []int{7,3,9,8,10,4}
	for _, val := range value{
		tree.Insert(val)
	}

	assert.Equal(t, len(value), tree.Size())
}


func TestAVLTree_Empty1(t *testing.T) {
	tree := newTree(comparator.IntComparator)
	tree.Insert(5)

	tree.Erase(5)

	assert.Equal(t, true, tree.Empty())
}

func TestAVLTree_Empty2(t *testing.T) {
	tree := newTree(comparator.IntComparator)
	tree.Insert(5)

	assert.Empty(t, false, tree.Empty())
}


func TestAVLTree_Iterator(t *testing.T) {
	tree := newTree(comparator.IntComparator)
	value := []int{7,3,9,8,10,4}
	for _, val := range value{
		tree.Insert(val)
	}

	iterator := tree.Iterator()
	sort.Sort(sort.IntSlice(value))
	for i := 0; iterator.HasNext(); i++{
		assert.Equal(t, value[i], iterator.Next())
	}
}

func TestAVLTree_InsertSkewed(t *testing.T) {
	tree := newTree(comparator.IntComparator)
	value := []int{10,20,30,40,50,25}
	for _, val := range value{
		tree.Insert(val)
	}

	iterator := tree.preorderIterator()
	value = []int{30,20,10,25,40,50}
	for i:=0; iterator.HasNext(); i++{
		assert.Equal(t, value[i], iterator.Next())
	}
}

func TestAVLTree_Lower(t *testing.T) {
	tree := newTree(comparator.IntComparator)
	value := []int{7,3,9,8,10,4}
	for _, val := range value{
		tree.Insert(val)
	}
	sort.Sort(sort.IntSlice(value))

	for i:=0; i<len(value); i++{
		lower := tree.Lower()
		assert.Equal(t, value[i], lower)
		tree.Erase(lower)
	}
}

func TestAVLTree_Upper(t *testing.T) {
	tree := newTree(comparator.IntComparator)
	value := []int{7,3,9,8,10,4}
	for _, val := range value{
		tree.Insert(val)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(value)))

	for i:=0; i<len(value); i++{
		upper := tree.Upper()
		assert.Equal(t, value[i], upper)
		tree.Erase(upper)
	}
}