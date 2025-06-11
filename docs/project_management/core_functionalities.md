# Core Functionalities

This document details the core functionalities of the GitHub Project Management Feature.

## 1. Projects

*   **Ability to create new projects within the tool.**
    *   Users can initiate new projects directly, providing a dedicated space for organizing work.
*   **Option to link/sync with existing GitHub Projects (Projects V2).**
    *   Seamlessly connect with and import data from existing GitHub Projects (Projects V2).
    *   This includes importing boards, columns, and items (issues, PRs, draft issues).
    *   Synchronization should be two-way where feasible.
*   **Project properties:**
    *   **Name:** A unique and descriptive name for the project.
    *   **Description:** A detailed explanation of the project's goals and scope.
    *   **Start Date:** The planned or actual start date of the project.
    *   **End Date:** The planned or actual end date of the project.
    *   **Owner/Lead:** The primary person responsible for the project.
*   **Project views:**
    *   **Kanban board:** A visual board to manage tasks in different stages (e.g., To Do, In Progress, Done). Columns should be customizable.
    *   **List view:** A tabular view of tasks, allowing for sorting, filtering, and bulk editing.
    *   **Calendar view:** A view to visualize tasks with due dates on a calendar.

## 2. Tasks

*   **Create new tasks within a project.**
    *   Tasks can be created directly within a project in the tool.
*   **Tasks can be standalone or linked to GitHub Issues.**
    *   If linked to a GitHub Issue:
        *   Sync Title.
        *   Sync Description.
        *   Sync Labels.
        *   Sync Assignees.
        *   Sync Status (mapping tool's task status to GitHub Issue state, e.g., Open/Closed).
*   **Task properties:**
    *   **Title:** A clear and concise title for the task.
    *   **Description:** A detailed description of what the task entails.
    *   **Assignee(s):** One or more users responsible for completing the task.
    *   **Status:** The current stage of the task (e.g., To Do, In Progress, Review, Done). Statuses should be customizable per project.
    *   **Priority:** The urgency or importance of the task (e.g., Low, Medium, High).
    *   **Due Date:** The target completion date for the task.
    *   **Estimated effort:** A measure of the work required (e.g., story points, hours).
*   **Subtasks:**
    *   Ability to break down complex tasks into smaller, more manageable subtasks.
    *   Subtasks can have their own properties (assignee, due date, status).
*   **Dependencies:**
    *   Define dependencies between tasks.
    *   Example: "Task B cannot start until Task A is complete."
    *   Visualize dependencies (e.g., Gantt chart or links in task views).

## 3. Collaboration

*   **Comments and discussions on tasks.**
    *   Allow users to comment on tasks to ask questions, provide updates, or discuss details.
    *   Support for rich text formatting in comments.
*   **@mentions to notify team members.**
    *   Ability to mention other GitHub users in comments and descriptions to send them notifications.
*   **Activity feed for projects and tasks.**
    *   A chronological log of all significant actions and updates within a project or on a specific task.
    *   Examples: task creation, status changes, new comments, assignee changes.

## 4. Time Tracking (Optional - Future Enhancement)

*   **Ability to log time spent on tasks.**
    *   Users can manually enter time spent on individual tasks.
    *   Basic reporting on time logged per task or project.

## 5. Reporting & Analytics (Optional - Future Enhancement)

*   **Burndown charts:**
    *   Visualize the remaining work over time against the planned timeline.
*   **Task completion reports:**
    *   Track the number of tasks completed, overdue tasks, and tasks by status or assignee.
*   **Project progress overview:**
    *   A dashboard or report summarizing key project metrics, including overall progress, budget (if applicable), and resource allocation.
