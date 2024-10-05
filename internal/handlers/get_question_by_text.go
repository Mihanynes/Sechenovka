package handlers

import (
	"Sechenovka/internal/models/quiz"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

const answersDBPath = "storage/user_answers/"

type QuizController struct {
	Quiz *quiz.Quiz
}

func NewQuizController(q *quiz.Quiz) *QuizController {
	return &QuizController{Quiz: q}
}

func (qc *QuizController) GetQuestion() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		op := "handlers/GetQuestion"
		var answer quiz.Answer

		body := ctx.Body()
		err := json.Unmarshal(body, &answer)
		if err != nil {
			log.Println(op + err.Error())
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Invalid request"})
		}

		userInfo := ctx.Get("Authorization")
		parts := strings.Split(userInfo, ":")
		userInfo = parts[0]
		userAnswersDBPath := answersDBPath + userInfo + ".yaml"

		nextQuestionId := answer.NextQuestionId
		currentQuestionId := answer.QuestionId
		currentQuestionText := quiz.GetQuestionById(qc.Quiz, currentQuestionId).Question.QuestionText
		if currentQuestionId == -1 {
			err = appendAnswerToYAMLFile(currentQuestionText, answer, userAnswersDBPath)
			if err != nil {
				log.Println(op + err.Error())
				return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Internal server error"})
			}
			return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": "Test is over"})
		}

		nextQuestion := quiz.GetQuestionById(qc.Quiz, nextQuestionId)
		if nextQuestion == nil {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "Question not found"})
		}

		if currentQuestionId == 0 {
			initTest := quiz.Test{
				Timestamp: time.Now().Format(time.RFC3339),
				Questions: []quiz.QA{},
			}

			err = initYAMLFile(initTest, userAnswersDBPath)
			if err != nil {
				log.Println(op + err.Error())
				return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Internal server error"})
			}
			return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": nextQuestion})
		}

		err = appendAnswerToYAMLFile(currentQuestionText, answer, userAnswersDBPath)
		if err != nil {
			log.Println(op + err.Error())
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Internal server error"})
		}

		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": nextQuestion})
	}
}

func appendAnswerToYAMLFile(questionText string, answer quiz.Answer, fileName string) error {
	return appendToYAMLFile(fileName, func(data *quiz.Data) {
		lastTestIndex := len(data.History) - 1
		if lastTestIndex >= 0 {
			data.History[lastTestIndex].Questions = append(data.History[lastTestIndex].Questions, quiz.QA{
				QuestionText: questionText,
				AnswerText:   answer.AnswerText,
				AnswerScore:  answer.AnswerScore,
			})
		}
	})
}

func initYAMLFile(initTest quiz.Test, fileName string) error {
	yamlData, err := ioutil.ReadFile(fileName)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("ошибка чтения файла: %w", err)
	}

	var data quiz.Data
	if yamlData != nil {
		err := yaml.Unmarshal(yamlData, &data)
		if err != nil {
			return fmt.Errorf("ошибка демаршализации YAML данных: %w", err)
		}
	}

	data.History = append(data.History, initTest)

	yamlData, err = yaml.Marshal(&data)
	if err != nil {
		return fmt.Errorf("ошибка маршализации данных: %w", err)
	}

	err = ioutil.WriteFile(fileName, yamlData, 0644)
	if err != nil {
		return fmt.Errorf("ошибка записи файла: %w", err)
	}

	return nil
}

func appendToYAMLFile(fileName string, modify func(*quiz.Data)) error {
	data, err := readYAMLFile(fileName)
	if err != nil {
		return err
	}

	modify(&data)

	yamlData, err := yaml.Marshal(&data)
	if err != nil {
		return fmt.Errorf("ошибка маршализации данных: %w", err)
	}
	return ioutil.WriteFile(fileName, yamlData, 0644)
}

func readYAMLFile(fileName string) (quiz.Data, error) {
	var data quiz.Data
	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return data, nil // если файл не существует, возвращаем пустую структуру
		}
		return data, fmt.Errorf("ошибка чтения YAML файла: %w", err)
	}
	err = yaml.Unmarshal(yamlFile, &data)
	if err != nil {
		return data, fmt.Errorf("ошибка демаршализации YAML данных: %w", err)
	}
	return data, nil
}
