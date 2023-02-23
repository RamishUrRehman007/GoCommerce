package routes

import (
	"Ramish_GO_Fiber/database"
	"Ramish_GO_Fiber/dto"
	"math"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetDeliveredAmount(order_id int32) float64 {
	var delivered_amount float64
	sum_count := []map[string]interface{}{}
	database.Database.DB.Table("deliveries").Select("sum(deliveries.delivered_quantity)*order_items.price_per_unit as delivered_amount").Joins("JOIN order_items on order_items.id=deliveries.order_item_id").Where("order_items.order_id=?", order_id).Group("order_items.price_per_unit").Find(&sum_count)
	for _, sm_ct := range sum_count {
		delivered_amount = delivered_amount + sm_ct["delivered_amount"].(float64)
	}

	return delivered_amount
}

func GetOrders(c *fiber.Ctx) error {
	results := []map[string]interface{}{}

	database.Database.DB.Table("order_items").Select("sum(order_items.quantity*order_items.price_per_unit) as total_amount, orders.order_name, customer_companies.company_name, customers.name, orders.created_at, orders.id").Joins("JOIN orders on orders.id=order_items.order_id").Joins("JOIN customers on customers.user_id=orders.customer_id").Joins("JOIN customer_companies on customer_companies.id=customers.company_id").Group("order_items.order_id, orders.order_name, customer_companies.company_name, customers.name, orders.created_at, orders.id").Find(&results)

	orders := []dto.Orders{}
	for _, item := range results {
		order := dto.Orders{
			OrderName:       item["order_name"].(string),
			CustomerCompany: item["company_name"].(string),
			CustomerName:    item["name"].(string),
			OrderDate:       item["created_at"].(time.Time),
			DeliveredAmount: math.Floor(GetDeliveredAmount(item["id"].(int32))*100) / 100,
			TotalAmount:     math.Floor(item["total_amount"].(float64)*100) / 100,
		}
		orders = append(orders, order)
	}

	return c.Status(200).JSON(orders)
}
