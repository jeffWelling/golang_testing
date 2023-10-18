package main

import "golang_testing/submodule"

func main() {
	f := submodule.Backend("myplugin")
	f.IAmAPlugin()
}
