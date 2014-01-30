package nagiosplugin_test

import (
	"github.com/fractalcat/nagiosplugin"
)

func ExampleExit() {
	nagiosplugin.Exit(nagiosplugin.CRITICAL, "Badness over 9000!")
}
