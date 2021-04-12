mod file;


fn main() {
    let file: String  = "./src/config.yml".to_string();
    let content = file::read_file(&file);
    println!("{}", content);
    let domains = file::get_domains(content);
    for d in domains{
        println!("You'll check the next domain: {}", d)
    }
}
