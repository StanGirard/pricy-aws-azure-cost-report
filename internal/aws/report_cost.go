package aws

import (
	"flag"
	"fmt"
	"sort"

	"github.com/stangirard/pricy/internal/format"
)

var (
	details = flag.Bool("details", false, "print details")
)

func PrintCostByService(ServicesPrices []format.ServicePrice) {
	//Sort by Price
	sort.Slice(ServicesPrices, func(i, j int) bool {
		return ServicesPrices[i].Cost > ServicesPrices[j].Cost
	})

	for _, price := range ServicesPrices {
		price.Print()
	}
}

func PrintTotalCost(totalCost float64) {
	fmt.Println("Total Cost:", totalCost)
}

func reportGenerate(costByServices []format.ServicePrice) {
	if *details {
		PrintCostByService(costByServices)
	}
	PrintTotalCost(TotalCostUsage(costByServices))
}