package util

import (
	"fmt"
    // "github.com/gookit/color"
)

func Initialize() {
	print_banner()
	load_plugins()
}

/*
To print banner when programme is executed
*/
func print_banner() {
    fmt.Println(" _ _  _ ____ _ ____ _  _ ___ ")
    fmt.Println(" | |\\ | [__  | | __ |__|  |  ")
    fmt.Println(" | | \\| ___] | |__] |  |  |  ")
    fmt.Println("\n A penetration testing tool by @muffled.io\n ")
}

/*
To compile plugins at 'core/plugins' directory to *.so
*/
func load_plugins() {
	command := fmt.Sprintf("find core/plugins -name *.sh -exec {} \\ ;")
	if ok, _ := Contains("Loaded", Execute(command)); ok {
		fmt.Println("[*] Plugins Loaded")
	} else {
		fmt.Println("[x] Error Loading Plugins")
	}
}
