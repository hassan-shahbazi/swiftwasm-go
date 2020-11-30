package main

import (
	"testing"

	"github.com/wasmerio/go-ext-wasm/wasmer"
)

var instance wasmer.Instance

func TestMain(m *testing.M) {
	_, _, instance = instantiate(swiftBinaryPath)
	m.Run()
}

func TestStartBinary_Swift(t *testing.T) {
	if start(&instance).GetType() != wasmer.TypeVoid {
		t.Error("error in starting the project")
	}
}

func TestExportedFunction_Swift(t *testing.T) {
	result := sum(&instance)
	if result != 3 {
		t.Error("expected value is: 3, but got:", result)
	}
}

func TestImportedFunction_Swift(t *testing.T) {
	result := fetchCodeOnBinary(&instance, 2)
	if result != 4 {
		t.Error("expected value is: 4, but got:", result)
	}
}

func TestStartBinary_Rust(t *testing.T) {
	if start(&instance).GetType() != wasmer.TypeVoid {
		t.Error("error in starting the project")
	}
}

func TestExportedFunction_Rust(t *testing.T) {
	result := sum(&instance)
	if result != 3 {
		t.Error("expected value is: 3, but got:", result)
	}
}

func TestImportedFunction_Rust(t *testing.T) {
	result := fetchCodeOnBinary(&instance, 2)
	if result != 4 {
		t.Error("expected value is: 4, but got:", result)
	}
}
