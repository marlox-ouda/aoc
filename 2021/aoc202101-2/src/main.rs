use std::fs::File;
use std::io::{BufRead, BufReader};

fn read_input(filename: &str) -> Vec<u16> {
    let file = File::open(filename)
        .expect("Fichier non trouvé");
    let reader = BufReader::new(file);
    let mut numbers: Vec<u16> = Vec::new();

    for line in reader.lines() {
        let number : u16 = line
            .unwrap()
            .parse()
            .expect("Read line is not a number");
        numbers.push(number);
    }
    return numbers;
}

fn main() {
    let filename = "input.txt";
    let numbers = read_input(&filename);
    let numbers: Vec<u16> = numbers.windows(3).map(|window| window.iter().sum()).collect();
    let mut counter = 0;
    for window in numbers.windows(2) {
        let (prev, next) = (window[0], window[1]);
        if prev < next {
            counter += 1;
        }
    }
    println!("{}", counter);
}
