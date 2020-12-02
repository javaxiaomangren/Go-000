package service

import (
	"github.com/GO-000/week02/model"
	"github.com/GO-000/week02/dao"
	"github.com/pkg/errors"
)


func GetCommentById(id int) (comment model.Comment, e error) {
	// Todo connect 
	raw, err := dao.GetCommentById(id)
	if err != nil {
		return nil, err
	}

	comment := raw to Comment model 
	
	return 
}