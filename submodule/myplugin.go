package submodule

import "log"

type myStruct struct {
}

func (plugin *myStruct) IAmAPlugin() bool {
	log.Println("myplugin called")
	return true
}

func init() {
	s := new(myStruct)
	RegisterPlugin("myplugin", s)
}
