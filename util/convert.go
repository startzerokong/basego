package util

import "strconv"

func IntToString(key int) string {
	value := strconv.Itoa(key)
	return value
}

func StringToInt(key string) (int, error) {
	value, err := strconv.Atoi(key)
	if err != nil {
		return 0, err
	}
	return value, nil
}

func Int64ToString(key int64) string {
	value := strconv.FormatInt(key,10)
	return value
}

func StringToInt64(key string) (int64, error) {
	value, err := strconv.ParseInt(key, 10, 64)
	if err != nil {
		return 0, err
	}
	return value, nil
}

