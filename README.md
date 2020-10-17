# Swift + Web Assembly

This repository contains all code related to the tutorial: https://medium.com/@h.shahbazi/the-power-of-swift-web-assembly-part-3-e583c6ab8afe



## Requirements

- Ubuntu 20.04 and above
- Go
- swift-wasm



## Installation

- To install **swiftenv**, follow instructions on: https://swiftenv.fuller.li/en/latest/installation.html

- Then, download the latest version of **swiftwasm** from the release page: https://github.com/swiftwasm/swift/releases

- Install some libraries as dependencies for running Swift on Linux

  ```bash
  $ apt update 
  $ apt install binutils git gnupg2 libc6-dev libcurl4 libedit2 libgcc-9-dev libpython2.7 libsqlite3-0 libstdc++-9-dev libxml2 libz3-dev pkg-config tzdata zlib1g-dev curl lsb-release
  ```

- Install **wasmer** by following instructions on: https://wasmer.io/



## Running

1. Set the correct version of Swiftwasm

   ```bash
   $ swiftenv global 
   ```

2. Build Swift project and generate wasm binary:

   ```bash
   $ TOOLCHAIN_PATH=$(cd $(dirname "$(swiftenv which swiftc)") && cd ../share && pwd)
   $ swift build --triple wasm32-unknown-wasi -c release --toolchain $TOOLCHAIN_PATH
   ```

3.  Copy wasm to the project's directory

   ```bash
   $ cp swiftwasm/.build/release/swiftwasm ../binary.wasm
   ```

4. Run the Go application

   ```bash
   $ go run ./
   > 4
   > 10
   ```

   