package server 

import (
	"github.com/GO-000/week02/service"
)

func GetCommentDetailRequest(req Request) {
	cmt, err = service.GetCommentById(req.id)

	if err != nil {
		return http 404 
	}
	....
}