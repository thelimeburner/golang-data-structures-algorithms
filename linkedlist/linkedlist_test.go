package linkedlist

import "testing"

//TestNew tests that the printer function works
func TestNew(t *testing.T) {

	list1 := New()

	//check that next is nil
	if list1.Next != nil {
		t.Error("Error: Failed to create new list")
	}

}

//TestAdd adds items and tests they were successfully added
func TestAdd(t *testing.T) {

	tables := []struct {
		vals   []int
		result []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, []int{1, 2, 3, 4, 5, 6, 7}},
		{[]int{1}, []int{1}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 1, 2, 3, 4, 5, 6, 7, 1, 2, 3, 4, 5, 6, 7}, []int{1, 2, 3, 4, 5, 6, 7, 1, 2, 3, 4, 5, 6, 7, 1, 2, 3, 4, 5, 6, 7}},
		// {[]int{}, []int{}},
	}

	for _, tc := range tables {
		list1 := New()
		//fill up the list
		for _, val := range tc.vals {
			list1.Append(val)
		}

		//check the list for values
		currentPtr := list1.Next
		testPtr := 0
		for currentPtr.Next != nil {
			if currentPtr.Data != tc.result[testPtr] {
				t.Error("Error: LinkedList Append did not return expected values, expected ", tc.result[testPtr], " got ", currentPtr.Data)
				break
			}
			currentPtr = currentPtr.Next
			testPtr++
		}
	}
}

//TestGet adds items and then tests if we can retrieve them
func TestGet(t *testing.T) {

	tables := []struct {
		vals   []int
		result []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, []int{1, 2, 3, 4, 5, 6, 7}},
		{[]int{1}, []int{1}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 1, 2, 3, 4, 5, 6, 7, 1, 2, 3, 4, 5, 6, 7}, []int{1, 2, 3, 4, 5, 6, 7, 1, 2, 3, 4, 5, 6, 7, 1, 2, 3, 4, 5, 6, 7}},
		// {[]int{}, []int{}},
	}

	for _, tc := range tables {
		list1 := New()
		//fill up the list
		for _, val := range tc.vals {
			list1.Append(val)
		}

		for index, resultVal := range tc.result {
			res, _ := list1.Get(index)
			if res != resultVal {
				t.Error("Error: LinkedList Get did not return expected values, expected ", resultVal, " got ", res)
				break
			}

		}
	}

	//test if error is returned on overflow
	input := []int{1, 2, 3, 4, 5, 6, 7, 1, 2, 3, 4, 5, 6, 7, 1, 2, 3, 4, 5, 6, 7}
	list2 := New()

	//add items to list
	for _, val := range input {
		list2.Append(val)
	}

	//fetch items
	res, err := list2.Get(len(input))
	if res != -1 && err == nil {
		t.Error("Overflow Case failed. Returned: ", res)
	}
}

//TestDel tests delete functionality
func TestDel(t *testing.T) {
	tables := []struct {
		vals   []int
		dels   []int
		result []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, []int{1}, []int{2, 3, 4, 5, 6, 7}},
		{[]int{1, 6, 3, 9, 10}, []int{3, 10}, []int{1, 6, 9}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 1, 2, 3, 4, 5, 6, 7, 1, 2, 3, 4, 5, 6, 7}, []int{4, 4, 4}, []int{1, 2, 3, 5, 6, 7, 1, 2, 3, 5, 6, 7, 1, 2, 3, 5, 6, 7}},
		// {[]int{}, []int{}},
	}

	for _, tc := range tables {
		list1 := New()
		//fill up the list
		for _, val := range tc.vals {
			list1.Append(val)
		}

		for _, dVal := range tc.dels {
			err := list1.Del(dVal)
			if err != nil {
				t.Error("Error: LinkedList Delete returned an error for delete: ", dVal)
			}
		}

		for index, resultVal := range tc.result {
			res, _ := list1.Get(index)
			if res != resultVal {
				t.Error("Error: LinkedList Get did not return expected values, expected ", resultVal, " got ", res)
				break
			}

		}

	}

	//test for delete on empty list
	list2 := New()

	err := list2.Del(5)
	if err == nil {
		t.Error("Error: Expected error on delete of empty list, got nil")
	}

	//test for delete of bad value
	list3 := New()
	list3.Append(5)
	list3.Append(6)
	list3.Append(7)

	err = list3.Del(1)
	if err == nil {
		t.Error("Error: Expected error on delete of bad value list, got nil")
	}

}
