/*
Special interface that I made to make system more reusable.
So if you need in future another db it should satisfy this interface.
Generic T represents field of database.
In this approach DB must be open everytime.

Why interface? For example you create a function that executes something on database
and you need to execute this command automaticly for any DB passed which contains service.Post.
So you will define it like this
func exec_fix_for_db(repo service.Repository[Post]) {
	repo.ExecSpecific("Very important fix!!")
}
*/

package pkg

import "database/sql"

type Repository[T any] interface {
	GetAllValues() ([]T, error)
	GetValueByID(id int) (T, error)
	InsertValue(T) (id int, err error)
	DeleteValueByID(id int) error
	UpdateValue(id int, object T) error
	ExecSpecific(SQL_command string) (sql.Result, error)
	QuerySpecific(SQL_command string) (*sql.Rows, error)
}
