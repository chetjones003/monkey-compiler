use crate::lexer::Lexer;
use crate::lexer::Token;

const PROMPT:&str = ">> ";

pub fn start(start_in: std::io::Stdin, start_out: std::io::Stdout) {
    let mut scan = String::new();
    let scanner = std::io::stdin()
        .read_line(&mut scan)
        .expect("Failed to read line");
    let scanner = start_in;
    println!("{:?}", scanner);
}
