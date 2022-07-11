package hackerrank

import (
	"testing"
)

func TestAngryProfessor(t *testing.T) {
	if angryProfessor(3, []int32{-1, -3, 4, 2}) != "YES" {
		t.Error("Expected YES, got ", angryProfessor(3, []int32{-1, -3, 4, 2}))
	}
	if angryProfessor(2, []int32{-0, -1, 2, 1}) != "NO" {
		t.Error("Expected NO, got ", angryProfessor(3, []int32{-1, -3, 4, 2}))
	}
	if angryProfessor(4, []int32{-93, -86, 49, -62, -90, -63, 40, 72, 11, 67}) != "NO" {
		t.Error("Expected NO, got ", angryProfessor(3, []int32{-1, -3, 4, 2}))
	}
	if angryProfessor(10, []int32{23, -35, -2, 58, -67, -56, -42, -73, -19, 37}) != "YES" {
		t.Error("Expected NO, got ", angryProfessor(3, []int32{-1, -3, 4, 2}))
	}
	if angryProfessor(9, []int32{13, 91, 56, -62, 96, -5, -84, -36, -46, -13}) != "YES" {
		t.Error("Expected NO, got ", angryProfessor(3, []int32{-1, -3, 4, 2}))
	}
	if angryProfessor(7, []int32{26, 94, -95, 34, 67, -97, 17, 52, 1, 86}) != "YES" {
		t.Error("Expected NO, got ", angryProfessor(3, []int32{-1, -3, 4, 2}))
	}
	if angryProfessor(2, []int32{19, 29, -17, -71, -75, -27, -56, -53, 65, 83}) != "NO" {
		t.Error("Expected NO, got ", angryProfessor(3, []int32{-1, -3, 4, 2}))
	}
	if angryProfessor(10, []int32{-32, -3, -70, 8, -40, -96, -18, -46, -21, -79}) != "YES" {
		t.Error("Expected NO, got ", angryProfessor(3, []int32{-1, -3, 4, 2}))
	}
	if angryProfessor(9, []int32{-50, 0, 64, 14, -56, -91, -65, -36, 51, -28}) != "YES" {
		t.Error("Expected NO, got ", angryProfessor(3, []int32{-1, -3, 4, 2}))
	}
	if angryProfessor(6, []int32{-58, -29, -35, -18, 43, -28, -76, 43, -13, 6}) != "NO" {
		t.Error("Expected NO, got ", angryProfessor(3, []int32{-1, -3, 4, 2}))
	}
	if angryProfessor(1, []int32{88, -17, -96, 43, 83, 99, 25, 90, -39, 86}) != "NO" {
		t.Error("Expected NO, got ", angryProfessor(3, []int32{-1, -3, 4, 2}))
	}
}
