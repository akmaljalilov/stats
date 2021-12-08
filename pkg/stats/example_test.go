package stats

import (
	"fmt"
	"github.com/akmaljalilov/bank/v2/pkg/types"
)

func ExampleAvg() {
	fmt.Println(Avg([]types.Payment{
		{Amount: 10},
		{Amount: 20},
		{Amount: 30},
		{Amount: 60},
	}))
	// Output:
	// 30

}

func ExampleTotalInCategory() {
	fmt.Println(TotalInCategory([]types.Payment{
		{Amount: 10, Category: "A"},
		{Amount: 20, Category: "B"},
		{Amount: 30, Category: "B"},
		{Amount: 60, Category: "A"},
	}, "A"))
	// Output:
	// 70

}
