package config

import (
	"Sechenovka/internal/model"
	"gopkg.in/yaml.v3"
	"os"
)

const (
	SelfCheckQuiz         = 1
	RecommendationsQuiz   = 2
	TakingMedicationsQuiz = 3
)

var PathQuizIdMap = map[string]int{
	"config/quiz_questions/self_check.yaml":         SelfCheckQuiz,
	"config/quiz_questions/recommendations.yaml":    RecommendationsQuiz,
	"config/quiz_questions/taking_medications.yaml": TakingMedicationsQuiz,
}

func GetQuestionsConfig() (map[int][]*model.Question, error) {
	generalConfig := make(map[int][]*model.Question)
	for path, quizId := range PathQuizIdMap {
		config, err := parseYAML(path)
		if err != nil {
			return nil, err
		}
		generalConfig[quizId] = config
	}
	return generalConfig, nil
}

func parseYAML(filename string) ([]*model.Question, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var questions []*model.Question
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&questions); err != nil {
		return nil, err
	}

	return questions, nil
}
