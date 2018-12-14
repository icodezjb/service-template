package util

import "context"

func GetContextStringValue(ctx context.Context, keyName string) string {
	if value, ok := ctx.Value(keyName).(string); ok {
		return value
	}

	return ""
}

func GetContextInt64Value(ctx context.Context, keyName string) int64 {
	if value, ok := ctx.Value(keyName).(int64); ok {
		return value
	}

	return 0
}
