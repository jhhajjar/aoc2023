use std::{
    collections::HashMap,
    fs::File,
    io::{BufRead, BufReader},
};

fn find_number(input: String) -> i32 {
    let things_to_find = [
        "1", "2", "3", "4", "5", "6", "7", "8", "9", "one", "two", "three", "four", "five", "six",
        "seven", "eight", "nine",
    ];
    let num_map = HashMap::from([
        ("one", "1"),
        ("two", "2"),
        ("three", "3"),
        ("four", "4"),
        ("five", "5"),
        ("six", "6"),
        ("seven", "7"),
        ("eight", "8"),
        ("nine", "9"),
    ]);
    let mut last_number = String::from("0");
    let mut first_number = String::from("0");
    for thing in things_to_find {
        println!("yup {}", input.find(thing).is_some());
        if input.find(thing).is_some() {
            let parsed = num_map.get(thing).unwrap_or(&thing).to_owned();
            first_number = String::from(parsed);
        }
    }
    // println!("{}", first_number);
    let number = format!("{}{}", first_number, last_number);
    let number = number.parse().unwrap();

    number
}

fn main() {
    // read in the file line by line
    let file = File::open("i2.txt").expect("could not open file");
    let buf = BufReader::new(file);
    // for each line, loop through string, perform algo
    let r = buf.lines().map(|l| find_number(l.unwrap()));
    // add it all up
    let sum: i32 = r.sum();
    println!("{}", sum)
}
