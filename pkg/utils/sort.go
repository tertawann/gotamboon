package utils

import (
	"sort"

	"github.com/gotamboon/modules/entities"
)

func SortDonatorsByTotal(rankingMaps map[string]*entities.DonatorRanking) []*entities.DonatorRanking {
	// Create a slice to store the sorted donators
	donators := make([]*entities.DonatorRanking, 0, len(rankingMaps))

	// Append donators to the slice
	for _, donator := range rankingMaps {
		donators = append(donators, donator)
	}

	// Sort the donators slice by Total field
	sort.Slice(donators, func(i, j int) bool {
		return donators[i].Total > donators[j].Total
	})

	return donators
}
