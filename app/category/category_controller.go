package category

import (
	"github.com/beego/beego/v2/server/web/context"
	"zhiqu/infrastructure/controller"
	"zhiqu/infrastructure/response"
)

// Controller operations for Category
type Controller struct {
	controller.BaseController
	service *Service
}

func NewController(service *Service) *Controller {
	return &Controller{service: service}
}

// Post ...
// @Title Post
// @Description create Category
// @Param	body		body 	models.Category	true		"body for Category content"
// @Success 201 {int} models.Category
// @Failure 403 body is empty
// @router / [post]
func (c *Controller) Post(context *context.Context) {
	//var v models.Category
	//if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
	//	if _, err := models.AddCategory(&v); err == nil {
	//		c.Ctx.Output.SetStatus(201)
	//		c.Data["xjson"] = v
	//	} else {
	//		c.Data["xjson"] = err.Error()
	//	}
	//} else {
	//	c.Data["xjson"] = err.Error()
	//}
	//c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Category by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Category
// @Failure 403 :id is empty
// @router /:id [get]
func (c *Controller) GetOne(context *context.Context) {
}
func (c *Controller) Get(ctx *context.Context) {
	c.WriteJson(response.Ok("消息正常"))
}

// GetAll ...
// @Title Get All
// @Description get Category
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Category
// @Failure 403
// @router / [get]
func (c *Controller) GetAll(context *context.Context) {
}

// Put ...
// @Title Put
// @Description update the Category
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Category	true		"body for Category content"
// @Success 200 {object} models.Category
// @Failure 403 :id is not int
// @router /:id [put]
func (c *Controller) Put(context *context.Context) {
	//s
}

// Delete ...
// @Title Delete
// @Description delete the Category
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *Controller) Delete(context *context.Context) {
	//idStr := c.Ctx.Input.Param(":id")
	//id, _ := strconv.Atoi(idStr)
	//if err := models.DeleteCategory(id); err == nil {
	//	c.Data["xjson"] = "OK"
	//} else {
	//	c.Data["xjson"] = err.Error()
	//}
	//c.ServeJSON()
}
