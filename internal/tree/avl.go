package tree

import  "github.com/Number1Tsar/myGoDS/pkg/comparator"

type AVLTree struct {
	root       *node
	comparator comparator.Comparator
	size       uint
}

func NewAVLTree(comp comparator.Comparator) *AVLTree {
	return &AVLTree{
		root:       nil,
		comparator: comp,
	}
}

func (tree *AVLTree) Size() int{
	return int(tree.size)
}

func (tree *AVLTree) Empty() bool{
	return tree.size == 0
}

func (tree *AVLTree) Insert(value interface{}){
	tree.root = tree.insertNode(tree.root, value)
}

func (tree *AVLTree) Find(value interface{}) bool{
	return tree.findNode(tree.root, value)
}

func (tree *AVLTree) Erase(value interface{}){
	tree.root = tree.eraseNode(tree.root, value)
}

func (tree *AVLTree) insertNode(treeNode *node, value interface{}) *node {
	if treeNode == nil{
		tree.size++
		return &node{
			value:value,
			left:nil,
			right:nil,
			height:1,
		}
	}
	compValue := tree.comparator(value, treeNode.value)
	if compValue == 0{
		treeNode.value = value
	} else if compValue > 0{
		treeNode.left = tree.insertNode(treeNode.left, value)
	} else{
		treeNode.right = tree.insertNode(treeNode.right, value)
	}

	treeNode.height = 1 + max(height(treeNode.left), height(treeNode.right))
	balanceFactor := balance(treeNode)

	if balanceFactor > 1 && balance(treeNode.left) > 0{  //left heavy
		return rightRotation(treeNode)
	} else if balanceFactor > 1 && balance(treeNode.left) < 0{ //left-right
		treeNode.left = leftRotation(treeNode.left)
		return rightRotation(treeNode)
	} else if balanceFactor < -1 && balance(treeNode.right) < 0{ // right heavy
		return leftRotation(treeNode)
	} else if balanceFactor < -1 && balance(treeNode.right) > 0{ // right-left
		treeNode.right = rightRotation(treeNode.right)
		return leftRotation(treeNode)
	}

	return treeNode
}

func max(x, y int) int{
	if x > y {
		return x
	}
	return y
}

func leftRotation(treeNode *node) *node {
	y := treeNode.right
	x := y.left

	y.left = treeNode
	treeNode.right = x

	treeNode.height = 1 + max(height(treeNode.left), height(treeNode.right))
	y.height = 1 + max(height(y.left), height(y.right))

	return y
}

func rightRotation(treeNode *node) *node {
	y := treeNode.left
	x := y.right

	y.right = treeNode
	treeNode.left = x

	treeNode.height = 1 + max(height(treeNode.left), height(treeNode.right))
	y.height = 1 + max(height(y.left), height(y.right))

	return y
}

func height(treeNode *node) int{
	if treeNode == nil{
		return 0
	}
	return treeNode.height
}

func balance(treeNode *node) int{
	if treeNode == nil{
		return 0
	}
	return height(treeNode.left) - height(treeNode.right)
}

func (tree *AVLTree) findNode(treeNode *node, value interface{}) bool{
	if treeNode == nil{
		return false
	}
	compValue := tree.comparator(value, treeNode.value)
	if compValue == 0{
		return true
	} else if compValue > 0{
		return tree.findNode(treeNode.left, value)
	} else{
		return tree.findNode(treeNode.right, value)
	}
}

func (tree *AVLTree) eraseNode(treeNode *node, value interface{}) *node {
	if treeNode == nil{
		return nil
	}
	compValue := tree.comparator(value, treeNode.value)
	if compValue == 0{
		tree.size--
		if treeNode.left == nil && treeNode.right == nil{
			return nil
		} else if treeNode.right == nil{
			return treeNode.left
		} else {
			ptr := treeNode.right
			for ptr.left != nil{
				ptr = ptr.left
			}
			treeNode.value = ptr.value
			treeNode.right = tree.eraseNode(treeNode.right, ptr.value)
		}
	} else if compValue > 0{
		treeNode.left = tree.eraseNode(treeNode.left, value)
	} else{
		treeNode.right = tree.eraseNode(treeNode.right, value)
	}

	treeNode.height = 1 + max(height(treeNode.left), height(treeNode.right))
	balanceFactor := balance(treeNode)

	if balanceFactor > 1 && balance(treeNode.left) >= 0{  //left heavy
		return rightRotation(treeNode)
	} else if balanceFactor > 1 && balance(treeNode.left) < 0{ //left-right
		treeNode.left = leftRotation(treeNode.left)
		return rightRotation(treeNode)
	} else if balanceFactor < -1 && balance(treeNode.right) <= 0{ // right heavy
		return leftRotation(treeNode)
	} else if balanceFactor < -1 && balance(treeNode.right) > 0{ // right-left
		treeNode.right = rightRotation(treeNode.right)
		return leftRotation(treeNode)
	}

	return treeNode
}

func (tree *AVLTree) Iterator() Iterator {
	return newIterator(tree.root)
}

func (tree *AVLTree) preorderIterator() Iterator {
	return newPreorderIterator(tree.root)
}

func (tree *AVLTree) Lower() interface{}{
	ptr := tree.root
	for ptr.left != nil{
		ptr = ptr.left
	}
	return ptr.value
}

func (tree *AVLTree) Upper() interface{}{
	ptr := tree.root
	for ptr.right != nil{
		ptr = ptr.right
	}
	return ptr.value
}

type avlIterator struct{
	stack []*node
}


func newIterator(root *node) Iterator {
	list := make([]*node, 0)
	for root != nil{
		list = append(list, root)
		root = root.left
	}

	return &avlIterator{stack:list}
}


func(iterator *avlIterator) Next() interface{}{
	if len(iterator.stack) == 0{
		return nil
	}
	top := iterator.stack[len(iterator.stack)-1]
	value := top.value
	iterator.stack = iterator.stack[0:len(iterator.stack)-1]
	top = top.right
	for top != nil{
		iterator.stack = append(iterator.stack, top)
		top = top.left
	}
	return value
}


func (iterator *avlIterator) HasNext() bool{
	return len(iterator.stack) > 0
}


type avlPreorderIterator struct {
	stack []*node
}

func newPreorderIterator(root *node) Iterator {
	list := make([]*node, 0)
	if root != nil{
		list = append(list, root)
	}
	return &avlPreorderIterator{stack:list}
}


func (iterator *avlPreorderIterator) Next() interface{}{
	if len(iterator.stack) == 0{
		return nil
	}
	top := iterator.stack[len(iterator.stack)-1]
	iterator.stack = iterator.stack[0:len(iterator.stack)-1]
	value := top.value
	if top.right != nil{
		iterator.stack = append(iterator.stack, top.right)
	}
	if top.left != nil{
		iterator.stack = append(iterator.stack, top.left)
	}
	return value
}

func (iterator *avlPreorderIterator) HasNext() bool{
	return len(iterator.stack) > 0
}