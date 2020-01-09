package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/globalsign/mgo"
	"github.com/k0kubun/pp"
	"github.com/karrick/godirwalk"
)

func main() {

	optVerbose := flag.Bool("verbose", false, "Print file system entries.")
	flag.Parse()

	i := 1
	err := godirwalk.Walk("../dataset", &godirwalk.Options{
		Callback: func(osPathname string, de *godirwalk.Dirent) error {
			if *optVerbose {
				fmt.Printf("%s %s\n", de.ModeType(), osPathname)
			}
			if de.ModeType() != os.ModeDir {
				convertQuizz(osPathname, i)
			}
			i++
			return nil
		},
		ErrorCallback: func(osPathname string, err error) godirwalk.ErrorAction {
			if *optVerbose {
				fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
			}

			// For the purposes of this example, a simple SkipNode will suffice,
			// although in reality perhaps additional logic might be called for.
			return godirwalk.SkipNode
		},
		Unsorted: true, // set true for faster yet non-deterministic enumeration (see godoc)
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func convertQuizz(filePath string, id int) {
	// Open our jsonFile
	jsonFile, err := os.Open(filePath)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened file: ", filePath)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var data Data
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		log.Fatalf("cannot unmarshal data: %v\n", err)
	}
	//pp.Println(data.Kahoot)

	database := "AskDB"
	collection := "Questions"
	session, err := mgo.Dial("mongodb://127.0.0.1:27017/" + database)
	if err != nil {
		log.Fatalf("cannot connect to mongodb host: %v\n", err)
	}

	/*
		c := session.DB(database).C(collection)
		query := ""
		err := c.Find(query).One(&result)
		if err != nil {
			log.Fatalf("cannot connect to mongodb host: %v", err)
		}
	*/

	quiz := &Quiz{
		ID:    id,
		Name:  data.Card.Title,
		Image: data.Card.Cover,
	}

	for _, q := range data.Kahoot.Questions {
		question := &QuizQuestion{
			Question: q.Question,
		}

		var correctAnswer string
		for i, c := range q.Choices {
			question.Answers = append(question.Answers, c.Answer)
			if c.Correct {
				correctAnswer = strconv.Itoa(i)
			}
		}
		question.Correct = correctAnswer
		question.Image = q.Image
		question.Time = q.Time
		question.PointsMultiplier = q.PointsMultiplier

		quiz.Questions = append(quiz.Questions, question)
	}

	coll := session.DB(database).C(collection)
	err = coll.Insert(quiz)
	if err != nil {
		log.Fatalln(err)
	}

	pp.Println(quiz)

	// Use session as normal
	session.Close()

}
