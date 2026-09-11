package logic

import (
	"reflect"
	"testing"
)

func TestRankAndPagePostIDs(t *testing.T) {
	ids := []int64{1, 2, 3, 4, 5}
	scores := map[int64]float64{1: 10, 2: 30, 3: 20, 4: 30}

	got := rankAndPagePostIDs(ids, scores, 1, 3)
	want := []int64{4, 2, 3}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("rankAndPagePostIDs() = %v, want %v", got, want)
	}

	got = rankAndPagePostIDs(ids, scores, 2, 3)
	want = []int64{1}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("rankAndPagePostIDs() second page = %v, want %v", got, want)
	}
}
