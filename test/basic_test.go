package test

import (
	"testing"

	"github.com/jimeh/go-tmuxtheme/pkg/theme"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBasicTheme(t *testing.T) {
	th := theme.New()

	err := th.Load("../basic.tmuxtheme")
	require.NoError(t, err)

	err = th.Execute()
	require.NoError(t, err)

	assert.Equal(t, map[string]string{}, th.ServerOptions)
	assert.Equal(
		t,
		map[string]string{
			"clock-mode-colour":            "red",
			"clock-mode-style":             "24",
			"display-panes-active-colour":  "default",
			"display-panes-colour":         "default",
			"message-command-style":        "bg=default,fg=default",
			"message-style":                "bg=default,fg=default",
			"mode-style":                   "bg=red,fg=default",
			"pane-active-border-style":     "bg=default,fg=green",
			"pane-border-style":            "bg=default,fg=default",
			"status-interval":              "1",
			"status-justify":               "centre",
			"status-left":                  "#S #[fg=white]» #[fg=yellow]#I #[fg=cyan]#P",
			"status-left-length":           "40",
			"status-left-style":            "bg=black,fg=green",
			"status-right":                 "#H #[fg=white]« #[fg=yellow]%H:%M:%S #[fg=green]%d-%b-%y",
			"status-right-length":          "40",
			"status-right-style":           "bg=black,fg=cyan",
			"status-style":                 "bg=black,fg=cyan",
			"window-status-activity-style": "bg=black,fg=yellow",
			"window-status-current-format": " #I:#W#F ",
			"window-status-current-style":  "bg=red,fg=black",
			"window-status-format":         " #I:#W#F ",
			"window-status-separator":      "",
		},
		th.GlobalSessionOptions,
	)
	assert.Equal(t, map[string]string{}, th.SessionOptions)
	assert.Equal(t, map[string]string{}, th.GlobalWindowOptions)
	assert.Equal(t, map[string]string{}, th.WindowOptions)
}
