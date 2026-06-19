package view

import (
	"github.com/cassianobraz/crudGo/src/controller/model/response"
	"github.com/cassianobraz/crudGo/src/model"
)

func ConvertDomainToResponse(
	UseDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		Id:    UseDomain.GetID(),
		Email: UseDomain.GetEmail(),
		Name:  UseDomain.GetName(),
		Age:   UseDomain.GetAge(),
	}
}
