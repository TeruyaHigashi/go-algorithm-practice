package sort

import (
	"math/rand"
	"testing"
	"time"
)

const testDataAbs = 100

func TestInsertionSort(t *testing.T) {
	sortingTest(InsertionSort, t)
}

func TestBubbleSort(t *testing.T) {
	sortingTest(BubbleSort, t)
}

func TestSelectionSort(t *testing.T) {
	sortingTest(SelectionSort, t)
}

func sortingTest(sortingFunc func([]int) []int, t *testing.T) {

	testData := createTestData()
	for _, unsorted := range testData {

		actual := sortingFunc(unsorted)

		if len(actual) != len(unsorted) {
			t.Errorf("Length is different.")
			t.Errorf("expected:%d", len(unsorted))
			t.Errorf("actual  :%d", len(actual))
		}

		if len(actual) == 1 {
			if actual[0] != unsorted[0] {
				t.Errorf("Value has changed")
				t.Errorf("not sorted")
				t.Errorf("actual : %v", actual)
			}
		}

		if len(actual) > 1 {
			for i := 1; i < len(actual); i++ {
				if actual[i-1] > actual[i] {
					t.Errorf("Not sorted")
					t.Errorf("actual : %v", actual)
				}
			}
		}
	}
}

func createTestData() [][]int {

	var testData [][]int
	testData = append(testData, []int{}, []int{1})

	rand.Seed(time.Now().UnixNano())
	var onlyPositives []int
	for i := 0; i < testDataAbs; i++ {
		onlyPositives = append(onlyPositives, int(rand.Intn(testDataAbs)))
	}
	testData = append(testData, onlyPositives)

	rand.Seed(time.Now().UnixNano())
	var includedNegatives []int
	for i := 0; i < testDataAbs; i++ {
		includedNegatives = append(includedNegatives, int(rand.Intn(testDataAbs+testDataAbs)-testDataAbs))
	}
	testData = append(testData, includedNegatives)

	return testData
}
