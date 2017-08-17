package main

import (
	"github.com/wsxiaoys/terminal"
	"bufio"
	"os"
	"github.com/howeyc/gopass"
	"strings"
)

func printInfo(msg string) {
	terminal.Stdout.
	Color("c").Print(msg).Reset().Nl()
}

func printQuestion(question string) {
	terminal.Stdout.
	Nl().
	Color("g").Print("? ").Color("y").Print(question).Nl().
	Reset().Print("> ")
}

func askProperty(question string) string {
	printQuestion(question)

	answer, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	return strings.TrimSpace(answer)
}

func askSecret(question string) string {
	printQuestion(question)

	secretBytes, err := gopass.GetPasswd()
	if err != nil {
		panic(err)
	}

	return string(secretBytes[:])
}
