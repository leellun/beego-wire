package user

import (
	"fmt"
	"zhiqu/infrastructure/controller"
)

// Controller operations for User
type Controller struct {
	controller.BaseController
	service *Service
}

func NewController(service *Service) *Controller {
	a := &Controller{service: service}
	fmt.Printf("====%v\n", a)
	return a
}

// Post ...
// @Title Post
// @Description create User
// @Param	body		body 	model.User	true		"body for User content"
// @Success 201 {int} model.User
// @Failure 403 body is empty
// @router / [post]
func (c *Controller) Post() {
}

// GetOne ...
// @Title Get One
// @Description get User by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} model.User
// @Failure 403 :id is empty
// @router /:id [get]
func (c *Controller) GetOne() {
	//fmt.Printf("====%v\n", c)
	//idStr := c.Ctx.Input.Param(":id")
	//id, _ := strconv.Atoi(idStr)
	//v := c.service.GetUserById(id)
	//c.WriteJson(v)
	//c.WriteJson(response.Error(""))
}

// GetAll ...
// @Title Get All
// @Description get User
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} model.User
// @Failure 403
// @router / [get]
func (c *Controller) GetAll() {
}

// Put ...
// @Title Put
// @Description update the User
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	model.User	true		"body for User content"
// @Success 200 {object} model.User
// @Failure 403 :id is not int
// @router /:id [put]
func (c *Controller) Put() {
}

// Delete ...
// @Title Delete
// @Description delete the User
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *Controller) Delete() {
}
