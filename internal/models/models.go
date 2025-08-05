package models

import(
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

//Para o cadastro do usuário na plataforma
type User struct{
	//É criado um ID do tipo especial do driver do mongoDB, pois garante formatação, rapidez, e há métodos incorporado nesse tipo(como transformar em string)
	//Formatado em bson, para garantir mapeamento os campos do banco (como é utilizado o Binary JSON no mongoDB)
	ID				primitive.ObjectID		`json:"_id bson:"_id"`
	//Foi utlizado o tipo de ponteiro para string para diferenciar entre espaço vazio e nulo.
	//A adição da tag validate garantirá que os dados da requisição atenda os parâmetros
	//como o require(é obrigatório retorno), e min e max como quantidade de caracter
	First_Name		*string					`json:"first_name" validate="required,min=2,max=20"`
	Last_Name		*string					`json:"last_name" validate="requiere,min=2,max=30"`
	Password		*string					`json:"password" validate="required,min=2"`
	Email			*string					`json:"email" validate="email, required"` 
	Phone			*string					`json:"phone" validate="required"` 
	Token			*string					`json:"token"`
	Refresh_Token	*string					`json:"refresh_token"`
	Created_At		time.Time				`json:"created_at"`
	Updated_At		time.Time				`json:"updated_at"`
	User_ID			string					`json:"user_id"`
	//A struct usuário terá um relacionamento com essas outras estruturas (composição).
	UserCart		[]ProductUser			`json:"usercart" bson:"usercart"`
	Address_Details	[]Address				`json:"address_details" bson:"address_details"`
	Order_Status	[]Order					`json:"order_status" bson:"order_status"` 
}
//Todos produtos disponíveis na plataforma
type Product struct{
	Product_ID		primitive.ObjectID		`bson:"_id"`
	Product_Name	*string					`json:"product_name"`
	//Definido uint64 para garantir valores positivos e mais longos
	Price			*uint64					`json:"price"`
	Rating			*uint8					`json:"rating"`
	Image			*string					`json:"image"`
}
//Produtos que estão sendo comprados pelo usuario
type ProductUser struct{
	Product_ID		primitive.ObjectID		`bson:"_id"`
	Product_Name	*string					`json:"product_name bson:"product_name`
	//Como o carrinho sempre terá valor, não foi preciso adicionar o tipo como ponteiro para int, e é util caso haja mudança de preço
	Price			int						`json:"price" bson"price"`
	Rating			*uint					`json:"rating" bson:"rating"`
	Image			*string					`json:"image" bson:"image"`

}
// Local de entrega 
type Address struct {
	Address_ID		primitive.ObjectID		`bson:"_id"`
	Number			*string					`json:"number" bson:"numer"`
	Street			*string					`json:"street" bson:"number`
	City			*string					`json:"street bson:"street"`
	CEP				*string					`json:"cep" bson:"cep"`
}
//Coompras finalizadas(pedido);
type Order struct {
	Order_ID		primitive.ObjectID		`bson:"_id"`
	//Obtem a lista de produtos do comprador em carrinho
	Order_Cart		[]ProductUser			`json:"order_cart" bson:"order_cart`
	Ordered_At		time.Time				`json:"ordered_at" bson:"ordered_at`
	Price			int						`json:"price_total" bson:"price_total"`
	Discount		*int					`json:"discount" bson:"discount"`
	Payment_Method	Payment					`json:"payment_method" bson:"payment_method`
}
//método de pagamento escolhido
type Payment struct {
	//Digital para pagamentos digitais como cartão, pix.
	Digital bool
	// "Cash on Delivery" para pagamentos quando for entregar o pedido
	COD		bool


}