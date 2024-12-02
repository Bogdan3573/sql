package app

import  "github.com/BohdanBoriak/boilerplate-go-back"

type TaskService struct {
	taskRepo database. TaskRepository
}

func NewTaskService(tr database. TaskRepository) TaskService {
    return TaskService{
		taskRepo: tr,
	}

}

func (s TaskService) Save(t domain. Task) (domain. Task, error) {
	tsk, err := s. taskRepo.Save(t)
    if err nil {
		log.Printf("TaskService - > Save -> s.taskRepo.Save: %s",
		return domain. Task{}, err
	}
	return tsk, nil
}