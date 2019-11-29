package test

import (
	"testing"

	"github.com/jimeh/go-tmuxtheme/pkg/theme"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDefaultTheme(t *testing.T) {
	th := theme.New()

	err := th.Load("../default.tmuxtheme")
	require.NoError(t, err)

	err = th.Execute()
	require.NoError(t, err)

	assert.Equal(t, map[string]string{}, th.ServerOptions)
	assert.Equal(
		t,
		map[string]string{
			"clock-mode-colour":            "blue",
			"clock-mode-style":             "24",
			"display-panes-active-colour":  "red",
			"display-panes-colour":         "blue",
			"message-command-style":        "bg=green,fg=black",
			"message-style":                "bg=yellow,fg=black",
			"mode-style":                   "bg=yellow,fg=black",
			"pane-active-border-style":     "bg=default,fg=green",
			"pane-border-style":            "bg=default,fg=white",
			"status-interval":              "15",
			"status-justify":               "left",
			"status-left":                  "[#S]",
			"status-left-length":           "10",
			"status-left-style":            "bg=green,fg=black",
			"status-right":                 "\"#H\" %H:%M %d-%b-%y",
			"status-right-length":          "40",
			"status-right-style":           "bg=green,fg=black",
			"status-style":                 "bg=green,fg=black",
			"window-status-activity-style": "bg=black,fg=green",
			"window-status-current-format": "#I:#W#F",
			"window-status-current-style":  "bg=green,fg=black",
			"window-status-format":         "#I:#W#F",
			"window-status-separator":      " ",
		},
		th.GlobalSessionOptions,
	)
	assert.Equal(t, map[string]string{}, th.SessionOptions)
	assert.Equal(t, map[string]string{}, th.GlobalWindowOptions)
	assert.Equal(t, map[string]string{}, th.WindowOptions)
}
