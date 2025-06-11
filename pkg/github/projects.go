package github

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/github/github-mcp-server/pkg/translations"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// GetProject creates a tool to retrieve a GitHub project by its ID.
func GetProject(getClient GetClientFn, t translations.TranslationHelperFunc) (mcp.Tool, server.ToolHandlerFunc) {
	return mcp.NewTool(
			"get_project",
			mcp.WithDescription(t("TOOL_GET_PROJECT_DESCRIPTION", "Get details of a GitHub project.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_GET_PROJECT_USER_TITLE", "Get project"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithNumber("project_id",
				mcp.Required(),
				mcp.Description("ID of the project"),
			),
		), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			projectID, err := RequiredInt(request, "project_id")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}

			req, err := client.NewRequest("GET", fmt.Sprintf("projects/%d", projectID), nil)
			if err != nil {
				return nil, fmt.Errorf("failed to create request: %w", err)
			}
			req.Header.Set("Accept", "application/vnd.github.inertia-preview+json")

			var project map[string]any
			resp, err := client.Do(ctx, req, &project)
			if err != nil {
				return nil, fmt.Errorf("failed to get project: %w", err)
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != http.StatusOK {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to get project: %s", string(body))), nil
			}

			r, err := json.Marshal(project)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal project: %w", err)
			}

			return mcp.NewToolResultText(string(r)), nil
		}
}

// ListRepositoryProjects lists projects for a repository.
func ListRepositoryProjects(getClient GetClientFn, t translations.TranslationHelperFunc) (mcp.Tool, server.ToolHandlerFunc) {
	return mcp.NewTool(
			"list_repository_projects",
			mcp.WithDescription(t("TOOL_LIST_REPOSITORY_PROJECTS_DESCRIPTION", "List projects in a GitHub repository.")),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				Title:        t("TOOL_LIST_REPOSITORY_PROJECTS_USER_TITLE", "List repository projects"),
				ReadOnlyHint: ToBoolPtr(true),
			}),
			mcp.WithString("owner",
				mcp.Required(),
				mcp.Description("Repository owner"),
			),
			mcp.WithString("repo",
				mcp.Required(),
				mcp.Description("Repository name"),
			),
			mcp.WithString("state",
				mcp.Description("Filter by state"),
				mcp.Enum("open", "closed", "all"),
			),
			WithPagination(),
		), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			owner, err := RequiredParam[string](request, "owner")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			repo, err := RequiredParam[string](request, "repo")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			state, err := OptionalParam[string](request, "state")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			pagination, err := OptionalPaginationParams(request)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			client, err := getClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get GitHub client: %w", err)
			}
			url := fmt.Sprintf("repos/%s/%s/projects", owner, repo)
			req, err := client.NewRequest("GET", url, nil)
			if err != nil {
				return nil, fmt.Errorf("failed to create request: %w", err)
			}
			req.Header.Set("Accept", "application/vnd.github.inertia-preview+json")

			q := req.URL.Query()
			if state != "" {
				q.Add("state", state)
			}
			q.Add("page", fmt.Sprintf("%d", pagination.page))
			q.Add("per_page", fmt.Sprintf("%d", pagination.perPage))
			req.URL.RawQuery = q.Encode()

			var projects []map[string]any
			resp, err := client.Do(ctx, req, &projects)
			if err != nil {
				return nil, fmt.Errorf("failed to list projects: %w", err)
			}
			defer func() { _ = resp.Body.Close() }()
			if resp.StatusCode != http.StatusOK {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("failed to read response body: %w", err)
				}
				return mcp.NewToolResultError(fmt.Sprintf("failed to list projects: %s", string(body))), nil
			}
			r, err := json.Marshal(projects)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal response: %w", err)
			}
			return mcp.NewToolResultText(string(r)), nil
		}
}
