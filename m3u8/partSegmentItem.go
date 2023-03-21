package m3u8

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// PartSegmentItem represents partial segment attributes with a duration and URI and
// potentially a program date time
type PartSegmentItem struct {
	Duration    float64
	Uri         string
	Independent bool
}

// NewPartSegmentItem parses a text line and returns a *PartSegmentItem
func NewPartSegmentItem(text string) (*PartSegmentItem, error) {
	var ps PartSegmentItem
	line := strings.Replace(text, PartSegmentItemTag+":", "", -1)
	line = strings.Replace(line, "\n", "", -1)

	attributes := ParseAttributes(line)
	if val, ok := attributes["URI"]; ok {
		ps.Uri = val
	} else {
		return nil, errors.New("Missing URI attribute")
	}

	if val, ok := attributes["DURATION"]; ok {
		if duration, err := strconv.ParseFloat(val, 64); err == nil {
			ps.Duration = duration
		} else {
			return nil, errors.New("Invalid DURATION value")

		}
	} else {
		return nil, errors.New("Missing DURATION attribute")
	}

	if val, ok := attributes["INDEPENDENT"]; ok {
		if val == "true" {
			ps.Independent = true
		} else {
			return nil, errors.New("Invalid INDEPENDENT value")
		}
	} else {
		ps.Independent = false
	}

	return &ps, nil
}

func (ps *PartSegmentItem) String() string {
	tag := fmt.Sprintf(`%s:%s=%.3f,%s="%v"`, PartSegmentItemTag, DurationTag, ps.Duration, URITag, ps.Uri)

	if ps.Independent {
		return fmt.Sprintf("%s,%s=true", tag, IndepententTag)
	}

	return tag
}
