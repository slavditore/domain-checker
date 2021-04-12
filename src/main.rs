use std::fs;
use std::path::Path;

fn main() {
    let filename = Path::new("./src/config.yml");
    let contents = fs::read_to_string(filename)
        .expect("Something went wrong reading the file");
    println!("{}", contents)
}
