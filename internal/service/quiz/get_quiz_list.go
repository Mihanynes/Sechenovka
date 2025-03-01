package quiz

import (
	"Sechenovka/config"
	"Sechenovka/internal/handlers/quiz"
	"Sechenovka/internal/model"
	"time"
)

func (s *Service) GetQuizList(userId model.UserId) (quiz.QuizList, error) {
	var quizList quiz.QuizList
	for _, quizInfo := range config.QuizInfo {
		lastUserQuiz, err := s.userResponseStorage.GetLastUserResponse(userId, quizInfo.QuizId)
		if err != nil {
			return quizList, err
		}
		var isAvailable bool
		var nextTimeCan time.Time

		if lastUserQuiz == nil {
			isAvailable = true
			nextTimeCan = time.Now()
		} else {
			isAvailable = time.Now().After(lastUserQuiz.UpdatedAt.Add(quizInfo.TimeToPassAgain))
			nextTimeCan = lastUserQuiz.UpdatedAt.Add(quizInfo.TimeToPassAgain)
		}

		quizList.List = append(quizList.List, quiz.Quiz{
			QuizId:      quizInfo.QuizId,
			Name:        quizInfo.Name,
			Description: quizInfo.Description,
			IsAvailable: isAvailable,
			NextTimeCan: nextTimeCan,
		})
	}
	return quizList, nil
}
