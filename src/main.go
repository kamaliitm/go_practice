package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(knapsackIterative([][]int32{
		{65, 20}, {35, 8}, {245, 60}, {195, 55}, {65, 40}, {150, 70}, {275, 85}, {155, 25},
		{120, 30}, {320, 65}, {75, 75}, {40, 10}, {200, 95}, {100, 50}, {220, 40}, {99, 10},
	}, 130))

	// slBoard := NewSnakeAndLadderBoard()
	// slBoard.addSnake(62, 5)
	// slBoard.addSnake(33, 6)
	// slBoard.addSnake(49, 9)
	// slBoard.addSnake(88, 16)
	// slBoard.addSnake(41, 20)
	// slBoard.addSnake(56, 53)
	// slBoard.addSnake(98, 64)
	// slBoard.addSnake(93, 73)
	// slBoard.addSnake(95, 75)

	// slBoard.addLadder(2, 37)
	// slBoard.addLadder(27, 46)
	// slBoard.addLadder(10, 32)
	// slBoard.addLadder(51, 68)
	// slBoard.addLadder(61, 79)
	// slBoard.addLadder(65, 84)
	// slBoard.addLadder(71, 91)
	// slBoard.addLadder(81, 100)

	// slGame := NewSnakeAndLadderGame([]SLPlayer{
	// 	{name: "Gaurav"},
	// 	{name: "Sagar"},
	// }, *slBoard)

	// slGame.startGame()

	// const (
	// 	u1 = "u1"
	// 	u2 = "u2"
	// 	u3 = "u3"
	// 	u4 = "u4"
	// )
	// splitwiseGroup := NewSplitwiseGroup([]User{
	// 	{userId: u1},
	// 	{userId: u2},
	// 	{userId: u3},
	// 	{userId: u4},
	// })

	// splitwiseGroup.showBalancesForAll()
	// splitwiseGroup.addExpense(u1, 1000, 4, []string{u1, u2, u3, u4}, 0, nil)
	// splitwiseGroup.showBalancesForAll()

	// tests := [][]int{
	// 	{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
	// 	{4, 2, 0, 3, 2, 5},
	// }
	// for _, t := range tests {
	// 	fmt.Println(waterTrapped(t))
	// }

	// tests := [][]int32{
	// 	{7, 15, 12, 3},
	// 	{2, 5, 3, 1},
	// 	{3, 4, 2, 5, 1},
	// 	{656085744, 592976686, 3037922, 82266352, 17574000, 344340000, 8406892, 591292449, 569625472, 899357375, 251327440, 301303036, 281400020, 77370228, 15516426, 82859300, 88364436, 39767760, 148417500, 306863056, 10926174, 118195200, 310408774, 309307894, 200852782, 82193280, 424056750, 249277128, 180368880, 477624006, 86748948, 7434336, 48882310, 112635040, 6614541, 503907132, 598034610, 160500171, 70444416, 72752680, 271416096, 30521205, 529365648, 399367584, 129849984, 263500556, 76737948, 464269640, 613416088, 162724716, 163420800, 720512988, 1217212920, 727647624, 383190420, 8350904, 3456024, 289141064, 123384024, 158867856, 82005504, 36225521, 533012608, 54370440, 17671500, 53627000, 26597644, 855638940, 55343960, 57828624, 108025344, 21431808, 1182600, 265643950, 30054300, 219553915, 74203500, 658061160, 7502400, 931691763, 295769136, 107345925, 80109000, 130922055, 33544944, 65280, 452996453, 301655430, 7828912, 425016000, 297635898, 26861016, 739961600, 928116, 19645470, 8691456, 5123880, 596100015, 2735436, 25596483, 173620720, 202797504, 161748972, 30122300, 11082820, 574006860, 426732182, 71136825, 105659136, 1808140278, 450779280, 286831620, 104683008, 938781480, 175050736, 255681692, 67096152, 119518575, 15449840, 25273458, 165048960, 5642130, 85199958, 354920488, 446786340, 131214816, 41533296, 25766518, 90782304, 59007600, 700369740, 122021794, 56982366, 238027920, 434370816, 223677580, 72463156, 355858300, 144914616, 145950, 13822570, 19914930, 11072100, 21450528, 124958730, 105156800, 20843784, 781192188, 448358850, 6139822, 95694780, 78713888, 677177472, 9963972, 366432768, 181113408, 725292862, 473528052, 12864000, 518355540, 55832070, 318508876, 89963781, 1796290329, 844308846, 428627693, 276255100, 123609720, 440449488, 27589680, 426614166, 110068200, 408846096, 620309228, 186565236, 49051552, 561738897, 650114105, 32646556, 7174400, 275364045, 945364797, 3674160, 66314292, 11073770, 14885370, 245088324, 669628848, 33110250, 971699976, 324099072, 259496060, 492110802, 52206516, 508725376, 1534995792, 148078816, 57993375, 121071195, 381960195, 12496176, 23728250, 159836835, 712982980, 160098622, 909675852, 110400300, 485423372, 30637838, 339925836, 285371600, 13618242, 80809650, 92375040, 265612788, 1151909241, 234661320, 16962144, 213417000, 269646860, 1015650090, 117439476, 53164566, 6946134, 89506800, 305469360, 13406432, 292353000, 9969642, 43982198, 23887296, 67730660, 16384192, 218824704, 1082660306, 1679473908, 136336179, 39265884, 1077096020, 464272368, 87048192, 56752487, 156212388, 360621525, 8472816, 17613600, 62143172, 82696127, 79939536, 155805468, 159499963, 262072360, 39827904, 179532360, 2455040, 327280740, 148409340, 73738980, 3394872, 3082560, 225009981, 188912256, 179487784, 340340, 315171072, 96680664, 621206280, 536379504, 391714074, 65055960, 105570465, 365721408, 106712025, 286764660},
	// }
	// for _, t := range tests {
	// 	fmt.Println(lilysHomework(t))
	// }

	// fmt.Println(arrayManipulation(5, [][]int32{{1, 2, 100}, {2, 5, 100}, {3, 4, 100}}))
	// fmt.Println(largestRectangle([]int32{8979, 4570, 6436, 5083, 7780, 3269, 5400, 7579, 2324, 2116}))
	// fmt.Println(distances(10, 6, [][]int32{{3, 1}, {10, 1}, {10, 1}, {3, 1}, {1, 8}, {5, 2}}, 3))

	// noPrefix([]string{"aab", "defgab", "abcde", "aabcde", "cedaaa", "bbbbbbbbbb", "jabjjjad"})

	// fmt.Println(findNumOfWaysToBottomRightInMatrix(5, 5))
	// fmt.Println(findLevenshteinDistance("Carthorse", "Orchestra"))
	// fmt.Println(findNumOfScoreCombinations([]int32{2, 3, 7}, 12))
	// fmt.Println(gridChallenge([]string{"eabcd", "fghij", "olkmn", "trpqs", "xywuv"}))

	// fmt.Println(powerSet([]int32{0, 1, 2}))
	// intervals := []Interval{
	// 	{left: EndPoint{0, false}, right: EndPoint{3, false}},
	// 	{left: EndPoint{1, true}, right: EndPoint{1, true}},
	// 	{left: EndPoint{2, true}, right: EndPoint{4, true}},
	// 	{left: EndPoint{3, true}, right: EndPoint{4, false}},
	// 	{left: EndPoint{5, true}, right: EndPoint{7, false}},
	// 	{left: EndPoint{7, true}, right: EndPoint{8, false}},
	// 	{left: EndPoint{8, true}, right: EndPoint{11, false}},
	// 	{left: EndPoint{9, false}, right: EndPoint{11, true}},
	// 	{left: EndPoint{12, true}, right: EndPoint{14, true}},
	// 	{left: EndPoint{12, false}, right: EndPoint{16, true}},
	// 	{left: EndPoint{13, false}, right: EndPoint{15, false}},
	// 	{left: EndPoint{16, false}, right: EndPoint{17, false}},
	// }
	// fmt.Println(unionOfIntervals(intervals))

	// fmt.Println(characterFrequencyInText("AMAZ amaz"))
	// fmt.Println(intersectionOfSortedArrays([]int32{2, 3, 3, 5, 5, 6, 7, 7, 8, 12}, []int32{5, 5, 6, 8, 8, 9, 10, 10}))
	// cache := Cache{}
	// fmt.Println(cache.get("a"))
	// cache.set("a", "30")
	// fmt.Println(cache.get("a"))
	// fmt.Println(cache.get("b"))
	// cache.set("b", "31")
	// fmt.Println(cache.get("b"))
	// cache.set("b", "32")
	// fmt.Println(cache.get("b"))

	// fmt.Println(findKthLargestWithInPlaceOperations([]int{3, 2, 1, 5, 4}, 1))
	// fmt.Println(findSmallestInCyclicSortedArray([]int{378, 478, 550, 631, 103, 203, 220, 234, 279, 368}))
	// fmt.Println(firstOccurenceOfK([]int{-14, -10, 2, 108, 108, 243, 285, 285, 285, 401}, 285))
	// fmt.Println(kSortArr([]int{3, -1, 2, 6, 4, 5, 8}, 2))
	// fmt.Println(mergeSortedArrays([][]int{{1, 5, 9, 15}, {2, 6, 7, 17}, {4, 8, 11, 14}}))

	// stack := Stack{}
	// stack.push(4)
	// stack.push(5)
	// stack.push(5)
	// stack.push(3)
	// fmt.Println("max: ", stack.max())
	// stack.push(8)
	// fmt.Println("max: ", stack.max())
	// fmt.Println("pop: ", stack.pop())
	// fmt.Println("pop: ", stack.pop())
	// fmt.Println("max: ", stack.max())
	// fmt.Println("pop: ", stack.pop())
	// fmt.Println("max: ", stack.max())

	// n1 := Node{15, nil}
	// n2 := Node{8, &n1}
	// n3 := Node{5, &n2}
	// n4 := Node{2, &n3}
	// list1 := &n4

	// n11 := Node{16, nil}
	// n12 := Node{6, &n11}
	// n13 := Node{4, &n12}
	// n14 := Node{1, &n13}
	// list2 := &n14

	// fmt.Println(mergeSortedLinkedLists(list1, list2))
	// fmt.Println(lookAndSay(6))
	// fmt.Println(reverseWordsInSentence("Alice likes Bob and Sam"))
	// fmt.Println(replaceAndRemove([]string{"a", "c", "d", "b", "b", "c", "a", "b", "e", "-"}))
	// fmt.Println(baseConvert("615", 7, 13))
	// fmt.Println(intToString(314))
	// fmt.Println(stringToInt("-5548"))
	// fmt.Println(sample([]int32{1, 3, 5, 6, 8, 9, 10}, 3))
	// fmt.Println(maxProfit([]int32{310, 315, 275, 295, 260, 270, 290, 230, 255, 250}))
	// cache := Cache{}
	// cache.configure(2)
	// cache.set("a", "")
	// value := cache.get("a")
	// fmt.Println(value)
	// fmt.Println(deleteDuplicates([]int32{2, 3, 5, 5, 7, 11, 11, 11, 13}))
	// fmt.Println(plusOne([]int32{9, 9, 9}))
	// fmt.Println(reverse(-524))
	// arr := []int64{6, 1, 2, 4, 3, 4, 1, 4, 7, 9}
	// fmt.Println("init arr: ", arr)
	// dutchNationalFlag(arr, 3)
	// fmt.Println("final arr: ", arr)
}

