package entity

// Question is the question that will be asked in the game
type Question struct {
	Id              int
	Text            string
	PossibleAnswers []PossibleAnswer
	CorrectAnswer   uint // Correct possible answer choice ID
	Difficulty      QuestionDifficulty
	CategoryID      uint
}

// PossibleAnswerChoice will have 4 instances for each question and 1 of those will be the correct answer
type PossibleAnswerChoice uint8

const (
	PossibleAnswerA PossibleAnswerChoice = iota + 1
	PossibleAnswerB
	PossibleAnswerC
	PossibleAnswerD
)

// IsValid checks if the possible answer choice is valid
func (p PossibleAnswerChoice) IsValid() bool {
	return p <= PossibleAnswerD
}

// each Question instance will have one instance of QuestionDifficulty
type QuestionDifficulty uint8

const (
	QuestionDifficultyEasy QuestionDifficulty = iota + 1
	QuestionDifficultyMedium
	QuestionDifficultyHard
)

func (q QuestionDifficulty) IsValid() bool {
	return q <= QuestionDifficultyHard
}

type PossibleAnswer struct {
	ID     uint
	Text   string
	Choice uint8
}
