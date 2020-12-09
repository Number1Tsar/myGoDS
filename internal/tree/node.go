package tree


type node struct {
	value interface{}
	left *node
	right *node
	height int
}


type Iterator interface {
	Next() interface{}
	HasNext() bool
}



