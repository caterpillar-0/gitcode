// //---------------猜字谜游戏-----------------------------

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
	//设置随机数种子
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	//fmt.Println("The secret number is ", secretNumber)

	//用户输入数值
	fmt.Println("Please input your guess")
	//将文件通过newReader转化为一个reader变量
	reader := bufio.NewReader(os.Stdin)

	for {
		//通过reader变量的流操作，读取一行
		input, err := reader.ReadString('\n')
		if err != nil {
			//错误信息
			fmt.Println("error on input, please try again")
			return
		}
		//readstring返回结果包含“\n"，将其去掉，修剪后缀
		input = strings.Trim(input, "\r\n")
		fmt.Println(input)
		//字符串转化
		guess, err1 := strconv.Atoi(input)
		//fmt.Println("your guess is", guess)
		if err1 != nil {
			fmt.Println("error input, please enter an integer value")
			return
		}
		fmt.Println("your guess is", guess)
		//增加判断逻辑
		if guess > secretNumber {
			fmt.Println("bigger, please guess again")
		} else if guess < secretNumber {
			fmt.Println("smaller, please guess again")
		} else {
			fmt.Println("correct")
			break
		}

	}

}

// //----------------------------------猜字谜游戏v5-------------------
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"math/rand"
// 	"os"
// 	"strconv"
// 	"strings"
// 	"time"
// )

// func main() {
// 	maxNum := 100
// 	rand.Seed(time.Now().UnixNano())
// 	secretNumber := rand.Intn(maxNum)
// 	// fmt.Println("The secret number is ", secretNumber)

// 	fmt.Println("Please input your guess")
// 	reader := bufio.NewReader(os.Stdin)
// 	for {
// 		input, err := reader.ReadString('\n')
// 		if err != nil {
// 			fmt.Println("An error occured while reading input. Please try again", err)
// 			continue
// 		}
// 		input = strings.Trim(input, "\r\n")

// 		guess, err := strconv.Atoi(input)
// 		if err != nil {
// 			fmt.Println("Invalid input. Please enter an integer value")
// 			continue
// 		}
// 		fmt.Println("You guess is", guess)
// 		if guess > secretNumber {
// 			fmt.Println("Your guess is bigger than the secret number. Please try again")
// 		} else if guess < secretNumber {
// 			fmt.Println("Your guess is smaller than the secret number. Please try again")
// 		} else {
// 			fmt.Println("Correct, you Legend!")
// 			break
// 		}
// 	}
// }

// // -------------------v3------------------------------
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"math/rand"
// 	"os"
// 	"strconv"
// 	"strings"
// 	"time"
// )

// func main() {
// 	maxNum := 100
// 	rand.Seed(time.Now().UnixNano())
// 	secretNumber := rand.Intn(maxNum)
// 	fmt.Println("The secret number is ", secretNumber)

// 	fmt.Println("Please input your guess")
// 	reader := bufio.NewReader(os.Stdin)
// 	input, err := reader.ReadString('\n')
// 	if err != nil {
// 		fmt.Println("An error occured while reading input. Please try again", err)
// 		return
// 	}
// 	input = strings.Trim(input, "\r\n")

// 	guess, err := strconv.Atoi(input)
// 	if err != nil {
// 		fmt.Println("Invalid input. Please enter an integer value")
// 		return
// 	}
// 	fmt.Println("You guess is", guess)
// }
