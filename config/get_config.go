package config

import (
	"Sechenovka/internal/models"
	"gopkg.in/yaml.v3"
	"os"
)

func GetQuestionsConfig() ([]*models.Question, error) {
	config, err := parseYAML("config/questions_v2.yaml")
	if err != nil {
		return nil, err
	}
	return config, nil
}

func parseYAML(filename string) ([]*models.Question, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var questions []*models.Question
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&questions); err != nil {
		return nil, err
	}

	return questions, nil
}
