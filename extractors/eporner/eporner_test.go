package eporner

import (
	"testing"

	"github.com/wanyuqin/lux/extractors"
	"github.com/wanyuqin/lux/test"
)

func TestDownload(t *testing.T) {
	tests := []struct {
		name string
		args test.Args
	}{
		{
			name: "normal test",
			args: test.Args{
				URL:     "https://www.eporner.com/video-mbubfvXYFip/dirtywivesclub-becky-bandini/",
				Quality: "1080p",
				Size:    1525510307,
				Title:   "DirtyWivesClub - Becky Bandini - EPORNER",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := New().Extract(tt.args.URL, extractors.Options{})
			test.CheckError(t, err)
			test.Check(t, tt.args, data[0])
		})
	}
}
