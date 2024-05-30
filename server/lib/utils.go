package lib

func InvertMap[K comparable, V comparable](m map[K]V) map[V]K {
	out := make(map[V]K)
	for k, v := range m {
		out[v] = k
	}
	return out
}

func InterfaceArrayToArray[T any](arr []interface{}) []T {
	out := make([]T, len(arr))
	for i, v := range arr {
		out[i] = v.(T)
	}
	return out
}

func ArrayToInterfaceArray[T any](arr []T) []interface{} {
	out := make([]interface{}, len(arr))
	for i, v := range arr {
		out[i] = v
	}
	return out
}
