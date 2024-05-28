package lib

func InvertMap[K comparable, V comparable](m map[K]V) map[V]K {
	out := make(map[V]K)
	for k, v := range m {
		out[v] = k
	}
	return out
}
