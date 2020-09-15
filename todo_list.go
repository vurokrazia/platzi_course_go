package main

import "fmt"

type task struct {
	name        string
	description string
	completed   bool
}

func (t *task) markComplete() {
	t.completed = true
}

func (t *task) updateDescription(description string) {
	t.description = description
}

func (t *task) updateName(name string) {
	t.name = name
}

type taskList struct {
	tasks []*task
}

func (t *taskList) addTaskToLisk(item *task) {
	t.tasks = append(t.tasks, item)
}

func (t *taskList) deleteList(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
}

func (t *taskList) printTaskList() {
	for _, task := range t.tasks {
		fmt.Println(task.name, task.description)
	}
}

func (t *taskList) printTaskListCompleted() {
	for _, task := range t.tasks {
		if task.completed {
			fmt.Println(task.name, task.description)
		}
	}
}

func main() {
	t1 := &task{
		name:        "Arriba",
		description: "docker!!",
		completed:   false,
	}
	t2 := &task{
		name:        "Arriba",
		description: "rails!!",
		completed:   false,
	}
	//fmt.Println(t)
	//fmt.Printf("%+v\n", t)
	//t.markComplete()
	//t.updateName("Test")
	//t.updateDescription("Test")
	//fmt.Printf("%+v\n", t)

	list := &taskList{
		tasks: []*task{
			t1, t2,
		},
	}

	list.addTaskToLisk(t1)
	fmt.Printf("%+v\n", *list.tasks[0])
	fmt.Println(len(list.tasks))
	list.deleteList(0)
	fmt.Println(len(list.tasks))
	list.printTaskList()
	list.tasks[0].markComplete()
	list.printTaskListCompleted()

	mapTasks := make(map[string]*taskList)
	mapTasks["Yisus"] = list

	t3 := &task{
		name:        "Arriba",
		description: "Java!!",
		completed:   false,
	}
	t4 := &task{
		name:        "Arriba",
		description: "C!!",
		completed:   false,
	}

	list2 := &taskList{
		tasks: []*task{
			t3, t4,
		},
	}

	mapTasks["Jss"] = list2
	fmt.Println("Task Yisus")
	mapTasks["Yisus"].printTaskList()
	fmt.Println("Task Jss")
	mapTasks["Jss"].printTaskList()

}
