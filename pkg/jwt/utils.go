package jwt

import "context"

type PayloadCtxKey struct{}

func SetPayloadToContext(ctx context.Context, payload Payload) context.Context {
	return context.WithValue(ctx, PayloadCtxKey{}, payload)
}

func GetPayloadFromContext(ctx context.Context) (Payload, bool) {
	payload, ok := ctx.Value(PayloadCtxKey{}).(Payload)
	return payload, ok
}
