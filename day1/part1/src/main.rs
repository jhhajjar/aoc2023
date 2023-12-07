use std::{
    fs::File,
    io::{BufRead, BufReader},
};

fn find_number(input: String) -> i32 {
    let mut looking_at_first_number = true;
    let mut last_number = String::from("0");
    let mut first_number = String::from("0");
    for letter in input.chars() {
        if letter.is_digit(10) && looking_at_first_number {
            looking_at_first_number = false;
            first_number = String::from(letter);
            last_number = String::from(letter);
        } else if letter.is_digit(10) {
            last_number = String::from(letter);
        }
    }
    let number = format!("{}{}", first_number, last_number);
    let number = number.parse().unwrap();

    number
}

fn main() {
    // read in the file line by line
    let file = File::open("input.txt").expect("could not open file");
    let buf = BufReader::new(file);
    // for each line, loop through string, perform algo
    let r = buf.lines().map(|l| find_number(l.unwrap()));
    // add it all up
    let sum: i32 = r.sum();
    println!("{}", sum)
}
