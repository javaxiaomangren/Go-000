package dao

import (
	"github.com/GO-000/week02/model"
	"github.com/pkg/errors"
)

type Connection struct {
	....
}

var conn Connection 
....

func GetCommentById(id int) (comment model.Comment, e error) {
	// Todo connect 
	raw, err := conn.DoQuery("select * from comment where id = %d", id)
	if err != nil {
		return nil, errors.Wrapf(err, "query coment failed")
	}

	comment := raw to Comment model 
	
	return 
}