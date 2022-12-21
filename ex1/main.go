package main 

import (
	"github.com/gocarina/gocsv"
	"fmt"
	"os"
	"flag"
	"time"
	"math/rand"
)
type Problem struct {
	Task string `csv:"task"`
	Solution string `csv:"solution"`
}

func main() {
	filename := flag.String("file", "problems.csv", "the filename, str")
	timeLimit := flag.Int("limit", 30, "Time limit to answer in seconds")
	shuffle := flag.Bool("shuffle", false, "Shuffle input array") 
	flag.Parse()

	in, err:= os.Open(*filename)
	if err != nil {
		panic(err)
	}
	defer in.Close()

	problems := []*Problem{}
	if err := gocsv.UnmarshalFile(in, &problems); err != nil {
		panic(err)
	}
	if *shuffle==true {
		rand.Shuffle(len(problems), func(i, j int) {problems[i], problems[j] = problems[j], problems[i]})
	}

	points := 0

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	for _, problem := range problems {
		fmt.Print(problem.Task, " = ")
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer

		}()
		select {
		case <- timer.C:
			fmt.Println()
			fmt.Println("You took too long!")
			fmt.Printf("You reached %d points \n", points)
			return
		case answer := <- answerCh:
			if answer== problem.Solution {
				points++
				fmt.Println("Correct.")
			} else {
				fmt.Println("You're wrong. Game over.")
				fmt.Printf("You reached %d points \n", points)
				return
			}
		}

	}
	fmt.Printf("You answered all questions correctly and reached %d points \n", points)

}