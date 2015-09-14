package hash

import (
	"github.com/stretchr/testify/assert"
	//"reflect"
	"testing"
)

func TestBucket(t *testing.T) {

	key := "abcde"
	n := 13
	result := Bucket(key, n)
	//assert.True(t, reflect.DeepEqual(result, 2), "Expected zero")
	assert.True(t, result < n, "result must be less than n")

	assert.True(t, Bucket("abc", 41) != Bucket("abcd", 41), "Different keys provide different buckets")
}