// func main() {
// 	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

// 	// stdout, err := os.Create(os.Getenv("/Users/kamal/pani/learn/interviews/go_practice/out"))
// 	// checkError(err)

// 	// defer stdout.Close()

// 	// writer := bufio.NewWriterSize(stdout, 16*1024*1024)

// 	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
// 	checkError(err)
// 	t := int32(tTemp)

// 	for tItr := 0; tItr < int(t); tItr++ {
// 		s := readLine(reader)

// 		result := isBalanced(s)

// 		// fmt.Fprintf(writer, "%s\n", result)
// 		fmt.Println(result)
// 	}

// 	// writer.Flush()
// }

// func readLine(reader *bufio.Reader) string {
// 	str, _, err := reader.ReadLine()
// 	if err == io.EOF {
// 		return ""
// 	}

// 	return strings.TrimRight(string(str), "\r\n")
// }

// func checkError(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }

// type Stack struct {
// 	data []int64
// 	size int64
// }

// func (stack *Stack) push(e int64) {
// 	if stack.data == nil {
// 		stack.data = []int64{}
// 	}

// 	stack.data = append(stack.data, e)
// 	stack.size += 1

// 	fmt.Println("stack after push", stack)
// }

// func (stack *Stack) pop() int64 {
// 	if stack.data == nil || len(stack.data) == 0 {
// 		return -1
// 	}

