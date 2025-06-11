package github

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/github/github-mcp-server/pkg/translations"
	"github.com/google/go-github/v72/github"
	"github.com/migueleliasweb/go-github-mock/src/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_GetProject(t *testing.T) {
	mockClient := github.NewClient(nil)
	tool, _ := GetProject(stubGetClientFn(mockClient), translations.NullTranslationHelper)
	assert.Equal(t, "get_project", tool.Name)
	assert.NotEmpty(t, tool.Description)
	assert.Contains(t, tool.InputSchema.Properties, "project_id")
	assert.ElementsMatch(t, tool.InputSchema.Required, []string{"project_id"})

	mockProject := map[string]any{"id": 1, "name": "proj"}

	tests := []struct {
		name           string
		mockedClient   *http.Client
		requestArgs    map[string]interface{}
		expectError    bool
		expectedResult map[string]any
		expectedErrMsg string
	}{
		{
			name: "success",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatch(
					mock.GetProjectsByProjectId,
					mockProject,
				),
			),
			requestArgs:    map[string]interface{}{"project_id": float64(1)},
			expectError:    false,
			expectedResult: mockProject,
		},
		{
			name: "not found",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.GetProjectsByProjectId,
					mockResponse(t, http.StatusNotFound, `{"message": "not found"}`),
				),
			),
			requestArgs:    map[string]interface{}{"project_id": float64(2)},
			expectError:    true,
			expectedErrMsg: "failed to get project",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			client := github.NewClient(tc.mockedClient)
			_, handler := GetProject(stubGetClientFn(client), translations.NullTranslationHelper)
			request := createMCPRequest(tc.requestArgs)
			result, err := handler(context.Background(), request)

			if tc.expectError {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedErrMsg)
				return
			}

			require.NoError(t, err)
			text := getTextResult(t, result)
			var returned map[string]any
			err = json.Unmarshal([]byte(text.Text), &returned)
			require.NoError(t, err)
			assert.EqualValues(t, tc.expectedResult["id"], returned["id"])
		})
	}
}

func Test_ListRepositoryProjects(t *testing.T) {
	mockClient := github.NewClient(nil)
	tool, _ := ListRepositoryProjects(stubGetClientFn(mockClient), translations.NullTranslationHelper)
	assert.Equal(t, "list_repository_projects", tool.Name)
	assert.NotEmpty(t, tool.Description)
	assert.Contains(t, tool.InputSchema.Properties, "owner")
	assert.Contains(t, tool.InputSchema.Properties, "repo")
	assert.Contains(t, tool.InputSchema.Properties, "state")
	assert.ElementsMatch(t, tool.InputSchema.Required, []string{"owner", "repo"})

	mockProjects := []map[string]any{{"id": 1, "name": "p"}}

	tests := []struct {
		name           string
		mockedClient   *http.Client
		requestArgs    map[string]interface{}
		expectError    bool
		expectedResult []map[string]any
		expectedErrMsg string
	}{
		{
			name: "minimal",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatch(
					mock.GetReposProjectsByOwnerByRepo,
					mockProjects,
				),
			),
			requestArgs:    map[string]interface{}{"owner": "o", "repo": "r"},
			expectError:    false,
			expectedResult: mockProjects,
		},
		{
			name: "error",
			mockedClient: mock.NewMockedHTTPClient(
				mock.WithRequestMatchHandler(
					mock.GetReposProjectsByOwnerByRepo,
					mockResponse(t, http.StatusNotFound, `{"message":"no"}`),
				),
			),
			requestArgs:    map[string]interface{}{"owner": "o", "repo": "r"},
			expectError:    true,
			expectedErrMsg: "failed to list projects",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			client := github.NewClient(tc.mockedClient)
			_, handler := ListRepositoryProjects(stubGetClientFn(client), translations.NullTranslationHelper)
			request := createMCPRequest(tc.requestArgs)
			result, err := handler(context.Background(), request)

			if tc.expectError {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedErrMsg)
				return
			}

			require.NoError(t, err)
			text := getTextResult(t, result)
			var returned []map[string]any
			err = json.Unmarshal([]byte(text.Text), &returned)
			require.NoError(t, err)
			assert.Equal(t, len(tc.expectedResult), len(returned))
		})
	}
}
