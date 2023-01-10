package utils

import "context"

const userIDKey = "userID"

func UserIDToContext(ctx context.Context, userID interface{}) context.Context {
	return context.WithValue(ctx, userIDKey, userID)
}

func GetUserID(ctx context.Context) string {
	userID := ctx.Value(userIDKey)
	userIDParsed, ok := userID.(string)
	if ok {
		return userIDParsed
	}

	return ""
}
