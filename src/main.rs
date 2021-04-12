mod file;

fn main() {
    let file: String  = "./src/config.yml".to_string();
    let content = file::read_file(&file);
    println!("{}", content)
}
