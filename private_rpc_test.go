package main

import "testing"

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
