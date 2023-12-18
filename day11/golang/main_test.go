package main

import "testing"

// Test problemOne on Sample Input
func TestProblemOneWithTestInput(t *testing.T) {
	answer := problemOne("test.txt")
	want := 374
	if answer != want {
		t.Fatalf(`problemOne("test.txt") = %q, %v\n`, want, answer)
	}
}

// Test problemOne on actual Input
func TestProblemOneWithActualInput(t *testing.T) {
	answer := problemOne("input.txt")
	want := 9214785
	if answer != want {
		t.Fatalf(`problemOne("input.txt") = %q, %v\n`, want, answer)
	}
}
func TestProblemTwoWithTestInput(t *testing.T) {
	answer := problemTwo("test.txt")
	want := 82000210
	if answer != want {
		t.Fatalf(`problemTwo("test.txt") = %q, %v\n`, want, answer)
	}
}

// Test problemOne on actual Input
func TestProblemTwoWithActualInput(t *testing.T) {
	answer := problemTwo("input.txt")
	want := 613686987427
	if answer != want {
		t.Fatalf(`problemTwo("input.txt") = %q, %v\n`, want, answer)
	}
}
