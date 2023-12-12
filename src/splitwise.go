package main

import "fmt"

type BalanceMaxHeap struct {
	data []UserBalance
	size int
}

func NewBalanceMaxHeap() *BalanceMaxHeap {
	return &BalanceMaxHeap{
		data: []UserBalance{},
		size: 0,
	}
}

func (h *BalanceMaxHeap) isEmpty() bool {
	return h.size <= 0
}

func (h *BalanceMaxHeap) add(userBalance UserBalance) {
	h.size += 1
	h.data = append(h.data, userBalance)
	h.heapifyUp(h.size - 1)
}

func (h *BalanceMaxHeap) poll() *UserBalance {
	if h.size <= 0 {
		return nil
	}

	userBalance := h.data[0]
	h.size -= 1
	h.data[0] = h.data[h.size]
	h.heapifyDown(0)

	return &userBalance
}

func (h *BalanceMaxHeap) peek() *UserBalance {
	if h.size <= 0 {
		return nil
	}

	return &h.data[0]
}

func (h *BalanceMaxHeap) heapifyUp(index int) {
	if h.size == 0 && !h.isValidIndex(index) {
		return
	}

	parentInd := parent(index)
	if parentInd >= 0 && h.data[index].balance > h.data[parentInd].balance {
		h.swap(parentInd, index)
		h.heapifyUp(parentInd)
	}
}

func (h *BalanceMaxHeap) heapifyDown(index int) {
	if h.size == 0 || !h.isValidIndex(index) {
		return
	}

	maxIndex := index
	maxBalance := h.data[index].balance
	if leftChildInd := leftChild(index); h.isValidIndex(leftChildInd) {
		if maxBalance < h.data[leftChildInd].balance {
			maxBalance = h.data[leftChildInd].balance
			maxIndex = leftChildInd
		}
	}

	if rightChildInd := rightChild(index); h.isValidIndex(rightChildInd) {
		if maxBalance < h.data[rightChildInd].balance {
			maxBalance = h.data[rightChildInd].balance
			maxIndex = rightChildInd
		}
	}

	if maxIndex != index {
		h.swap(index, maxIndex)
		h.heapifyDown(maxIndex)
	}

}

func (h *BalanceMaxHeap) isValidIndex(index int) bool {
	return index >= 0 && index < h.size
}

func parent(index int) int {
	if index <= 0 {
		return -1
	}

	return (index - 1) / 2
}

func leftChild(index int) int {
	return (2 * index) + 1
}

func rightChild(index int) int {
	return (2 * index) + 2
}

func (h *BalanceMaxHeap) swap(i int, j int) {
	if h.size <= 0 || i <= 0 || j <= 0 || i >= h.size || j >= h.size {
		return
	}

	h.data[i], h.data[j] = h.data[j], h.data[i]
}

// User
type User struct {
	userId string
	name   string
	email  string
	phone  int64
}

type UserBalance struct {
	userId  string
	balance float32
}

// Graph
type UserBalanceEdge struct {
	userId  string
	balance float32 // amount this user (userId) owes the other user
}

type BalanceGraph struct {
	users []string
	edges map[string][]UserBalanceEdge
}

func NewBalanceGraph(users []string) *BalanceGraph {
	return &BalanceGraph{
		users: users,
		edges: map[string][]UserBalanceEdge{},
	}
}

func (g *BalanceGraph) addEdge(userId1 string, userId2 string, balance float32) {
	user1Edges, ok := g.edges[userId1]
	if ok {
		user1Edges = append(user1Edges, UserBalanceEdge{userId: userId2, balance: balance})
	} else {
		user1Edges = []UserBalanceEdge{{userId: userId2, balance: balance}}
	}

	g.edges[userId1] = user1Edges
}

func isUserPresent(users []string, userId string) bool {
	for _, u := range users {
		if u == userId {
			return true
		}
	}

	return false
}

const (
	EQUAL   = iota
	EXACT   = iota
	PERCENT = iota
)

type SplitwiseStore interface {
	// create
}

