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
	var err error

	if task.Id > 0 {
		err = l.db.Where(taskManagementEntity.Task{Id: task.Id}).
			Assign(&task).
			FirstOrCreate(&taskManagementEntity.Task{}).Error
	} else {
		err = l.db.Create(&task).Error
	}

	return *task, err
}

func (l *taskRepo) GetAllTasks(filter *sharedModel.CriteriaFilter) ([]taskDto.TaskRecord, int64, error) {
	tasks := []taskDto.TaskRecord{}

	query := l.db.Table("tasks t").
		Joins("JOIN task_status ts ON t.status_id = ts.id").
		Joins("JOIN users u ON t.created_by_id = u.id").
		Joins("LEFT JOIN users au ON t.assigned_to_id = au.id").
		Select(`
			t.*, ts.name as status, 
			ts.color, 
			CONCAT(u.name, ' ', u.last_name) as created_by,
			CONCAT(au.name, ' ', au.last_name) as assigned_to
		`)

	if search, ok := (filter.Filters)["search"]; ok && search != "" {
		query = query.Where("CONCAT(t.title, ' ', t.description) LIKE ?", "%"+search.(string)+"%")
	}

	// Only showing to the user the tasks created by the logged user or assigned to them
	if loggedUser, ok := (filter.Filters)["loggedUser"]; ok && loggedUser != "" {
		query = query.Where("(t.assigned_to_id = ? OR t.created_by_id = ?)", loggedUser, loggedUser)
	}

	// Filter by task status
	if status, ok := (filter.Filters)["status"]; ok && status != "" {
		query = query.Where("t.status_id = ?", status)
	}

	// Filter for user creation(custom filter)
	if createdBy, ok := (filter.Filters)["createdBy"]; ok && createdBy != "" {
		query = query.Where("t.created_by_id = ?", createdBy)
	}

	// Filter for user assignation(custom filter)
	if assignedTo, ok := (filter.Filters)["assignedTo"]; ok && assignedTo != "" {
		query = query.Where("t.assigned_to_id = ?", assignedTo)
	}

	total := gormUtil.GetCount(l.db, query)

	query = query.Limit(int(filter.Limit))
	query = query.Offset(int(filter.Skip))

	query = query.Order("t.id desc")
	l.db.Raw("?", query).Scan(&tasks)

	return tasks, total, nil
}

func (l *taskRepo) GetTaskRecordById(taskId uint) (taskDto.TaskRecord, error) {
	task := taskDto.TaskRecord{}

	err := l.db.Table("tasks t").
		Joins("JOIN task_status ts ON t.status_id = ts.id").
		Joins("JOIN users u ON t.created_by_id = u.id").
		Joins("LEFT JOIN users au ON t.assigned_to_id = au.id").
		Select(`
			t.*, ts.name as status, 
			ts.color, 
			CONCAT(u.name, ' ', u.last_name) as created_by,
			CONCAT(au.name, ' ', au.last_name) as assigned_to
		`).
		Where("t.id = ?", taskId).
		First(&task).Error

	return task, err
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

func (l *taskRepo) DeleteTask(taskId uint) error {
	err := l.db.
		Where("id = ?", taskId).
		Delete(taskManagementEntity.Task{}).Error

	return err
}

func NewTaskRepo(db *gorm.DB) taskManagementContract.TaskRepository {
	return &taskRepo{
		db: db,
	}
}
