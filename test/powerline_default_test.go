package test

import (
	"testing"

	"github.com/jimeh/go-tmuxtheme/pkg/theme"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPowerlineDefaultThemes(t *testing.T) {
	tests := []struct {
		filename string
		color1   string
		color2   string
	}{
		{
			filename: "../powerline/default/blue.tmuxtheme",
			color1:   "colour24",
			color2:   "colour33",
		},
		{
			filename: "../powerline/default/cyan.tmuxtheme",
			color1:   "colour39",
			color2:   "colour81",
		},
		{
			filename: "../powerline/default/gray.tmuxtheme",
			color1:   "colour245",
			color2:   "colour250",
		},
		{
			filename: "../powerline/default/green.tmuxtheme",
			color1:   "colour100",
			color2:   "colour190",
		},
		{
			filename: "../powerline/default/magenta.tmuxtheme",
			color1:   "colour125",
			color2:   "colour127",
		},
		{
			filename: "../powerline/default/orange.tmuxtheme",
			color1:   "colour130",
			color2:   "colour166",
		},
		{
			filename: "../powerline/default/purple.tmuxtheme",
			color1:   "colour90",
			color2:   "colour129",
		},
		{
			filename: "../powerline/default/red.tmuxtheme",
			color1:   "colour88",
			color2:   "colour160",
		},
		{
			filename: "../powerline/default/yellow.tmuxtheme",
			color1:   "colour227",
			color2:   "colour227",
		},
	}

	for _, tt := range tests {
		th := theme.New()

		err := th.Load(tt.filename)
		require.NoError(t, err)

		err = th.Execute()
		require.NoError(t, err)

		assert.Equal(t, map[string]string{}, th.ServerOptions)
		assert.Equal(
			t,
			map[string]string{
				"clock-mode-colour":            tt.color1,
				"clock-mode-style":             "24",
				"display-panes-active-colour":  "colour245",
				"display-panes-colour":         "colour233",
				"message-command-style":        "bg=" + tt.color1 + ",fg=black",
				"message-style":                "bg=" + tt.color1 + ",fg=black",
				"mode-style":                   "bg=" + tt.color1 + ",fg=black",
				"pane-active-border-style":     "bg=default,fg=" + tt.color1,
				"pane-border-style":            "bg=default,fg=colour238",
				"status-interval":              "1",
				"status-justify":               "centre",
				"status-left":                  "#[fg=colour233,bg=" + tt.color1 + ",bold] #S #[fg=" + tt.color1 + ",bg=colour240,nobold]\ue0b0#[fg=colour233,bg=colour240] #(whoami) #[fg=colour240,bg=colour235]\ue0b0#[fg=colour240,bg=colour235] #I:#P #[fg=colour235,bg=colour233,nobold]\ue0b0",
				"status-left-length":           "40",
				"status-left-style":            "bg=colour233,fg=colour243",
				"status-right":                 "#[fg=colour235,bg=colour233]\ue0b2#[fg=colour240,bg=colour235] %H:%M:%S #[fg=colour240,bg=colour235]\ue0b2#[fg=colour233,bg=colour240] %d-%b-%y #[fg=colour245,bg=colour240]\ue0b2#[fg=colour233,bg=colour245,bold] #H ",
				"status-right-length":          "150",
				"status-right-style":           "bg=colour233,fg=colour243",
				"status-style":                 "fg=colour240,bg=colour233",
				"window-status-activity-style": "bg=colour233,fg=colour245",
				"window-status-current-format": "#[fg=colour233,bg=black]\ue0b0#[fg=" + tt.color2 + ",nobold] #I:#W#F #[fg=colour233,bg=black,nobold]\ue0b2",
				"window-status-current-style":  "bg=colour100,fg=colour235",
				"window-status-format":         "  #I:#W#F  ",
				"window-status-separator":      "",
			},
			th.GlobalSessionOptions,
		)
		assert.Equal(t, map[string]string{}, th.SessionOptions)
		assert.Equal(t, map[string]string{}, th.GlobalWindowOptions)
		assert.Equal(t, map[string]string{}, th.WindowOptions)
	}
}
