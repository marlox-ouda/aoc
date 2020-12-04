package main

import "testing"

func TestExtractPassportLines(t *testing.T) {
  var (
    given_result []string
    expected_result string
  )
  given_result = extractPassportLines(example_passwords_input)
  if len(given_result) != 4 {
    t.Errorf("len(extractPassportLines(<example>)), %v given, 4 expected", len(given_result))
  } else {
    expected_result = "ecl:gry pid:860033327 eyr:2020 hcl:#fffffd\nbyr:1937 iyr:2017 cid:147 hgt:183cm"
    if given_result[0] != expected_result {
      t.Errorf("extractPassportLines(<example>)[0], %s given, %s expected", given_result[0], expected_result)
    }
    expected_result = "iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884\nhcl:#cfa07d byr:1929"
    if given_result[1] != expected_result {
      t.Errorf("extractPassportLines(<example>)[1], %s given, %s expected", given_result[1], expected_result)
    }
    expected_result = "hcl:#ae17e1 iyr:2013\neyr:2024\necl:brn pid:760753108 byr:1931\nhgt:179cm"
    if given_result[2] != expected_result {
      t.Errorf("extractPassportLines(<example>)[2], %s given, %s expected", given_result[2], expected_result)
    }
    expected_result = "hcl:#cfa07d eyr:2025 pid:166559648\niyr:2011 ecl:brn hgt:59in"
    if given_result[3] != expected_result {
      t.Errorf("extractPassportLines(<example>)[3], %s given, %s expected", given_result[3], expected_result)
    }
  }
}

func TestExtractPassportDataOne(t *testing.T) {
  var (
    given_result *passport
    err error
  )
  given_result, err = extractPassportData("ecl:gry pid:860033327 eyr:2020 hcl:#fffffd\nbyr:1937 iyr:2017 cid:147 hgt:183cm")
  if err != nil {
    t.Errorf("extractPassport(<1>) raise error: %s", err)
  } else {
    if given_result.byr != "1937" {
        t.Errorf("extractPassport(<1>).byr, %s given, \"1937\" expected", given_result.byr)
    }
    if given_result.iyr != "2017" {
        t.Errorf("extractPassport(<1>).iyr, %s given, \"2017\" expected", given_result.iyr)
    }
    if given_result.eyr != "2020" {
        t.Errorf("extractPassport(<1>).eyr, %s given, \"2020\" expected", given_result.eyr)
    }
    if given_result.hgt != "183cm" {
        t.Errorf("extractPassport(<1>).hgt, %s given, \"183cm\" expected", given_result.hgt)
    }
    if given_result.hcl != "#fffffd" {
        t.Errorf("extractPassport(<1>).hcl, %s given, \"#fffffd\" expected", given_result.hcl)
    }
    if given_result.ecl != "gry" {
        t.Errorf("extractPassport(<1>).ecl, %s given, \"gry\" expected", given_result.ecl)
    }
    if given_result.pid != "860033327" {
        t.Errorf("extractPassport(<1>).pid, %s given, \"860033327\" expected", given_result.pid)
    }
    if given_result.cid != "147" {
        t.Errorf("extractPassport(<1>).cid, %s given, \"147\" expected", given_result.cid)
    }
  }
}

func TestExtractPassportDataTwo(t *testing.T) {
  var (
    given_result *passport
    err error
  )
  given_result, err = extractPassportData("iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884\nhcl:#cfa07d byr:1929")
  if err != nil {
    t.Errorf("extractPassport(<2>) raise error: %s", err)
  } else {
    if given_result.byr != "1929" {
        t.Errorf("extractPassport(<2>).byr, %s given, \"1929\" expected", given_result.byr)
    }
    if given_result.iyr != "2013" {
        t.Errorf("extractPassport(<2>).iyr, %s given, \"2013\" expected", given_result.iyr)
    }
    if given_result.eyr != "2023" {
        t.Errorf("extractPassport(<2>).eyr, %s given, \"2023\" expected", given_result.eyr)
    }
    if given_result.hgt != "" {
        t.Errorf("extractPassport(<2>).hgt, %s given, \"\" expected", given_result.hgt)
    }
    if given_result.hcl != "#cfa07d" {
        t.Errorf("extractPassport(<2>).hcl, %s given, \"#cfa07d\" expected", given_result.hcl)
    }
    if given_result.ecl != "amb" {
        t.Errorf("extractPassport(<2>).ecl, %s given, \"amb\" expected", given_result.ecl)
    }
    if given_result.pid != "028048884" {
        t.Errorf("extractPassport(<2>).pid, %s given, \"028048884\" expected", given_result.pid)
    }
    if given_result.cid != "350" {
        t.Errorf("extractPassport(<2>).cid, %s given, \"350\" expected", given_result.cid)
    }
  }
}

