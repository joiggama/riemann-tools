package utils

func ZeroedCollection(collection []int64) bool {
	zeroed := true

	for _, value := range collection {
		if value != 0 {
			zeroed = false
			break
		}
	}

	return zeroed
}
