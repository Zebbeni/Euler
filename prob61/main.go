package main

import "fmt"

// a map of maps to store all 4-digit polygonal numbers. Allows looking up
// values by their first two digits and polygonal type (or family)
// (ie. polymap[firstTwoDigits][polygonalFamily])
var polymap map[int]map[int][]int

func main() {
	// generate all 4-digit polygonal numbers
	fmt.Println("generating polygonals...")
	generatePolygonals()

	fmt.Println("finding set...")
	findSet()
}

func findSet() {
	for _, familymap := range polymap {
		for family, values := range familymap {
			for _, value := range values {
				valueList := []int{value}              // the ordered list of numbers
				families := map[int]bool{family: true} // the map of families included
				if found := findMatch(valueList, families); found {
					return
				}
			}
		}
	}
}

func findMatch(valueList []int, families map[int]bool) bool {
	// check wraparound match if 6 values found
	if len(valueList) == 6 {
		if valueList[5]%100 == valueList[0]/100 {
			fmt.Println("Found! values:", valueList, " families:", families)
			return true
		}
		return false
	}
	// build new value lists and family maps with all valid values that match
	// the current tail and call findMatch
	lastDigits := valueList[len(valueList)-1] % 100
	for family, matchingValues := range polymap[lastDigits] {
		// skip if a number in this family already exists in the set we're building
		if _, ok := families[family]; ok {
			continue
		}
		for _, value := range matchingValues {
			newList := append(valueList, value)
			newFamilies := copyAndAddToMap(families, family)
			if found := findMatch(newList, newFamilies); found {
				return true
			}
		}
	}
	return false
}

func copyAndAddToMap(familyMap map[int]bool, family int) map[int]bool {
	newFamilyMap := make(map[int]bool)
	for k := range familyMap {
		newFamilyMap[k] = true
	}
	newFamilyMap[family] = true
	return newFamilyMap
}

func generatePolygonals() {
	polymap = make(map[int]map[int][]int)
	for i := 0; i < 6; i++ {
		for n := 1; n < 1000000; n++ {
			x := n * (n*(i+1) + (1 - i)) / 2
			if x >= 1000 && x < 10000 {
				addPolygonal(x, i)
			} else if x >= 10000 {
				break
			}
		}
	}
}

func addPolygonal(value, family int) {
	firstDigits := value / 100
	if _, ok := polymap[firstDigits]; !ok {
		polymap[firstDigits] = make(map[int][]int)
	}
	if _, ok := polymap[firstDigits][family]; !ok {
		polymap[firstDigits][family] = make([]int, 0)
	}
	polymap[firstDigits][family] = append(polymap[firstDigits][family], value)
}
