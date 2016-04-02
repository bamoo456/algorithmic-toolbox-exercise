package impl

import "testing"
import "fmt"

import "math/rand"
import "time"

func Test_MaximumPair(t *testing.T) {
	// to control test value will from [0-testValueScope]
	var testValueScope int64 = 100000
	// to control how many test values in slice
	testNums := 10
	timestamp := time.Now().UnixNano()
	r := rand.New(rand.NewSource(timestamp))
	testValues := make([]int64, testNums, testNums)
	var correctV1, correctV2 int64
	for i := 0; i < testNums; i++ {
		testValues[i] = r.Int63n(testValueScope)
		// setup correct value
		if testValues[i] > correctV1 {
			correctV2 = correctV1
			correctV1 = testValues[i]
		} else if testValues[i] > correctV2 {
			correctV2 = testValues[i]
		}
	}
	fmt.Println("Current test values are:", testValues)
	v1, v2 := MaximumPair(testValues)
	fmt.Println("The correct values are", correctV1, correctV2)
	fmt.Println("The maximum pair are:", v1, v2)
	if correctV1 != v1 || correctV2 != v2 {
		t.Error("Current Test fail")
	} else {
		t.Log("Test Pass")
	}
}
