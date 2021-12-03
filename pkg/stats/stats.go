package stats

import "github.com/akmaljalilov/pkg/types"

func Avg(payments []types.Payment) types.Money {
	ser := types.Money(0)
	for _, payment := range payments {
		ser += payment.Amount
	}
	return types.Money(int(ser) / len(payments))
}
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sum := types.Money(0)
	for _, payment := range payments {
		if payment.Category == category {
			sum += payment.Amount
		}
	}
	return sum
}
