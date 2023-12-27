package main

type CoolDownTimeNode struct {
	task          string
	minTimeToPick int
}

func leastInterval(tasks []byte, n int) int {
	processedTasks := map[string]int{}
	timer := 0
	pendingTasks := []CoolDownTimeNode{}

	for _, task := range tasks {
		taskStr := string(task)
		processedTime, ok := processedTasks[taskStr]
		if !ok || (timer-processedTime) >= n {
			timer += 1
			processedTasks[taskStr] = timer
			continue
		}

		pendingTasks = append(pendingTasks, CoolDownTimeNode{task: taskStr, minTimeToPick: timer + n})
	}

	i := 0
	for i < len(pendingTasks) {
		node := pendingTasks[i]
		if node.minTimeToPick >= timer {
			timer = node.minTimeToPick + 1
		} else {
			node.minTimeToPick = timer + n
			pendingTasks = append(pendingTasks, node)
		}

		i += 1
	}

	return timer
}
