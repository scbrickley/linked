package linked

import (
	"testing"
	"github.com/google/go-cmp/cmp"
)

const checkMark = '\u2713'
const ballotX = '\u2717'

func TestCreateAndPush(t *testing.T) {
	list := EmptyList()

	list.PushNode(Node {3, nil})
	t.Log("list should have a value of 3")
	if list.value == 3 {
		t.Logf("list.value == %d %c", list.value, checkMark)
	} else {
		t.Errorf("list.value == %d %c", list.value, ballotX)
	}
}

func TestLengthgthAndIndex(t *testing.T) {
	list := EmptyList()
	t.Log("Length of list should equal 0")
	if list.Length() == 0 {
		t.Logf("Got: %v %c", list.Length(), checkMark)
	} else {
		t.Errorf("Got: %v %c", list.Length(), ballotX)
	}

	list.Push(23)
	t.Log("Length of list should equal 1")
	if list.Length() == 1 {
		t.Logf("Got: %v %c", list.Length(), checkMark)
	} else {
		t.Errorf("Got: %v %c", list.Length(), ballotX)
	}

	list.Push(75)
	list.Push(67)
	list.Push(33)
	list.Push(78)
	list.Push(45)
	t.Log("Length of list should equal 6")
	if list.Length() == 6 {
		t.Logf("Got: %v %c", list.Length(), checkMark)
	} else {
		t.Errorf("Got: %v %c", list.Length(), ballotX)
	}

	sixth := list.Index(5)
	second := list.Index(1)

	t.Log("The sixth element of the list should be 45")
	if sixth == 45 {
		t.Logf("Got: %v %c", sixth, checkMark)
	} else {
		t.Errorf("Got: %v %c", sixth, ballotX)
	}

	t.Log("The second element of the list should be 75")
	if second == 75 {
		t.Logf("Got: %v %c", second, checkMark)
	} else {
		t.Errorf("Got: %v %c", second, ballotX)
	}
}

func TestNewList(t *testing.T) {
	list := NewList(1, 2, 3, 4, 5, 6, "seven")

	t.Log("List length should be 7")
	if list.Length() == 7 {
		t.Logf("Got: %v %c", list.Length(), checkMark)
	} else {
		t.Errorf("Got: %v %c", list.Length(), ballotX)
	}
}

func TestToSlice(t *testing.T) {
	list := NewList(1, 2, 3, 4, 5)
	slice := []interface{}{1, 2, 3, 4, 5}

	t.Log("NewList(1, 2, 3, 4, 5).ToSlice() should equal [1 2 3 4 5]")
	listToSlice := list.ToSlice()
	if cmp.Equal(slice, listToSlice) {
		t.Logf("Got: %v %c", listToSlice, checkMark)
	} else {
		t.Errorf("Got: %v %c", listToSlice, ballotX)
	}
}

func TestReverse(t *testing.T) {
	list := NewList(1, 2, 3, 4)

	t.Logf("list == %v", list.ToSlice())
	t.Log("Calling list.Reverse()")
	list.Reverse()

	t.Log("list[0] should equal 4")
	if list.Index(0) == 4 {
		t.Logf("Got: %d %c", list.Index(0), checkMark)
	} else {
		t.Errorf("Got: %d %c", list.Index(0), ballotX)
	}

	t.Log("list[1] should equal 3")
	if list.Index(1) == 3 {
		t.Logf("Got: %d %c", list.Index(1), checkMark)
	} else {
		t.Errorf("Got: %d %c", list.Index(1), ballotX)
	}

	t.Log("list[2] should equal 2")
	if list.Index(2) == 2 {
		t.Logf("Got: %d %c", list.Index(2), checkMark)
	} else {
		t.Errorf("Got: %d %c", list.Index(2), ballotX)
	}

	t.Log("list[3] should equal 1")
	if list.Index(3) == 1 {
		t.Logf("Got: %d %c", list.Index(3), checkMark)
	} else {
		t.Errorf("Got: %d %c", list.Index(3), ballotX)
	}

	t.Log("Length of list should still be 4")
	if list.Length() == 4 {
		t.Logf("Got: %d %c", list.Length(), checkMark)
	} else {
		t.Errorf("Got: %d %c", list.Length(), ballotX)
	}
	t.Logf("list == %v", list.ToSlice())
}
