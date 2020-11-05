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
