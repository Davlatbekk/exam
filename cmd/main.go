package main

import (
	"app/config"
	"app/controller"
	"app/storage/jsonDb"
	"log"
)

func main() {
	cfg := config.Load()

	jsonDb, err := jsonDb.NewFileJson(&cfg)
	if err != nil {
		log.Fatal("error while connecting to database")
	}
	defer jsonDb.CloseDb()

	c := controller.NewController(&cfg, jsonDb)

	// name, res := c.Skidka(&models.UserPrimaryKey{Id: "27457ac2-74dd-4656-b9b0-0d46b1af10dc"})

	// fmt.Println(name, res)

	// res, err := c.ActiveUSer()

	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// fmt.Println(res)

	// res, err := c.Statistika()

	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// fmt.Println(res)

	err = c.DaysTopProduct()

	if err != nil {
		log.Println(err)
		return
	}

	// err = c.Last10()

	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// err = c.Top10()

	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// res, err := c.StatistikaProduct()

	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// fmt.Println(res)

	// name, allmoney := c.AllMoney(&models.UserPrimaryKey{Id: "0c7e40db-9948-4349-aade-a8378862de9c"})

	// fmt.Println(name, allmoney)

	// name, history := c.UserHistory(&models.UserPrimaryKey{Id: "0c7e40db-9948-4349-aade-a8378862de9c"})

	// fmt.Println(name, history)

	// res, err := c.Filter("2022-08-15", "")

	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// fmt.Println(res)

	// err = c.DeleteOrder(&models.OrderPrimaryKey{Id: "25613c77-247d-44f2-a316-3c96359f210e"})
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// fmt.Println(order)

	// c.CreateOrder(&models.CreateOrder{
	// 	User_id:          "657a41b6-1bdc-47cc-bdad-1f85eb8fb98c",
	// 	Cutomer_name:     "DAvlatbek",
	// 	Customer_address: "UZB",
	// 	Customer_phone:   "99 082 84 83",
	// 	OrderItems: []models.CreateOrderItems{

	// 		{
	// 			Product_id: "38292285-4c27-497b-bc5f-dfe418a9f959",
	// 			Count:      4,
	// 		},
	// 	},
	// })

	// c.UpdateOrder(&models.UpdateOrder{
	// 	Cutomer_name:     "Abdu",
	// 	Customer_address: "UZB",
	// 	Customer_phone:   "99 008 74 33",
	// },"c14e6a91-aff4-44b8-ad64-4a85b5bf243b")

	// c.CreateProduct(&models.CreateProduct{
	// 	Name:       "Smartfon vivo V25 8/256 GB",
	// 	Price:      4_860_000,
	// 	CategoryID: "6325b81f-9a2b-48ef-8d38-5cef642fed6b",
	// })

	// product, err := c.GetByIdProduct(&models.ProductPrimaryKey{Id: "38292285-4c27-497b-bc5f-dfe418a9f959"})

	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// c.GetAllProduct(
	// 	offset:
	// 	limit:
	// 	categoryid: "38292285-4c27-497b-bc5f-dfe418a9f959"
	// )

	// fmt.Printf("%+v\n", product)

	// fmt.Println(c.GetAllProduct(&models.GetListRequestProduct{
	// 	Offset:     0,
	// 	Limit:      1,
	// 	CategoryID: "6325b81f-9a2b-48ef-8d38-5cef642fed6b",
	// }))

}

// func Category(c *controller.Controller) {
// 	// c.CreateCategory(&models.CreateCategory{
// 	// 	Name:     "Smartfonlar va telefonlar",
// 	// 	ParentID: "eed2e676-1f17-429f-b75c-899eda296e65",
// 	// })

// 	category, err := c.GetByIdCategory(&models.CategoryPrimaryKey{Id: "eed2e676-1f17-429f-b75c-899eda296e65"})
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	fmt.Println(category)

// }

// func User(c *controller.Controller) {

// 	sender := "bbda487b-1c0f-4c93-b17f-47b8570adfa6"
// 	receiver := "657a41b6-1bdc-47cc-bdad-1f85eb8fb98c"
// 	err := c.MoneyTransfer(sender, receiver, 500_000)
// 	if err != nil {
// 		log.Println(err)
// 	}
// }
