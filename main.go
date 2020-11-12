package main

import (
  "encoding/csv"
  "math/rand"
  "fmt"
  "log"
  "os"
  "time"
  "strings"
)

func main(){
  problems := "problems.csv"
  file, err := os.Open(problems)
  if err != nil {
    log.Fatal(err)
  }
  r := csv.NewReader(file)
  data, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
  question := parseLines(data)

  correct := 0
  count := 0
  random := rand.New(rand.NewSource(time.Now().UnixNano()))
  for {
    i := random.Intn(len(data))
    fmt.Printf("What is %s=?> ", question[i].q)
    var answer string
    fmt.Scanf("%s ", &answer)
    if err != nil {
      fmt.Println(err)
    } else {
      if answer == "q" {
        break
      } else if answer == question[i].a {
        fmt.Println("correct")
        correct++
      } else {
        fmt.Println("wrong")
      }
    }
    count++
  }
  fmt.Printf("You scored %d out of %d.\n", correct, count)
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}