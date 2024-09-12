package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UserController ...
var (
	UserController UserControllerInterface = &userController{}
)

type UserControllerInterface interface {
	Create(c *gin.Context)
	GetOne(c *gin.Context)
	GetAll(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type userController struct {
	service UserServiceInterface
}

func NewUserController(ser UserServiceInterface) UserControllerInterface {
	return &userController{
		ser,
	}
}

// ///////controllers/////////////////
// Create godoc
// @Summary Create a user
// @Description Create a new user item
// @Tags users
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param name formData string true "User's Name" example:"John"
// @Param email formData string true "User's Email" example:"john@example.com"
// @Param age formData string true "User's age" example:23
// @Success 201 {object} User
// @Failure 400 {object} error
// @Router /register [post]
func (controller userController) Create(c *gin.Context) {
	user := &User{}
	user.Name = c.PostForm("name")
	user.Email = c.PostForm("email")
	age, err := strconv.ParseInt(c.PostForm("age"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed to parse age"})
		return
	}
	user.Age = int(age)
	u, err1 := controller.service.Create(user)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err1.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": u})
}

// GetOne godoc
// @Summary GetOne a user
// @Description GetOne a new user item
// @Tags users
// @Accept json
// @Produce json
// @Param   id     path    string     true        "id"
// @Success 200 {object} User
// @Failure 400 {error} error
// @Router /api/users/{id} [get]
func (controller userController) GetOne(c *gin.Context) {
	code := c.Param("id")
	user, problem := controller.service.GetOne(code)
	if problem != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": problem.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

// GetAll godoc
// @Summary GetAll a user
// @Description GetAll a new user item
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} User
// @Failure 400 {error} error
// @Router /api/users [get]
func (controller userController) GetAll(c *gin.Context) {
	users, problem := controller.service.GetAll()
	if problem != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": problem.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}

// Update godoc
// @Summary Update a user
// @Description Update a new user item
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {string} Updated successifuly
// @Failure 400 {error} error
// @Router /api/users/{code} [put]
func (controller userController) Update(c *gin.Context) {
	user := &User{}
	user.Name = c.PostForm("name")
	user.Email = c.PostForm("email")
	age, err := strconv.ParseInt(c.PostForm("age"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed to parse age"})
		return
	}
	user.Age = int(age)
	code := c.Param("id")
	_, problem := controller.service.Update(code, user)
	if problem != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": problem.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Updated successifuly"})
}

// Delete godoc
// @Summary Delete a user
// @Description Delete a new user item
// @Tags users
// @Accept json
// @Produce json
// @Param   code     path    string     true        "code"
// @Success 200 {string} deleted successfully
// @Failure 400 {error} error
// @Router /api/users/{id} [delete]
func (controller userController) Delete(c *gin.Context) {
	id := c.Param("id")
	success, failure := controller.service.Delete(id)
	if failure != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": failure.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": success})

}
