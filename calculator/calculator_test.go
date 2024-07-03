package calculator

import "testing"

func TestAdd(t *testing.T){
	result := Add(2,3)
	expected := 5

	if result != expected {
		t.Errorf("Add(2,3)=%d,want%d",result,expected)
	}
}

func TestSub(t *testing.T){
	result := Sub(9,2)
	expected := 7

	if result != expected{
		t.Errorf("Sub(9,2)=%d,want%d",result,expected)
	}
}

func TestMultplication(t *testing.T){
	result := Multipication(5,5)
	expected := 25

	if result != expected{
		t.Errorf("Multipication(5,5)=%d,want%d",result,expected)
	}
}

func TestDivision(t *testing.T){
	result := Division(25,5)
	expected := 5

	if result != expected {
		t.Errorf("Division(25,5)=%d,want%d",result,expected)
	}
}