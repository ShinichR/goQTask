package main

import (
	"fmt"
	"github.com/petar/GoLLRB/llrb"
)

//func lessInt(a, b interface{}) bool { return a.(int) < b.(int) }

func main() {
	tree := llrb.New()
	tree.ReplaceOrInsert(llrb.Int(1))
	tree.InsertNoReplace(llrb.Int(1))
	tree.ReplaceOrInsert(llrb.Int(2))
	tree.ReplaceOrInsert(llrb.Int(3))
	tree.InsertNoReplace(llrb.Int(3))
	tree.InsertNoReplace(llrb.Int(2))
	tree.ReplaceOrInsert(llrb.Int(4))
	tree.DeleteMin()
	tree.Delete(llrb.Int(4))
	tree.AscendGreaterOrEqual(tree.Min(), func(item llrb.Item) bool {
		i, ok := item.(llrb.Int)
		if !ok {
			return false
		}
		fmt.Println(int(i))
		return true
	})
}
