package main

import (
	"fmt"

	"github.com/wasmerio/go-ext-wasm/wasmer"
)

func main() {
	// Reads the WebAssembly module as bytes.
	bytes, err := wasmer.ReadBytes("binary.wasm")
	if err != nil {
		panic(err)
	}

	// Compile bytes into wasm binary
	module, err := wasmer.Compile(bytes)
	if err != nil {
		panic(err)
	}

	// Get current wasi version and corresponded import objects
	wasiVersion := wasmer.WasiGetVersion(module)
	if wasiVersion == 0 {
		// wasiVersion is unknow, use Latest instead
		wasiVersion = wasmer.Latest
	}
	importObject := wasmer.NewDefaultWasiImportObjectForVersion(wasiVersion)

	// Instantiates the WebAssembly module using derived import objects.
	instance, err := module.InstantiateWithImportObject(importObject)
	if err != nil {
		panic(err)
	}
	defer importObject.Close()
	defer instance.Close()

	start(&instance)
}

func start(instance *wasmer.Instance) {
	// Gets start function from the WebAssembly instance.
	start := instance.Exports["_start"]

	// Calls that exported function with Go standard values. The WebAssembly types are inferred and values are casted automatically.
	result, err := start()
	if err != nil {
		panic(err)
	}

	// To ensure the start function doesn't return any values
	fmt.Println(result.GetType() == wasmer.TypeVoid)
}
