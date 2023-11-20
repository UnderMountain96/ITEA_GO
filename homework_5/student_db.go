package main

import "fmt"

type Student struct {
	Name    string
	Surname string
	Lessons []Lesson
}

func NewStudent(name string, surname string) Student {
	return Student{
		Name:    name,
		Surname: surname,
	}
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

func NewLesson(title string, score float64) Lesson {
	return Lesson{
		Title: title,
		Score: score,
	}
}

func main() {
	students := make([]Student, 0)

	std1 := NewStudent("Ivan", "Sirko")
	std1.AddLesson(NewLesson("lesson 1", 4))
	std1.AddLesson(NewLesson("lesson 2", 5))
	std1.AddLesson(NewLesson("lesson 3", 3))
	std1.AddLesson(NewLesson("lesson 4", 4))

	std2 := NewStudent("Bohdan", "Khmelnytsky")
	std2.AddLesson(NewLesson("lesson 1", 5))
	std2.AddLesson(NewLesson("lesson 2", 5))
	std2.AddLesson(NewLesson("lesson 3", 5))
	std2.AddLesson(NewLesson("lesson 4", 3))

	std3 := NewStudent("Ivan", "Bohun")
	std3.AddLesson(NewLesson("lesson 1", 5))
	std3.AddLesson(NewLesson("lesson 2", 4))
	std3.AddLesson(NewLesson("lesson 3", 0))
	std3.AddLesson(NewLesson("lesson 4", 4))

	std4 := NewStudent("Ivan", "Mazepa")
	std4.AddLesson(NewLesson("lesson 1", 5))
	std4.AddLesson(NewLesson("lesson 2", 4))
	std4.AddLesson(NewLesson("lesson 3", 5))
	std4.AddLesson(NewLesson("lesson 4", 5))

	students = append(students, std2, std1, std3, std4)

	fmt.Println("Students:\n")

	for _, s := range students {
		fmt.Printf("FullName:\t%s\n", s.FullName())
		fmt.Printf("Average score:\t%.1f\n\n", s.AverageScore())
	}
}
