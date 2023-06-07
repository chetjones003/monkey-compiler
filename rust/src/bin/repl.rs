#![allow(dead_code)]

use monkey_compiler::lexer::lexer::{Lexer, Token};
use anyhow::Result;

fn main() -> Result<()> {
    println!("Hello and welcome to the Monkey Language REPL");
    println!("Feel free to type in commands");
    std::io::stdin().lines().for_each(|line| {
        if let Ok(line) = line {
            let mut tokenizer = Lexer::new(line);

            while let Ok(token) = tokenizer.next_token() {
                println!("{} ", token);
                if let Token::EOF = token {
                    break;
                }
            }
        }
    });
    return Ok(());
}
