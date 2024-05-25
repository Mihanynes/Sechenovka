package models

type QuestionAnswer struct {
	Question Question `yaml:"question"`
	Answer   []Answer `yaml:"answers"`
}
type Question struct {
	QuestionId   int    `yaml:"question_id"`
	QuestionText string `yaml:"question_text"`
}

type Answer struct {
	AnswerId       int    `yaml:"answer_id"`
	QuestionId     int    `yaml:"question_id"`
	AnswerText     string `yaml:"answer_text"`
	AnswerScore    int    `yaml:"answer_score"`
	NextQuestionId int    `yaml:"next_question_id"`
}

type Quiz struct {
	Questions []QuestionAnswer `yaml:"quiz"`
}

func GetQuestionById(q *Quiz, questionId int) *QuestionAnswer {
	if questionId == 0 {
		return &q.Questions[0]
	}
	for _, question := range q.Questions {
		if question.Question.QuestionId == questionId {
			return &question
		}
	}
	return nil
}