// type Balance struct {
// 	user1  string
// 	user2  string
// 	amount float32 // amount > 0 => user1 owes user2
// 	// amount < 0 => user2 owes user1
// }

type Expense struct {
	balances map[string]float32
}

type SplitwiseGroup struct {
	users []User
	// balances []Balance
	expenses []Expense
}

func NewSplitwiseGroup(users []User) *SplitwiseGroup {
	return &SplitwiseGroup{
		users: users,
		// balances: []Balance{},
		// balancesMap: map[string]map[string]float32{},
		expenses: []Expense{},
	}
}

func (sg *SplitwiseGroup) addExpense(
	payer string,
	amount float32,
	numUsers int,
	users []string,
	expenseType int,
	splitValues []float32,
) {
	if expenseType == EQUAL {
		sg.splitAndAddEqualExpense(
			payer,
			amount,
			numUsers,
			users,
		)
	}
}

func (sg *SplitwiseGroup) splitAndAddEqualExpense(
	payer string,
	amount float32,
	numUsers int,
	users []string,
) {
	shareAmount := amount / float32(numUsers)
	expense := Expense{balances: map[string]float32{}}
	for _, userId := range users {

		balance := -shareAmount
		if userId == payer {
			balance = amount - shareAmount
		}

		expense.balances[userId] = balance
	}

	sg.expenses = append(sg.expenses, expense)
}

func (sg *SplitwiseGroup) showBalancesForAll() {

	finalBalances := map[string]float32{}
	for _, expense := range sg.expenses {
		for userId, balance := range expense.balances {
			if existingBalance, ok := finalBalances[userId]; !ok {
				finalBalances[userId] = balance
			} else {
				finalBalances[userId] = existingBalance + balance
			}
		}
	}

	if len(finalBalances) == 0 {
		fmt.Println("No balances")
		return
	}

	balanceGraph := getUserBalancesGraph(finalBalances)
	for userId, userEdges := range balanceGraph.edges {
		for _, edge := range userEdges {
			fromTo := "to"
			getsOwes := "owes"
			balance := edge.balance
			if edge.balance < 0 {
				fromTo = "from"
				getsOwes = "gets back"
				balance = -balance
			}
			fmt.Printf("%s %s %f %s %s\n", userId, getsOwes, balance, fromTo, edge.userId)
		}
	}
	// fmt.Println(balanceGraph)
}

func getUserBalancesGraph(balancesMap map[string]float32) *BalanceGraph {
	positiveBalHeap := &BalanceMaxHeap{}
	negativeBalHeap := &BalanceMaxHeap{}

	for userId, balance := range balancesMap {
		if balance > 0 {
			positiveBalHeap.add(UserBalance{
				userId:  userId,
				balance: balance,
			})
		} else {
			negativeBalHeap.add(UserBalance{
				userId:  userId,
				balance: -balance,
			})

		}
	}

	users := make([]string, 0, len(balancesMap))
	for k := range balancesMap {
		users = append(users, k)
	}

	balanceGraph := NewBalanceGraph(users)

	for !positiveBalHeap.isEmpty() && !negativeBalHeap.isEmpty() {
		posBal := positiveBalHeap.poll()
		negBal := negativeBalHeap.poll()

		diff := posBal.balance - negBal.balance
		if diff > 0 {
			balanceGraph.addEdge(negBal.userId, posBal.userId, negBal.balance)
			balanceGraph.addEdge(posBal.userId, negBal.userId, -negBal.balance)
			positiveBalHeap.add(UserBalance{userId: negBal.userId, balance: diff})
		} else if diff < 0 {
			balanceGraph.addEdge(negBal.userId, posBal.userId, posBal.balance)
			balanceGraph.addEdge(posBal.userId, negBal.userId, -posBal.balance)
			negativeBalHeap.add(UserBalance{userId: negBal.userId, balance: -diff})
		} else {
			balanceGraph.addEdge(negBal.userId, posBal.userId, posBal.balance)
			balanceGraph.addEdge(posBal.userId, negBal.userId, -posBal.balance)
		}
	}

	return balanceGraph
}

func (sg *SplitwiseGroup) showBalancesForUser(userId string) {
}
