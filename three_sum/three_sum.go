package three_sum

import (
	"fmt"
	"sort"
)

type Index struct {
	Val   int
	Inner []*Index
}

func threeSum(nums []int) [][]int {
	oneThird := len(nums) / 3
	first, second := oneThird, oneThird
	third := oneThird + len(nums)%3

	var store []*Index
	var result [][]int

	for i := 0; i < first; i++ {
		println("first")

		for j := 0; j < second; j++ {
			println("second")
			for k := 0; k < third; k++ {
				println("third")
				triple := []int{nums[i], nums[j+oneThird], nums[k+oneThird*2]}
				sort.Slice(triple, func(i, j int) bool {
					return triple[i] < triple[j]
				})

				tripleIndex := []int{i, j + oneThird, k + oneThird*2}
				sort.Slice(tripleIndex, func(i, j int) bool {
					return tripleIndex[i] < tripleIndex[j]
				})

				fmt.Printf("%v", triple)
				// sort before search
				if isTarget(triple) {

					if !containsIndex(tripleIndex, &store) {
						addIndex(tripleIndex, &store)
					}
					if !containsSameTriplet(triple, result) {
						addValues(triple, &result)
					}
				}
			}
		}
	}

	return result
}

func isTarget(triple []int) bool {
	return triple[0]+triple[1]+triple[2] == 0
}

func containsIndex(tripleIndex []int, store *[]*Index) bool {
	var found1 *Index
	var found2 *Index
	var found3 *Index

	for i := 0; i < len(*store); i++ {
		if (*store)[i].Val == tripleIndex[0] {
			found1 = (*store)[i]
			break
		}
	}
	if found1 != nil {
		for i := 0; i < len(found1.Inner); i++ {
			if found1.Inner[i].Val == tripleIndex[1] {
				found2 = found1.Inner[i]
				break
			}
		}
		if found2 != nil {
			for i := 0; i < len(found2.Inner); i++ {
				if found2.Inner[i].Val == tripleIndex[2] {
					found3 = found2.Inner[i]
					break
				}
			}

			if found3 != nil {
				return true
			} else {
				return false
			}
		} else {
			return false
		}
	} else {
		return false
	}
}

func addIndex(tripleIndex []int, store *[]*Index) {
	var found1 *Index
	var found2 *Index
	var found3 *Index
	for i := 0; i < len(*store); i++ {
		if (*store)[i].Val == tripleIndex[0] {
			found1 = (*store)[i]
			break
		}
	}
	if found1 != nil {
		// second search
		for i := 0; i < len(found1.Inner); i++ {
			if found1.Inner[i].Val == tripleIndex[1] {
				found2 = found1.Inner[i]
				break
			}
		}
		if found2 != nil {
			// third search
			for i := 0; i < len(found2.Inner); i++ {
				if found2.Inner[i].Val == tripleIndex[2] {
					found3 = found2.Inner[i]
					break
				}
			}
			if found3 != nil {
				return // already exists
			} else {
				found2.Inner = append(found2.Inner, &Index{Val: tripleIndex[2]})
			}
		} else {
			found1.Inner = append(found1.Inner, &Index{Val: tripleIndex[1], Inner: []*Index{{Val: tripleIndex[2]}}})
		}
	} else {
		*store = append(*store, &Index{Val: tripleIndex[0], Inner: []*Index{{Val: tripleIndex[1], Inner: []*Index{{Val: tripleIndex[2]}}}}})
	}
}

func addValues(triple []int, result *[][]int) {
	*result = append(*result, triple)
}

func containsSameTriplet(triple []int, storeValue [][]int) bool {
	for i := 0; i < len(storeValue); i++ {
		if storeValue[i][0] == triple[0] && storeValue[i][1] == triple[1] && storeValue[i][2] == triple[2] {
			return true
		}
	}

	return false
}
