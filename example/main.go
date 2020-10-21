package main

func main() {
	if false {
		announcements()
		instanceMeta()
		stats()

		antenna()

		driveEndpoints()
	}

	driveMutableFlow()
}

func boolStatusToString(v bool) string {
	if v {
		return "enabled"
	}

	return "disabled"
}
