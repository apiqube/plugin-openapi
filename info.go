package main

// plugin-openapi is a Generator plugin, not an Executor.
// It has no protocols and no manifest fields — it is invoked via
// `qube generate --from openapi.json` and produces YAML test files.
//
// Generator plugins use a different subset of the plugin contract:
// plugin_info() + generate(source) instead of plugin_info() + execute(test).
func info() PluginInfo {
	return PluginInfo{
		Name:        "openapi",
		Version:     "0.1.0",
		Description: "Generates qube test files from OpenAPI/Swagger specifications.",
		Protocols:   nil, // generator plugins have no protocols
		Fields:      nil, // generator plugins do not add manifest fields
	}
}
