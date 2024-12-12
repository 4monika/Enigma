package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"unicode"
)

const (
	AlpSize = 26
	nRolls  = 7
)

type Roll struct {
	Data       map[string]string
	ReversData map[string]string
}

func main() {
	alphabet := strings.Split("abcdefghijklmnopqrstuvwxyz", "") //определяем алфавит
	Keys := "ahreeda"                                           //Определяем ключи для роллеров
	Rolls := MakeRolls(alphabet)                                //создаём роллеры(необходимое кол-во можно задать в константах)
	Word := transformWord()                                     //преобразование слова к необходимому виду
	ChangedWord := Enigma(Word, Rolls, Keys)                    //Шифруем слово
	fmt.Println(ChangedWord)                                    //вывод в консоль
	fmt.Println(Enigma(ChangedWord, Rolls, Keys))               //заодно выводим в консоль результат проверки
}
func MakeRolls(alphabet []string) []Roll {
	NewRolls := [nRolls]Roll{}
	curMap := make(map[string]string, AlpSize)
	curReversMap := make(map[string]string, AlpSize)
	for j := 0; j < AlpSize; j++ {
		curMap[alphabet[j]] = alphabet[j]
		curReversMap[alphabet[j]] = alphabet[j]
	}
	NewRolls[0].Data = curMap
	NewRolls[0].ReversData = curReversMap
	for i := 1; i < nRolls-1; i++ {
		curKey := rand.Perm(AlpSize)
		curMap := make(map[string]string, AlpSize)
		curReversMap := make(map[string]string, AlpSize)
		for j := 0; j < AlpSize; j++ {
			curMap[alphabet[j]] = alphabet[curKey[j]]
			curReversMap[alphabet[curKey[j]]] = alphabet[j]
		}
		NewRolls[i].Data = curMap
		NewRolls[i].ReversData = curReversMap
	}
	curKey := rand.Perm(AlpSize / 2)
	curMap = make(map[string]string, AlpSize)
	curReversMap = make(map[string]string, AlpSize)
	for j := 0; j < AlpSize/2; j++ {
		curMap[alphabet[j]] = alphabet[curKey[j]+AlpSize/2]
		curReversMap[alphabet[curKey[j]+AlpSize/2]] = alphabet[j]
	}
	for j := 0; j < AlpSize/2; j++ {
		curMap[alphabet[j+AlpSize/2]] = curReversMap[alphabet[j+AlpSize/2]]
		curReversMap[alphabet[j]] = curMap[alphabet[j]]
	}
	NewRolls[nRolls-1].Data = curMap
	NewRolls[nRolls-1].ReversData = curReversMap
	return NewRolls[:]
}
func Enigma(Word []string, Rolls []Roll, Keys string) []string {
	answer := make([]string, len(Word))
	for i, letter := range Word {
		swap := Keys[0]
		if letter != "" {
			curRoll := Rolls[1]
			letter = curRoll.Data[string((letter[0]-97+Keys[1]+byte(i)-swap+26)%AlpSize+97)]
			swap = Keys[1]
			for j := 1; j < len(Rolls)-2; j++ {
				curRoll = Rolls[j+1]
				letter = curRoll.Data[string((letter[0]-97+Keys[j+1]-swap+26)%AlpSize+97)]
				swap = Keys[j+1]
			}
			curRoll = Rolls[nRolls-1]
			letter = curRoll.Data[string((letter[0]-97+Keys[nRolls-1]-byte(i)-swap+26)%AlpSize+97)]
			swap = Keys[nRolls-1]
			curRoll = Rolls[nRolls-2]
			letter = curRoll.ReversData[string((letter[0]-97+Keys[nRolls-2]+byte(i)-swap+26)%AlpSize+97)]
			swap = Keys[nRolls-2]
			for j := 1; j < len(Rolls)-1; j++ {
				curRoll := Rolls[len(Rolls)-j-2]
				letter = curRoll.ReversData[string((letter[0]-97+Keys[len(Keys)-j-2]-swap+26)%AlpSize+97)]
				swap = Keys[len(Keys)-j-2]
			}
			curRoll = Rolls[0]
			letter = curRoll.ReversData[string((letter[0]-97+Keys[0]-byte(i)-swap+26)%AlpSize+97)]
			swap = Keys[0]
			answer[i] = letter
		} else {
			answer[i] = ""
		}

	}
	return answer
}
func transformWord() []string {
	var word string
	fmt.Fscanln(os.Stdin, &word) //необходимо вводить предложение без пробелов
	word = strings.ToLower(word)
	sb := make([]string, 0, len(word))
	for _, ch := range word {
		if unicode.Is(unicode.Latin, ch) {
			sb = append(sb, string(ch))
		}
	}
	return sb
}
