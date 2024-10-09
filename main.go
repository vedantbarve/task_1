package main

import (
	"fmt"

	"github.com/vedantbarve/task_1/helpers"
)

func main() {
	// amazon
	result := helpers.GetApprovers("E:/MISC/go/task_1/data/amazon.csv Department=Finance Level=L2")
	fmt.Println(result)

	// amazon
	result = helpers.GetApprovers("E:/MISC/go/task_1/data/chewy.csv Country=India Department=Engineering Level=L6")
	fmt.Println(result)
}
