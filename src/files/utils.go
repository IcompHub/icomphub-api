package files

func contains(slice []string, val string) bool {
	for _, t := range slice {
		if t == val {
			return true
		}
	}
	return false
}
