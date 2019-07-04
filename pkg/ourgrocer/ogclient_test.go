package ourgrocer

import (
	"strings"
	"testing"
)

func TestExtractTeamId(t *testing.T) {
	stream := strings.NewReader("var g_teamId = \"E0KAegvBF9SOQ78b9vhlYr\"")
	teamID, err := extractTeamID(stream)
	ok(t, err)
	equals(t, "E0KAegvBF9SOQ78b9vhlYr", teamID)
}
