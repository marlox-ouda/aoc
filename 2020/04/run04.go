package main

import (
  "io/ioutil"
  "fmt"
  "os"
  "strconv"
  "strings"
  "unicode"
)

type passport struct {
  byr string
  iyr string
  eyr string
  hgt string
  hcl string
  ecl string
  pid string
  cid string
}

var example_passwords_input = `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in
`

func extractPassportLines(passports_blob string) []string {
  var (
    index int
    passport_lines []string
    passport_line string
  )
  passport_lines = strings.Split(passports_blob, "\n\n")
  for index, passport_line = range passport_lines {
    passport_lines[index] = strings.Trim(passport_line, "\n")
  }
  return passport_lines
}

func extractPassportData(password_line string) (*passport, error) {
  var (
    pass passport
    password_items []string
    key_value []string
    item string
    key string
    value string
  )
  password_line = strings.Trim(password_line, "\n")
  password_line = strings.ReplaceAll(password_line, "\n", " ")
  password_items = strings.Split(password_line, " ")
  for _, item = range password_items {
    key_value = strings.SplitN(item, ":", 2)
    if len(key_value) != 2 {
      return &pass, fmt.Errorf("Fail to extract key:value pair from '%s' in %s", item, password_line)
    }
    value = key_value[1]
    switch key = key_value[0]; key {
    case "byr":
      pass.byr = value
    case "iyr":
      pass.iyr = value
    case "eyr":
      pass.eyr = value
    case "hgt":
      pass.hgt = value
    case "hcl":
      pass.hcl = value
    case "ecl":
      pass.ecl = value
    case "pid":
      pass.pid = value
    case "cid":
      pass.cid = value
    default:
      return &pass, fmt.Errorf("Unkown key '%s' with value '%s' from '%s'", key, value, password_line)
    }
  }
  return &pass, nil
}

func checkPassportRequiredField(pass *passport) bool {
  if pass.byr != "" && pass.iyr != "" && pass.eyr != "" && pass.hgt != "" && pass.hcl != "" && pass.ecl != "" && pass.pid != "" {
    return true
  }
  return false
}

func checkByr(byr_value string) bool {
  var (
    byr_int int
    err error
  )
  byr_int, err = strconv.Atoi(byr_value)
  if err == nil && 1920 <= byr_int && byr_int <= 2002 {
    return true
  }
  return false
}

func checkIyr(iyr_value string) bool {
  var (
    iyr_int int
    err error
  )
  iyr_int, err = strconv.Atoi(iyr_value)
  if err == nil && 2010 <= iyr_int && iyr_int <= 2020 {
    return true
  }
  return false
}

func checkEyr(eyr_value string) bool {
  var (
    eyr_int int
    err error
  )
  eyr_int, err = strconv.Atoi(eyr_value)
  if err == nil && 2020 <= eyr_int && eyr_int <= 2030 {
    return true
  }
  return false
}

func checkHgt(hgt_value string) bool {
  var (
    hgt_int int
    err error
  )
  if strings.HasSuffix(hgt_value, "in") {
    hgt_int, err = strconv.Atoi(hgt_value[:len(hgt_value)-2])
    if err == nil && 59 <= hgt_int && hgt_int <= 76 {
      return true
    }
  } else if strings.HasSuffix(hgt_value, "cm") {
    hgt_int, err = strconv.Atoi(hgt_value[:len(hgt_value)-2])
    if err == nil && 150 <= hgt_int && hgt_int <= 193 {
      return true
    }
  }
  return false
}

func checkHcl(hcl_value string) bool {
  var err error
  if strings.HasPrefix(hcl_value, "#") {
    _, err = strconv.ParseInt(hcl_value[1:], 16, 64)
    if err == nil {
      return true
    }
  }
  return false
}

var allowed_ecl_values = []string{
  "amb",
  "blu",
  "brn",
  "gry",
  "grn",
  "hzl",
  "oth",
}

func checkEcl(ecl_value string) bool {
  var checked_ecl_value string
  for _, checked_ecl_value = range allowed_ecl_values {
    if checked_ecl_value == ecl_value {
      return true
    }
  }
  return false
}

func checkPid(pid_value string) bool {
  var (
    char rune
  )
  if len(pid_value) == 9 {
    for _, char = range pid_value {
      if !unicode.IsDigit(char) {
        return false
      }
    }
    return true
  }
  return false
}


func commonOne(passports_blob string) int {
  var (
    valid_passports_number int
    passport_lines []string
    passport_line string
    pass *passport
    err error
  )
  passport_lines = extractPassportLines(passports_blob)
  for _, passport_line = range passport_lines {
    pass, err = extractPassportData(passport_line)
    if err == nil {
      if checkPassportRequiredField(pass) {
        valid_passports_number += 1
      }
    } else {
      fmt.Printf("Error: %s\n", err)
    }
  }
  return valid_passports_number
}

func exampleOne() {
  var valid_passports_number int
  valid_passports_number = commonOne(example_passwords_input)
  fmt.Printf("ExampleOne: %v\n", valid_passports_number)
}


func mainOne() {
  var (
    valid_passports_number int
    input []byte
    err error
  )
  input, err = ioutil.ReadFile("input4.txt")
  if err == nil {
    valid_passports_number = commonOne(string(input))
    fmt.Printf("MainOne: %v\n", valid_passports_number)
  } else {
    fmt.Printf("MainOne: I/O erro : %s\n", err)
  }
}

func commonTwo(passports_blob string) int {
  var (
    valid_passports_number int
    passport_lines []string
    passport_line string
    pass *passport
    err error
  )
  passport_lines = extractPassportLines(passports_blob)
  for _, passport_line = range passport_lines {
    pass, err = extractPassportData(passport_line)
    if err == nil {
      if checkPassportRequiredField(pass) {
        valid_passports_number += 1
      }
    } else {
      fmt.Printf("Error: %s\n", err)
    }
  }
  return valid_passports_number
}

func exampleTwo() {
  var valid_passports_number int
  valid_passports_number = commonOne(example_passwords_input)
  fmt.Printf("ExampleTwo: %v\n", valid_passports_number)
}


func mainTwo() {
  var (
    valid_passports_number int
    input []byte
    err error
  )
  input, err = ioutil.ReadFile("input4.txt")
  if err == nil {
    valid_passports_number = commonOne(string(input))
    fmt.Printf("MainTwo: %v\n", valid_passports_number)
  } else {
    fmt.Printf("MainTwo: I/O erro : %s\n", err)
  }
}

func main() {
  request := "all"
  if len(os.Args) > 1 {
    request = os.Args[1]
  }
  if request == "1" || request == "example1" || request == "all" {
    exampleOne()
  }
  if request == "1" || request == "input1" || request == "all" {
    mainOne()
  }
  if request == "2" || request == "example2" || request == "all" {
    exampleTwo()
  }
  if request == "2" || request == "input2" || request == "all" {
    mainTwo()
  }
}
