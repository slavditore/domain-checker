mod file;
extern crate yaml_rust;
use yaml_rust::{YamlLoader};


fn main() {
    let file: String  = "./src/config.yml".to_string();
    let content = file::read_file(&file);
    println!("{}", content);
    let doc = &(YamlLoader::load_from_str(&content).unwrap())[0];
    for g in doc["groups"].as_vec().unwrap(){
        for d in g["domains"].as_vec().unwrap(){
            println!("{}", d.as_str().unwrap())
        }
    }
}
