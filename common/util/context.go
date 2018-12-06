package util

import "context"

func GetContextStringValue(ctx context.Context, keyName string) string {
	if value, ok := ctx.Value(keyName).(string); ok {
		return value
	}

	return ""
}
