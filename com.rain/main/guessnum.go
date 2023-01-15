package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	guessNumber := rand.Intn(maxNum)

	fmt.Println("输入你猜的数字:")
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("error", err)
			continue
		}
		input = strings.TrimSuffix(input, "\r\n")

		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("error", err)
			continue
		}
		fmt.Println("你猜的数字是:", guess)
		if guess > guessNumber {
			fmt.Println("猜大了")
		} else if guess < guessNumber {
			fmt.Println("猜小了")
		} else {
			fmt.Println("猜对了")
			break
		}
	}
}
