# GitHub Projects Management

This document proposes adding initial support for managing GitHub Projects within the MCP server.

## Goals

- Allow tools to fetch individual project information.
- Provide listing of projects for a repository so they can be surfaced by clients.
- Integrate with the existing toolset mechanism so projects management can be enabled alongside issues and pull requests.

## Non Goals

- Mutating project state such as moving items or editing fields.
- Support for the new ProjectsV2 beta APIs.

## Proposed API

Two new tools are introduced:

1. **get_project** – Returns basic information about a project given its ID.
2. **list_repository_projects** – Lists projects for a repository with optional state filtering and pagination.

Both tools will live in a new `projects` toolset. When enabled, the toolset exposes read‑only functionality similar to the existing `issues` toolset.

## Testing

Unit tests will verify schema snapshots and basic success/error handling using mocked `go-github` clients. Linter and `go test ./...` should continue to pass.
