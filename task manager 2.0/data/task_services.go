package data

import (
	"task_manager/models"
)

type Task_manager struct{
	Tasks []*models.Task
	NextId int
}

func (t *Task_manager) Get_task(id string)(task models.Task, ok bool){
	for i := range len(t.Tasks){
		if t.Tasks[i].ID == id {
			task = *t.Tasks[i]
			ok = true
			return
		}
	}

	task = models.Task{}
	ok = false
	return
}

func (t *Task_manager) Delete_task(id string) bool{
	for i := range len(t.Tasks){
		if (t.Tasks)[i].ID == id{
			t.Tasks = append((t.Tasks)[:i], (t.Tasks)[i + 1:]...)
			return true
		}
	}

	return false
}

func (t *Task_manager) Update_task(id string, task models.Task) (models.Task, bool){
	for i := range len(	t.Tasks){
		if (t.Tasks)[i].ID == id{
			t.Tasks[i] = &task
			t.Tasks[i].ID = id
			return *t.Tasks[i], true
		}
	}

	return models.Task{}, false
}