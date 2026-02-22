package models

import (
	"path/filepath"
	"strings"
	"time"
)

// Dependency represents a relationship between two issues (from bd list JSON)
type Dependency struct {
	IssueID     string `json:"issue_id"`
	DependsOnID string `json:"depends_on_id"`
	Type        string `json:"type"`
}

// IsParentChild returns true if this dependency represents a parent-child
// relationship. Matches both "parent-child" and "parent" type strings.
func (d Dependency) IsParentChild() bool {
	return strings.HasPrefix(d.Type, "parent")
}

// Task represents a beads issue
type Task struct {
	ID                 string       `json:"id"`
	Title              string       `json:"title"`
	Description        string       `json:"description,omitempty"`
	Notes              string       `json:"notes,omitempty"`
	Design             string       `json:"design,omitempty"`
	AcceptanceCriteria string       `json:"acceptance_criteria,omitempty"`
	Status             string       `json:"status"`
	Priority           int          `json:"priority"`
	Type               string       `json:"issue_type"`
	Labels             []string     `json:"labels,omitempty"`
	Assignee           string       `json:"assignee,omitempty"`
	Owner              string       `json:"owner,omitempty"`
	CreatedAt          time.Time    `json:"created_at"`
	CreatedBy          string       `json:"created_by,omitempty"`
	UpdatedAt          time.Time    `json:"updated_at"`
	ClosedAt           *time.Time   `json:"closed_at,omitempty"`
	CloseReason        string       `json:"close_reason,omitempty"`
	DueDate            *time.Time   `json:"due_date,omitempty"`
	DeferUntil         *time.Time   `json:"defer_until,omitempty"`
	BlockedBy          []string     `json:"blocked_by,omitempty"`
	Blocks             []string     `json:"blocks,omitempty"`
	Dependencies       []Dependency `json:"dependencies,omitempty"`
	DependencyCount    int          `json:"dependency_count,omitempty"`
	DependentCount     int          `json:"dependent_count,omitempty"`
}

// PriorityString returns a short priority label
func (t Task) PriorityString() string {
	switch t.Priority {
	case 0:
		return "P0"
	case 1:
		return "P1"
	case 2:
		return "P2"
	case 3:
		return "P3"
	case 4:
		return "P4"
	default:
		return "P?"
	}
}

// StatusIcon returns a status indicator
func (t Task) StatusIcon() string {
	switch t.Status {
	case "open":
		return "○"
	case "in_progress":
		return "◐"
	case "closed":
		return "●"
	default:
		return "?"
	}
}

// IsBlocked returns true if task has blockers
func (t Task) IsBlocked() bool {
	return len(t.BlockedBy) > 0
}

// GetParentID returns the parent task ID. It checks explicit parent-child
// dependencies first, then falls back to ID naming convention (dot notation).
func (t Task) GetParentID() string {
	for _, dep := range t.Dependencies {
		if dep.IsParentChild() {
			return dep.DependsOnID
		}
	}
	return ParentID(t.ID)
}

// FilePath returns the path to the task's markdown file
func (t Task) FilePath() string {
	return filepath.Join(".beads", "issues", t.ID+".md")
}

// Comment represents a comment on an issue
type Comment struct {
	ID        int       `json:"id"`
	IssueID   string    `json:"issue_id"`
	Author    string    `json:"author"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}
