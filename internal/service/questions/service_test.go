package questions

import (
	"Sechenovka/internal/model"
	"testing"
)

func Test_GetOptionsByQuestionText(t *testing.T) {
	testCases := []struct {
		name            string
		questionText    string
		expectedOptions []model.Option
		expectedError   bool
	}{
		{
			name:         "Существующий вопрос: Когда возникают боли?",
			questionText: "Когда возникают боли?",
			expectedOptions: []model.Option{
				{Answer: "При покое", Points: 2, NextQuestionText: "Как долго длится боль?"},
				{Answer: "При ходьбе", Points: 1, NextQuestionText: "Как долго длится боль?"},
				{Answer: "При поворотах тела", Points: 0, NextQuestionText: "Как долго длится боль?"},
			},
			expectedError: false,
		},
		{
			name:         "Существующий вопрос: Есть ли у вас боль или дискомфорт в грудной клетке?",
			questionText: "Есть ли у вас боль или дискомфорт в грудной клетке?",
			expectedOptions: []model.Option{
				{Answer: "да", Points: 0, NextQuestionText: "Когда возникают боли?"},
				{Answer: "нет", Points: 0},
			},
			expectedError: false,
		},
		{
			name:            "Несуществующий вопрос",
			questionText:    "Неизвестный вопрос",
			expectedOptions: nil,
			expectedError:   true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := New(fakeQuestions)
			options, err := service.GetOptionsByQuestionText(tc.questionText)

			if (err != nil) != tc.expectedError {
				t.Errorf("Ожидали ошибку: %v, получили: %v", tc.expectedError, err)
				return
			}

			// Если ожидалась ошибка, пропускаем проверку опций
			if tc.expectedError {
				return
			}

			// Проверяем количество и содержание опций
			if len(options) != len(tc.expectedOptions) {
				t.Errorf("Ожидали %d опций, получили %d", len(tc.expectedOptions), len(options))
				return
			}

			for i, opt := range options {
				if opt != tc.expectedOptions[i] {
					t.Errorf("Ожидали опцию '%v', получили '%v'", tc.expectedOptions[i], opt)
				}
			}
		})
	}
}

var fakeQuestions = []*model.Question{
	{
		Text: "Есть ли у вас боль или дискомфорт в грудной клетке?",
		Options: []model.Option{
			{Answer: "да", Points: 0, NextQuestionText: "Когда возникают боли?"},
			{Answer: "нет", Points: 0},
		},
	},
	{
		Text: "Когда возникают боли?",
		Options: []model.Option{
			{Answer: "При покое", Points: 2, NextQuestionText: "Как долго длится боль?"},
			{Answer: "При ходьбе", Points: 1, NextQuestionText: "Как долго длится боль?"},
			{Answer: "При поворотах тела", Points: 0, NextQuestionText: "Как долго длится боль?"},
		},
	},
}
