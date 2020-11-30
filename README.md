# Swift + Web Assembly
This repository contains all code related to a 3-part tutorial published on Medium: [Part 1](https://medium.com/@h.shahbazi/the-power-of-swift-web-assembly-part-1-fdfa4e9134ee), [Part 2](https://medium.com/@h.shahbazi/the-power-of-swift-web-assembly-part-2-30b6c4619c27), and [Part 3](https://medium.com/@h.shahbazi/the-power-of-swift-web-assembly-part-3-e583c6ab8afe)


## Requirements
- Ubuntu >= 18.04. (macOS naturally works fine, but I have not tested it yet)
- Go >= 1.15
- swift-wasm >= wasm-5.3.1-RELEASE


## Installation
- To install **swiftenv**, follow instructions on: https://swiftenv.fuller.li/en/latest/installation.html
- Download the latest version of **swiftwasm** from the release page: https://github.com/swiftwasm/swift/releases
- Install dependencies for running Swift on Linux
  ```bash
  $ apt update
  $ apt install binutils git gnupg2 libc6-dev libcurl4 libedit2 libgcc-9-dev libpython2.7 libsqlite3-0 libstdc++-9-dev libxml2 libz3-dev pkg-config tzdata zlib1g-dev curl lsb-release
  ```
- Install **wasmer** by following instructions on: https://wasmer.io/


## Running
1. Set the correct version of Swiftwasm

   ```bash
   $ swiftenv global wasm-5.3.1-RELEASE
   ```

2. Build Swift project and generate wasm binary:
   ```bash
   $ TOOLCHAIN_PATH=$(cd $(dirname "$(swiftenv which swiftc)") && cd ../share && pwd)
   $ swift build --triple wasm32-unknown-wasi -c release --toolchain $TOOLCHAIN_PATH -Xlinker --export=allocate -Xlinker --export=deallocate -Xlinker --export=hello -Xlinker --export=sum -Xlinker --export=concatenate -Xlinker --export=fetch -Xlinker --allow-undefined
   ```

3.  Copy wasm to the project's directory
   ```bash
   $ cp swiftwasm/.build/release/swiftwasm.wasm ../binary.wasm
   ```

4. Run the Go application
   ```bash
   $ go run ./
   > Hello World!
   > '_start' called, but returned void: true
   > 'save' exported function: 3
   > 'concatenate' exported function with string parameters: World World!
   > 'fetch_code imported function input: 2' 4
   ```

5. Run tests
  ```bash
  $ binary=swift go test ./... -v -race -count 1 -run _Swift

  > === RUN   TestStartBinary_Swift
  > Hello World!
  > --- PASS: TestStartBinary_Swift (0.00s)
  > === RUN   TestExportedFunction_Swift
  > --- PASS: TestExportedFunction_Swift (0.00s)
  > === RUN   TestImportedFunction_Swift
  > --- PASS: TestImportedFunction_Swift (0.00s)
  > PASS
  > ok  	github.com/hassan-shahbazi/swiftwasi/src	6.416s
  ```

## Benchmark
With the first version of swiftwasm _(wasm-5.3.1-RELEASE)_, a swift generated binary offers a larger binary with a significantly poor performance comparing to Rust.

### Size
The following comparison shows the binary size for a same set of functions in Rust and Swift, 1.7M and 9.8M respectively. The Swift binary is about 5 times larger than the Rust one.

```bash
$ ls -lh rust | grep wasm
-rwxrwxr-x 1 hassan hassan 1.7M Oct 30 20:33 binary.wasm

$ ls -lh swiftwasm | grep wasm
-rwxrwxr-x 1 hassan hassan 9.8M Oct 26 17:08 binary.wasm
```

### Performance
Larger binary size is not the only issue with Swift generated binaries. The more important issue is where _wasmer_ is almost 32 times slower in running Swift binaries comparing to Rust binaries with identical functions.

```bash
$ binary=rust go test ./... -race -count 1 -run _Rust
ok  	github.com/hassan-shahbazi/swiftwasi/src	0.214s

$ binary=swift go test ./... -race -count 1 -run _Swift
ok  	github.com/hassan-shahbazi/swiftwasi/src	6.899s
```

#### Optimization
Discussed in [#2135](https://github.com/swiftwasm/swift/issues/2135), using [`wasm-opt -0s`](https://github.com/webassembly/binaryen) can significantly improve Swift binary size **from 9.8M to 4.4M** as well as the performance **from 6.899s to 3.547s**.

```bash
$ ~/binaryen/bin/wasm-opt swiftwasm/binary.wasm -o swiftwasm/binary.wasm -Os

$ ls -lh swiftwasm | grep wasm
-rwxrwxr-x 1 hassan hassan 4.4M Nov 30 16:12 binary.wasm

$ binary=swift go test ./... -race -count 1 run _Swift
ok  	github.com/hassan-shahbazi/swiftwasi/src	3.547s
```
