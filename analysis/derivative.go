package analysis

func Derivative(s *Source) *Source {
	var result Source
	keys := s.Keys()
	if len(keys) > 1 {
		previousKey := keys[0]
		previousValue, _ := s.LastAt(previousKey)

		for _, currentKey := range keys[1:] {
			currentValue, _ := s.FirstAt(currentKey)
			diffValues := currentValue - previousValue
			diffKeys := currentKey - previousKey
			d := diffValues / diffKeys

			result.Add(previousKey, d)
			result.Add(currentKey, d)

			previousKey = currentKey
			previousValue, _ = s.LastAt(previousKey)
		}
	}

	return &result
}
