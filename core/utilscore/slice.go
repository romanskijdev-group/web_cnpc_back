package utilscore

func ConvertToStringSlice(slice []*string) []string {
	result := make([]string, len(slice))
	for i, v := range slice {
		if v != nil {
			result[i] = *v
		}
	}
	return result
}
