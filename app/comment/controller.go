package comment

// CommentController operations for Comment
type CommentController struct {
}

// Post ...
// @Title Post
// @Description create Comment
// @Param	body		body 	models.Comment	true		"body for Comment content"
// @Success 201 {int} models.Comment
// @Failure 403 body is empty
// @router / [post]
func (c *CommentController) Post() {
}

// GetOne ...
// @Title Get One
// @Description get Comment by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Comment
// @Failure 403 :id is empty
// @router /:id [get]
func (c *CommentController) GetOne() {
}

// GetAll ...
// @Title Get All
// @Description get Comment
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Comment
// @Failure 403
// @router / [get]
func (c *CommentController) GetAll() {
}

// Put ...
// @Title Put
// @Description update the Comment
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Comment	true		"body for Comment content"
// @Success 200 {object} models.Comment
// @Failure 403 :id is not int
// @router /:id [put]
func (c *CommentController) Put() {
}

// Delete ...
// @Title Delete
// @Description delete the Comment
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *CommentController) Delete() {
}