// 	e := stack.data[stack.size-1]
// 	stack.data = stack.data[:stack.size-1]
// 	stack.size -= 1

// 	fmt.Println("stack after pop", stack)

// 	return e
// }

// func (stack *Stack) peek() int64 {
// 	if stack.data == nil || len(stack.data) == 0 {
// 		return -1
// 	}

// 	return stack.data[stack.size-1]
// }

// type Queue struct {
// 	stack1 *Stack // for push
// 	stack2 *Stack // for pop
// }

// func (queue *Queue) enqueue(e int64) {
// 	if queue.stack1 == nil {
// 		queue.stack1 = &Stack{data: []int64{}}
// 	}

// 	queue.stack1.push(e)
// }

// func (queue *Queue) dequeue() int64 {
// 	if queue.stack1 == nil && queue.stack2 == nil {
// 		return -1
// 	}

// 	if queue.stack2 != nil && queue.stack2.size > 0 {
// 		return queue.stack2.pop()
// 	}

// 	if queue.stack2 == nil {
// 		queue.stack2 = &Stack{data: []int64{}}
// 	}

// 	stack1Size := queue.stack1.size
// 	for i := int64(0); i < stack1Size; i++ {
// 		queue.stack2.push(queue.stack1.pop())
// 	}

// 	return queue.stack2.pop()
// }

