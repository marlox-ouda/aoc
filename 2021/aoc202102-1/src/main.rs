use std::fs::File;
use std::io::{BufRead, BufReader};
use regex::Regex;
use lazy_static::lazy_static;

lazy_static! {
    static ref COMMAND_RE: Regex = Regex::new(r"(?x)^
        (?P<command>forward|up|down)
        \s+
        (?P<range>\d+)
        $")
        .unwrap();
}

struct Position {
    horizontal: i32,
    depth: i32,
}

impl Position {
    fn new() -> Position {
        Position {
            horizontal: 0,
            depth: 0,
        }
    }

    fn follow(&mut self, command: Command) {
        match command {
            Command::Up(range) => {self.depth -= range;},
            Command::Down(range) => {self.depth += range;},
            Command::Forward(range) => {self.horizontal += range;},
        };
    }

    fn result(&self) -> i32 {
        self.horizontal * self.depth
    }
}

enum Command {
    Forward(i32),
    Up(i32),
    Down(i32),
}

fn extract_command(line: &str) -> Command {
    let captures = COMMAND_RE
        .captures(&line)
        .expect("Regex does not match a line");
        //.expect(&format!("Regex does not match a line \"{}\"", &line));
    let range : i32 = captures
        .get(2)
        .unwrap()
        .as_str()
        .parse()
        .unwrap();
    let command = captures
        .get(1)
        .unwrap()
        .as_str();
    match command {
        "up" => Command::Up(range),
        "down" => Command::Down(range),
        "forward" => Command::Forward(range),
        _ => {
            panic!{"Unknown instruction"};
        },
    }
}


fn read_input(filename: &str) -> Vec<Command> {
    let file = File::open(filename)
        .expect("Fichier non trouvé");
    let reader = BufReader::new(file);
    let mut commands: Vec<Command> = Vec::new();

    for line in reader.lines() {
        let line = line.unwrap();
        commands.push(extract_command(&line));
    }
    return commands;
}

fn main() {
    let filename = "input.txt";
    let commands = read_input(&filename);
    let mut position = Position::new();
    for command in commands {
        position.follow(command);
    }
    println!("{}", position.result());
}
