package invite_test

import (
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
	"github.com/yitsushi/go-misskey/services/admin/invite"
	"github.com/yitsushi/go-misskey/test"
)

func TestInviteCreate(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		count        int
		expiredAt    time.Time
		respFileFunc func(t *testing.T) func(in interface{}) string
		wantErr      assert.ErrorAssertionFunc
	}{
		"short min": {
			count:     0,
			expiredAt: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			respFileFunc: func(_ *testing.T) func(in interface{}) string {
				return func(_ interface{}) string { return "" }
			},
			wantErr: assert.Error,
		},
		"min": {
			count:     1,
			expiredAt: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			respFileFunc: func(t *testing.T) func(in interface{}) string {
				t.Helper()

				return func(in interface{}) string {
					req, ok := in.(*invite.CreateRequest)
					require.True(t, ok)

					b, err := json.Marshal(req)
					require.NoError(t, err)
					assert.Equal(t, "{\"count\":1,\"expiresAt\":\"2024-01-01T00:00:00Z\"}", string(b))

					return "invite-create.json"
				}
			},
			wantErr: assert.NoError,
		},
		"max": {
			count:     100,
			expiredAt: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			respFileFunc: func(t *testing.T) func(in interface{}) string {
				t.Helper()

				return func(in interface{}) string {
					req, ok := in.(*invite.CreateRequest)
					require.True(t, ok)

					b, err := json.Marshal(req)
					require.NoError(t, err)
					assert.Equal(t, "{\"count\":100,\"expiresAt\":\"2024-01-01T00:00:00Z\"}", string(b))

					return "invite-create.json"
				}
			},
			wantErr: assert.NoError,
		},
		"exceed max": {
			count:     101,
			expiredAt: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			respFileFunc: func(_ *testing.T) func(in interface{}) string {
				return func(_ interface{}) string { return "" }
			},
			wantErr: assert.Error,
		},
		"no expired time": {
			count:     1,
			expiredAt: time.Time{},
			respFileFunc: func(t *testing.T) func(in interface{}) string {
				t.Helper()

				return func(in interface{}) string {
					req, ok := in.(*invite.CreateRequest)
					require.True(t, ok)

					b, err := json.Marshal(req)
					require.NoError(t, err)
					assert.Equal(t, "{\"count\":1}", string(b))

					return "invite-create.json"
				}
			},
			wantErr: assert.NoError,
		},
	}

	for tn, tc := range testCases {
		t.Run(tn, func(t *testing.T) {
			t.Parallel()

			cli := test.MakeMockClient(test.SimpleMockOptions{
				Endpoint:         "/api/admin/invite/create",
				RequestData:      &invite.CreateRequest{},
				ResponseFileFunc: tc.respFileFunc(t),
				StatusCode:       http.StatusOK,
			})

			_, err := cli.Admin().Invite().Create(tc.count, tc.expiredAt)
			tc.wantErr(t, err)
		})
	}

	t.Run("invalid expired time", func(t *testing.T) {
		t.Parallel()

		req := invite.CreateRequest{
			Count:     1,
			ExpiresAt: "invalid",
		}
		assert.Error(t, req.Validate())
	})
}

func TestInviteCreate_ResponseParse(t *testing.T) {
	t.Parallel()

	jsonFixedTime, err := time.Parse(time.RFC3339, "2019-08-24T14:15:22Z")
	require.NoError(t, err)

	cli := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/admin/invite/create",
		RequestData:  &invite.CreateRequest{},
		ResponseFile: "invite-create.json",
		StatusCode:   http.StatusOK,
		Type:         test.JSONMockType,
	})
	response, err := cli.Admin().Invite().Create(1, time.Now())
	require.NoError(t, err)

	diff := cmp.Diff([]*models.Invite{
		{
			ID:        "xxxxxxxxxx",
			Code:      "GR6S02ERUA5VR",
			ExpiredAt: jsonFixedTime,
			CreatedAt: jsonFixedTime,
			CreatedBy: &models.User{
				ID:             "xxxxxxxxxx",
				Name:           "藍",
				Username:       "ai",
				Host:           core.NewString("misskey.example.com"),
				AvatarURL:      "string",
				AvatarBlurhash: core.NewString("string"),
				IsBot:          true,
				IsCat:          true,
			},
			UsedBy: &models.User{
				ID:             "xxxxxxxxxx",
				Name:           "藍",
				Username:       "ai",
				Host:           core.NewString("misskey.example.com"),
				AvatarURL:      "string",
				AvatarBlurhash: core.NewString("string"),
				IsBot:          true,
				IsCat:          true,
			},
			UsedAt: jsonFixedTime,
			Used:   true,
		},
	},
		response,
	)
	assert.Empty(t, diff)
}
