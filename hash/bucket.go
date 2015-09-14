package hash

func Bucket(key string, numberBuckets int) int {
	// Super Simple Hash Function (TM)
	// Just adds the integer values of each character in the string
	//
	// key is a string to be hashed
	// n is the number of buckets
	sum := 0
	for _, character := range []byte(key) {
		sum += int(character)
	}

	bucket := sum % numberBuckets
	return bucket
}
