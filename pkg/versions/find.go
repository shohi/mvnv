package versions

// Find checks whether ver is in the versions.
// TODO: support partial match
func Find(versions []string, ver string) bool {
	if ver == "" {
		return false
	}

	for _, v := range versions {
		if v == ver {
			return true
		}
	}

	return false
}
