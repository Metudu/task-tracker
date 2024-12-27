package task

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type task struct {
	ID int `json:"id"`
	Description string `json:"description"`
	Status string `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

const (
	JSON_PATH = "./task/task.json"
)

func writeJSON(tasks []*task) error {
	file, err := os.OpenFile(JSON_PATH, os.O_CREATE|os.O_TRUNC|os.O_RDWR, os.ModeDevice)
	if err != nil {
		return fmt.Errorf("an error occured when opening the file: %v", err.Error())
	}
	defer file.Close()

	data, err := json.Marshal(tasks)
	if err != nil {
		return fmt.Errorf("an error occured when marshaling the task: %v", err.Error())
	}

	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("an error occured when writing the task to json file: %v", err.Error())
	}
	return nil
}

func readJSON() ([]*task, error) {
	data, err := os.ReadFile(JSON_PATH)
	if err != nil {
		return nil, fmt.Errorf("an error occured when reading the json file: %v", err.Error())
	}

	if len(data) == 0 {
		return []*task{}, nil
	}

	var tasks []*task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, fmt.Errorf("an error occured when unmarshaling the json: %v", err.Error())
	}

	return tasks, nil
}

func createID() (int, error) {
	tasks, err := readJSON()
	if err != nil {
		log.Fatal(err)	
	}

	if len(tasks) == 0 {
		return 1, nil
	}

	return tasks[len(tasks) - 1].ID + 1, nil
}

func AddTask(description string) {
	tasks, err := readJSON()
	if err != nil {
		log.Fatal(err)
	}

	id, err := createID()
	if err != nil {
		log.Fatal(err)
	}

	t := &task{ID: id, Description: description, Status: "todo", CreatedAt: time.Now(), UpdatedAt: time.Now()}
	tasks = append(tasks, t)

	if err := writeJSON(tasks); err != nil {
		log.Fatal(err)
	}

	log.Printf("Task added successfully (ID: %v)\n", id)
}

func UpdateTask(id int, description string) {
	tasks, err := readJSON()
	if err != nil {
		log.Fatal(err)
	}

	if len(tasks) == 0 {
		log.Fatal("There is no task to update!")
	}

	for _, task := range tasks {
		if task.ID == id {
			task.Description = description
			task.UpdatedAt = time.Now()
			if err := writeJSON(tasks); err != nil {
				log.Fatal(err)
			}
			log.Printf("Task with id %v has updated!", id)
			return
		}
	}

	log.Fatalf("No matching task ID for %v\n", id)
}

func DeleteTask(id int) {
	tasks, err := readJSON()
	if err != nil {
		log.Fatal(err)
	}

	if len(tasks) == 0 {
		log.Fatal("There is no task to delete!")
	}

	var updatedTasks []*task
	for _, task := range tasks {
		if task.ID == id {
			continue
		}
		updatedTasks = append(updatedTasks, task)
	}

	if len(updatedTasks) == len(tasks) {
		log.Println("Nothing has removed!")
		return
	}

	if err := writeJSON(updatedTasks); err != nil {
		log.Fatal(err)
	}
}

func mark(id int, status string) {
	tasks, err := readJSON()
	if err != nil {
		log.Fatal(err)
	}

	if len(tasks) == 0 {
		log.Fatal("There is no task to mark!")
	}

	for _, task := range tasks {
		if task.ID == id {
			if task.Status == status {
				log.Printf("The task with id %v has already marked as %v\n", id, status)
				return
			}
			task.Status = status
			if err := writeJSON(tasks); err != nil {
				log.Fatal(err)
			}
			log.Printf("The task with id %v has marked as %v", id, status)
			return
		}
	}

	log.Fatalf("No matching task ID for %v\n", id)
}

func MarkAsInProgress(id int) {
	mark(id, "in-progress")
}

func MarkAsDone(id int) {
	mark(id, "done")
}

func (t *task) printTask() {
	fmt.Println("===============")
	fmt.Printf("Task ID: %v\n", t.ID)
	fmt.Printf("Task Description: %v\n", t.Description)
	fmt.Printf("Task Status: %v\n", t.Status)
	fmt.Println("===============")
}

func ListTasks(filter string) {
	tasks, err := readJSON()
	if err != nil {
		log.Fatal(err)
	}

	if len(tasks) == 0 {
		log.Println("There is no task to print!")
		return
	}

	if filter == "" {
		for _, task := range tasks {
			task.printTask()
		}
	} else {
		for _, task := range tasks {
			if task.Status == filter {
				task.printTask()
			}
		}
	}
}