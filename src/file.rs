use std::fs;
use std::path::Path;

pub fn read_file(file: &String) -> String {
    let filename = Path::new(&file);
    let contents = fs::read_to_string(filename)
        .expect("Something went wrong reading the file");
    return contents
}