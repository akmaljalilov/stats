package stats

import "github.com/akmaljalilov/bank/v2/pkg/types"

func Avg(payments []types.Payment) types.Money {
	ser := types.Money(0)
	for _, payment := range payments {
		if payment.Status != "StatusFail" {
			ser += payment.Amount
		}
	}
	return types.Money(int(ser) / len(payments))
}
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sum := types.Money(0)
	for _, payment := range payments {
		if payment.Category == category {
			if payment.Status != "StatusFail" {
				sum += payment.Amount
			}
		}
	}
	return sum
}

func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	res := map[types.Category]types.Money{}
	countCategory := map[types.Category]int{}
	for _, payment := range payments {
		res[payment.Category] += payment.Amount
		countCategory[payment.Category] += 1
	}
	for category, count := range countCategory {
		res[category] = res[category] / types.Money(count)
	}
	return res
}
