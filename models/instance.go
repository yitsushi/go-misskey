package models

// Instance has information about a single instance given a host.
type Instance struct {
	ID                      string `json:"id"`
	CaughtAt                string `json:"caughtAt"`
	Host                    string `json:"host"`
	UserCount               int    `json:"userCount"`
	NotesCount              int    `json:"notesCount"`
	FollowingCount          int    `json:"followingCount"`
	FollowersCount          int    `json:"followersCount"`
	LatestStatus            int    `json:"latestStatus"`
	DriveUsage              int64  `json:"driveUsage"`
	DriveFiles              int64  `json:"driveFiles"`
	LatestRequestSentAt     string `json:"latestRequestSentAt"`
	LatestRequestReceivedAt string `json:"latestRequestReceivedAt"`
	LastCommunicatedAt      string `json:"lastCommunicatedAt"`
	SoftwareName            string `json:"softwareName"`
	SoftwareVersion         string `json:"softwareVersion"`
	Name                    string `json:"name"`
	Description             string `json:"description"`
	MaintainerName          string `json:"maintainerName"`
	MaintainerEmail         string `json:"maintainerEmail"`
	IconURL                 string `json:"iconUrl"`
	FaviconURL              string `json:"faviconUrl"`
	InfoUpdatedAt           string `json:"infoUpdatedAt"`
	ThemeColor              string `json:"86b300"`
	IsNotResponding         bool   `json:"isNotResponding"`
	IsSuspended             bool   `json:"isSuspended"`
	OpenRegistrations       bool   `json:"openRegistrations"`
}

// ServerInfo has server information about a single instance.
type ServerInfo struct {
	Machine string `json:"machine"`
	OS      string `json:"os"`
	Node    string `json:"node"`
	PSQL    string `json:"psql"`
	CPU     struct {
		Model string  `json:"model"`
		Cores float64 `json:"cores"`
	} `json:"cpu"`
	Mem struct {
		Total float64 `json:"total"`
	} `json:"mem"`
	FS struct {
		Total float64 `json:"total"`
		Used  float64 `json:"used"`
	} `json:"fs"`
	Net struct {
		Interface string `json:"interface"`
	} `json:"net"`
}

// TableStats has server instance information about the database tables.
type TableStats struct {
	Migrations struct {
		Count int `json:"count"`
		Size  int `json:"size"`
	} `json:"migrations"`
}

// TableStats has server instance information about the database tables.
type IndexStats []struct {
	Schemaname string  `json:"schemaname"`
	Tablename  string  `json:"tablename"`
	Indexname  string  `json:"indexname"`
	Tablespace *string `json:"tablespace"`
	Indexdef   string  `json:"indexdef"`
}
