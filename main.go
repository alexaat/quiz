package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var correntAnswers = 0
var data = map[string]int{}

func main() {

	fileName := flag.String("path", "problems.csv", "The path to input file")
	timer := flag.Int("time", 30, "Set game timer in seconds")
	flag.Parse()

	data = readFile(*fileName)

	fmt.Println("Press Enter to start")
	fmt.Scanln()

	go Timer(*timer)

	for key, value := range data {
		var answer string
		fmt.Printf("What is %v?\n", key)
		fmt.Scan(&answer)
		v, err := strconv.Atoi(answer)
		if err != nil {
			v = value + 1
		}
		if v == value {
			correntAnswers++
		}

	}

	fmt.Printf("Your score is %v out of %v\n", correntAnswers, len(data))
}

func readFile(fileName string) map[string]int {
	f, err := os.Open(fileName)
	if err != nil {
		exit(fmt.Sprintf("Unable to open %v\n", fileName))
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		exit(fmt.Sprintf("Unable to read %v\n", fileName))
	}

	var mapResult = map[string]int{}

	for _, row := range data {
		keyIndex := 0
		for i, field := range row {
			if strings.TrimSpace(field) != "" {
				keyIndex = i
				break
			}
		}

		key := strings.TrimSpace(row[keyIndex])
		value := strings.TrimSpace(row[keyIndex+1])
		if key != "" && value != "" {
			v, err := strconv.Atoi(value)
			if err != nil {
				exit(fmt.Sprint("Unable to convert string to int"))
			}
			mapResult[key] = v
		}

	}
	return mapResult
}

func Timer(t int) {

	timer := time.NewTimer(time.Duration(t) * time.Second)
	<-timer.C
	var scoreMsg = fmt.Sprintf("Your score is %v out of %v\n", correntAnswers, len(data))
	exit("Time is out", scoreMsg)

}

func exit(msgs ...string) {
	for _, msg := range msgs {
		fmt.Println(msg)
	}
	os.Exit(0)
}
