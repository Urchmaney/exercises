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
  "time"
)

func  main()  {
	file_flag := flag.String("file_name", "problems.csv", "Full File Name of CSV")
  time_flag := flag.Int("time", 30, "Quiz time span")
  flag.Parse()
  file_name := *file_flag
  time := *time_flag
  if _, err := os.Stat(file_name); err != nil {
    fmt.Println("File does not exist in folder.")
    return
  }
  content, err := ioutil.ReadFile(file_name)
  if err != nil {
    fmt.Println("Error while trying to open the file.")
    return
  }
  questions := ProcessQuestionsFromBytes(content)
  answers := [][]string {}
  time_channel := make(chan bool)
  answer_channel := make(chan []string)
  
  go timer(time, time_channel)
  go GetUserAnswers(questions, answer_channel)

  done := false

  for !done {
    select {
    case answer := <- answer_channel:
      if answer == nil {
        done = true
      } else {
        answers = append(answers, answer)
      }
    case time_up := <- time_channel:
      done = time_up
    }
  }

  score := 0
  for _, v := range answers {
    if strings.ToLower(strings.TrimSpace(v[0])) == strings.ToLower(strings.TrimSpace(v[1])) {
      score++
    }
  }

  fmt.Printf("\n %d/%d \n", score, len(questions))
}

func GetUserAnswers(questions [][]string, c chan []string) {
  input_reader := bufio.NewReader(os.Stdin)
  question_no := 0
  for _, v := range questions {
    fmt.Printf("%d. %s ?\n", question_no, v[0])
    answer, _ := input_reader.ReadString('\n')
    question_no ++
    c <- []string { v[1], answer }
  }
  c <- nil
}

func ProcessQuestionsFromBytes (content []byte) [][]string {
  questions := [][]string {}
  reader := csv.NewReader(strings.NewReader(string(content)))
  for {
    record, err := reader.Read()
    if err == io.EOF {
      break
    }
    if err != nil {
      log.Fatal(err)
    }
    questions = append(questions, []string { record[0], record[1] })
  }
  return questions
}

func timer (seconds int, c chan bool) {
  time.Sleep(time.Duration(seconds) * time.Second)
  c <- true
}