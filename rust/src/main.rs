mod lexer;
mod repl;

fn main() {
    println!("Hello and welcome to the Monkey Language REPL");
    println!("Feel free to type in commands");
    repl::start(std::io::stdin(), std::io::stdout());
}
