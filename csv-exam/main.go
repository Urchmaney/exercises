package main

import (
  "fmt"
  "io/ioutil"
  "os"
  "encoding/csv"
  "log"
  "bufio"
  "strings"
  "io"
  "flag"
)

func  main()  {
	file_flag := flag.String("file_name", "problems.csv", "Full File Name of CSV")
  flag.Parse()
  file_name := *file_flag
  if _, err := os.Stat(file_name); err != nil {
    fmt.Println("File does not exist in folder.")
    return
  }
  content, err := ioutil.ReadFile(file_name)
  if err != nil {
    fmt.Println("Error while trying to open the file.")
    return
  }
  reader := csv.NewReader(strings.NewReader(string(content)))
  input_reader := bufio.NewReader(os.Stdin)
  answers := [][]string {}
  question_no := 1

  for {
    record, err := reader.Read()
    if err == io.EOF {
      break
    }
    if err != nil {
      log.Fatal(err)
    }

    fmt.Printf("%d. %s ?\n", question_no, record[0])
    answer, err := input_reader.ReadString('\n')
    answers = append(answers, []string { record[0], record[1], answer })
    question_no ++
  }

  score := 0
  for _, v := range answers {
    if strings.ToLower(strings.TrimSpace(v[2])) == strings.ToLower(strings.TrimSpace(v[1])) {
      score++
    }
  }

  fmt.Printf("\n %d/%d \n", score, len(answers))
}