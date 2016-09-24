package err

import (
	"testing"
	"errors"
)

var (

	err1str string = "Error1 Message."
	err2str string = "Error2 Message."
	err3str string = "Error3 Message."

	err1 error = errors.New(err1str)
	err2 error = errors.New(err2str)
	err3 error = errors.New(err3str)
)
func TestAll(t *testing.T) {
	err := JoinError(err1, err2, err3)
	t.Log("JoinError : " + err.Error())

	err = JoinStringError("JoinStringError : ", err1str, err2str, err3str)
	t.Log(err.Error())

	err = JoinVarError("JoinVarError : ", 1024, 3.141592653, err1, err2, err3)
	t.Log(err.Error())
	t.Log("Last Error : " + GetLastError(err))
}

func TestSetErrorSeparator(t *testing.T) {
	SetErrorSeparator(" | ")
	TestAll(t)
}

func BenchmarkJoinError(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JoinError(err1, err2, err3)
	}
}

func BenchmarkJoinStringError(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JoinStringError(err1str, err2str, err3str)
	}
}

func BenchmarkJoinVarError(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JoinVarError(err1str, 1024, err3)
	}
}

func BenchmarkGetLastError(b *testing.B) {
	b.StopTimer()
	err := JoinError(err1, err2, err3)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		GetLastError(err)
	}
}