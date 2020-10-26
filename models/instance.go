package models

// Instance has information about a single instance given a host
type Instance struct {
	ID                      string `json:"id"`
	CaughtAt                string `json:"caughtAt"`
	Host                    string `json:"host"`
	UserCount               int    `json:"userCount"`
	NotesCount              int    `json:"notesCount"`
	FollowingCount          int    `json:"followingCount"`
	FollowersCount          int    `json:"followersCount"`
	DriveUsage              int64  `json:"driveUsage"`
	DriveFiles              int64  `json:"driveFiles"`
	LatestRequestSentAt     string `json:"latestRequestSentAt"`
	LatestStatus            int    `json:"latestStatus"`
	LatestRequestReceivedAt string `json:"latestRequestReceivedAt"`
	LastCommunicatedAt      string `json:"lastCommunicatedAt"`
	IsNotResponding         bool   `json:"isNotResponding"`
	IsSuspended             bool   `json:"isSuspended"`
	SoftwareName            string `json:"softwareName"`
	SoftwareVersion         string `json:"softwareVersion"`
	OpenRegistrations       bool   `json:"openRegistrations"`
	Name                    string `json:"name"`
	Description             string `json:"description"`
	MaintainerName          string `json:"maintainerName"`
	MaintainerEmail         string `json:"maintainerEmail"`
	IconURL                 string `json:"iconUrl"`
	InfoUpdatedAt           string `json:"infoUpdatedAt"`
}
