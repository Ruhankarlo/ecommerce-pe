package main

import(
	"github.com/gin-gonic/gin"
	"net/http"
	"ecommerce-pe/internal/controllers"
	"ecommerce-pe/internal/database"
	"ecommerce-pe/internal/middleware"
	"ecommerce-pe/internal/routes"
	"log"
	"os"
)

func main() {
	//Será obtido o valor de PORT da varivável de ambente, se tiver.
	port := os.Getenv("PORT");
	//Define o valor como 8080, caso não existir a variável de ambiente citada acima.
	if port==""{
		port="8000";
	}

	//É instanciado a aplicação com acesso aos dados
	//Acesso a tebelas de produtos e as tabelas de usuários no banco
	app := controllers.NewApplication(
		database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users")
		);

	// Instanciamos um router do gin, mais simples que o padrão, para utilizar somente o que precisasmos.
	router := gin.New();
	//Esse router irá utilizar o logger(middleware) que registra os detalhes da requisição
	router.Use(gin.Logger());
  	//É instanciado a função UserRoute e atribuído ao router princial, que está como parametro, as rotas de entradas que foram configuradas.
	routes.UserRoutes(router);
	//Após ter definido as rotas publicas anterior, é adicionado outro middleware (global), esse será o de autenticação.
	//Para verificar o token e retornará as rotas
	router.Use(middleware.Authentication());
	//Se a autenticação do token estiver válida, obterá as seguintes rotas com funções.
	router.GET("/addtocart", app.AddToCart()); //adicionar item ao carrinho
	router.GET("/removeitem", app.RemoveItem()); //remover item do carrinho
	router.GET("/cartcheckout", app.BuyFromCart()); // finalizar as compras do carrinho
	router.GET("/instantbuy", app.InstantBuy()); // comprar um item apenas (sem carrinho)

	
	// inicio o servidor na porta que foi configurada
	// é utilizado o Fatal para caputurar erro do servidor e logar
	log.Fatal(router.Run(":", port));


}