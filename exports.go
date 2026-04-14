package main

import "encoding/json"

// plugin_info returns metadata as JSON.
//
//export plugin_info
func pluginInfo() uint64 {
	data, _ := json.Marshal(info())
	return toPointer(data)
}

// plugin_init runs one-time initialization.
//
//export plugin_init
func pluginInit(cfgPtr uint64) uint64 {
	// TODO: read cfgPtr, apply any generator settings
	_ = cfgPtr
	return 0
}

// generate produces test files from an OpenAPI spec.
//
//export generate
func generateFromSpec(inputPtr uint64) uint64 {
	// TODO:
	// 1. readBytes(inputPtr) → GenerateInput
	// 2. Call generate(in)
	// 3. Marshal GenerateOutput to JSON
	// 4. Return toPointer(data)
	_ = inputPtr
	return 0
}

// plugin_destroy releases generator resources.
//
//export plugin_destroy
func pluginDestroy() {}

// toPointer packs a byte slice into a (ptr << 32) | len uint64.
func toPointer(data []byte) uint64 {
	// TODO: WASM implementation via unsafe
	return uint64(len(data))
}

// readBytes reads bytes at the given packed location.
func readBytes(packed uint64) []byte {
	// TODO: WASM implementation
	_ = packed
	return nil
}
