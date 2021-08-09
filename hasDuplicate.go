package main

func HasDuplicate(toCheck []interface{}) (has bool, duplicatedItems []interface{}) {
	clearArray := make(map[interface{}]bool)
	for _, val := range toCheck {
		if existValue := clearArray[val]; !existValue {
			clearArray[val] = true
		} else {
			duplicatedItems = append(duplicatedItems, val)
		}
	}

	return len(duplicatedItems) > 0, duplicatedItems
}
