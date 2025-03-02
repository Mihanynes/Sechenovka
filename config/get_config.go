package config

import (
	"Sechenovka/internal/model"
	"gopkg.in/yaml.v3"
	"os"
)

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
