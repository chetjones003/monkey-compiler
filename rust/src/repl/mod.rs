use crate::lexer::Lexer;
use crate::lexer::Token;

const PROMPT:&str = ">> ";

pub fn start(start_in: String, start_out: String) {
    let mut scanner = String::new();
    scanner = start_in;
    println!("{scanner}");
}
