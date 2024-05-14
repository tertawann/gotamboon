package utils

import (
	"sort"

	"github.com/gotamboon/modules/entities"
)

// current coding : find the best pratice to manage sort then write unit test.
func SortDonatorsByTotal(rankingMaps map[string]*entities.DonatorRanking) []*entities.DonatorRanking {

	donators := make([]*entities.DonatorRanking, 0, len(rankingMaps))

	for _, donator := range rankingMaps {
		donators = append(donators, donator)
	}

	sort.Slice(donators, func(i, j int) bool {
		return donators[i].Total > donators[j].Total
	})

	return donators
}
