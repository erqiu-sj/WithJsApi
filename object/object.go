package object

// Keys 返回一个map中所有key 组合成的string切片
func Keys[V any](ownMap map[string]V) []string {
	index := 0
	collection := make([]string, len(ownMap), len(ownMap))
	for k := range ownMap {
		collection[index] = k
		index++
	}
	return collection
}

// Values 返回一个map中所有value 组合成的value切片
func Values[V any](ownMap map[string]V) []V {
	index := 0
	collection := make([]V, len(ownMap), len(ownMap))
	for key := range ownMap {
		collection[index] = ownMap[key]
		index++
	}
	return collection
}
