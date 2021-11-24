package sort

import (
	"math/rand"
	"testing"
	"time"
)

const testDataRange = 100

func TestInsertionSort(t *testing.T) {
	sortingTest(InsertionSort, t)
}

func sortingTest(sortingFunc func([]int) []int, t *testing.T) {
	expected := createTestData()
	actual := sortingFunc(expected)

	if len(actual) != len(expected) {
		t.Errorf("Length is different.")
		t.Errorf("expected:%d", len(expected))
		t.Errorf("actual  :%d", len(actual))
	}

	if len(actual) > 1 {
		for i := 1; i < len(actual); i++ {
			if actual[i-1] > actual[i] {
				t.Errorf("not sorted")
				t.Errorf("actual : %v", actual)
			}
		}
	}
}

func createTestData() []int {
	rand.Seed(time.Now().UnixNano())
	var testData []int
	for i := 0; i < testDataRange; i++ {
		testData = append(testData, int(rand.Intn(testDataRange+testDataRange)-testDataRange))
	}
	return testData
}
