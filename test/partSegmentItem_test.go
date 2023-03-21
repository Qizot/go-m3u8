package test

import (
	"testing"

	"github.com/Qizot/go-m3u8/m3u8"
	"github.com/stretchr/testify/assert"
)

func TestPartSegmentItem_Parse(t *testing.T) {
	item := &m3u8.PartSegmentItem{
		Duration:    10.991,
		Uri:         "test.ts",
		Independent: true,
	}

	assert.Equal(t, `#EXT-X-PART:DURATION=10.991,URI="test.ts",INDEPENDENT=true`, item.String())

	item.Independent = false
	assert.Equal(t, `#EXT-X-PART:DURATION=10.991,URI="test.ts"`, item.String())
}
