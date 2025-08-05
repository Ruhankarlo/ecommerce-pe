package routes

import (
	"ecommerce-pe/internal/controllers"
	"github.com/gin-gonic/gin"
)

// função UserRoute parametrizando o gin.Engine, para configurar várias rotas específica para usuários
func UserRoutes(incomingRoutes *gin.Engine){
	//Com o ponteiro solicitando o método, é parametrizado as URL e as funções controllers que serão 
	//processadas caso houver uma requisição nessas rotas de entrada.
	incomingRoutes.POST("/users/signup", controllers.SingUp());
	incomingRoutes.POST("/users/login", controllers.Login());
	incomingRoutes.POST("/admin/addproduct", controllers.ProductViewerAdmin());
	incomingRoutes.GET("/users/productview", controller.SearchProduct());
	incomingRoutes.GET("/users/search", controller.SearchProductByQuery());
}