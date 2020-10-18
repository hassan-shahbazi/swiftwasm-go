import c_header

print("Hello World!")

@_cdecl("sum")
func sumFromHost(x: Int32, y: Int32) -> Int32 {
  return x + y
}

@_cdecl("allocate")
func allocate(size: Int) -> UnsafeMutableRawPointer {
  return UnsafeMutableRawPointer.allocate(byteCount: size, alignment: MemoryLayout<UInt8>.alignment)
}
@_cdecl("deallocate")
func deallocate(pointer: UnsafeMutableRawPointer, size: Int) {
  pointer.deallocate()
}

@_cdecl("concatenate")
func concatenateFromHost(s1: UnsafePointer<CChar>, s2: UnsafePointer<CChar>) -> UnsafeMutablePointer<CChar> {
  let str1 = String(cString: s1)
  let str2 = String(cString: s2)
  let result = str1 + " " + str2

  return UnsafeMutablePointer<CChar>(mutating: result)
}

@_cdecl("fetch")
func fetchCodefromHost(input: Int32) -> Int32 {
  return c_fetchCode(input);
}
