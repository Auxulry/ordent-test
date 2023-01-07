// Package colorize is used for log colorize
package colorize

import (
	"fmt"

	"github.com/MochamadAkbar/ordent-test/common/constants"
)

func MessageColorized(color, m string) string {
	return fmt.Sprintf("%v%v%v", color, m, constants.DefaultColor)
}
