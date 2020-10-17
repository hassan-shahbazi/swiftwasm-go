// swift-tools-version:5.3
// The swift-tools-version declares the minimum version of Swift required to build this package.

import PackageDescription

let package = Package(
    name: "swiftwasm",
    targets: [
      .target(name: "c_header", dependencies: []),
      .target(name: "swiftwasm", dependencies: ["c_header"]),
    ]
)
