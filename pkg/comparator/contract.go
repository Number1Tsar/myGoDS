package comparator


type Comparator func(lhs, rhs interface{}) int


func IntComparator(lhs, rhs interface{}) int{
	return rhs.(int) - lhs.(int)
}