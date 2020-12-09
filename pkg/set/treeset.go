package set

import (
	"github.com/Number1Tsar/myGoDS/internal/tree"
	"github.com/Number1Tsar/myGoDS/pkg/comparator"
)

type treeSet struct {
	bst *tree.AVLTree
}

func New(comp comparator.Comparator) *treeSet {
	tree := tree.NewAVLTree(comp)
	return &treeSet{bst: tree}
}

func(set *treeSet) Insert(values ...interface{}){
	for _, val := range values{
		set.bst.Insert(val)
	}
}

func (set *treeSet) Erase(values ...interface{}){
	for _, val := range values{
		set.bst.Erase(val)
	}
}

func (set *treeSet) Size() int{
	return set.bst.Size()
}

func (set *treeSet) Lower() interface{}{
	return set.bst.Lower()
}

func (set *treeSet) Upper() interface{}{
	return set.bst.Upper()
}

func (set *treeSet) Iterator() tree.Iterator{
	return set.bst.Iterator()
}

func (set *treeSet) Empty() bool{
	return set.bst.Empty()
}

func (set *treeSet) Find(value interface{}) bool{
	return set.bst.Find(value)
}