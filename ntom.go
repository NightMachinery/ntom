/// 2>/dev/null ; gorun "$0" "$@" ; exit $?

package main

import "os"

func main() {
	println("Hello world!")

	os.Exit(0)
}
