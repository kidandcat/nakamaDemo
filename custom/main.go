package main

import (
	"context"
	"database/sql"
	"regexp"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
)

var (
	errBadInput           = runtime.NewError("input contained invalid data", 3)
	errGuildAlreadyExists = runtime.NewError("guild name is in use", 6)
)

func InitModule(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	logger.Info("Hello World!")
	initializer.RegisterRpc("create_guild", CreateGuildRpc)
	initializer.RegisterBeforeAuthenticateCustom(BeforeAuthenticateCustom)
	return nil
}

func CreateGuildRpc(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {
	// ... check if a guild already exists and set value of `alreadyExists` accordingly
	var alreadyExists bool = true

	if alreadyExists {
		return "", errGuildAlreadyExists
	}

	return "", nil
}

func BeforeAuthenticateCustom(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AuthenticateCustomRequest) (*api.AuthenticateCustomRequest, error) {
	// Only match custom Id in the format "cid-000000"
	pattern := regexp.MustCompile("^cid-([0-9]{6})$")

	if !pattern.MatchString(in.Account.Id) {
		return nil, errBadInput
	}

	return in, nil
}
