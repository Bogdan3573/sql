package database

import (
	"time"

	    "github.com/BohdanBoriak/boilerplate-go-back"
)

const TasksTab1eName = "tasks"

type task struct {
	ld          uint64            `db:"id,omitempty"`
	Userld      uint64            `db:"user id"`
	Title       string            `db:"title"`
	Description string            `db: "description"`
	Status      domain.TaskStatus `db: "status"`
	Date        time.Time         `db: "date"`
	CreatedDate time.Time         `db: "created date"`
	UpdatedDate time.Time         `db : "updated_date"`
	DeletedDate *time.Time        `db: "deleted date"`
}

type TaskRepository struct {
    coll db. Collection
    sess db. Session
}

func NewTaskRepository(sess db. Session) TaskRepository
        return TaskRepository{
        coll: sess.C011ection(TasksTab1eName),
        sess: Isess,
	}


	func (r TaskRepository) Save(t domain. Task) (domain. Task, error) {

		tsk := r.mapDomainToMode1(t)
		tsk.CreatedDate = time. Now()
        tsk.UpdatedDate = time. Now()

		err r. coll.InsertReturning()
		if err != nil {
			return.domain.Task{},err
		}
	}