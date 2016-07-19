// Source : https://leetcode.com/problems/zigzag-conversion/
// Author : Code.CC
// Date   : 2016-07-19

/**********************************************************************************
*
* The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this:
* (you may want to display this pattern in a fixed font for better legibility)
*
* P   A   H   N
* A P L S I I G
* Y   I   R
*
* And then read line by line: "PAHNAPLSIIGYIR"
*
* Write the code that will take a string and make this conversion given a number of rows:
*
* string convert(string text, int nRows);
*
* convert("PAYPALISHIRING", 3) should return "PAHNAPLSIIGYIR"
*
*
**********************************************************************************/

package main

import (
	"fmt"
)

func zigZagConversion(text string, nRows int) string {
	ctext := make([][]string, nRows)
	row := 0
	step := 1
	for i := 0; i < len(text); i++ {
		if row == nRows-1 {
			step = -1
		}
		if row == 0 {
			step = 1
		}
		ctext[row] = append(ctext[row], string(text[i]))
		row += step
	}

	var result string
	for i := 0; i < nRows; i++ {
		for j := 0; j < len(ctext[i]); j++ {
			result = result + string(ctext[i][j])
		}
	}

	return result

}

func main() {
	text := "PAYPALISHIRING"
	nRows := 3
	fmt.Println(zigZagConversion(text, nRows))
}
