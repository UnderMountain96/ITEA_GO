package main

import (
	"errors"
	"fmt"
)

const (
	minScore = 0
	maxScore = 5
)

type Student struct {
	Name    string
	Surname string
	Lessons []Lesson
}

func NewStudent(name string, surname string) (Student, error) {
	if name == "" {
		return Student{}, errors.New("name must not be empty")
	}

	if surname == "" {
		return Student{}, errors.New("surname must not be empty")
	}

	student := Student{
		Name:    name,
		Surname: surname,
	}

	return student, nil
}

func (s *Student) FullName() string {
	return s.Name + " " + s.Surname
}

func (s *Student) AverageScore() float64 {
	var sumScore float64
	for _, l := range s.Lessons {
		sumScore += l.Score
	}

	if sumScore == 0 {
		return 0
	}

	return sumScore / float64(len(s.Lessons))
}

func (s *Student) AddLesson(l Lesson) {
	s.Lessons = append(s.Lessons, l)
}

type Lesson struct {
	Title string
	Score float64
}

func NewLesson(title string, score float64) (Lesson, error) {
	if title == "" {
		return Lesson{}, errors.New("title must not be empty")
	}

	if score < 0 {
		return Lesson{}, fmt.Errorf("body must be grater then %d", minScore)
	}

	if score > 5 {
		return Lesson{}, fmt.Errorf("body must be lower then %d", maxScore)
	}

	lesson := Lesson{
		Title: title,
		Score: score,
	}
	return lesson, nil
}

func main() {
	lessons := make([]Lesson, 0)

	std, err := NewStudent("Ivan", "Sirko")
	if err != nil {
		fmt.Println(err)
		return
	}

	lsn1, err := NewLesson("lesson 1", 4)
	if err != nil {
		fmt.Println(err)
		return
	}
	lsn2, err := NewLesson("lesson 2", 5)
	if err != nil {
		fmt.Println(err)
		return
	}
	lsn3, err := NewLesson("lesson 3", 3)
	if err != nil {
		fmt.Println(err)
		return
	}
	lsn4, err := NewLesson("lesson 4", 4)
	if err != nil {
		fmt.Println(err)
		return
	}

	lessons = append(lessons, lsn1, lsn2, lsn3, lsn4)

	for _, l := range lessons {
		std.AddLesson(l)
	}

	fmt.Printf("FullName:\t%s\n", std.FullName())
	fmt.Printf("Average score:\t%.1f\n\n", std.AverageScore())
}
