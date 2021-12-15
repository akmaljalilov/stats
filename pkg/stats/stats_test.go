package stats

import (
	"github.com/akmaljalilov/bank/pkg/types"
	"reflect"
	"testing"
)

func TestCategoriesAvg(t *testing.T) {
	res := CategoriesAvg([]types.Payment{
		{ID: 1, Category: "auto", Amount: 1_000_00},
		{ID: 2, Category: "food", Amount: 2_000_00},
		{ID: 3, Category: "auto", Amount: 3_000_00},
		{ID: 4, Category: "auto", Amount: 4_000_00},
		{ID: 4, Category: "fun", Amount: 5_000_00},
	})
	expected := map[types.Category]types.Money{
		"auto": 266666,
		"food": 200000,
		"fun":  500000,
	}
	if !reflect.DeepEqual(expected, res) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, res)
	}
}
