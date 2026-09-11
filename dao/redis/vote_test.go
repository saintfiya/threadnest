package redis

import (
	"context"
	"errors"
	"testing"
)

func TestVoteForPostRejectsInvalidValueBeforeRedisCall(t *testing.T) {
	err := VoteForPost(context.Background(), "1", "2", 2)
	if !errors.Is(err, ErrInvalidVote) {
		t.Fatalf("VoteForPost() error = %v, want ErrInvalidVote", err)
	}
}
