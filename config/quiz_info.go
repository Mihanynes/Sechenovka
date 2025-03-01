package config

import "time"

type Quiz struct {
	QuizId          int
	Name            string
	Description     string
	TimeToPassAgain time.Duration
}

var QuizInfo = map[int]Quiz{
	SelfCheckQuiz: {
		QuizId:          SelfCheckQuiz,
		Name:            "Опрос состояния здоровья",
		Description:     "Зададим вопросы о вашем текущем состоянии здоровья",
		TimeToPassAgain: 5 * time.Second,
	},
	RecommendationsQuiz: {
		QuizId:          RecommendationsQuiz,
		Name:            "Уведомления о приеме препаратов",
		Description:     "Спросим какие препараты вы принимаете",
		TimeToPassAgain: 5 * time.Second,
	},
	TakingMedicationsQuiz: {
		QuizId:          TakingMedicationsQuiz,
		Name:            "Уведомление о рекомендациях",
		Description:     "Спросим о ваших привычках и дадим рекомендации",
		TimeToPassAgain: 5 * time.Second,
	},
}
