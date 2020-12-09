package set

import (
	"testing"

	"github.com/Number1Tsar/myGoDS/pkg/comparator"
	"github.com/stretchr/testify/assert"
)

func newSet(comp comparator.Comparator) *treeSet{
	return New(comp)
}

func TestTreeSet_Insert(t *testing.T) {
	value := []int{1,2,3,4,5,1}
	set := newSet(comparator.IntComparator)

	set.Insert(1,2,3,4,5,1)
	for _, i := range value{
		assert.Equal(t, true, set.Find(i))
	}
}

func TestTreeSet_Erase(t *testing.T) {
	set := newSet(comparator.IntComparator)
	set.Insert(2,3,4,5)

	set.Erase(4,5)
	assert.Equal(t, false, set.Find(4))
	assert.Equal(t, false, set.Find(5))
}

func TestTreeSet_Size(t *testing.T) {
	set := newSet(comparator.IntComparator)
	set.Insert(1,2,3,4,1,2,3,4)

	assert.Equal(t, 4, set.Size())
}

func TestTreeSet_Empty(t *testing.T) {
	set := newSet(comparator.IntComparator)
	set.Insert(1)

	assert.Equal(t, false, set.Empty())
	set.Erase(1)
	assert.Equal(t, true, set.Empty())
}

func TestTreeSet_Lower(t *testing.T) {
	set := newSet(comparator.IntComparator)
	set.Insert(1,2,3,4,5)

	small := set.Lower()

	assert.Equal(t, 1, small)
}

func TestTreeSet_Upper(t *testing.T) {
	set := newSet(comparator.IntComparator)
	set.Insert(1,2,3,4,5)

	large := set.Upper()

	assert.Equal(t, 5, large)
}

func TestTreeSet_Iterator(t *testing.T) {
	set := newSet(comparator.IntComparator)
	value := []int{1,2,3,4,5}
	set.Insert(1,2,3,4,5)

	iterator := set.Iterator()

	for i:=0; iterator.HasNext(); i++{
		assert.Equal(t, value[i], iterator.Next())
	}
}