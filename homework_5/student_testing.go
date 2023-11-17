package main

import "fmt"

type StudentTestProvider interface {
	GetStudentTest() Test
}

type StudentTest interface {
	StudentTestProvider
}

type Test struct {
	Questions      []Question
	CorrectAnswers []int
}

func (t *Test) GetStudentTest() {
	t.Questions = []Question{
		NewQuestion(
			"When did the term \"Cossack\" first come into use?",
			map[int]string{
				1: "15th century",
				2: "16th century",
				3: "17th century",
				4: "18th century",
			},
			2,
		),
		NewQuestion(
			"Which historical figure is considered the founder of the Zaporizhian Sich, a major Cossack host?",
			map[int]string{
				1: "Bohdan Khmelnytsky",
				2: "Ivan Mazepa",
				3: "Petro Doroshenko",
				4: "Dmytro Vyshnevetsky",
			},
			3,
		),
		NewQuestion(
			"What was the primary role of Ukrainian Cossacks in the 16th-18th centuries?",
			map[int]string{
				1: "Peasantry",
				2: "Religious leadership",
				3: "Military service and defense",
				4: "Trade and commerce",
			},
			3,
		),
		NewQuestion(
			"Which treaty solidified the Cossack Hetmanate as an autonomous state within the Polish-Lithuanian Commonwealth?",
			map[int]string{
				1: "Treaty of Hadiach",
				2: "Treaty of Pereyaslav",
				3: "Treaty of Kucuk Kaynarca",
				4: "Treaty of Andrusovo",
			},
			1,
		),
		NewQuestion(
			"Which Cossack leader famously led the uprising against Polish rule in the mid-17th century?",
			map[int]string{
				1: "Ivan Sirko",
				2: "Ivan Mazepa",
				3: "Bohdan Khmelnytsky",
				4: "Pylyp Orlyk",
			},
			3,
		),
		NewQuestion(
			"The Cossack capital, Sich, was situated on which river?",
			map[int]string{
				1: "Dnieper River",
				2: "Don River",
				3: "Volga River",
				4: "Danube River",
			},
			1,
		),
		NewQuestion(
			"What was the \"Koliivshchyna\" uprising related to in Ukrainian Cossack history?",
			map[int]string{
				1: "Fight against Ottoman Empire",
				2: "Internal conflict among Cossack factions",
				3: "Religious tensions and persecution",
				4: "War against the Russian Empire",
			},
			3,
		),
		NewQuestion(
			"Which monarch abolished the Zaporizhian Sich in 1775?",
			map[int]string{
				1: "Catherine the Great of Russia",
				2: "Peter the Great of Russia",
				3: "Frederick the Great of Prussia",
				4: "Maria Theresa of Austria",
			},
			1,
		),
		NewQuestion(
			"Which Battle marked the significant defeat of the Cossack Hetmanate by the Russian Empire in the early 18th century?",
			map[int]string{
				1: "Battle of Poltava",
				2: "Battle of Berestechko",
				3: "Battle of Zhovti Vody",
				4: "Battle of Konotop",
			},
			1,
		),
		NewQuestion(
			"What was the significance of the Cossack Code of Law known as \"The Articles\"?",
			map[int]string{
				1: "It outlined the social hierarchy within the Cossack society.",
				2: "It dictated religious practices and ceremonies.",
				3: "It established military strategies and tactics.",
				4: "It governed the conduct and behavior of the Cossacks.",
			},
			4,
		),
	}
	t.CorrectAnswers = make([]int, 0)
}

func (t *Test) AddCorrectAnswer(idx int) {
	t.CorrectAnswers = append(t.CorrectAnswers, idx)
}

type Question struct {
	Text          string
	AnswerOptions map[int]string
	CorrectAnswer int
}

func NewQuestion(text string, answerOptions map[int]string, correctAnswer int) Question {
	return Question{
		Text:          text,
		AnswerOptions: answerOptions,
		CorrectAnswer: correctAnswer,
	}
}

func main() {
	test := Test{}
	test.GetStudentTest()

	for n, question := range test.Questions {
		questionNumner := n + 1

		fmt.Printf("Question %d:\t%s\n\n", questionNumner, question.Text)
		for idx, answer := range question.AnswerOptions {
			fmt.Printf("%d) %s\n", idx, answer)
		}
		fmt.Println()

		fmt.Print("Entry your answer: ")
		var stdAnswer int
		fmt.Scan(&stdAnswer)

		if stdAnswer == question.CorrectAnswer {
			test.AddCorrectAnswer(questionNumner)
		}

		fmt.Println()
	}

	NumberCorrectAnswers := len(test.CorrectAnswers)
	NumberWrongAnswers := len(test.Questions) - NumberCorrectAnswers

	fmt.Printf("Number of correct answers: \t%d\n", NumberCorrectAnswers)
	fmt.Printf("Number of wrong answers: \t%d\n", NumberWrongAnswers)

}
