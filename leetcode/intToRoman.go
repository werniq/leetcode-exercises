package main

func intToRoman(num int) string {
    in := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
    rm := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
    // romans[1] = "I"
    // romans[4] = "IV"
    // romans[5] = "V"
    // romans[9] = "IX"
    // romans[10] = "X"
    // romans[40] = "XL"
    // romans[50] = "L"
    // romans[90] = "XC"
    // romans[100] = "C"
    // romans[400] = "CD"
    // romans[500] = "D"
    // romans[900] = "CM"
    // romans[1000] = "M"
    answ := ""
    for a := 12; a >= 0; a-- {
        for num >= in[a] {
            answ += rm[a]
            num -= in[a]
        }
    }
    return answ
    // notes
    // 5312 mod 5 == 2 => II
    // 5312 mod 10 == 2 => II 
    // 5312 mod 50 == 12 => XII
    // 5312 mod 500 == 312 => CCCXII
    // 5312 mod 1000 => 312 => CCCXII
    // 5312 mod 10000 => 5312 => MMMMMCCCXII
}
