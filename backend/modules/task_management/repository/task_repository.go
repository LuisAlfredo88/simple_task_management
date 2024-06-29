package repository

import (
	taskManagementContract "stm/modules/task_management/domain/contract"
	taskDto "stm/modules/task_management/domain/dto"
	taskManagementEntity "stm/modules/task_management/domain/entity"
	sharedModel "stm/shared/model"
	gormUtil "stm/shared/utility/gorm"

	"gorm.io/gorm"
)

type taskRepo struct {
	db *gorm.DB
}

func (l *taskRepo) Save(task *taskManagementEntity.Task) (taskManagementEntity.Task, error) {

	err := l.db.Where(taskManagementEntity.Task{Id: task.Id}).
		Assign(task).
		FirstOrCreate(&taskManagementEntity.Task{}).Error

	return *task, err
}

func (l *taskRepo) GetAllTasks(filter *sharedModel.CriteriaFilter) ([]taskDto.TaskRecord, int64, error) {
	tasks := []taskDto.TaskRecord{}

	query := l.db.Table("tasks t").
		Joins("JOIN task_status ts ON t.status_id = ts.id").
		Joins("JOIN users u ON t.created_by_id = u.id").
		Select(`
		t.*, ts.name as status, 
		ts.color, 
		CONCAT(u.name, ' ', u.last_name) as created_by
	`)

	if search, ok := (filter.Filters)["search"]; ok && search != "" {
		query = query.Where("CONCAT(t.title, ' ', t.description) LIKE ?", "%"+search.(string)+"%")
	}

	if status, ok := (filter.Filters)["status"]; ok && status != "" {
		query = query.Where("t.status_id = ?", status)
	}

	total := gormUtil.GetCount(l.db, query)

	query = query.Limit(int(filter.Limit))
	query = query.Offset(int(filter.Skip))

	query = query.Order("t.id desc")
	l.db.Raw("?", query).Scan(&tasks)

	return tasks, total, nil
}

func (l *taskRepo) GetTaskById(taskId uint) (taskManagementEntity.Task, error) {
	task := taskManagementEntity.Task{}

	err := l.db.Where(&taskManagementEntity.Task{
		Id: taskId,
	}).Find(&task).Error

	return task, err
}

func (l *taskRepo) ChangeTaskStatus(taskId uint, status uint) error {
	task := taskManagementEntity.Task{}

	err := l.db.Model(&task).
		Where("id = ?", taskId).
		Update("status_id", status).Error

	return err
}

func (l *taskRepo) GetAllTasksStatus() ([]taskManagementEntity.TaskStatus, error) {
	taskStatus := []taskManagementEntity.TaskStatus{}
	err := l.db.Find(&taskStatus).Error
	return taskStatus, err
}

func NewTaskRepo(db *gorm.DB) taskManagementContract.TaskRepository {
	return &taskRepo{
		db: db,
	}
}
