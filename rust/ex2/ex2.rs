// https://doc.rust-lang.org/std/fmt/

fn main() {
    let x = 5;
    println!("x: {} ", x);
    let s = format!("y: {y:b}", y=8);
    println!("s: {0} ", s);
    println!("y: {y:0>5}", y=7);
    eprintln!("error!");

    struct Foo(i32);

    // fmt::Display
    // fmt::Debug
}
