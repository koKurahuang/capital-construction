package main

import (
	"fmt"
	"github.com/koKurahuang/capital-construction/ds"
	"strconv"
)

func main() {
	/*defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()*/
	fmt.Println("红黑树演示， 输入quit退出")
	var tree ds.RedBlackTree

	for {
		fmt.Println("插入的数值")
		var sign string
		fmt.Scanln(&sign)

		switch sign {
		case "quit", "QUIT":
			return
		default:
			value, err := strconv.Atoi(sign)
			if  err != nil {
				fmt.Println("瞎捷豹打")
				continue
			}
			if value < 0 || value > 999 {
				fmt.Println("给0到999的")
				continue
			}
			fmt.Println("-----------------------------------------------------------------------------------")
			tree.Insert(value)
			tree.Print()
			fmt.Println()
			fmt.Println()
			fmt.Println()
		}


	}
}
