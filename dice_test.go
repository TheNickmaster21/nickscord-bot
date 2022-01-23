package main

import "testing"

func TestInvalidParseDiceArguments(t *testing.T) {
	_, err := ParseDiceArguments("3d12 q")
	if err == nil {
		t.Error("Expected failure for 1")
	}
}

func TestParseDiceArguments(t *testing.T) {
	results, err := ParseDiceArguments(" d6 + 2d4+d12")
	if err != nil {
		t.Errorf("Error not expected: %v", err)
	}
	length := len(results)
	if length != 4 {
		t.Errorf("Expected 4 results, got %v instead", length)
	}
	if results[0] != 6 {
		t.Errorf("Expected first die to be 6, got %v instead", length)
	}
	if results[1] != 4 {
		t.Errorf("Expected second die to be 4, got %v instead", length)
	}
	if results[2] != 4 {
		t.Errorf("Expected third die to be 4, got %v instead", length)
	}
	if results[3] != 12 {
		t.Errorf("Expected forth die to be 12, got %v instead", length)
	}
}

func TestInvalidRollDice(t *testing.T) {
	_, err := RollDice([]int{22, -3})
	if err == nil {
		t.Error("Expected failure for -3")
	}
}

func TestRollDice(t *testing.T) {
	results, err := RollDice([]int{22, 3, 6})
	if err != nil {
		t.Errorf("Error not expected: %v", err)
	}
	length := len(results)
	if length != 3 {
		t.Errorf("Expected 3 results, got %v instead", length)
	}
}

func TestInvalidRoll(t *testing.T) {
	_, err := Roll(-1)
	if err == nil {
		t.Error("Expected failure for -1")
	}
}

func TestOneRoll(t *testing.T) {
	res, err := Roll(1)
	if err != nil {
		t.Errorf("Error not expected: %v", err)
	}
	if res.Result != 1 {
		t.Errorf("Expected result of 1; got %v instead", res.Result)
	}
	if !res.Bold {
		t.Errorf("Expected result to be bold; got %v instead", res.Bold)
	}
}

func TestBigRoll(t *testing.T) {
	res, err := Roll(100)
	if err != nil {
		t.Errorf("Error not expected: %v", err)
	}
	if res.Result < 1 || res.Result > 100 {
		t.Errorf("Expected result between 1 and 100; got %v instead", res.Result)
	}
	if res.Bold != (res.Result == 100) {
		t.Errorf("Wrong behavior for bold; got %v", res.Bold)
	}
}
