package go_cover_support

func IntersectString(x []string, y []string) []string {
	set := make([]string, 0)
	if len(x) == 0 || len(y) == 0 {
		return set
	}

	hash := map[string]struct{}{}

	for _, v := range x {
		hash[v] = struct{}{}
	}

	for _, v := range y {
		_, ok := hash[v]
		if ok {
			set = append(set, v)
		}
	}

	return set
}

func UniqUint16(a []uint16) []uint16 {
	length := len(a)

	seen := make(map[uint16]struct{}, length)
	j := 0

	for i := 0; i < length; i++ {
		v := a[i]

		if _, ok := seen[v]; ok {
			continue
		}

		seen[v] = struct{}{}
		a[j] = v
		j++
	}

	return a[0:j]
}

func MergeUint16(data ...interface{}) []uint16 {
	result := make([]uint16, 0)
	for _, datum := range data {
		v, ok := datum.([]uint16)
		if !ok {
			break
		}
		for _, u := range v {
			result = append(result, u)
		}
	}
	return result
}

func ContainsString(s []string, v string) bool {
	for _, vv := range s {
		if vv == v {
			return true
		}
	}
	return false
}

func ContainsUint16(s []uint16, v uint16) bool {
	for _, vv := range s {
		if vv == v {
			return true
		}
	}
	return false
}

func DifferenceString(x []string, y []string) ([]string, []string) {
	leftSlice := []string{}
	rightSlice := []string{}

	for _, v := range x {
		if ContainsString(y, v) == false {
			leftSlice = append(leftSlice, v)
		}
	}

	for _, v := range y {
		if ContainsString(x, v) == false {
			rightSlice = append(rightSlice, v)
		}
	}

	return leftSlice, rightSlice
}

func DifferenceUint16(x []uint16, y []uint16) []uint16 {
	leftSlice := []uint16{}
	rightSlice := []uint16{}

	for _, v := range x {
		if ContainsUint16(y, v) == false {
			leftSlice = append(leftSlice, v)
		}
	}

	for _, v := range y {
		if ContainsUint16(x, v) == false {
			rightSlice = append(rightSlice, v)
		}
	}

	for _, u := range rightSlice {
		leftSlice = append(leftSlice, u)
	}

	return UniqUint16(leftSlice)
}
