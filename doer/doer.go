package doer

//go:generate mockgen -destination=../mocks/mock_doer.go -package=mocks playground/doer Doer

// Doer do something awesome
type Doer interface {
	DoSomething(int, string) error
}
