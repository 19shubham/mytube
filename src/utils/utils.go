package utils

import "sort"

//for float value
type PairInt struct {
	Key   string
	Value int
}

type PairListInt []PairInt

func (p PairListInt) Len() int           { return len(p) }
func (p PairListInt) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairListInt) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairListInt) DescendingOrder() []string {
	keyList := make([]string, 0, len(p))
	for i := len(p) - 1; i >= 0; i-- {
		keyList = append(keyList, p[i].Key)
	}
	return keyList
}
func (p PairListInt) AscendingOrder() []string {
	keyList := make([]string, 0, len(p))
	for _, pair := range p {
		keyList = append(keyList, pair.Key)
	}
	return keyList
}

// SortMapByIntValue - for sorting map by integer value
func SortMapByIntValue(inputMap map[string]int) PairListInt {
	pairList := make(PairListInt, 0, len(inputMap))
	for k, v := range inputMap {
		pairList = append(pairList, PairInt{k, v})
	}
	sort.Sort(pairList)
	return pairList
}
