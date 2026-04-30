package main

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func InitModule(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	logger.Info("estoty nakama runtime module loaded")
	if err := loadAndCacheGameConfig(); err != nil {
		return err
	}
	if err := initializer.RegisterRpc(updateUserMetadataRPCID, updateUserMetadataRPC); err != nil {
		return err
	}
	if err := initializer.RegisterRpc(getGameConfigRPCID, getGameConfigRPC); err != nil {
		return err
	}
	if err := initializer.RegisterRpc(privateHealthCheckRPCID, privateHealthCheckRPC); err != nil {
		return err
	}
	return nil
}
