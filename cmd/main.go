package main

import (
	"fmt"
	"github.com/Number1Tsar/myGoDS/pkg/comparator"
	"github.com/Number1Tsar/myGoDS/pkg/set"
)

func main(){
	//Two sum problem test
	values := []int{2,7,11,15}
	target := 9
	set := set.New(comparator.IntComparator)
	for _, val := range values{
		if set.Find(target-val){
			fmt.Printf("[%v %v]", val, target-val)
			break
		}
		set.Insert(val)
	}
}
