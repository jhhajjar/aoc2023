use std::fs;

fn main() {
    // Read the calibration document from a file
    let input = fs::read_to_string("calibration_document.txt").expect("Unable to read file");

    // Calculate the sum of calibration values
    let sum: u64 = input
        .lines()
        .map(|line| calculate_calibration_value(line))
        .sum();

    println!("Sum of calibration values: {}", sum);
}

fn calculate_calibration_value(line: &str) -> u64 {
    // Find the first and last numbers in the line and concatenate them
    let mut result = String::new();
    let mut first_digit_found = false;

    for c in line.chars() {
        if c.is_digit(10) {
            result.push(c);
            first_digit_found = true;
        } else if first_digit_found && (c.is_alphabetic() || c.is_whitespace()) {
            // Stop when an alphabetic character or whitespace is encountered after the first digit
            break;
        }
    }

    // Reverse the line and repeat the process to find the last number
    let mut reversed_line = line.chars().rev();
    let mut last_digit_found = false;

    for c in reversed_line.by_ref() {
        if c.is_digit(10) {
            result.push(c);
            last_digit_found = true;
        } else if last_digit_found && (c.is_alphabetic() || c.is_whitespace()) {
            // Stop when an alphabetic character or whitespace is encountered after the last digit
            break;
        }
    }

    // Reverse the result string to get the correct order of digits
    let reversed_result: String = result.chars().rev().collect();

    // Convert the concatenated result to u64
    reversed_result.parse().unwrap_or(0)
}

// use std::{
//     collections::HashMap,
//     fs::File,
//     io::{BufRead, BufReader},
// };

// fn find_number(input: String) -> i32 {
//     let things_to_find = [
//         "1", "2", "3", "4", "5", "6", "7", "8", "9", "one", "two", "three", "four", "five", "six",
//         "seven", "eight", "nine",
//     ];
//     let num_map = HashMap::from([
//         ("one", "1"),
//         ("two", "2"),
//         ("three", "3"),
//         ("four", "4"),
//         ("five", "5"),
//         ("six", "6"),
//         ("seven", "7"),
//         ("eight", "8"),
//         ("nine", "9"),
//     ]);
//     let mut last_number = String::from("0");
//     let mut first_number = String::from("0");
//     for thing in things_to_find {
//         println!("yup {}", input.find(thing).is_some());
//         if input.find(thing).is_some() {
//             let parsed = num_map.get(thing).unwrap_or(&thing).to_owned();
//             first_number = String::from(parsed);
//         }
//     }
//     // println!("{}", first_number);
//     let number = format!("{}{}", first_number, last_number);
//     let number = number.parse().unwrap();

//     number
// }

// fn main() {
//     // read in the file line by line
//     let file = File::open("i2.txt").expect("could not open file");
//     let buf = BufReader::new(file);
//     // for each line, loop through string, perform algo
//     let r = buf.lines().map(|l| find_number(l.unwrap()));
//     // add it all up
//     let sum: i32 = r.sum();
//     println!("{}", sum)
// }
