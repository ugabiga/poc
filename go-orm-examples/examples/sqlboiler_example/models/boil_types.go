// Code generated by SQLBoiler 4.11.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"strconv"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/strmangle"
)

// M type is for providing columns and column values to UpdateAll.
type M map[string]interface{}

// ErrSyncFail occurs during insert when the record could not be retrieved in
// order to populate default value information. This usually happens when LastInsertId
// fails or there was a primary key configuration that was not resolvable.
var ErrSyncFail = errors.New("models: failed to synchronize data after insert")

type insertCache struct {
	query        string
	retQuery     string
	valueMapping []uint64
	retMapping   []uint64
}

type updateCache struct {
	query        string
	valueMapping []uint64
}

func makeCacheKey(cols boil.Columns, nzDefaults []string) string {
	buf := strmangle.GetBuffer()

	buf.WriteString(strconv.Itoa(cols.Kind))
	for _, w := range cols.Cols {
		buf.WriteString(w)
	}

	if len(nzDefaults) != 0 {
		buf.WriteByte('.')
	}
	for _, nz := range nzDefaults {
		buf.WriteString(nz)
	}

	str := buf.String()
	strmangle.PutBuffer(buf)
	return str
}

type ProjectStatus string

// Enum values for ProjectStatus
const (
	ProjectStatusTodo       ProjectStatus = "todo"
	ProjectStatusInProgress ProjectStatus = "in_progress"
	ProjectStatusDone       ProjectStatus = "done"
)

func AllProjectStatus() []ProjectStatus {
	return []ProjectStatus{
		ProjectStatusTodo,
		ProjectStatusInProgress,
		ProjectStatusDone,
	}
}

func (e ProjectStatus) IsValid() error {
	switch e {
	case ProjectStatusTodo, ProjectStatusInProgress, ProjectStatusDone:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e ProjectStatus) String() string {
	return string(e)
}

type TaskStatus string

// Enum values for TaskStatus
const (
	TaskStatusTodo       TaskStatus = "todo"
	TaskStatusInProgress TaskStatus = "in_progress"
	TaskStatusDone       TaskStatus = "done"
)

func AllTaskStatus() []TaskStatus {
	return []TaskStatus{
		TaskStatusTodo,
		TaskStatusInProgress,
		TaskStatusDone,
	}
}

func (e TaskStatus) IsValid() error {
	switch e {
	case TaskStatusTodo, TaskStatusInProgress, TaskStatusDone:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e TaskStatus) String() string {
	return string(e)
}
