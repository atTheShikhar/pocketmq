package core

import (
	"bytes"
	"strconv"
	"testing"
)

func TestQueue(t *testing.T) {
	q := NewQueue[int](5)

	t.Run("enqueue", func(t *testing.T) {
		for i := 1; i < 6; i++ {
			q.Enqueue(i)
		}
		buf := bytes.Buffer{}
		q.Traverse(&buf)
		want := "12345"
		got := buf.String()
		assertValueError(t, got, want)
	})

	t.Run("dequeue", func(t *testing.T) {
		wantPopped := "12"
		wantRemains := "345"

		var gotPopped string
		for i := 0; i < 2; i++ {
			val, _ := q.Dequeue()
			gotPopped += strconv.Itoa(val)
		}

		buf := bytes.Buffer{}
		q.Traverse(&buf)
		gotRemains := buf.String()
		assertValueError(t, wantPopped, gotPopped)
		assertValueError(t, wantRemains, gotRemains)
	})

	t.Run("enqueue error", func(t *testing.T) {
		wantError := "max capacity reached"
		var gotError error
		for i := 6; i < 9; i++ {
			gotError = q.Enqueue(i)
			if gotError != nil {
				break
			}
		}
		if gotError == nil {
			t.Errorf("expected an error, but haven't received one")
		} else {
			assertValueError(t, gotError.Error(), wantError)
		}
	})

	t.Run("dequeue error", func(t *testing.T) {
		wantError := "queue is empty"
		var gotError error
		for i := 1; i < 7; i++ {
			_, gotError = q.Dequeue()
			if gotError != nil {
				break
			}
		}
		if gotError == nil {
			t.Errorf("expected an error, but haven't received one")
		} else {
			assertValueError(t, gotError.Error(), wantError)
		}
	})

}

// func assertQueueError(t testing.TB, err error) {
// 	t.Helper()

// 	if err != nil {
// 		t.Errorf("error: %v", err)
// 	}
// }

func assertValueError(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("expected '%v', got '%v'", got, want)
	}
}
