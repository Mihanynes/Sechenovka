package config

import (
	"Sechenovka/internal/model"
	"Sechenovka/internal/utils/pointer"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_ParseYAML(t *testing.T) {
	questions, err := parseYAML("./testdata/questions.yaml")
	if err != nil {
		t.Fatalf("Ошибка при парсинге YAML: %v", err)
	}

	expectedQuestions := []*model.Question{
		{
			Text:        "Есть ли у вас боль или дискомфорт в грудной клетке?",
			ScoreToFail: nil, // Поле не указано в YAML, поэтому здесь nil
			Options: []model.Option{
				{
					Answer:           "да",
					Points:           0,
					NextQuestionText: "Когда возникают боли?",
				},
				{
					Answer:           "нет",
					Points:           0,
					NextQuestionText: "",
				},
			},
		},
		{
			Text:        "Когда возникают боли?",
			ScoreToFail: pointer.Get(2),
			Options: []model.Option{
				{
					Answer:           "При покое",
					Points:           2,
					NextQuestionText: "Как долго длится боль?",
				},
				{
					Answer:           "При ходьбе",
					Points:           1,
					NextQuestionText: "Как долго длится боль?",
				},
				{
					Answer:           "При поворотах тела",
					Points:           0,
					NextQuestionText: "Как долго длится боль?",
				},
			},
		},
	}

	require.ElementsMatch(t, expectedQuestions, questions)
}
