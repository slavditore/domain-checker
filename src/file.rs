use std::fs;
use std::path::Path;
extern crate yaml_rust;
use yaml_rust::{YamlLoader};

pub fn read_file(file: &String) -> String {
    let filename = Path::new(&file);
    let contents = fs::read_to_string(filename)
        .expect("Something went wrong reading the file");
    return contents
}

pub fn get_domains(content: String){
    let doc = &(YamlLoader::load_from_str(&content).unwrap())[0];
    for g in doc["groups"].as_vec().unwrap(){
        for d in g["domains"].as_vec().unwrap(){
            println!("{}", d.as_str().unwrap())
        }
    }
}