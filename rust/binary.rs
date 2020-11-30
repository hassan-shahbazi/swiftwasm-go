fn main() {
    println!("Hello, world!");
}

#[no_mangle]
pub extern fn sum(x: i32, y: i32) -> i32 {
    return x + y
}

extern {
    fn fetch_code(input: i32) -> i32;
}
#[no_mangle]
pub extern fn fetch(input: i32) -> i32 {
    unsafe { fetch_code(input) }
}
