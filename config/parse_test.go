package config

import (
	"testing"
)

func Test_ParseYAML(t *testing.T) {
	questions, err := ParseYAML("./testdata/questions.yaml")
	if err != nil {
		t.Fatalf("Ошибка при парсинге YAML: %v", err)
	}

	// Проверка количества вопросов
	if len(questions) != 2 {
		t.Fatalf("Ожидали 2 вопроса, получили %d", len(questions))
	}

	// Проверка текстов вопросов
	expectedQuestions := []string{
		"Есть ли у вас боль или дискомфорт в грудной клетке?",
		"Когда возникают боли?",
	}
	for i, q := range questions {
		if q.Text != expectedQuestions[i] {
			t.Errorf("Ожидали вопрос '%s', получили '%s'", expectedQuestions[i], q.Text)
		}
	}
}