// func (queue *Queue) print() {
// 	if queue.stack1 == nil && queue.stack2 == nil {
// 		return
// 	}

// 	if queue.stack2 != nil && queue.stack2.size > 0 {
// 		fmt.Println(queue.stack2.peek())
// 		return
// 	}

// 	if queue.stack2 == nil {
// 		queue.stack2 = &Stack{data: []int64{}}
// 	}

// 	stack1Size := queue.stack1.size
// 	for i := int64(0); i < stack1Size; i++ {
// 		queue.stack2.push(queue.stack1.pop())
// 	}

// 	fmt.Println(queue.stack2.peek())
// }

// func main() {
// 	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

// 	stdout, err := os.Create("/Users/kamal/pani/learn/interviews/go_practice/out")
// 	checkError(err)

// 	defer stdout.Close()

// 	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

// 	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

// 	qInt, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
// 	checkError(err)

// 	queue := Queue{}
// 	for i := int64(0); i < qInt; i++ {
// 		queryTokens := strings.Split(strings.TrimSpace(readLine(reader)), " ")
// 		queryId, err := strconv.ParseInt(queryTokens[0], 10, 64)
// 		checkError(err)

// 		switch queryId {
// 		case 1:
// 			entry, err := strconv.ParseInt(queryTokens[1], 10, 64)
// 			checkError(err)
// 			queue.enqueue(entry)
// 			fmt.Println("\nenqueue", queue.stack1, queue.stack2)
// 		case 2:
// 			queue.dequeue()
// 			fmt.Println("\ndequeue", queue.stack1, queue.stack2)
// 		default:
// 			queue.print()
// 		}
// 	}

// 	writer.Flush()
// }

// func readLine(reader *bufio.Reader) string {
// 	str, _, err := reader.ReadLine()
// 	if err == io.EOF {
// 		return ""
// 	}

// 	return strings.TrimRight(string(str), "\r\n")
// }

// func checkError(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }

func reverse(x int64) int64 {
	isNegative := x < 0
	x = int64(math.Abs(float64(x)))
	var res int64 = 0

	for x > 0 {
		res = (res * 10) + (x % 10)
		x /= 10
	}

	if isNegative {
		res = -res
	}

	return res

}

func power(x int64, y int64) int64 {
	if y == 0 {
		return 1
	}

	var res int64 = 1

	for y > 0 {
		if y&1 == 1 {
			res *= x
		}

		y >>= 1
		x *= x
	}

	return res
}

func parity(x int64) int64 {
	x ^= x >> 32
	x ^= x >> 16
	x ^= x >> 8
	x ^= x >> 4
	x ^= x >> 2
	x ^= x >> 1
	return x & 1
}
