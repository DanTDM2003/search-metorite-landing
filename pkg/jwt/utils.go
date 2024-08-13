package jwt

import (
	"context"

	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
)

type PayloadCtxKey struct{}
type UserCtxKey struct{}

func SetPayloadToContext(ctx context.Context, payload Payload) context.Context {
	return context.WithValue(ctx, PayloadCtxKey{}, payload)
}

func GetPayloadFromContext(ctx context.Context) (Payload, bool) {
	payload, ok := ctx.Value(PayloadCtxKey{}).(Payload)
	return payload, ok
}

func SetUserToContext(ctx context.Context, user models.User) context.Context {
	return context.WithValue(ctx, UserCtxKey{}, user)
}

func GetUserFromContext(ctx context.Context) (models.User, bool) {
	user, ok := ctx.Value(UserCtxKey{}).(models.User)
	return user, ok
}
