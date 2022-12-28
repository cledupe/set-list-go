package setlist

import "testing"

type Test struct {
	value int
}

func TestSetListWithInt(t *testing.T) {
	setList := New[int]()
	t.Log("Test Initial value")
	if setList.Size() != 0 {
		t.Fail()
	}

	t.Log("Test insertion one value")
	setList.Insert(1)
	if setList.Size() != 1 {
		t.Fail()
	}

	t.Log("Test unique value")
	setList.Insert(1)
	if setList.Size() != 1 {
		t.Fail()
	}

	t.Log("Test insert new Value")
	setList.Insert(2)
	if setList.Size() != 2 {
		t.Fail()
	}

	t.Log("Test array generator")
	array := setList.AsSlice()

	if len(array) != 2 {
		t.Fail()
	}

	t.Log("Test remove value")
	setList.Remove(1)
	if setList.Size() != 1 {
		t.Fail()
	}

	t.Log("Test Has function false return ")
	if setList.Has(1) != false {
		t.Fail()
	}

	t.Log("Test Has function with true return")
	if setList.Has(2) != true {
		t.Fail()
	}
}

func TestSetListWithStruct(t *testing.T) {
	setList := New[Test]()
	t.Log("Test Initial value")
	if setList.Size() != 0 {
		t.Fail()
	}

	t.Log("Test insertion one value")
	setList.Insert(Test{value: 1})
	if setList.Size() != 1 {
		t.Fail()
	}

	t.Log("Test unique value")
	setList.Insert(Test{value: 1})
	if setList.Size() != 1 {
		t.Fail()
	}

	t.Log("Test insert new Value")
	setList.Insert(Test{value: 2})
	if setList.Size() != 2 {
		t.Fail()
	}

	t.Log("Test array generator")
	array := setList.AsSlice()

	if len(array) != 2 {
		t.Fail()
	}

	t.Log("Test remove value")
	setList.Remove(Test{value: 1})
	if setList.Size() != 1 {
		t.Fail()
	}

	t.Log("Test Has function false return ")
	if setList.Has(Test{value: 1}) != false {
		t.Fail()
	}

	t.Log("Test Has function with true return")
	if setList.Has(Test{value: 2}) != true {
		t.Fail()
	}
}

func TestInsetSlice(t *testing.T) {
	testStruct := []Test{{value: 1}, {value: 2}, {value: 3}}
	setList := New[Test]()
	setList.InsertSlice(testStruct)
	t.Log("Test Size Structure")
	if setList.Size() != 3 {
		t.Fail()
	}

	sliceStructure := setList.AsSlice()
	t.Log("Test Size array")
	if len(sliceStructure) != 3 {
		t.Fail()
	}

}
