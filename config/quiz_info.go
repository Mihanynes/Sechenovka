package config

import "time"

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

type Quiz struct {
	QuizId          int           `json:"quiz_id"`
	Name            string        `json:"name"`
	Description     string        `json:"description"`
	TimeToPassAgain time.Duration `json:"time_to_pass_again"`
}

var QuizInfo = map[int]Quiz{
	SelfCheckQuiz: {
		QuizId:          SelfCheckQuiz,
		Name:            "Опрос состояния здоровья",
		Description:     "Зададим вопросы о вашем текущем состоянии здоровья",
		TimeToPassAgain: 5 * time.Second,
	},
	TakingMedicationsQuiz: {
		QuizId:          TakingMedicationsQuiz,
		Name:            "Уведомления о приеме препаратов",
		Description:     "Спросим какие препараты вы принимаете",
		TimeToPassAgain: 5 * time.Second,
	},
	RecommendationsQuiz: {
		QuizId:          RecommendationsQuiz,
		Name:            "Уведомление о рекомендациях",
		Description:     "Спросим о ваших привычках и дадим рекомендации",
		TimeToPassAgain: 5 * time.Second,
	},
}
