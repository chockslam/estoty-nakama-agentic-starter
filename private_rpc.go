package main

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/heroiclabs/nakama-common/runtime"
)

const privateHealthCheckRPCID = "private_health_check"

var errPrivateRPCRequiresServerToServer = errors.New("private_health_check is only callable via server-to-server/runtime HTTP key")

func privateHealthCheckRPC(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {
	if !isPrivateRPCAllowed(runtimeUserID(ctx)) {
		return "", runtime.NewError(errPrivateRPCRequiresServerToServer.Error(), grpcCodePermissionDenied)
	}

	return `{"success":true}`, nil
}

func runtimeUserID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}

	userID, _ := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
	return strings.TrimSpace(userID)
}

func isPrivateRPCAllowed(userID string) bool {
	return strings.TrimSpace(userID) == ""
}
