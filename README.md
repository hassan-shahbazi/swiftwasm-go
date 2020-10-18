# Swift + Web Assembly
This repository contains all code related to a 3-part tutorial published on Medium: [Part 1](https://medium.com/@h.shahbazi/the-power-of-swift-web-assembly-part-1-fdfa4e9134ee), [Part 2](https://medium.com/@h.shahbazi/the-power-of-swift-web-assembly-part-2-30b6c4619c27), and [Part 3](https://medium.com/@h.shahbazi/the-power-of-swift-web-assembly-part-3-e583c6ab8afe)



## Requirements
- Ubuntu >= 18.04. (macOS naturally works fine, but I have not tested it yet)
- Go >= 1.15
- swift-wasm >= DEVELOPMENT-SNAPSHOT-2020-10-17-a


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
   $ swiftenv global wasm-DEVELOPMENT-SNAPSHOT-2020-10-17-a
   ```

2. Build Swift project and generate wasm binary:
   ```bash
   $ TOOLCHAIN_PATH=$(cd $(dirname "$(swiftenv which swiftc)") && cd ../share && pwd)
   $ swift build --triple wasm32-unknown-wasi -c release --toolchain $TOOLCHAIN_PATH -Xlinker --export=allocate -Xlinker --export=deallocate -Xlinker --export=hello -Xlinker --export=sum -Xlinker --export=concatenate -Xlinker --export=fetch -Xlinker --allow-undefined
   ```

3.  Copy wasm to the project's directory
   ```bash
   $ cp swiftwasm/.build/release/swiftwasm ../binary.wasm
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
