package main

import (
	"bufio"
	"os"
)

func main() {
	// file, _ := os.Open("./file.txt")
	// scanner := bufio.NewScanner(file)
	// scanner.Split(bufio.ScanLines)

	// var lines []string

	// for scanner.Scan() {
	// 	lines = append(lines, scanner.Text())
	// }
	// file.Close()
	// var line []string

	// for _, x := range lines {
	// 	line = append(line, x)
	// }
	// fmt.Println(line)
	// for i, j := 0, len(line)-1; i < j; i, j = i+1, j-1 {
	// 	line[i], line[j] = line[j], line[i]
	// }
	// fmt.Println(line)

	lines := []string{"RTR1 1.1.1.1", "RTR2 2.2.2.2", "RTR3 3.3.3.3", "RTR4 4.4.4.4"}
	file, _ := os.OpenFile("./file.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	writer := bufio.NewWriter(file)

	for _, line := range lines {
		_, _ = writer.WriteString(line + "\n")

	}
	writer.Flush()
	file.Close()

}
