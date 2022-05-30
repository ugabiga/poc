// Code generated by SQLBoiler 4.11.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// ProjectTask is an object representing the database table.
type ProjectTask struct {
	ProjectID int64     `boil:"project_id" json:"project_id" toml:"project_id" yaml:"project_id"`
	TaskID    int64     `boil:"task_id" json:"task_id" toml:"task_id" yaml:"task_id"`
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`

	R *projectTaskR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L projectTaskL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ProjectTaskColumns = struct {
	ProjectID string
	TaskID    string
	UpdatedAt string
	CreatedAt string
}{
	ProjectID: "project_id",
	TaskID:    "task_id",
	UpdatedAt: "updated_at",
	CreatedAt: "created_at",
}

var ProjectTaskTableColumns = struct {
	ProjectID string
	TaskID    string
	UpdatedAt string
	CreatedAt string
}{
	ProjectID: "project_tasks.project_id",
	TaskID:    "project_tasks.task_id",
	UpdatedAt: "project_tasks.updated_at",
	CreatedAt: "project_tasks.created_at",
}

// Generated where

type whereHelperint64 struct{ field string }

func (w whereHelperint64) EQ(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint64) NEQ(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint64) LT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint64) LTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint64) GT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint64) GTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint64) IN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint64) NIN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var ProjectTaskWhere = struct {
	ProjectID whereHelperint64
	TaskID    whereHelperint64
	UpdatedAt whereHelpertime_Time
	CreatedAt whereHelpertime_Time
}{
	ProjectID: whereHelperint64{field: "\"project_tasks\".\"project_id\""},
	TaskID:    whereHelperint64{field: "\"project_tasks\".\"task_id\""},
	UpdatedAt: whereHelpertime_Time{field: "\"project_tasks\".\"updated_at\""},
	CreatedAt: whereHelpertime_Time{field: "\"project_tasks\".\"created_at\""},
}

// ProjectTaskRels is where relationship names are stored.
var ProjectTaskRels = struct {
	Project string
	Task    string
}{
	Project: "Project",
	Task:    "Task",
}

// projectTaskR is where relationships are stored.
type projectTaskR struct {
	Project *Project `boil:"Project" json:"Project" toml:"Project" yaml:"Project"`
	Task    *Task    `boil:"Task" json:"Task" toml:"Task" yaml:"Task"`
}

// NewStruct creates a new relationship struct
func (*projectTaskR) NewStruct() *projectTaskR {
	return &projectTaskR{}
}

func (r *projectTaskR) GetProject() *Project {
	if r == nil {
		return nil
	}
	return r.Project
}

func (r *projectTaskR) GetTask() *Task {
	if r == nil {
		return nil
	}
	return r.Task
}

// projectTaskL is where Load methods for each relationship are stored.
type projectTaskL struct{}

var (
	projectTaskAllColumns            = []string{"project_id", "task_id", "updated_at", "created_at"}
	projectTaskColumnsWithoutDefault = []string{}
	projectTaskColumnsWithDefault    = []string{"project_id", "task_id", "updated_at", "created_at"}
	projectTaskPrimaryKeyColumns     = []string{"project_id", "task_id"}
	projectTaskGeneratedColumns      = []string{}
)

type (
	// ProjectTaskSlice is an alias for a slice of pointers to ProjectTask.
	// This should almost always be used instead of []ProjectTask.
	ProjectTaskSlice []*ProjectTask
	// ProjectTaskHook is the signature for custom ProjectTask hook methods
	ProjectTaskHook func(context.Context, boil.ContextExecutor, *ProjectTask) error

	projectTaskQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	projectTaskType                 = reflect.TypeOf(&ProjectTask{})
	projectTaskMapping              = queries.MakeStructMapping(projectTaskType)
	projectTaskPrimaryKeyMapping, _ = queries.BindMapping(projectTaskType, projectTaskMapping, projectTaskPrimaryKeyColumns)
	projectTaskInsertCacheMut       sync.RWMutex
	projectTaskInsertCache          = make(map[string]insertCache)
	projectTaskUpdateCacheMut       sync.RWMutex
	projectTaskUpdateCache          = make(map[string]updateCache)
	projectTaskUpsertCacheMut       sync.RWMutex
	projectTaskUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var projectTaskAfterSelectHooks []ProjectTaskHook

var projectTaskBeforeInsertHooks []ProjectTaskHook
var projectTaskAfterInsertHooks []ProjectTaskHook

var projectTaskBeforeUpdateHooks []ProjectTaskHook
var projectTaskAfterUpdateHooks []ProjectTaskHook

var projectTaskBeforeDeleteHooks []ProjectTaskHook
var projectTaskAfterDeleteHooks []ProjectTaskHook

var projectTaskBeforeUpsertHooks []ProjectTaskHook
var projectTaskAfterUpsertHooks []ProjectTaskHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ProjectTask) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectTaskAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ProjectTask) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectTaskBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ProjectTask) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectTaskAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ProjectTask) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectTaskBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ProjectTask) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectTaskAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ProjectTask) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectTaskBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ProjectTask) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectTaskAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ProjectTask) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectTaskBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ProjectTask) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectTaskAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddProjectTaskHook registers your hook function for all future operations.
func AddProjectTaskHook(hookPoint boil.HookPoint, projectTaskHook ProjectTaskHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		projectTaskAfterSelectHooks = append(projectTaskAfterSelectHooks, projectTaskHook)
	case boil.BeforeInsertHook:
		projectTaskBeforeInsertHooks = append(projectTaskBeforeInsertHooks, projectTaskHook)
	case boil.AfterInsertHook:
		projectTaskAfterInsertHooks = append(projectTaskAfterInsertHooks, projectTaskHook)
	case boil.BeforeUpdateHook:
		projectTaskBeforeUpdateHooks = append(projectTaskBeforeUpdateHooks, projectTaskHook)
	case boil.AfterUpdateHook:
		projectTaskAfterUpdateHooks = append(projectTaskAfterUpdateHooks, projectTaskHook)
	case boil.BeforeDeleteHook:
		projectTaskBeforeDeleteHooks = append(projectTaskBeforeDeleteHooks, projectTaskHook)
	case boil.AfterDeleteHook:
		projectTaskAfterDeleteHooks = append(projectTaskAfterDeleteHooks, projectTaskHook)
	case boil.BeforeUpsertHook:
		projectTaskBeforeUpsertHooks = append(projectTaskBeforeUpsertHooks, projectTaskHook)
	case boil.AfterUpsertHook:
		projectTaskAfterUpsertHooks = append(projectTaskAfterUpsertHooks, projectTaskHook)
	}
}

// One returns a single projectTask record from the query.
func (q projectTaskQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ProjectTask, error) {
	o := &ProjectTask{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for project_tasks")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all ProjectTask records from the query.
func (q projectTaskQuery) All(ctx context.Context, exec boil.ContextExecutor) (ProjectTaskSlice, error) {
	var o []*ProjectTask

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ProjectTask slice")
	}

	if len(projectTaskAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all ProjectTask records in the query.
func (q projectTaskQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count project_tasks rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q projectTaskQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if project_tasks exists")
	}

	return count > 0, nil
}

// Project pointed to by the foreign key.
func (o *ProjectTask) Project(mods ...qm.QueryMod) projectQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.ProjectID),
	}

	queryMods = append(queryMods, mods...)

	return Projects(queryMods...)
}

// Task pointed to by the foreign key.
func (o *ProjectTask) Task(mods ...qm.QueryMod) taskQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.TaskID),
	}

	queryMods = append(queryMods, mods...)

	return Tasks(queryMods...)
}

// LoadProject allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (projectTaskL) LoadProject(ctx context.Context, e boil.ContextExecutor, singular bool, maybeProjectTask interface{}, mods queries.Applicator) error {
	var slice []*ProjectTask
	var object *ProjectTask

	if singular {
		object = maybeProjectTask.(*ProjectTask)
	} else {
		slice = *maybeProjectTask.(*[]*ProjectTask)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &projectTaskR{}
		}
		args = append(args, object.ProjectID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &projectTaskR{}
			}

			for _, a := range args {
				if a == obj.ProjectID {
					continue Outer
				}
			}

			args = append(args, obj.ProjectID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`projects`),
		qm.WhereIn(`projects.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Project")
	}

	var resultSlice []*Project
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Project")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for projects")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for projects")
	}

	if len(projectTaskAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Project = foreign
		if foreign.R == nil {
			foreign.R = &projectR{}
		}
		foreign.R.ProjectTasks = append(foreign.R.ProjectTasks, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ProjectID == foreign.ID {
				local.R.Project = foreign
				if foreign.R == nil {
					foreign.R = &projectR{}
				}
				foreign.R.ProjectTasks = append(foreign.R.ProjectTasks, local)
				break
			}
		}
	}

	return nil
}

