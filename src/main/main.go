package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// 카드 값 초기화.
	cardList := make([]int, 20)
	for i := 0; i < 20; i++ {
		cardValue := (i + 1)
		if cardValue > 10 {
			cardValue -= 10
		}
		cardList[i] = cardValue
	}
	fmt.Println("카드 정보 : ", cardList)

	totalGameCount := 100
	userCount := 5 // 5 를 넘어갈 수 없음
	selectedCardSize := userCount * 2

	userWin := make([]int, userCount)
	noWinner := 0

	for gameCount := 0; gameCount < totalGameCount; gameCount++ {
		fmt.Println("===== Start ", gameCount+1, "th")

		// 매번 새로운 랜덤 값 생성을 위해 사용.
		rand.Seed(time.Now().UnixNano())

		// 카드를 선택.
		selectedCardList := make([]int, selectedCardSize)
		for i := 0; i < selectedCardSize; i++ {
			duplicated := 0
			pickedCard := cardList[cardPick(20)]
			for _, value := range selectedCardList {
				if value == pickedCard {
					duplicated++
				}
				if duplicated >= 2 {
					// fmt.Println("2장이상 겹치는 경우 발견!!", value, " : ", pickedCard)
					i--
					break
				} else {
					selectedCardList[i] = pickedCard
				}
			}
		}
		fmt.Println("뽑은 카드 : ", selectedCardList)

		// 카드 계산.
		gameResultList := make([]int, userCount)
		for i := 0; i < userCount; i++ {
			gameResultList[i] = (selectedCardList[i*2] + selectedCardList[(i*2)+1]) % 10
		}

		winValue := 0
		winIndex := 0
		for index, value := range gameResultList {
			if value >= winValue {
				winValue = value
				winIndex = index
			}
		}

		checkDrawGame := 0
		for i := 0; i < userCount; i++ {
			if gameResultList[i] == winValue {
				checkDrawGame++
			}
		}

		fmt.Println("게임 결과 : ", gameResultList)
		if checkDrawGame > 1 {
			noWinner++
			fmt.Println("Draw Game!!!")
		} else {
			userWin[winIndex] = userWin[winIndex] + 1
			fmt.Println("Current Result : ", userWin)
		}
	}

	fmt.Println("=================================")
	fmt.Println("Total Game : ", totalGameCount)
	fmt.Println("Game Result: ", userWin)
	fmt.Println("Draw Game Count : ", noWinner)
}

func cardPick(cardSize int) int {
	return rand.Intn(cardSize)
}
