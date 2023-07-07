package rumble

import (
	"testing"

	"github.com/wanyuqin/lux/extractors"
	"github.com/wanyuqin/lux/test"
)

func TestRumble(t *testing.T) {
	tests := []struct {
		name string
		args test.Args
	}{
		{
			name: "normal test",
			args: test.Args{
				URL:   "https://rumble.com/v24swn0-just-say-yes-to-climate-lockdowns.html",
				Title: "Just Say YES to Climate Lockdowns!",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			New().Extract(tt.args.URL, extractors.Options{})
		})
	}
}