func TestCheckPassportRequiredField(t * testing.T) {
  var (
    given_result bool
    pass passport
  )
  pass = passport{"1937", "2017", "2020", "183cm", "#fffffd", "gry", "860033327", "147"}
  given_result = checkPassportRequiredField(&pass)
  if !given_result {
    t.Error("checkPassportRequiredField(<0>), false given, true expected")
  }
  pass = passport{"1929", "2013", "2023", "", "#cfa07d", "amb", "028048884", "350"}
  given_result = checkPassportRequiredField(&pass)
  if given_result {
    t.Error("checkPassportRequiredField(<1>), true given, false expected")
  }
  pass = passport{"1931", "2013", "2024", "179cm", "#ae17e1", "brn", "760753108", ""}
  given_result = checkPassportRequiredField(&pass)
  if !given_result {
    t.Error("checkPassportRequiredField(<2>), false given, true expected")
  }
  pass = passport{"", "2011", "2025", "59in", "#cfa07d", "brn", "166559648", ""}
  given_result = checkPassportRequiredField(&pass)
  if given_result {
    t.Error("checkPassportRequiredField(<3>), true given, false expected")
  }
}

func TestCheckByr(t * testing.T) {
  var given_result bool
  given_result = checkByr("2002")
  if !given_result {
    t.Error("checkByr(\"2002\"), false given, true expected")
  }
  given_result = checkByr("2003")
  if given_result {
    t.Error("checkByr(\"2003\"), true given, false expected")
  }
}

func TestCheckHgt(t * testing.T) {
  var given_result bool
  given_result = checkHgt("60in")
  if !given_result {
    t.Error("checkHgt(\"60in\"), false given, true expected")
  }
  given_result = checkHgt("190cm")
  if !given_result {
    t.Error("checkHgt(\"190cm\"), false given, true expected")
  }
  given_result = checkHgt("190in")
  if given_result {
    t.Error("checkHgt(\"190in\"), true given, false expected")
  }
  given_result = checkHgt("190")
  if given_result {
    t.Error("checkHgt(\"190\"), true given, false expected")
  }
}

func TestCheckHcl(t * testing.T) {
  var given_result bool
  given_result = checkHcl("#123abc")
  if !given_result {
    t.Error("checkHcl(\"#123abc\"), false given, true expected")
  }
  given_result = checkHcl("#123abz")
  if given_result {
    t.Error("checkHcl(\"#123abz\"), true given, false expected")
  }
  given_result = checkHcl("123abc")
  if given_result {
    t.Error("checkHcl(\"123abc\"), true given, false expected")
  }
}

func TestCheckEcl(t * testing.T) {
  var given_result bool
  given_result = checkEcl("brn")
  if !given_result {
    t.Error("checkEcl(\"brn\"), false given, true expected")
  }
  given_result = checkEcl("wat")
  if given_result {
    t.Error("checkEcl(\"wat\"), true given, false expected")
  }
}

func TestCheckPid(t * testing.T) {
  var given_result bool
  given_result = checkPid("000000001")
  if !given_result {
    t.Error("checkPid(\"000000001\"), false given, true expected")
  }
  given_result = checkPid("0123456789")
  if given_result {
    t.Error("checkPid(\"0123456789\"), true given, false expected")
  }
}
