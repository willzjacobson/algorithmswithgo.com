package main

import "fmt"

/*
	given a list of tasks that:
	 - each take unit time
	 - each have a deadline of when they should be complete
	 - each have a penalty for if they are not completed on time
	create an algorithm to schedule the tasks so as to incur the minimum total penalty.

	Basically, we're looking for the optimal subset of a "weighted matroid" (weighted because each task has a weight).
	This optimal set will be the set of tasks that are completed on time (and thus as a result, the penalties incurred for the laste tasks will be minimized)
	Since all weights are positive, the optimal subset is by definition also a maximal subset (as many tasks as can be crammed in).
	In this case, we say that a set of tasks A is independent if there is a way to scedule the tasks such that none are late.

	In this example, the tasks are in order of decreasing weights. If they were not already in this order, we'd have to order them that way in order to use a greedy algo.
*/

// Task : type describing a task
type Task struct {
	id      int
	dealine int
	penalty int
}

// GreedyTaskScheduler : Greedy algorithm to determine optimal order of tasks to incur loweat total penalty for late tasks
// returns: task order, number tasks completed on time, total late penalty incurred
func GreedyTaskScheduler(tasks []Task) ([]Task, int, int) {
	optimalIndependentSetOfTasks := []Task{}
	remainingTasks := []Task{}
	suOfLatePenalties := 0

	for _, t := range tasks {
		// only add this task to the list of optimal tasks if, when this task is added, all the tasks in the list can be completed by their deadlines
		timeSlotsAvailableIfTaskIsChosen := len(optimalIndependentSetOfTasks) + 1
		if t.dealine >= timeSlotsAvailableIfTaskIsChosen {
			// regardless of the deadlines of the other tasks in the optimal subset, we know this task can be added
			optimalIndependentSetOfTasks = append(optimalIndependentSetOfTasks, t)
			continue
		}

		fitsInOptimalSet := false
		for i, o := range optimalIndependentSetOfTasks {
			// check that a task already in the optimal set can be moved later,
			// and that the task we're trying to add can take its place without being late
			if o.dealine >= timeSlotsAvailableIfTaskIsChosen && t.dealine >= i+1 {
				fitsInOptimalSet = true
				break
			}
		}

		if fitsInOptimalSet {
			optimalIndependentSetOfTasks = append(optimalIndependentSetOfTasks, t)
		} else {
			// this task didn't make the cut
			remainingTasks = append(remainingTasks, t)
			suOfLatePenalties += t.penalty
		}
	}

	// concatinate the optimal set of tasks with the set of remaining tasks into an order of operations (order of the optimal matters, order of the remaining does not)
	taskOrder := append(optimalIndependentSetOfTasks, remainingTasks...)

	return taskOrder, len(optimalIndependentSetOfTasks), suOfLatePenalties
}

// GreedyTaskSchedulerFaster : another greedy implementation that performs slightly faster than GreedyTaskScheduler
func GreedyTaskSchedulerFaster(tasks []Task) ([]Task, int, int) {
	// initialize
	order := make([]Task, len(tasks))
	numOnTimeTasks := 0
	sumOfLatePenalties := 0
	for i := range tasks {
		order[i] = Task{id: -1}
	}

	// order tasks
	for _, t := range tasks {
		placed := false
		// starting the loop at 'deadline-1' rather than 'deadline' since we index slices starting at 0; can only fit 1 task in before t=1, not 2 tasks
		for j := t.dealine - 1; j >= 0 && placed == false; j-- {
			if order[j].id == -1 {
				order[j] = t
				placed = true
				numOnTimeTasks++
			}
		}

		if placed == false {
			for j := len(order) - 1; j >= t.dealine; j-- {
				if order[j].id == -1 {
					order[j] = t
					sumOfLatePenalties += t.penalty
					break
				}
			}
		}
	}

	return order, numOnTimeTasks, sumOfLatePenalties
}

func main() {
	tasks := []Task{
		Task{id: 1,
			dealine: 4,
			penalty: 70,
		},
		Task{id: 2,
			dealine: 2,
			penalty: 60,
		},
		Task{id: 3,
			dealine: 4,
			penalty: 50,
		},
		Task{id: 4,
			dealine: 3,
			penalty: 40,
		},
		Task{id: 5,
			dealine: 1,
			penalty: 30,
		},
		Task{id: 6,
			dealine: 4,
			penalty: 20,
		},
		Task{id: 7,
			dealine: 6,
			penalty: 10,
		},
	}

	order1, ontime1, penalty1 := GreedyTaskScheduler(tasks)
	fmt.Println("order: ", order1)
	fmt.Println("ontime: ", ontime1)
	fmt.Println("penalty: ", penalty1)

	fmt.Println("-=-=-")
	order2, ontime2, penalty2 := GreedyTaskSchedulerFaster(tasks)
	fmt.Println("order: ", order2)
	fmt.Println("ontime: ", ontime2)
	fmt.Println("penalty: ", penalty2)
}
