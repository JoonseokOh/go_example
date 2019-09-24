package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// startCardGame()

	// card := makeCard()
	// fmt.Println("origin card : ", card)
	// result := shuffleCard(card)
	// fmt.Println("suffled card : ", result)

	// a := make(map[int]int)
	// a[1] = 100
	// val, flag := a[1]
	// fmt.Println("val : ", val, ", flag : ", flag)
	// fmt.Println("a : ", a[1])

	// Create Dealer
	cardDealer := new(CardDealer)

	// Player Count.
	playerCount := 5

	// Assign player to dealer
	cardDealer.player = make([]Player, playerCount)
	for i := range cardDealer.player {
		cardDealer.player[i].coin = 10
	}

	drawGameRoundList := make([]int, 0)

	totalGameCount := 100
	drawGameCount := 0

	for gameRound := 0; gameRound < totalGameCount; gameRound++ {
		fmt.Println("===== Start : ", gameRound+1)

		// Ready Card
		card := makeCard()
		fmt.Println("origin card : ", card)

		cardDealer.cardList = shuffleCard(card)
		fmt.Println("suffled card : ", cardDealer.cardList)

		// init player gamaResult.
		for j := range cardDealer.player {
			cardDealer.player[j].cardResult = -1
		}

		// Player Count
		for j := range cardDealer.player {
			// Divide And Get Result.
			if cardDealer.player[j].coin > 0 {
				cardDealer.player[j].cardResult = (cardDealer.cardList[j*2] + cardDealer.cardList[(j*2)+1]) % 10
			}
		}

		winValue := 0
		winIndex := 0
		for j := range cardDealer.player {
			if cardDealer.player[j].cardResult > winValue {
				winValue = cardDealer.player[j].cardResult
				winIndex = j
			}
		}

		checkDrawGame := 0
		for j := range cardDealer.player {
			if cardDealer.player[j].cardResult == winValue {
				checkDrawGame++
			}
		}

		if checkDrawGame > 1 {
			drawGameCount++
			drawGameRoundList = append(drawGameRoundList, gameRound)
			fmt.Println("Draw Game : ", gameRound)
		} else {
			for j := range cardDealer.player {
				if cardDealer.player[j].coin > 0 {
					if winIndex != j {
						cardDealer.player[j].coin--
						cardDealer.player[winIndex].coin++

						if cardDealer.player[j].coin <= 0 {
							cardDealer.player[j].bankruptcyRound = gameRound
						}
					} else {
						cardDealer.player[winIndex].winGameCount++
					}
				}
			}
			fmt.Println("Result : ", cardDealer.player)
		}
	}

	fmt.Println("========================")
	fmt.Println("Total Game : ", totalGameCount)
	for j := range cardDealer.player {
		fmt.Println("Player ", j, " : coin (", cardDealer.player[j].coin, "), win(", cardDealer.player[j].winGameCount, "), bankruptcyRound(", cardDealer.player[j].bankruptcyRound, ")")
	}
	fmt.Println("Draw Game : ", drawGameCount, " : ", drawGameRoundList)
}

type CardDealer struct {
	player   []Player
	cardList []int
}

type Player struct {
	coin            int
	cardResult      int
	winGameCount    int
	bankruptcyRound int
}

func increase(number *int) {
	*number++
}

func startCardGame() {

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
			pickedCard := cardList[pickCard(20)]
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

func pickCard(cardSize int) int {
	return rand.Intn(cardSize)
}

func shuffleCard(card []int) []int {

	myCard := make([]int, len(card))
	for index := range myCard {
		s1 := rand.NewSource(time.Now().UnixNano())
		rand := rand.New(s1)
		randomNumber := rand.Intn(len(card))
		myCard[index] = card[randomNumber]
		card = append(card[:randomNumber], card[randomNumber+1:]...)
	}
	return myCard
}

func makeCard() []int {
	result := make([]int, 20)
	for i := 0; i < 20; i++ {
		result[i] = (i + 1) % 10
		if result[i] == 0 {
			result[i] = 10
		}
	}
	return result
}

func checkCard(card []int) int {
	flag := 0
	for _, temp := range card {
		flag = flag + temp
	}
	return flag
}
