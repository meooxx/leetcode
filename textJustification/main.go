package main

import "fmt"

func main() {
	fmt.Printf("%#v", fullJustify([]string{"he", "is", "world"}, 6))
	fmt.Printf("%#v", fullJustify([]string{
		"What", "must", "be", "acknowledgment", "shall", "be",
	}, 16))
	fmt.Printf("%#v", fullJustify([]string{
		"acknowledgment", "we",
	}, 16))
	fmt.Printf("%#v", fullJustify([]string{
		"This", "is", "an", "example", "of", "text", "justification.",
	}, 16))

}

/*
*   hello hello world
*  [hello  world]
*  [hello hello world]
 */

func fullJustify3(words []string, maxWidth int) []string {
	result := []string{}

	// 单行 words + space  `hello world`
	rowLen := 0
	left := 0
	row := make([]byte, maxWidth)
	// 单行 words长度 `hellworld`
	rowWordLen := 0
	for i := 0; i < len(words); i++ {
		rowLen += len(words[i])
		rowWordLen += len(words[i])
		// the last line word
		if i == len(words)-1 && rowLen <= maxWidth {
			rowIndex := 0
			for left <= i {
				// [hello] word => [hello word]
				if rowIndex > 0 {
					row[rowIndex] = ' '
					rowIndex++
				}
				for wIndex := 0; wIndex < len(words[left]); wIndex++ {
					row[rowIndex] = words[left][wIndex]
					rowIndex++
				}
				left++
			}
			// hello => helloxxxx
			for rowIndex < maxWidth {
				row[rowIndex] = ' '
				rowIndex++
			}
			result = append(result, string(row))
			break
		}

		if rowLen > maxWidth {
			// rowLen > maxWidth
			// back to i-1,  rowLen must be <= maxWidth
			rowWordLen -= len(words[i])
			i -= 1
			// firstW      lastW
			// hello       word
			// 2
			// hello hello word
			firstW := words[left]
			lastW := words[i]
			if i == left {
				lastW = ""
			}
			for wIndex := 0; wIndex < len(firstW); wIndex++ {
				row[wIndex] = firstW[wIndex]
			}

			wCount := i - left + 1
			left += 1

			// aleast 2 words
			if wCount >= 2 {
				for wIndex := 0; wIndex < len(lastW); wIndex++ {
					//    -1
					// [hell]   hell
					row[maxWidth-len(lastW)+wIndex] = lastW[wIndex]
				}
			}
			// 5 / 2  2...1  xxxhelloxxwold,  x for space
			spaces := maxWidth - rowWordLen
			preWordsSpaces := spaces
			restSpaces := 0
			if wCount-2 > 0 {
				preWordsSpaces = spaces / (wCount - 2 + 1)
				restSpaces = spaces % (wCount - 2 + 1)
			}
			rowStart := len(firstW)
			for rowStart < maxWidth-len(lastW) {
				for s := 0; s < preWordsSpaces; s++ {
					row[rowStart] = ' '
					rowStart++
				}
				if restSpaces > 0 {
					row[rowStart] = ' '
					rowStart++
					restSpaces--
				}
				if wCount-2 > 0 && left < i {
					for l := 0; l < len(words[left]); l++ {
						row[rowStart] = words[left][l]
						rowStart++
					}
					left++
				}

			}

			result = append(result, string(row))
			// rest all var
			row = make([]byte, maxWidth)
			rowLen = 0
			left = i + 1
			rowWordLen = 0
			continue

		}
		rowLen += 1
	}
	return result
}

func fullJustify(words []string, maxWidth int) []string {
	result := []string{}
	letterCount := 0 // word only
	sum := 0         // including spaces
	start := 0       // first word in row
	for wordIndex := 0; wordIndex < len(words); wordIndex++ {
		word := words[wordIndex]
		letterCount += len(word)
		// there're aleast one space before every word
		if sum > 0 {
			sum++
		}
		sum += len(word)

		if wordIndex == len(words)-1 && sum <= maxWidth {
			row := make([]rune, maxWidth)
			rowIndex := 0

			for i := start; i <= wordIndex; i++ {
				word := words[i]
				if i > start {
					row[rowIndex] = ' '
					rowIndex++
				}
				for _, c := range word {
					row[rowIndex] = c
					rowIndex++
				}

			}
			for rowIndex < maxWidth {
				row[rowIndex] = ' '
				rowIndex++
			}
			result = append(result, string(row))
		} else if sum > maxWidth {
			row := make([]rune, maxWidth)
			rowIndex := 0
			letterCount -= len(words[wordIndex])
			wordIndex--
			// n words split n-1 spans
			span := wordIndex - start + 1 - 1
			if span == 0 {
				span = 1
			}
			// spaces before the word
			spaces := (maxWidth - letterCount) / span
			remainSpaces := (maxWidth - letterCount) % span

			// place the first word of row
			for _, c := range words[start] {
				row[rowIndex] = c
				rowIndex++
			}
			for wIdx := start + 1; wIdx <= wordIndex; wIdx++ {
				for l := 0; l < spaces; l++ {
					row[rowIndex] = ' '
					rowIndex++
				}
				if remainSpaces > 0 {
					row[rowIndex] = ' '
					rowIndex++
					remainSpaces--
				}
				for _, c := range words[wIdx] {
					row[rowIndex] = c
					rowIndex++
				}
			}
			for rowIndex < maxWidth {
				row[rowIndex] = ' '
				rowIndex++
			}
			result = append(result, string(row))

			start = wordIndex + 1
			sum = 0
			letterCount = 0
		}

	}

	return result

}
