package pack

import (
	"sort"
)

type PackResult struct {
	PackSize int `json:"packSize"`
	Count    int `json:"count"`
}

func CalculatePackages(packSizes []int, items int) []PackResult {
	orderedPackSizes := sortPackSizes(packSizes)
	largePack := tryFillOneLargePack(orderedPackSizes, items)
	smallPacks := tryFillSmallPacks(orderedPackSizes, items)
	if largePack.Count == 0 {
		return smallPacks
	}
	if len(smallPacks) == 0 {
		return []PackResult{largePack}
	}

	totalPacksLarge, totalItemsLarge := countPacksAndItems([]PackResult{largePack})
	totalPacksSmall, totalItemsSmall := countPacksAndItems(smallPacks)
	if totalItemsSmall < totalItemsLarge {
		return smallPacks
	} else if totalItemsLarge < totalItemsSmall {
		return []PackResult{largePack}
	}
	if totalPacksLarge < totalPacksSmall {
		return []PackResult{largePack}
	}
	return smallPacks
}

func countPacksAndItems(selectedPacks []PackResult) (int, int) {
	totalPacks := 0
	totalItems := 0
	for _, pack := range selectedPacks {
		totalPacks += pack.Count
		totalItems += pack.PackSize * pack.Count
	}
	return totalPacks, totalItems
}

func sortPackSizes(packSizes []int) []int {
	sort.Slice(packSizes, func(i, j int) bool {
		return packSizes[i] > packSizes[j]
	})
	return packSizes
}

func tryFillOneLargePack(packSizes []int, items int) PackResult {
	for i := len(packSizes) - 1; i >= 0; i-- {
		if packSizes[i] >= items {
			return PackResult{PackSize: packSizes[i], Count: 1}
		}
	}
	return PackResult{}
}

func tryFillSmallPacks(packSizes []int, items int) []PackResult {
	smallPacks := []PackResult{}
	remainingItems := items
	for _, packSize := range packSizes {
		if remainingItems == 0 {
			break
		}
		requiredPacks := remainingItems / packSize
		if requiredPacks > 0 {
			smallPacks = append(smallPacks, PackResult{
				PackSize: packSize,
				Count:    requiredPacks,
			})
			remainingItems = remainingItems % packSize
		}
	}

	if remainingItems > 0 {
		for i := len(packSizes) - 1; i >= 0; i-- {
			if remainingItems <= packSizes[i] {
				smallPacks = append(smallPacks, PackResult{
					PackSize: packSizes[i],
					Count:    1,
				})
				break
			}
		}
	}
	return smallPacks
}
