package meta

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// InstanceMetaRequest represents an Announcement request.
type InstanceMetaRequest struct {
	Detail bool `json:"detail"`
}

// InstanceMetaResponse represents the response from the meta endpoint.
type InstanceMetaResponse struct {
	MaintainerName               core.String    `json:"maintainerName"`
	MaintainerEmail              core.String    `json:"maintainerEmail"`
	Version                      core.String    `json:"version"`
	Name                         core.String    `json:"name"`
	URI                          core.String    `json:"uri"`
	Description                  core.String    `json:"description"`
	Langs                        []core.String  `json:"langs"`
	ToSURL                       core.String    `json:"tosUrl"`
	RepositoryURL                core.String    `json:"repositoryUrl"`
	FeedbackURL                  core.String    `json:"feedbackUrl"`
	DriveCapacityPerLocalUserMb  uint64         `json:"driveCapacityPerLocalUserMb"`
	DriveCapacityPerRemoteUserMb uint64         `json:"driveCapacityPerRemoteUserMb"`
	HcaptchaSiteKey              core.String    `json:"hcaptchaSiteKey"`
	RecaptchaSiteKey             core.String    `json:"recaptchaSiteKey"`
	SwPublickey                  core.String    `json:"swPublickey"`
	MascotImageURL               core.String    `json:"mascotImageUrl"`
	BannerURL                    core.String    `json:"bannerUrl"`
	ErrorImageURL                core.String    `json:"errorImageUrl"`
	IconURL                      core.String    `json:"iconUrl"`
	MaxNoteTextLength            uint64         `json:"maxNoteTextLength"`
	Emojis                       []models.Emoji `json:"emojis"`
	Features                     Features       `json:"features"`
	RequireSetup                 bool           `json:"requireSetup"`
	EnableEmail                  bool           `json:"enableEmail"`
	EnableTwitterIntegration     bool           `json:"enableTwitterIntegration"`
	EnableGithubIntegration      bool           `json:"enableGithubIntegration"`
	EnableDiscordIntegration     bool           `json:"enableDiscordIntegration"`
	EnableServiceWorker          bool           `json:"enableServiceWorker"`
	EnableRecaptcha              bool           `json:"enableRecaptcha"`
	Secure                       bool           `json:"secure"`
	DisableRegistration          bool           `json:"disableRegistration"`
	DisableLocalTimeline         bool           `json:"disableLocalTimeline"`
	DisableGlobalTimeline        bool           `json:"disableGlobalTimeline"`
	CacheRemoteFiles             bool           `json:"cacheRemoteFiles"`
	ProxyRemoteFiles             bool           `json:"proxyRemoteFiles"`
	EnableHcaptcha               bool           `json:"enableHcaptcha"`
}

// Features lists all available features of the instance and their statuses.
type Features struct {
	Registration   bool `json:"registration"`
	LocalTimeLine  bool `json:"localTimeLine"`
	GlobalTimeLine bool `json:"globalTimeLine"`
	Elasticsearch  bool `json:"elasticsearch"`
	Hcaptcha       bool `json:"hcaptcha"`
	Recaptcha      bool `json:"recaptcha"`
	ObjectStorage  bool `json:"objectStorage"`
	Twitter        bool `json:"twitter"`
	Github         bool `json:"github"`
	Discord        bool `json:"discord"`
	ServiceWorker  bool `json:"serviceWorker"`
	MiAuth         bool `json:"miauth"`
}

// InstanceMeta is the endpoint to get metadata about the instance.
func (s *Service) InstanceMeta(details bool) (InstanceMetaResponse, error) {
	request := &InstanceMetaRequest{
		Detail: details,
	}

	var response InstanceMetaResponse
	err := s.Call(
		&core.BaseRequest{Request: request, Path: "/meta"},
		&response,
	)

	return response, err
}
