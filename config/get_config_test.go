package config

import (
	"fmt"
	"testing"
)

func Test_ParseYAML(t *testing.T) {
	questions, err := parseYAML("config/testdata/1_test.yaml")
	if err != nil {
		t.Fatalf("Ошибка при парсинге YAML: %v", err)
	}

	//expectedQuestions := []*model.Question{
	//	{
	//		QuestionText: "Есть ли у вас боль или дискомфорт в грудной клетке?",
	//		QuestionId:   1,
	//		Options: []*model.Option{
	//			{
	//				AnswerText:     "да",
	//				AnswerId:       1,
	//				QuestionId:     1,
	//				Points:         0,
	//				NextQuestionId: 2,
	//			},
	//			{
	//				AnswerText:     "нет",
	//				AnswerId:       2,
	//				QuestionId:     1,
	//				Points:         0,
	//				NextQuestionId: 2,
	//			},
	//		},
	//	},
	//	{
	//		QuestionId:   2,
	//		QuestionText: "Когда возникают боли?",
	//		Options: []*model.Option{
	//			{
	//				AnswerText:     "При покое",
	//				Points:         0,
	//				NextQuestionId: 3,
	//			},
	//			{
	//				AnswerText:     "При ходьбе",
	//				Points:         1,
	//				NextQuestionId: 4,
	//			},
	//			{
	//				AnswerText: "При поворотах тела",
	//				Points:     2,
	//				IsEnded:    true,
	//			},
	//		},
	//	},
	//}

	fmt.Println(*questions[0].ScoreToFail)
	//require.ElementsMatch(t, expectedQuestions, questions)
}
