package restapi
import (
	"github.com/gin-gonic/gin"
)
//InitializeRoutes for the API
func InitializeRoutes(r *gin.RouterGroup) {
	
	//books endpoint
	r.GET("/books", getAllBooks)
r.GET("/books/:id", getBook)
r.POST("/books", addNewBook)
r.PUT("/books/:id", updateBook)
r.DELETE("/books/:id", deleteBook)
	// users endpoint
	r.GET("/users", getAllUsers)
	r.GET("/users/:id", getUser)
	r.POST("/users", addNewUser)
	r.PUT("/users/:id", updateUser)
	r.DELETE("/users/:id", deleteUser)


	r.GET("/categories", getAllCategories)
	r.GET("/categories/:id", getCategory)
	r.POST("/categories", addNewCategory)
	r.PUT("/categories/:id", updateCategory)
	r.DELETE("/categories/:id", deleteCategory)
}