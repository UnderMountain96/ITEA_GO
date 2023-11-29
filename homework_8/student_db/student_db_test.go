package main

import (
	"fmt"
	"testing"
)

func TestNewStudent(t *testing.T) {
	expectedName := "Ivan"
	expectedSurname := "Bohun"

	student, err := NewStudent(expectedName, expectedSurname)
	if err != nil {
		t.Errorf("cannot create new student: %s", err)
	}

	if student.Name != expectedName {
		t.Errorf("invalid Name: got: %s, want: %s", student.Name, expectedName)
	}

	if student.Surname != expectedSurname {
		t.Errorf("invalid Surname: got: %s, want: %s", student.Surname, expectedSurname)
	}

	if len(student.Lessons) != 0 {
		t.Error("Lessons must be empty")
	}

	testCases := map[string]struct {
		nameValue    string
		surnameValue string
		errorReason  string
	}{
		"empty name": {
			nameValue:    "",
			surnameValue: "Bohun",
			errorReason:  "cannot set empty name",
		},
		"empty surname": {
			nameValue:    "Ivan",
			surnameValue: "",
			errorReason:  "cannot set empty surname",
		},
	}

	for tn, tc := range testCases {
		t.Run(tn, func(t *testing.T) {
			_, err := NewStudent(tc.nameValue, tc.surnameValue)
			if err == nil {
				t.Error(tc.errorReason)
			}
		})
	}
}

func TestFullName(t *testing.T) {
	expectedName := "Ivan"
	expectedSurname := "Bohun"
	expectedFullName := expectedName + " " + expectedSurname

	student, err := NewStudent(expectedName, expectedSurname)
	if err != nil {
		t.Errorf("cannot create new student: %s", err)
	}

	if student.FullName() != expectedFullName {
		t.Errorf("invalid FullName: got: %s, want: %s", student.FullName(), expectedFullName)
	}
}

func TestAddLesson(t *testing.T) {
	expectedName := "Ivan"
	expectedSurname := "Bohun"

	student, err := NewStudent(expectedName, expectedSurname)
	if err != nil {
		t.Errorf("cannot create new student: %s", err)
	}

	if len(student.Lessons) != 0 {
		t.Error("Lessons must be empty")
	}

	count := 5
	for i := 0; i < count; i++ {
		student.AddLesson(Lesson{
			fmt.Sprintf("lesson_%d", i),
			float64(i),
		})
	}

	if len(student.Lessons) != count {
		t.Errorf("invalid len Lessons: got: %d, want: %d", len(student.Lessons), count)
	}
}

func TestAverageScore(t *testing.T) {
	expectedName := "Ivan"
	expectedSurname := "Bohun"

	student, err := NewStudent(expectedName, expectedSurname)
	if err != nil {
		t.Errorf("cannot create new student: %s", err)
	}

	count := 5
	score := 5.0
	for i := 0; i < count; i++ {
		student.AddLesson(Lesson{
			fmt.Sprintf("lesson_%d", i),
			score,
		})
	}

	averageScore := (score * float64(count)) / float64(len(student.Lessons))

	if student.AverageScore() != averageScore {
		t.Errorf("invalid average score: got: %.1f, want: %.1f", student.AverageScore(), averageScore)
	}
}

func TestNewLesson(t *testing.T) {
	expectedTitle := "Lesson"
	expectedScore := 5.0

	lesson, err := NewLesson(expectedTitle, expectedScore)
	if err != nil {
		t.Errorf("cannot create new lesson: %s", err)
	}

	if lesson.Title != expectedTitle {
		t.Errorf("invalid Title: got: %s, want: %s", lesson.Title, expectedTitle)
	}

	if lesson.Score != expectedScore {
		t.Errorf("invalid Score: got: %.1f, want: %.1f", lesson.Score, expectedScore)
	}

	testCases := map[string]struct {
		titleValue  string
		scoreValue  float64
		errorReason string
	}{
		"empty title": {
			titleValue:  "",
			scoreValue:  5,
			errorReason: "cannot set empty title",
		},
		"score too low": {
			titleValue:  "Lesson",
			scoreValue:  -1,
			errorReason: fmt.Sprintf("cannot set score lower then %d", minScore),
		},
		"score too high": {
			titleValue:  "Lesson",
			scoreValue:  10,
			errorReason: fmt.Sprintf("cannot set score higher then %d", maxScore),
		},
	}

	for tn, tc := range testCases {
		t.Run(tn, func(t *testing.T) {
			_, err := NewLesson(tc.titleValue, tc.scoreValue)
			if err == nil {
				t.Error(tc.errorReason)
			}
		})
	}
}
