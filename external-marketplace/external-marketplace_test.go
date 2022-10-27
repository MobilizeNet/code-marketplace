package externalmarketplace_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/coder/code-marketplace/src/api"
	"github.com/coder/code-marketplace/src/database"
	externalmarketplace "github.com/coder/code-marketplace/src/external-marketplace"
	"github.com/stretchr/testify/require"
)

func TestExternalMarketplace(t *testing.T) {

	extMarketplace := &externalmarketplace.ExternalMarketplace{
		BaseUrl: "https://open-vsx.org/",
	}

	cases := []struct {
		Suite      string
		Name       string
		Url        string
		HttpMethod string
		Request    any
		Response   any
		Status     int
	}{
		{
			Suite:      "GetExtensions",
			Name:       "NoFilters",
			Url:        "vscode/gallery/extensionquery",
			Status:     http.StatusOK,
			HttpMethod: http.MethodPost,
			Request: &api.QueryRequest{
				Filters: []database.Filter{},
			},
			Response: &api.QueryResponse{
				Results: []api.QueryResult{{
					Extensions: []*database.Extension{},
					Metadata: []api.ResultMetadata{{
						Type: "ResultCount",
						Items: []api.ResultMetadataItem{{
							Count: 0,
							Name:  "TotalCount",
						}},
					}},
				}},
			},
		},
		{
			Suite:      "GetExtensions",
			Name:       "MethodNotAllowed",
			Url:        "vscode/gallery/extensionquery",
			Status:     http.StatusMethodNotAllowed,
			HttpMethod: http.MethodGet,
			Request: &api.QueryRequest{
				Filters: []database.Filter{},
			},
			Response: &api.QueryResponse{
				Results: []api.QueryResult{{
					Extensions: []*database.Extension{},
					Metadata: []api.ResultMetadata{{
						Type: "ResultCount",
						Items: []api.ResultMetadataItem{{
							Count: 0,
							Name:  "TotalCount",
						}},
					}},
				}},
			},
		},
		{
			Suite:      "GetExtensions",
			Name:       "BadRequest",
			Url:        "vscode/gallery/extensionquery",
			Status:     http.StatusBadRequest,
			HttpMethod: http.MethodPost,
			Request:    nil,
			Response: &api.QueryResponse{
				Results: []api.QueryResult{{
					Extensions: []*database.Extension{},
					Metadata: []api.ResultMetadata{{
						Type: "ResultCount",
						Items: []api.ResultMetadataItem{{
							Count: 0,
							Name:  "TotalCount",
						}},
					}},
				}},
			},
		},
		{
			Suite:      "GetExtensions",
			Name:       "BadUrl",
			Url:        "vscodegallery/extensionquery/wrong_url",
			Status:     http.StatusForbidden,
			HttpMethod: http.MethodPost,
			Request: &api.QueryRequest{
				Filters: []database.Filter{},
			},
			Response: &api.QueryResponse{
				Results: []api.QueryResult{{
					Extensions: []*database.Extension{},
					Metadata: []api.ResultMetadata{{
						Type: "ResultCount",
						Items: []api.ResultMetadataItem{{
							Count: 0,
							Name:  "TotalCount",
						}},
					}},
				}},
			},
		},
	}

	for _, c := range cases {
		c := c
		reqBody, reqError := json.Marshal(c.Request)
		bodyReader := bytes.NewReader(reqBody)
		if reqError == nil {
			status, response := extMarketplace.HttpRequest(c.HttpMethod, c.Url, bodyReader)
			require.Equal(t, c.Status, status)
			if status == http.StatusOK {
				var extMarketplaceResponse *api.QueryResponse
				err := json.Unmarshal([]byte(response), &extMarketplaceResponse)
				require.NoError(t, err)
				if err == nil {
					require.Greater(t,
						len(extMarketplaceResponse.Results),
						0,
					)
					require.Greater(t,
						len(extMarketplaceResponse.Results[0].Extensions),
						0,
					)
				}
			}
		}

	}
}
