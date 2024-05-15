package utils

import (
	"sort"

	"github.com/gotamboon/modules/entities"
)

// current coding : find the best pratice to manage sort then write unit test.
func SortDescDonatorsByTotal(rankingMaps map[string]*entities.DonatorRanking) []*entities.DonatorRanking {

	donators := ConvertMapToSlice(rankingMaps, len(rankingMaps))

	sort.Slice(donators, func(i, j int) bool {
		return donators[i].Total > donators[j].Total
	})

	return donators
}
