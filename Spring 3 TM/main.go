package main

import (
	defanginganipaddress "Spring_3_TM/defanging_an__ip_address"
	detectcapitaluse "Spring_3_TM/detectCapitalUse"
	lenghtOfLastWord "Spring_3_TM/lengthOfLastWord"
	mostwordsfound "Spring_3_TM/most_Words_Found"
	scoreofastring "Spring_3_TM/score_of_a_string"
	"fmt"
)

func main() {
	//Замена . на [.] в ip address
	var str string
	fmt.Scan(&str)
	fmt.Println(defanginganipaddress.DefangIPaddr(str))

	var str1 string
	fmt.Scan(&str1)
	fmt.Println(scoreofastring.ScoreOfString(str1))

	var str2 = []string{"alice and bob love CS2", "i can not think echkere", "this is great thanks very much"}
	fmt.Println(mostwordsfound.MostWordsFound(str2))

	var str3 string
	fmt.Scan(&str3)
	fmt.Println(detectcapitaluse.DetectCapitalUse(str3))

	var str4 string
	fmt.Scan(&str4)
	fmt.Println(lenghtOfLastWord.LengthOfLastWord(str4))
}
