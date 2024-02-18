package util

func DefaultIfNull[T any](value *T, defaultValue T) T {
	if value == nil {
		return defaultValue
	}

	return *value
}

func NilIfStringEmpty(value *string) *string {
	if *value == "" {
		return nil
	}

	return value
}
