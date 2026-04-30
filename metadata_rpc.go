package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"strings"

	"github.com/heroiclabs/nakama-common/runtime"
)

const updateUserMetadataRPCID = "update_user_metadata"

var (
	errAuthenticatedUserRequired = errors.New("authenticated user session required")
	errMetadataInvalidJSON       = errors.New("metadata payload must be valid JSON")
	errMetadataObjectRequired    = errors.New("metadata payload must be a JSON object")
	errMetadataReservedKey       = errors.New("metadata payload contains a reserved key")
	errMetadataLoadFailed        = errors.New("unable to load account metadata")
	errMetadataUpdateFailed      = errors.New("unable to update account metadata")
	errMetadataResponseFailed    = errors.New("unable to marshal metadata response")
)

type metadataUpdateSuccessResponse struct {
	Success  bool           `json:"success"`
	Metadata map[string]any `json:"metadata"`
}

func updateUserMetadataRPC(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {
	userID, incomingMetadata, err := prepareMetadataUpdate(ctx, payload)
	if err != nil {
		return "", metadataRPCError(err)
	}

	account, err := nk.AccountGetId(ctx, userID)
	if err != nil {
		return "", runtime.NewError(errMetadataLoadFailed.Error(), grpcCodeInternal)
	}

	if account == nil || account.GetUser() == nil {
		return "", runtime.NewError(errMetadataLoadFailed.Error(), grpcCodeInternal)
	}

	existingMetadata, err := decodeExistingMetadata(account.GetUser().GetMetadata())
	if err != nil {
		return "", runtime.NewError(errMetadataLoadFailed.Error(), grpcCodeInternal)
	}

	updatedMetadata := mergeMetadata(existingMetadata, incomingMetadata)
	// Empty profile fields are no-ops in AccountUpdateId, so this only updates metadata.
	if err := nk.AccountUpdateId(ctx, userID, "", updatedMetadata, "", "", "", "", ""); err != nil {
		return "", runtime.NewError(errMetadataUpdateFailed.Error(), grpcCodeInternal)
	}

	response, err := json.Marshal(metadataUpdateSuccessResponse{
		Success:  true,
		Metadata: updatedMetadata,
	})
	if err != nil {
		return "", runtime.NewError(errMetadataResponseFailed.Error(), grpcCodeInternal)
	}

	return string(response), nil
}

func prepareMetadataUpdate(ctx context.Context, payload string) (string, map[string]any, error) {
	userID, err := callerUserID(ctx)
	if err != nil {
		return "", nil, err
	}

	incomingMetadata, err := decodeMetadataObject(payload)
	if err != nil {
		return "", nil, err
	}
	if err := validateIncomingMetadata(incomingMetadata); err != nil {
		return "", nil, err
	}

	return userID, incomingMetadata, nil
}

func callerUserID(ctx context.Context) (string, error) {
	if ctx == nil {
		return "", errAuthenticatedUserRequired
	}

	userID, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
	if !ok || strings.TrimSpace(userID) == "" {
		return "", errAuthenticatedUserRequired
	}

	return userID, nil
}

func decodeMetadataObject(raw string) (map[string]any, error) {
	var value any
	if err := json.Unmarshal([]byte(raw), &value); err != nil {
		return nil, errMetadataInvalidJSON
	}

	object, ok := value.(map[string]any)
	if !ok {
		return nil, errMetadataObjectRequired
	}

	return object, nil
}

func validateIncomingMetadata(metadata map[string]any) error {
	for key := range metadata {
		switch key {
		case "userId", "username", "id", "user_id":
			return errMetadataReservedKey
		}
	}

	return nil
}

func decodeExistingMetadata(raw string) (map[string]any, error) {
	if strings.TrimSpace(raw) == "" {
		return map[string]any{}, nil
	}

	return decodeMetadataObject(raw)
}

// mergeMetadata returns a shallow merged copy where incoming top-level keys
// overwrite existing keys. Nested objects are not deep-merged or deep-copied,
// so callers must not mutate returned nested values if existing must remain unchanged.
func mergeMetadata(existing map[string]any, incoming map[string]any) map[string]any {
	merged := make(map[string]any, len(existing)+len(incoming))

	for key, value := range existing {
		merged[key] = value
	}

	for key, value := range incoming {
		merged[key] = value
	}

	return merged
}

func metadataRPCError(err error) error {
	switch {
	case errors.Is(err, errAuthenticatedUserRequired):
		return runtime.NewError(err.Error(), grpcCodeUnauthenticated)
	case errors.Is(err, errMetadataInvalidJSON), errors.Is(err, errMetadataObjectRequired), errors.Is(err, errMetadataReservedKey):
		return runtime.NewError(err.Error(), grpcCodeInvalidArgument)
	default:
		return runtime.NewError(err.Error(), grpcCodeInternal)
	}
}
