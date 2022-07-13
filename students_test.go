package coverage

import (
	"os"
	"testing"
	"time"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW
var people People
func TestMain(m *testing.M) {
	people = People{
		{"2", "3", time.Now()},
		{"4", "5", time.Now().Add(time.Hour)},
		{"6", "7", time.Now().Add(time.Hour * 2)},
	}
	m.Run()
}

func TestLen(t *testing.T) {
	if people.Len() != 3 {
		t.Errorf("want %d, got %d", 1, people.Len())
	}
}

func TestSwap(t *testing.T) {
	p0 := people[0]
	p1 := people[1]
	p2 := people[2]
	people.Swap(0, 2)
	if p0 != people[2] || p1 != people[1] || p2 != people[0] {
		t.Errorf("Swap failed")
	}
	people.Swap(0, 2)
}

func TestLess(t *testing.T) {
	if people.Less(0, 1) {
		t.Errorf("want false, got true")
	}
	buf1 := people[1].birthDay
	people[1].birthDay = people[0].birthDay
	if people.Less(1, 0) {
		t.Errorf("want false, got true")
	}
	buf2 := people[1].firstName
	people[1].firstName = people[0].firstName
	if people.Less(1, 0) {
		t.Errorf("want false, got true")
	}
	people[1].firstName = buf2
	people[1].birthDay = buf1
}

func TestMatrix(t *testing.T) {
	res, err := New("test string")
	if err == nil || res != nil {
		t.Errorf("error missing")
	}
	res, err = New("1 2 3\n1 2")
	if err == nil || res != nil {
		t.Errorf("error %s", err)
	}
	res, err = New("1 2 3")
	if err != nil || res == nil {
		t.Errorf("error %s", err)
	}
	rows := res.Rows()
	if rows[0][0] != 1 || rows[0][1] != 2 || rows[0][2] != 3 || len(rows) != 1 {
		t.Errorf("wrong rows")
	}
	cols := res.Cols()
	if cols[0][0] != 1 || cols[1][0] != 2 || cols[2][0] != 3 || len(cols) != 3 {
		t.Errorf("wrong cols")
	}
	if !res.Set(0, 0, 4) {
		t.Errorf("want true, got false")
	}
	if res.Set(4, 0, 4) {
		t.Errorf("want false, got true")
	}
}
