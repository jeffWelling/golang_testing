package submodule

var plugins = make(map[string]PluginInterface)

type PluginInterface interface {
	IAmAPlugin() bool
}

func RegisterPlugin(name string, plugin PluginInterface) {
	plugins[name] = plugin
}

func Backend(name string) PluginInterface {
	return plugins[name]
}
