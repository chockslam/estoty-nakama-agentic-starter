package main

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/heroiclabs/nakama-common/runtime"
)

func TestDecodeMetadataObject(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		raw     string
		want    map[string]any
		wantErr error
	}{
		{
			name: "accepts object",
			raw:  `{"favoriteHero":"warrior","tutorialCompleted":true}`,
			want: map[string]any{
				"favoriteHero":      "warrior",
				"tutorialCompleted": true,
			},
		},
		{
			name:    "rejects invalid json",
			raw:     `{"favoriteHero":`,
			wantErr: errMetadataInvalidJSON,
		},
		{
			name:    "rejects array",
			raw:     `[]`,
			wantErr: errMetadataObjectRequired,
		},
		{
			name:    "rejects string",
			raw:     `"hello"`,
			wantErr: errMetadataObjectRequired,
		},
		{
			name:    "rejects number",
			raw:     `42`,
			wantErr: errMetadataObjectRequired,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := decodeMetadataObject(tt.raw)
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("decodeMetadataObject() error = %v, want %v", err, tt.wantErr)
				}
				return
			}

			if err != nil {
				t.Fatalf("decodeMetadataObject() unexpected error = %v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("decodeMetadataObject() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestMergeMetadata(t *testing.T) {
	t.Parallel()

	existing := map[string]any{
		"level": 1,
		"profile": map[string]any{
			"title": "rookie",
		},
	}
	incoming := map[string]any{
		"level": 2,
		"hat":   "space_helmet",
	}

	got := mergeMetadata(existing, incoming)
	want := map[string]any{
		"level": 2,
		"profile": map[string]any{
			"title": "rookie",
		},
		"hat": "space_helmet",
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("mergeMetadata() = %#v, want %#v", got, want)
	}

	if existing["level"] != 1 {
		t.Fatalf("mergeMetadata() mutated existing map: got level %v, want 1", existing["level"])
	}
	if !reflect.DeepEqual(existing["profile"], map[string]any{"title": "rookie"}) {
		t.Fatalf("mergeMetadata() mutated nested existing data: %#v", existing["profile"])
	}
}

func TestPrepareMetadataUpdate(t *testing.T) {
	t.Parallel()

	ctxWithUserID := context.WithValue(context.Background(), runtime.RUNTIME_CTX_USER_ID, "user-123")

	tests := []struct {
		name         string
		ctx          context.Context
		payload      string
		wantUserID   string
		wantMetadata map[string]any
		wantErr      error
	}{
		{
			name:         "accepts caller user id and payload object",
			ctx:          ctxWithUserID,
			payload:      `{"favoriteHero":"warrior","tutorialCompleted":true}`,
			wantUserID:   "user-123",
			wantMetadata: map[string]any{"favoriteHero": "warrior", "tutorialCompleted": true},
		},
		{
			name:    "rejects missing user id",
			ctx:     context.Background(),
			payload: `{"favoriteHero":"warrior"}`,
			wantErr: errAuthenticatedUserRequired,
		},
		{
			name:    "rejects invalid json",
			ctx:     ctxWithUserID,
			payload: `{"favoriteHero":`,
			wantErr: errMetadataInvalidJSON,
		},
		{
			name:    "rejects non-object json",
			ctx:     ctxWithUserID,
			payload: `["warrior"]`,
			wantErr: errMetadataObjectRequired,
		},
		{
			name:         "payload cannot override target user id",
			ctx:          ctxWithUserID,
			payload:      `{"userId":"user-999","favoriteHero":"warrior"}`,
			wantUserID:   "user-123",
			wantMetadata: map[string]any{"userId": "user-999", "favoriteHero": "warrior"},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			gotUserID, gotMetadata, err := prepareMetadataUpdate(tt.ctx, tt.payload)
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("prepareMetadataUpdate() error = %v, want %v", err, tt.wantErr)
				}
				return
			}

			if err != nil {
				t.Fatalf("prepareMetadataUpdate() unexpected error = %v", err)
			}
			if gotUserID != tt.wantUserID {
				t.Fatalf("prepareMetadataUpdate() userID = %q, want %q", gotUserID, tt.wantUserID)
			}
			if !reflect.DeepEqual(gotMetadata, tt.wantMetadata) {
				t.Fatalf("prepareMetadataUpdate() metadata = %#v, want %#v", gotMetadata, tt.wantMetadata)
			}
		})
	}
}
