package main

import (
	"context"
	"testing"

	"github.com/heroiclabs/nakama-common/runtime"
)

func TestIsPrivateRPCAllowed(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		userID string
		want   bool
	}{
		{
			name:   "empty user id allowed",
			userID: "",
			want:   true,
		},
		{
			name:   "whitespace user id allowed",
			userID: "   ",
			want:   true,
		},
		{
			name:   "non-empty user id rejected",
			userID: "user-123",
			want:   false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := isPrivateRPCAllowed(tt.userID); got != tt.want {
				t.Fatalf("isPrivateRPCAllowed(%q) = %t, want %t", tt.userID, got, tt.want)
			}
		})
	}
}

func TestRuntimeUserID(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		ctx  context.Context
		want string
	}{
		{
			name: "nil context",
			ctx:  nil,
			want: "",
		},
		{
			name: "context without user id",
			ctx:  context.Background(),
			want: "",
		},
		{
			name: "non-string user id",
			ctx:  context.WithValue(context.Background(), runtime.RUNTIME_CTX_USER_ID, 123),
			want: "",
		},
		{
			name: "whitespace user id",
			ctx:  context.WithValue(context.Background(), runtime.RUNTIME_CTX_USER_ID, "   "),
			want: "",
		},
		{
			name: "normal user id",
			ctx:  context.WithValue(context.Background(), runtime.RUNTIME_CTX_USER_ID, "user-123"),
			want: "user-123",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := runtimeUserID(tt.ctx); got != tt.want {
				t.Fatalf("runtimeUserID() = %q, want %q", got, tt.want)
			}
		})
	}
}
