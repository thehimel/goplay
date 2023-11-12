// Reference: https://golangnote.com/map/sort-golang-map-by-keys-or-values

package main

import (
	"fmt"
	"sort"
)

func main() {
	// Initialize a map.
	myMap := make(map[string]float64)
	myMap["Orange"] = 10
	myMap["Banana"] = 12
	myMap["Apple"] = 8
	fmt.Println(myMap)

	// Create a slice of keys.
	keys := make([]string, 0, len(myMap))
	for key := range myMap {
		keys = append(keys, key)
	}
	fmt.Println(keys)

	// Sort the slice based on the value from the map.
	sort.SliceStable(keys, func(i, j int) bool {
		return myMap[keys[i]] < myMap[keys[j]]
	})
	fmt.Println(keys)

	// Create a new map using the sorted slice.
	myNewMap := map[string]float64{}
	for _, k := range keys {
		myNewMap[k] = myMap[k]
	}
	fmt.Println(myNewMap)
}
