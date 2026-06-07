package app

import (
	"fmt"
	"runtime"

	"github.com/ebitengine/purego"
)

func TryPureGo() {
	webkit, e := purego.Dlopen(getWebKit(), purego.RTLD_NOW|purego.RTLD_GLOBAL)

	if e != nil {
		panic(e)
	}

	major := callWebKitGetMajorVersion(webkit)
	minor := callWebKitGetMinorVersion(webkit)
	micro := callWebKitGetMicroVersion(webkit)
	fmt.Printf("Found webkitgtk-6.0 v%v.%v.%v\n", major, minor, micro)
}

func getWebKit() string {
	switch runtime.GOOS {
	case "linux":
		// Very specific to my system.
		return "/lib/x86_64-linux-gnu/webkitgtk-6.0/injected-bundle/libwebkitgtkinjectedbundle.so"
	default:
		panic(fmt.Errorf("GOOS=%s not supported", runtime.GOOS))
	}
}

func callWebKitGetMajorVersion(webkit uintptr) uint {
	var webkit_get_major_version func() uint
	purego.RegisterLibFunc(&webkit_get_major_version, webkit, "webkit_get_major_version")
	return webkit_get_major_version()
}

func callWebKitGetMinorVersion(webkit uintptr) uint {
	var webkit_get_minor_version func() uint
	purego.RegisterLibFunc(&webkit_get_minor_version, webkit, "webkit_get_minor_version")
	return webkit_get_minor_version()
}

func callWebKitGetMicroVersion(webkit uintptr) uint {
	var webkit_get_micro_version func() uint
	purego.RegisterLibFunc(&webkit_get_micro_version, webkit, "webkit_get_micro_version")
	return webkit_get_micro_version()
}