// LoadTask allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (projectTaskL) LoadTask(ctx context.Context, e boil.ContextExecutor, singular bool, maybeProjectTask interface{}, mods queries.Applicator) error {
	var slice []*ProjectTask
	var object *ProjectTask

	if singular {
		object = maybeProjectTask.(*ProjectTask)
	} else {
		slice = *maybeProjectTask.(*[]*ProjectTask)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &projectTaskR{}
		}
		args = append(args, object.TaskID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &projectTaskR{}
			}

			for _, a := range args {
				if a == obj.TaskID {
					continue Outer
				}
			}

			args = append(args, obj.TaskID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`tasks`),
		qm.WhereIn(`tasks.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Task")
	}

	var resultSlice []*Task
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Task")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for tasks")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for tasks")
	}

	if len(projectTaskAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Task = foreign
		if foreign.R == nil {
			foreign.R = &taskR{}
		}
		foreign.R.ProjectTasks = append(foreign.R.ProjectTasks, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.TaskID == foreign.ID {
				local.R.Task = foreign
				if foreign.R == nil {
					foreign.R = &taskR{}
				}
				foreign.R.ProjectTasks = append(foreign.R.ProjectTasks, local)
				break
			}
		}
	}

	return nil
}

// SetProject of the projectTask to the related item.
// Sets o.R.Project to related.
// Adds o to related.R.ProjectTasks.
func (o *ProjectTask) SetProject(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Project) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"project_tasks\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"project_id"}),
		strmangle.WhereClause("\"", "\"", 2, projectTaskPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ProjectID, o.TaskID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.ProjectID = related.ID
	if o.R == nil {
		o.R = &projectTaskR{
			Project: related,
		}
	} else {
		o.R.Project = related
	}

	if related.R == nil {
		related.R = &projectR{
			ProjectTasks: ProjectTaskSlice{o},
		}
	} else {
		related.R.ProjectTasks = append(related.R.ProjectTasks, o)
	}

	return nil
}

// SetTask of the projectTask to the related item.
// Sets o.R.Task to related.
// Adds o to related.R.ProjectTasks.
func (o *ProjectTask) SetTask(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Task) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"project_tasks\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"task_id"}),
		strmangle.WhereClause("\"", "\"", 2, projectTaskPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ProjectID, o.TaskID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.TaskID = related.ID
	if o.R == nil {
		o.R = &projectTaskR{
			Task: related,
		}
	} else {
		o.R.Task = related
	}

	if related.R == nil {
		related.R = &taskR{
			ProjectTasks: ProjectTaskSlice{o},
		}
	} else {
		related.R.ProjectTasks = append(related.R.ProjectTasks, o)
	}

	return nil
}

// ProjectTasks retrieves all the records using an executor.
func ProjectTasks(mods ...qm.QueryMod) projectTaskQuery {
	mods = append(mods, qm.From("\"project_tasks\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"project_tasks\".*"})
	}

	return projectTaskQuery{q}
}

// FindProjectTask retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindProjectTask(ctx context.Context, exec boil.ContextExecutor, projectID int64, taskID int64, selectCols ...string) (*ProjectTask, error) {
	projectTaskObj := &ProjectTask{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"project_tasks\" where \"project_id\"=$1 AND \"task_id\"=$2", sel,
	)

	q := queries.Raw(query, projectID, taskID)

	err := q.Bind(ctx, exec, projectTaskObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from project_tasks")
	}

	if err = projectTaskObj.doAfterSelectHooks(ctx, exec); err != nil {
		return projectTaskObj, err
	}

	return projectTaskObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ProjectTask) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no project_tasks provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(projectTaskColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	projectTaskInsertCacheMut.RLock()
	cache, cached := projectTaskInsertCache[key]
	projectTaskInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			projectTaskAllColumns,
			projectTaskColumnsWithDefault,
			projectTaskColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(projectTaskType, projectTaskMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(projectTaskType, projectTaskMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"project_tasks\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"project_tasks\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into project_tasks")
	}

	if !cached {
		projectTaskInsertCacheMut.Lock()
		projectTaskInsertCache[key] = cache
		projectTaskInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the ProjectTask.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ProjectTask) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	projectTaskUpdateCacheMut.RLock()
	cache, cached := projectTaskUpdateCache[key]
	projectTaskUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			projectTaskAllColumns,
			projectTaskPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update project_tasks, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"project_tasks\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, projectTaskPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(projectTaskType, projectTaskMapping, append(wl, projectTaskPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update project_tasks row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for project_tasks")
	}

	if !cached {
		projectTaskUpdateCacheMut.Lock()
		projectTaskUpdateCache[key] = cache
		projectTaskUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q projectTaskQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for project_tasks")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for project_tasks")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ProjectTaskSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), projectTaskPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"project_tasks\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, projectTaskPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in projectTask slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all projectTask")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ProjectTask) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no project_tasks provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(projectTaskColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	projectTaskUpsertCacheMut.RLock()
	cache, cached := projectTaskUpsertCache[key]
	projectTaskUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			projectTaskAllColumns,
			projectTaskColumnsWithDefault,
			projectTaskColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			projectTaskAllColumns,
			projectTaskPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert project_tasks, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(projectTaskPrimaryKeyColumns))
			copy(conflict, projectTaskPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"project_tasks\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(projectTaskType, projectTaskMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(projectTaskType, projectTaskMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert project_tasks")
	}

	if !cached {
		projectTaskUpsertCacheMut.Lock()
		projectTaskUpsertCache[key] = cache
		projectTaskUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single ProjectTask record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ProjectTask) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ProjectTask provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), projectTaskPrimaryKeyMapping)
	sql := "DELETE FROM \"project_tasks\" WHERE \"project_id\"=$1 AND \"task_id\"=$2"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from project_tasks")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for project_tasks")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q projectTaskQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no projectTaskQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from project_tasks")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for project_tasks")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ProjectTaskSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(projectTaskBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), projectTaskPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"project_tasks\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, projectTaskPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from projectTask slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for project_tasks")
	}

	if len(projectTaskAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *ProjectTask) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindProjectTask(ctx, exec, o.ProjectID, o.TaskID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ProjectTaskSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ProjectTaskSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), projectTaskPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"project_tasks\".* FROM \"project_tasks\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, projectTaskPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ProjectTaskSlice")
	}

	*o = slice

	return nil
}

// ProjectTaskExists checks if the ProjectTask row exists.
func ProjectTaskExists(ctx context.Context, exec boil.ContextExecutor, projectID int64, taskID int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"project_tasks\" where \"project_id\"=$1 AND \"task_id\"=$2 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, projectID, taskID)
	}
	row := exec.QueryRowContext(ctx, sql, projectID, taskID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if project_tasks exists")
	}

	return exists, nil
}