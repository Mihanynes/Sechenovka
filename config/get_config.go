package config

import (
	"Sechenovka/internal/model"
	"gopkg.in/yaml.v3"
	"os"
)

const testPath = "config/questions.yaml"

func GetQuestionsConfig() ([]*model.Question, error) {
	config, err := parseYAML(testPath)
	if err != nil {
		return nil, err
	}
	return config, nil
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
