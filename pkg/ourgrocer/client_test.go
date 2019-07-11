// +build small_test

package ourgrocer_test

import (
	"strings"
	"testing"

	"github.com/adammansfield/grocerer/pkg/ourgrocer"
)

// TODO: modify test to only call Login (and still test teamID extraction)
func TestExtractTeamId(t *testing.T) {
	stream := strings.NewReader("var g_teamId = \"E0KAegvBF9SOQ78b9vhlYr\"")
	teamID, err := ourgrocer.ExtractTeamID(stream)
	ok(t, err)
	equals(t, "E0KAegvBF9SOQ78b9vhlYr", teamID)
}
