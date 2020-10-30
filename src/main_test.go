package main

import (
	"github.com/wasmerio/go-ext-wasm/wasmer"
	"testing"
)

func TestStartBinary_Swift(t *testing.T) {
	_,_, instance := instantiate(swiftBinaryPath)

	if start(&instance).GetType() != wasmer.TypeVoid {
		t.Error("error in starting the project")
	}
}

func TestExportedFunction_Swift(t *testing.T) {
	_,_, instance := instantiate(swiftBinaryPath)

	result := sum(&instance)
	if result != 3 {
		t.Error("expected value is: 3, but got:", result)
	}
}

func TestImportedFunction_Swift(t *testing.T) {
	_,_, instance := instantiate(swiftBinaryPath)

	result := fetchCodeOnBinary(&instance, 2)
	if result != 4 {
		t.Error("expected value is: 4, but got:", result)
	}
}

func TestStartBinary_Rust(t *testing.T) {
	_,_, instance := instantiate(rustBinaryPath)

	if start(&instance).GetType() != wasmer.TypeVoid {
		t.Error("error in starting the project")
	}
}

func TestExportedFunction_Rust(t *testing.T) {
	_,_, instance := instantiate(rustBinaryPath)

	result := sum(&instance)
	if result != 3 {
		t.Error("expected value is: 3, but got:", result)
	}
}

func TestImportedFunction_Rust(t *testing.T) {
	_,_, instance := instantiate(rustBinaryPath)

	result := fetchCodeOnBinary(&instance, 2)
	if result != 4 {
		t.Error("expected value is: 4, but got:", result)
	}
}
