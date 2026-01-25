package ui

import (
	"sync"

	"github.com/charmbracelet/glamour"
)

var (
	// Cached renderer to avoid repeated terminal queries that corrupt input
	cachedRenderer     *glamour.TermRenderer
	cachedRendererOnce sync.Once
	cachedWidth        int
)

// getRenderer returns a cached glamour renderer, creating it on first use.
// This prevents repeated terminal color queries that interfere with keyboard input.
func getRenderer(width int) *glamour.TermRenderer {
	cachedRendererOnce.Do(func() {
		r, err := glamour.NewTermRenderer(
			glamour.WithAutoStyle(),
			glamour.WithWordWrap(width),
		)
		if err == nil {
			cachedRenderer = r
			cachedWidth = width
		}
	})
	return cachedRenderer
}

// RenderMarkdown renders markdown text to styled terminal output
func RenderMarkdown(text string, width int) string {
	if text == "" {
		return text
	}

	r := getRenderer(width)
	if r == nil {
		return text
	}

	rendered, err := r.Render(text)
	if err != nil {
		return text
	}

	return rendered
}
