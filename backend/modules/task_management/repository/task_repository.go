package repository

import (
	taskManagementContract "stm/modules/task_management/domain/contract"
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

func (l *taskRepo) GetAllTasks(filter *sharedModel.CriteriaFilter) ([]taskManagementEntity.Task, int64, error) {
	tasks := []taskManagementEntity.Task{}

	query := l.db.Model(taskManagementEntity.Task{})

	if search, ok := (filter.Filters)["search"]; ok && search != "" {
		query = query.Where("concat(title, ' ', description, ' ', last_name) LIKE ?", "%"+search.(string)+"%")
	}

	if status, ok := (filter.Filters)["status"]; ok && status != "" {
		query = query.Where("status_id = ?", status.(uint))
	}

	total := gormUtil.GetCount(l.db, query)

	query = query.Limit(int(filter.Limit))
	query = query.Offset(int(filter.Skip))

	query = query.Order("id desc")
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

func NewTaskRepo(db *gorm.DB) taskManagementContract.TaskRepository {
	return &taskRepo{
		db: db,
	}
}
