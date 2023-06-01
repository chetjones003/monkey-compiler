use std::io::BufRead;
use crate::lexer::Lexer;
use crate::lexer::Token;

const PROMPT:&str = ">> ";

pub fn start(start_in: String, start_out: String) {
    let mut scanner = std::io::stdin()
        .lock()
        .lines()
        .map(|x| x.unwrap())
        .collect();

    let mut lex = Lexer::new(scanner);

    let tokens = Lexer::next_token(&mut lex);
    while tokens != Token::EOF {
        println!("{start_out}, {tokens}");
    }
}
