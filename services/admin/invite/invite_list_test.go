package invite_test

import (
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
	"github.com/yitsushi/go-misskey/services/admin/invite"
	"github.com/yitsushi/go-misskey/test"
)

func TestInviteList(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		limit        int
		offset       int
		statusType   invite.StatusType
		order        invite.SortOrder
		respFileFunc func(_ *testing.T) func(in interface{}) string
		wantErr      assert.ErrorAssertionFunc
	}{
		"short limit": {
			limit:      0,
			offset:     0,
			statusType: invite.StatusTypeAll,
			order:      invite.SortOrderCreatedAtAsc,
			respFileFunc: func(_ *testing.T) func(in interface{}) string {
				return func(_ interface{}) string { return "" }
			},
			wantErr: assert.Error,
		},
		"min limit": {
			limit:      1,
			offset:     0,
			statusType: invite.StatusTypeAll,
			order:      invite.SortOrderCreatedAtAsc,
			respFileFunc: func(t *testing.T) func(in interface{}) string {
				t.Helper()

				return func(in interface{}) string {
					req, ok := in.(*invite.ListRequest)
					require.True(t, ok)

					b, err := json.Marshal(req)
					require.NoError(t, err)
					assert.Equal(t, "{\"limit\":1,\"offset\":0,\"type\":\"all\",\"sort\":\"-createdAt\"}", string(b))

					return "invite-list.json"
				}
			},
			wantErr: assert.NoError,
		},
		"max limit": {
			limit:      100,
			offset:     100,
			statusType: invite.StatusTypeExpired,
			order:      invite.SortOrderCreatedAtDesc,
			respFileFunc: func(t *testing.T) func(in interface{}) string {
				t.Helper()

				return func(in interface{}) string {
					req, ok := in.(*invite.ListRequest)
					require.True(t, ok)

					b, err := json.Marshal(req)
					require.NoError(t, err)
					assert.Equal(t, "{\"limit\":100,\"offset\":100,\"type\":\"expired\",\"sort\":\"+createdAt\"}", string(b))

					return "invite-list.json"
				}
			},
			wantErr: assert.NoError,
		},
		"exceed max limit": {
			limit:      101,
			offset:     0,
			statusType: invite.StatusTypeAll,
			order:      invite.SortOrderCreatedAtAsc,
			respFileFunc: func(_ *testing.T) func(in interface{}) string {
				return func(_ interface{}) string { return "" }
			},
			wantErr: assert.Error,
		},
		"invalid type": {
			limit:      1,
			offset:     0,
			statusType: invite.StatusType("invalid"),
			order:      invite.SortOrderCreatedAtAsc,
			respFileFunc: func(_ *testing.T) func(in interface{}) string {
				return func(_ interface{}) string { return "" }
			},
			wantErr: assert.Error,
		},
		"invalid order": {
			limit:      1,
			offset:     0,
			statusType: invite.StatusTypeAll,
			order:      invite.SortOrder("invalid"),
			respFileFunc: func(_ *testing.T) func(in interface{}) string {
				return func(_ interface{}) string { return "" }
			},
			wantErr: assert.Error,
		},
	}

	for tn, tc := range testCases {
		t.Run(tn, func(t *testing.T) {
			t.Parallel()

			cli := test.MakeMockClient(test.SimpleMockOptions{
				Endpoint:         "/api/admin/invite/list",
				RequestData:      &invite.ListRequest{},
				ResponseFileFunc: tc.respFileFunc(t),
				StatusCode:       http.StatusOK,
			})

			_, err := cli.Admin().Invite().List(tc.limit, tc.offset, tc.statusType, tc.order)
			tc.wantErr(t, err)
		})
	}
}

func TestInviteList_ResponseParse(t *testing.T) {
	t.Parallel()

	jsonFixedTime, err := time.Parse(time.RFC3339, "2019-08-24T14:15:22Z")
	require.NoError(t, err)

	cli := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/admin/invite/list",
		RequestData:  &invite.ListRequest{},
		ResponseFile: "invite-list.json",
		StatusCode:   http.StatusOK,
	})
	response, err := cli.Admin().Invite().List(2, 0, invite.StatusTypeAll, invite.SortOrderCreatedAtAsc)
	require.NoError(t, err)

	assert.EqualValues(t, []*models.Invite{
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
		{
			ID:        "yyyyyyyyyy",
			Code:      "GR6S02ERUA5VR",
			ExpiredAt: jsonFixedTime,
			CreatedAt: jsonFixedTime,
			CreatedBy: &models.User{
				ID:             "yyyyyyyyyy",
				Name:           "藍",
				Username:       "ai",
				Host:           core.NewString("misskey.example.com"),
				AvatarURL:      "string",
				AvatarBlurhash: core.NewString("string"),
				IsBot:          true,
				IsCat:          true,
			},
			UsedBy: &models.User{
				ID:             "yyyyyyyyyy",
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
}
