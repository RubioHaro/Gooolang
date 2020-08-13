package main

import "fmt"

type task struct {
	name        string
	description string
	done        bool
}

type taskList struct {
	// Los slices son mas poderosos que los arreglos.
	tasks []*task
}

func (list *taskList) addTask(t *task) {
	list.tasks = append(list.tasks, t)
}

func (list *taskList) removeTask(index int) {
	list.tasks = append(list.tasks[:index], list.tasks[index+1:]...)
}

func (t *task) setAsDone() {
	t.done = true
}

func (t *task) updateDescription(description string) {
	t.description = description
}

func (t *task) updateName(name string) {
	t.name = name
}

func (list *taskList) print() {
	// for i := 0; i < len(list.tasks); i++ {
	// 	fmt.Println("id: ", i, "Nombre:", list.tasks[i].name)
	// }

	// Forma mas elegante de imprimirla
	for index, tarea := range list.tasks {
		fmt.Println("id: ", index, "Nombre:", tarea.name)
	}
}

func main() {
	t1 := &task{
		name:        "Completar curso Go",
		description: "Realizar el curso de GO esta semana",
	}

	t2 := &task{
		name:        "Estudiar Teoria de la computacion",
		description: "Realizar guias, practicas y ejercicio sobre teoria de la computacion",
	}

	t3 := &task{
		name:        "Estudiar Calculo Aplicado",
		description: "Realizar guias, practicas y ejercicio sobre el calculo aplicado",
	}

	list := &taskList{
		tasks: []*task{
			t1, t2,
		},
	}
	list.addTask(t3)

	taskMap := make(map[string]*taskList)
	taskMap["Roy"] = list

	t4 := &task{
		name:        "Completar curso Python",
		description: "Realizar el curso de Python esta semana",
	}

	t5 := &task{
		name:        "Estudiar Teoria de la SCRUM",
		description: "Hacer SCRUMs",
	}

	list2 := &taskList{
		tasks: []*task{
			t4, t5,
		},
	}

	taskMap["George"] = list2

	fmt.Println("********Tareas de ROY*********")
	taskMap["Roy"].print()
	fmt.Println("*******Tareas de George*******")
	taskMap["George"].print()

}
