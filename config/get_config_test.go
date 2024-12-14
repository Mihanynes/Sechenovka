package config

import (
	"fmt"
	"testing"
)

func Test_ParseYAML(t *testing.T) {
	questions, err := parseYAML("./testdata/questions.yaml")
	if err != nil {
		t.Fatalf("Ошибка при парсинге YAML: %v", err)
	}

	//expectedQuestions := []*model.Question{
	//	{
	//		QuestionText: "Есть ли у вас боль или дискомфорт в грудной клетке?",
	//		QuestionID:   1,
	//		Options: []*model.Option{
	//			{
	//				AnswerText:     "да",
	//				AnswerID:       1,
	//				QuestionID:     1,
	//				Points:         0,
	//				NextQuestionID: 2,
	//			},
	//			{
	//				AnswerText:     "нет",
	//				AnswerID:       2,
	//				QuestionID:     1,
	//				Points:         0,
	//				NextQuestionID: 2,
	//			},
	//		},
	//	},
	//	{
	//		QuestionID:   2,
	//		QuestionText: "Когда возникают боли?",
	//		Options: []*model.Option{
	//			{
	//				AnswerText:     "При покое",
	//				Points:         0,
	//				NextQuestionID: 3,
	//			},
	//			{
	//				AnswerText:     "При ходьбе",
	//				Points:         1,
	//				NextQuestionID: 4,
	//			},
	//			{
	//				AnswerText: "При поворотах тела",
	//				Points:     2,
	//				IsEnded:    true,
	//			},
	//		},
	//	},
	//}

	fmt.Println(*questions[0])
	//require.ElementsMatch(t, expectedQuestions, questions)
}
