package simple

import (
	"fmt"
)

func ForLoop(iter int) {
	for i := 0; i < 5; i++ {
		fmt.Println("for: ", i)
	}
	nums := []int{10, 20, 30}

	for index, value := range nums {
		fmt.Printf("range index: %d, value: %d\n", index, value)
	}

}
