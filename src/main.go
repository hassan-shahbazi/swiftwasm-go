package main

// #include <stdlib.h>
//
// extern int fetchCode(void *, int);
import "C"
import (
	"fmt"
	"strings"
	"unsafe"

	"github.com/wasmerio/go-ext-wasm/wasmer"
)

const (
	swiftBinaryPath		= "../swiftwasm/binary.wasm"
	rustBinaryPath		= "../rust/binary.wasm"
)

func main() {
	importObject, imports, instance := instantiate(swiftBinaryPath)

	defer importObject.Close()
	defer imports.Close()
	defer instance.Close()

	fmt.Println("'_start' called, but returned void:", start(&instance).GetType() == wasmer.TypeVoid)
	fmt.Println("'save' exported function:", sum(&instance))
	fmt.Println("'concatenate' exported function with string parameters:", concatenate(&instance, "Hello", "World!"))
	fmt.Println("'fetch_code imported function input: 2'", fetchCodeOnBinary(&instance, 2))
}

func instantiate(path string) (*wasmer.ImportObject, *wasmer.Imports, wasmer.Instance) {
	// Reads the WebAssembly module as bytes.
	bytes, err := wasmer.ReadBytes(path)
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

	// Make new C imports
	imports := wasmer.NewImports()
	imports, err = imports.AppendFunction("fetch_code", fetchCode, C.fetchCode)
	if err != nil {
		panic(err)
	}

	// Extend the import objects with C imports
	importObject.Extend(*imports)

	// Instantiates the WebAssembly module using derived import objects.
	instance, err := module.InstantiateWithImportObject(importObject)
	if err != nil {
		panic(err)
	}
	return importObject, imports, instance
}

func start(instance *wasmer.Instance) wasmer.Value {
	// Gets start function from the WebAssembly instance.
	start := instance.Exports["_start"]

	// Calls that exported function with Go standard values. The WebAssembly types are inferred and values are casted automatically.
	result, err := start()
	if err != nil {
		panic(err)
	}

	// To ensure the start function doesn't return any values
	return result
}

func sum(instance *wasmer.Instance) int32 {
	// Gets sum function from the WebAssembly instance.
	sum := instance.Exports["sum"]

	// Calls that exported function with Go standard values. The WebAssembly types are inferred and values are casted automatically.
	result, err := sum(1, 2)
	if err != nil {
		panic(err)
	}
	return result.ToI32()
}

func concatenate(instance *wasmer.Instance, s1, s2 string) string {
	// Allocate memory function
	allocate := instance.Exports["allocate"]

	memoryPointerS1 := func() int32 {
		// run allocate for s2
		pointer, err := allocate(len(s1) + 1)
		if err != nil {
			panic(err)
		}

		// copy s2 to memory
		memoryS1 := instance.Memory.Data()[pointer.ToI32():]
		copy(memoryS1, s2)
		memoryS1[len(s1)] = 0

		return pointer.ToI32()
	}()

	memoryPointerS2 := func() int32 {
		// run allocate for s2
		pointer, err := allocate(len(s2) + 1)
		if err != nil {
			panic(err)
		}

		// copy s2 to memory
		memoryS2 := instance.Memory.Data()[pointer.ToI32():]
		copy(memoryS2, s2)
		memoryS2[len(s2)] = 0

		return pointer.ToI32()
	}()

	// Get concat function from the WebAssembly instance.
	concat := instance.Exports["concatenate"]

	// Calls that exported function with memory pointers.
	result, err := concat(memoryPointerS1, memoryPointerS2)
	if err != nil {
		panic(err)
	}

	// The result is another memory pointer contains the concated strings
	output, size := convertToString(instance, result)

	// deallocate memory
	deallocate := instance.Exports["deallocate"]

	// deallocate memoryS1
	_, err = deallocate(memoryPointerS1, len(s1))
	if err != nil {
		panic(err)
	}

	// deallocate memorySs
	_, err = deallocate(memoryPointerS2, len(s2))
	if err != nil {
		panic(err)
	}

	// deallocate output memory
	_, err = deallocate(result, size)
	if err != nil {
		panic(err)
	}

	return output
}

func convertToString(instance *wasmer.Instance, output wasmer.Value) (string, int32) {
	memory := instance.Memory.Data()[output.ToI32():]
	var builder strings.Builder
	counter := 0

	for memory[counter] != 0 {
		builder.WriteByte(memory[counter])
		counter++
	}
	return builder.String(), int32(counter)
}

func fetchCodeOnBinary(instance *wasmer.Instance, input int32) int32 {
	// Gets start function from the WebAssembly instance.
	fetch := instance.Exports["fetch"]

	// Calls that exported function with Go standard values. The WebAssembly types are inferred and values are casted automatically.
	result, err := fetch(input)
	if err != nil {
		panic(err)
	}

	return result.ToI32()
}

//export fetchCode
func fetchCode(context unsafe.Pointer, input C.int) C.int {
	return input * 2
}
