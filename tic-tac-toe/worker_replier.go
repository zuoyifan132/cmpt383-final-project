package main

import (
	"fmt"
	"math/rand"
	"time"
	"strings"
	"encoding/json"
	zmq "github.com/pebbe/zmq4"
	"os"
)


/*------------------------------------------helper------------------------------------------*/

// max function used by pMCTS
// return the max int out of a slice
func max(slice []int) int {
	if len(slice) < 0 {
		return 0
	}

	maxValue := slice[0]

	for _, val := range slice {
		if val > maxValue {
			maxValue = val
		}
	}

	return maxValue
}

// remove the target element from the slice
func remove(slice []int, value int) []int {
	for i, _ := range slice {
		if slice[i] == value {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return []int{}
}

/*------------------------------------------pMCTS-------------------------------------------*/

// determine if AI win the game
// return true if AI win, otherwise return false
func computer_win(l string) bool {
	if (l[0]== 120) && (l[1]== 120) && (l[2]== 120){
		return true
	} else if (l[3]== 120) && (l[4]== 120) && (l[5]== 120) {
		return true
	} else if (l[6]== 120) && (l[7]== 120) && (l[8]== 120) {
		return true
	} else if (l[0]== 120) && (l[3]== 120) && (l[6]== 120) {
		return true
	} else if (l[1]== 120) && (l[4]== 120) && (l[7]== 120) {
		return true
	} else if (l[2]== 120) && (l[5]== 120) && (l[8]== 120) {
		return true
	} else if (l[0]== 120) && (l[4]== 120) && (l[8]== 120) {
		return true
	} else if (l[2]== 120) && (l[4]== 120) && (l[6]== 120) {
		return true
	} else {
		return false
	}
}


// determine if player win the game
// return true if player win, otherwise return false
func player_win(l string) bool {
	if (l[0]== 111) && (l[1]== 111) && (l[2]== 111){
		return true
	} else if (l[3]== 111) && (l[4]== 111) && (l[5]== 111) {
		return true
	} else if (l[6]== 111) && (l[7]== 111) && (l[8]== 111) {
		return true
	} else if (l[0]== 111) && (l[3]== 111) && (l[6]== 111) {
		return true
	} else if (l[1]== 111) && (l[4]== 111) && (l[7]== 111) {
		return true
	} else if (l[2]== 111) && (l[5]== 111) && (l[8]== 111) {
		return true
	} else if (l[0]== 111) && (l[4]== 111) && (l[8]== 111) {
		return true
	} else if (l[2]== 111) && (l[4]== 111) && (l[6]== 111) {
		return true
	} else {
		return false
	}
}

// check if the status of game is draw
// return true if the game is draw, otherwsie retuen false
func draw(l string) bool {
	if (player_win(l) == false) && (computer_win(l) == false) {
		count := 0
		for i := 0; i < len(l); i++ {
			if (l[i] == 120) || (l[i] == 111) {
				count += 1
			}
		}
		if count == 9 {
			return true
		} else {
			return false
		}
	} else{
		return false
	}
}


// check if the game is over for player side
// return true if the game is over, otherwise return false
func end_game(l string) bool {
	if player_win(l) == true {
		return true
	} else if computer_win(l) == true {
		return true
	} else if draw(l) {
		return true
	} else {
		return false
	}
}


// randomly play the game or AI and choose the best score strategy for AI
// @l: the chess board
// @computer_choice: all possible option that AI could choose
// @choice: the option that AI want to choose for this move
// return the score for the @choice
func rand_playout(l string, computer_choice []int, choice int) int {
	rand.Seed(time.Now().UnixNano())
	// copy the board status
	l_temp := l

	// AI take this choice
	l_temp = strings.Replace(l_temp, string(l_temp[choice]), "x", 1)

	// copy all AI choice 
	computer_choice_temp := []int{}
	for i := 0; i < len(computer_choice); i++ {
		computer_choice_temp = append(computer_choice_temp, computer_choice[i])
	}
	// AI already take this, remove from slice
	computer_choice_temp = remove(computer_choice_temp, choice)

	/*-----------------------------Before playout the game already over-----------------------------*/
	if computer_win(l_temp) == true {
		return 2
	}

	if draw(l_temp) == true {
		return 1
	}

	if player_win(l_temp) == true {
		return -1
	}
	/*----------------------------------------------------------------------------------------------*/

	count_win := 0
	var fill int
	for {
		// randomly choose to player side
		fill = computer_choice_temp[rand.Intn(len(computer_choice_temp))]
		l_temp = strings.Replace(l_temp, string(l_temp[fill]), "o", 1)
		computer_choice_temp = remove(computer_choice_temp, fill)

		// check if after fill in "o" the game end
		if end_game(l_temp) == true {
			if player_win(l_temp) == true {
				count_win = -1
			} else if draw(l_temp) == true {
				count_win = 1
			}
			break
		}

		// choose random to "x"
		fill = computer_choice_temp[rand.Intn(len(computer_choice_temp))]
		l_temp = strings.Replace(l_temp, string(l_temp[fill]), "x", 1)
		computer_choice_temp = remove(computer_choice_temp, fill)

		// check if after fill in "x" the game end
		if end_game(l_temp) == true {
			if computer_win(l_temp) == true {
				count_win = 2
			} else {
				count_win = 0
			}
			break
		}
	}

	return count_win
}


// for current game status choose the best option
// @l: chess board
// return the best option
func pMCTS(l string, number_random_playout int, thread int) int {
	computer_choice := []int{}

	// list all posible computer choice
	for i := 0; i < len(l); i++ {
		// exclude 'x' and 'o'
		if (l[i] != 120) && (l[i] != 111) {
			computer_choice = append(computer_choice, int(l[i])-48)
		}
	}

	// make a map to store the score of randon playout
	dic := map[int]int{}
	for i := 0; i < len(computer_choice); i++ {
		dic[computer_choice[i]] = 0
	}

	// for every choice do the random playout
	for i := 0; i < len(computer_choice); i++ {
		/*for j := 0; j < number_random_playout; j++ {
			dic[computer_choice[i]] += rand_playout(l, computer_choice, computer_choice[i])
		}*/
		dic[computer_choice[i]] = worker_thread(l, computer_choice, computer_choice[i], number_random_playout, thread)
	}

	// print the map
	fmt.Println(dic)

	// find max in the dictionary value
	value := []int{}
	for i := 0; i < len(computer_choice); i++ {
		value = append(value, dic[computer_choice[i]])
	}
	maximum_value := max(value)

	for i := 0; i < len(computer_choice); i++ {
		if maximum_value == dic[computer_choice[i]] {
			return computer_choice[i]
		}
	}

	// error
	return -1
}


// worker_thread helper function in order to call it recursively
func worker_thread_helper(l string, computer_choice []int, choice int, number_random_playout int, thread int, res chan int) {
	res <- worker_thread(l, computer_choice, choice, number_random_playout, thread)
}


// separate the work to several thread to do the computation
func worker_thread(l string, computer_choice []int, choice int, number_random_playout int, thread int) int {
	treshhold := 500

	if  (number_random_playout < treshhold) || (thread <= 1) {
		count := 0
		for i := 0; i < number_random_playout; i++ {
			count += rand_playout(l, computer_choice, choice)
		}
		return count
	} else {	// run concurrent
		count := 0
		res := make(chan int, thread)
		for i := 0; i < thread; i++ {
			go worker_thread_helper(l, computer_choice, choice, int(number_random_playout/thread), thread, res)
		}
		for i := 0; i < thread; i++ {
			count += <- res 
		}
		return count
	}
}


// main function that build socket connect and send back response
func main() {

	/*--------------------------------socket setup--------------------------------*/
	fmt.Println("Replier setup...\n")

	requester, _ := zmq.NewSocket(zmq.REP)
	defer requester.Close()
	error := requester.Bind("tcp://*:5555")
	if error != nil {
		fmt.Println("replier bind failure")
		os.Exit(3)
	}

	fmt.Println("Replier setup done\n")
	/*---------------------------------------------------------------------------*/

	// recive number_random_playout and thread parameter
	var number_random_playout int = 100					// set default value 
	var thread int = 1									// set default value

	request1, _ := requester.RecvBytes(0)
	json.Unmarshal(request1, &number_random_playout)
	response1, _ := json.Marshal(true)					// acknowledgement
	requester.SendBytes(response1, 0)

	request2, _ := requester.RecvBytes(0)
	json.Unmarshal(request2, &thread)
	response2, _ := json.Marshal(true)					// acknowledgement
	requester.SendBytes(response2, 0)

	fmt.Println(number_random_playout)
	fmt.Println(thread)
	
	for {
		fmt.Println("Replier listening for requests")

		// receive request
		request, _ := requester.RecvBytes(0)
		var request_data string
		json.Unmarshal(request, &request_data)				// decode request to request_data

		fmt.Printf("Replier recives data: %v\n", request_data)

		// recive shutdown signal 
		if request_data == "!" {
			fmt.Println("Shutdown the replier...\n")
			break
		}

		// send response
		response_data := pMCTS(request_data, number_random_playout, thread)
		response, _ := json.Marshal(response_data)			// encode progressed data
		requester.SendBytes(response, 0)

		fmt.Println("Replier send data done")
	}

	fmt.Println("clsoe replier")
}

