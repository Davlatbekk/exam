package controller

import (
	"app/models"
	"errors"
)

func (c *Controller) AddShopCart(req *models.Add) (string, error) {
	_, err := c.store.User().GetByID(&models.UserPrimaryKey{Id: req.UserId})
	if err != nil {
		return "", err
	}

	_, err = c.store.Product().GetByID(&models.ProductPrimaryKey{Id: req.ProductId})
	if err != nil {
		return "", err
	}

	id, err := c.store.ShopCart().AddShopCart(req)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (c *Controller) RemoveShopCart(req *models.Remove) error {
	err := c.store.ShopCart().RemoveShopCart(req)
	if err != nil {
		return err
	}
	return err
}

func (c *Controller) CalculateTotal(req *models.UserPrimaryKey, status string, discount float64) (float64, error) {
	_, err := c.store.User().GetByID(req)
	if err != nil {
		return 0, err
	}

	users, err := c.store.ShopCart().GetUserShopCart(req)
	if err != nil {
		return 0, err
	}

	var total float64
	for _, v := range users {
		product, err := c.store.Product().GetByID(&models.ProductPrimaryKey{Id: v.ProductId})
		if err != nil {
			return 0, err
		}
		if status == "fixed" {
			total += float64(v.Count) * (product.Price - discount)
		} else if status == "percent" {
			if discount < 0 || discount > 100 {
				return 0, errors.New("Invalid discount range")
			}
			total += float64(v.Count) * (product.Price - (product.Price*discount)/100)
		} else {
			return 0, errors.New("Invalid status name")
		}
	}

	if total < 0 {
		return 0, nil
	}
	return total, nil
}

// func (c *Controller) Filter(from, to string) (models.ShopCart, error) {
// 	shopCart, err := c.store.ShopCart().GetAll()
// 	if err != nil {
// 		return models.ShopCart{}, err
// 	}

// 	if from == "" || to == "" {

// 		type kv struct {
// 			Key   int
// 			Value string
// 		}

// 		var top []kv

// 		for k, v := range shopCart.Products {
// 			top = append(top, kv{k, v.Time})
// 		}

// 		sort.Slice(top, func(i, j int) bool {
// 			return top[i].Value > top[j].Value
// 		})

// 		fmt.Println(top)

// 	} // fmt.Println(products.Products)
// 	res := models.ShopCart{}
// 	for _, val := range shopCart.Products {

// 		if from < val.Time && val.Time <= to {

// 			fmt.Println(val)
// 			// return val, nil
// 			res = val

// 		}
// 	}
// 	return res, nil

// }

// func (c *Controller) UserHistory(req *models.UserPrimaryKey) (string, []models.History) {

// 	statis := []models.History{}
// 	shopcart, err := c.store.ShopCart().GetAll()
// 	if err != nil {
// 		return "", []models.History{}
// 	}

// 	User, err := c.store.User().GetByID(&models.UserPrimaryKey{Id: req.Id})
// 	if err != nil {
// 		return "", []models.History{}
// 	}

// 	name := User.Name

// 	for _, shopcart := range shopcart.Products {
// 		if shopcart.Status == true && shopcart.UserId == req.Id {
// 			// fmt.Println("ok")

// 			product, err := c.store.Product().GetByID(&models.ProductPrimaryKey{shopcart.ProductId})
// 			if err != nil {
// 				return "", []models.History{}
// 			}

// 			statis = append(statis, models.History{
// 				Name:  product.Name,
// 				Price: product.Price,
// 				Count: shopcart.Count,
// 				Total: int(product.Price) * shopcart.Count,
// 				Time:  shopcart.Time,
// 			})

// 		}

// 	}

// 	return name, statis

// }

// func (c *Controller) AllMoney(req *models.UserPrimaryKey) (string, int) {

// 	name, val := c.UserHistory(req)

// 	allmoney := 0
// 	for _, v := range val {
// 		allmoney += v.Total
// 	}

// 	return name, allmoney

// }

// func (c *Controller) StatistikaProduct() (map[string]int, error) {
// 	statis := map[string]int{}
// 	shopcart, err := c.store.ShopCart().GetAll()
// 	if err != nil {
// 		return statis, err
// 	}

// 	// fmt.Println(products)
// 	for _, shopcart := range shopcart.Products {
// 		if shopcart.Status == true {

// 			product, err := c.store.Product().GetByID(&models.ProductPrimaryKey{shopcart.ProductId})
// 			if err != nil {
// 				return statis, err
// 			}
// 			statis[product.Name] += shopcart.Count
// 		}
// 	}

// 	return statis, nil
// }

// func (c *Controller) Top10() error {
// 	product, err := c.StatistikaProduct()

// 	if err != nil {
// 		return err
// 	}

// 	type kv struct {
// 		Key   string
// 		Value int
// 	}

// 	var top []kv

// 	for k, v := range product {
// 		top = append(top, kv{k, v})
// 	}

// 	sort.Slice(top, func(i, j int) bool {
// 		return top[i].Value > top[j].Value
// 	})

// 	for i := 0; i < 10; i++ {

// 		fmt.Println(top[i])

// 	}

// 	return nil
// }

// func (c *Controller) Last10() error {
// 	product, err := c.StatistikaProduct()

// 	if err != nil {
// 		return err
// 	}

// 	type kv struct {
// 		Key   string
// 		Value int
// 	}

// 	var top []kv

// 	for k, v := range product {
// 		top = append(top, kv{k, v})
// 	}

// 	sort.Slice(top, func(i, j int) bool {
// 		return top[i].Value < top[j].Value
// 	})

// 	for i := 0; i < 10; i++ {

// 		fmt.Println(top[i])

// 	}

// 	return nil
// }

// func (c *Controller) DaysTopProduct() error {
// 	prod := map[string]int{}
// 	shopcart, err := c.store.ShopCart().GetAll()
// 	if err != nil {
// 		return err
// 	}

// 	for _, element := range shopcart.Products {

// 		layout := "2006-01-02 15:04:05"
// 		date, err := time.Parse(layout, element.Time)
// 		if err != nil {
// 			fmt.Println(err)
// 			return err
// 		}
// 		formatted := date.Format("2006-01-02")
// 		// fmt.Println(formatted)

// 		prod[formatted] = prod[formatted] + element.Count
// 	}

// 	type kv struct {
// 		Key   string
// 		Value int
// 	}

// 	var top []kv

// 	for k, v := range prod {
// 		top = append(top, kv{k, v})
// 	}

// 	sort.Slice(top, func(i, j int) bool {
// 		return top[i].Value > top[j].Value
// 	})
// 	for i, val := range top {
// 		fmt.Printf(" %d Sana: %s count: %d\n", i+1, val.Key, val.Value)

// 	}

// 	return nil
// }

// func (c *Controller) Statistika() (map[string]int, error) {
// 	statis := map[string]int{}
// 	shopcart, err := c.store.ShopCart().GetAll()
// 	if err != nil {
// 		return statis, err
// 	}

// 	// fmt.Println(products)
// 	for _, shopcart := range shopcart.Products {
// 		if shopcart.Status == true {

// 			product, err := c.store.Product().GetByID(&models.ProductPrimaryKey{shopcart.ProductId})
// 			if err != nil {
// 				return statis, err
// 			}

// 			category, err := c.store.Category().GetAll(&models.GetListCategoryRequest{0, 11})

// 			if err != nil {
// 				return statis, err
// 			}
// 			for _, val := range category.Categories {
// 				if product.CategoryID == val.Id {
// 					statis[val.Name] += shopcart.Count
// 				}
// 			}

// 		}
// 	}

// 	return statis, nil
// }

// func (c *Controller) ActiveUSer() (string, error) {
// 	statis := map[string]int{}
// 	shopcart, err := c.store.ShopCart().GetAll()
// 	if err != nil {
// 		return "", err
// 	}

// 	users, err := c.store.User().GetAll(&models.GetListRequest{0, 13})

// 	if err != nil {
// 		return "", err
// 	}

// 	for _, shopcart := range shopcart.Products {
// 		if shopcart.Status == true {

// 			for _, val := range users.Users {
// 				if shopcart.UserId == val.Id {
// 					statis[val.Name] += shopcart.Count
// 				}
// 			}

// 		}
// 	}

// 	type kv struct {
// 		Key   string
// 		Value int
// 	}

// 	var top []kv

// 	for k, v := range statis {
// 		top = append(top, kv{k, v})
// 	}

// 	sort.Slice(top, func(i, j int) bool {
// 		return top[i].Value > top[j].Value
// 	})

// 	// fmt.Println(ss)

// 	for i := 0; i < len(top); i++ {

// 		fmt.Println(top[i])

// 	}

// 	return "", err

// }

// func (c *Controller) Skidka(req *models.UserPrimaryKey) (string, []models.History) {

// 	// shopcart, err := c.store.ShopCart().GetAll()
// 	// if err != nil {
// 	// 	return "", []models.History{}
// 	// }

// 	statis := []models.History{}

// 	// username := ""
// 	// for _, shopcart := range shopcart.Products {

// 	name, userhistory := c.UserHistory(&models.UserPrimaryKey{Id: req.Id})

// 	// fmt.Println(userhistory)

// 	// username = name
// 	for _, val := range userhistory {

// 		if val.Count > 9 {

// 			statis = append(statis, models.History{

// 				Name:  val.Name,
// 				Price: val.Price,
// 				Count: val.Count,
// 				Total: int(val.Price)*val.Count - int(val.Price),
// 				Time:  val.Time,
// 			})

// 		}
// 	}
// 	// }

// 	return name, statis

// }
