package facebook

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
				URL:     "https://www.facebook.com/100058251872436/videos/424557726111987",
				Title:   "Роман Грищук - Підтримка з Японії 🇯🇵 Гурт Yokohama Sisters 👏",
				Size:    1441128,
				Quality: "sd",
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
