package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gonextjs/database"
	"gonextjs/models"
	"gonextjs/util"
	"math"
	"strconv"
)

type ProductsList struct {
	Products    []models.Product
	FirstPage   int
	CurrentPage int
	LastPage    int
}

func GetAllProducts(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize", "10"))
	switch {
	case pageSize > 20:
		pageSize = 20
	case pageSize <= 0:
		pageSize = 10
	}
	offset := (page - 1) * pageSize

	var products ProductsList
	var totalRows int64
	var totalPage int

	database.DBConn.Model(models.Product{}).Count(&totalRows)
	totalPage = int(math.Ceil(float64(totalRows) / float64(pageSize)))

	products.FirstPage = 1
	products.CurrentPage = page
	products.LastPage = totalPage

	database.DBConn.Offset(offset).Limit(pageSize).Preload("User").Preload("Categories").Preload("Store").Preload("Currency").Find(&products.Products)
	return c.JSON(products)
}

func GetProduct(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	var product models.Product
	product.ID = uint(id)
	database.DBConn.Preload("User").Preload("Categories").Preload("Store").Preload("Currency").Find(&product)
	return c.JSON(product)
}

func UpdateProduct(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	product := models.Product{}
	if err := c.BodyParser(&product); err != nil {
		return err
	}
	product.ID = uint(id)
	database.DBConn.Model(&product).Updates(&product)
	return c.JSON(product)
}

func DeleteProduct(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	var product models.Product
	product.ID = uint(id)
	//database.DBConn.Model(&product).Delete(&product)
	database.DBConn.Delete(&product)
	return nil
}

func AddProducts(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	if data["title"] == "" {
		return c.Status(500).JSON(fiber.Map{
			"message": "title is required!",
		})
	}
	if data["description"] == "" {
		return c.Status(500).JSON(fiber.Map{
			"message": "description is required!",
		})
	}
	if data["weight"] == "" {
		return c.Status(500).JSON(fiber.Map{
			"message": "weight is required!",
		})
	}
	if data["length"] == "" {
		return c.Status(500).JSON(fiber.Map{
			"message": "length is required!",
		})
	}
	if data["width"] == "" {
		return c.Status(500).JSON(fiber.Map{
			"message": "width is required!",
		})
	}
	if data["height"] == "" {
		return c.Status(500).JSON(fiber.Map{
			"message": "height is required!",
		})
	}
	if data["price"] == "" {
		return c.Status(500).JSON(fiber.Map{
			"message": "price is required!",
		})
	}
	if data["additional_price"] == "" {
		return c.Status(500).JSON(fiber.Map{
			"message": "additional_price is required!",
		})
	}
	if data["qty"] == "" {
		return c.Status(500).JSON(fiber.Map{
			"message": "qty is required!",
		})
	}
	if data["user_id"] == "" {
		return c.Status(500).JSON(fiber.Map{
			"message": "user_id is required!",
		})
	}
	if data["store_id"] == "" {
		return c.Status(500).JSON(fiber.Map{
			"message": "store_id is required!",
		})
	}
	if data["categories_id"] == "" {
		return c.Status(500).JSON(fiber.Map{
			"message": "categories_id is required!",
		})
	}
	if data["currency_id"] == "" {
		return c.Status(500).JSON(fiber.Map{
			"message": "currency_id is required!",
		})
	}
	product := models.Product{
		Title:           data["title"],
		Description:     data["description"],
		Weight:          util.ParseFloat32(data["weight"]),
		Length:          util.ParseFloat32(data["length"]),
		Width:           util.ParseFloat32(data["width"]),
		Height:          util.ParseFloat32(data["height"]),
		Price:           util.ParseFloat32(data["price"]),
		AdditionalPrice: util.ParseFloat32(data["additional_price"]),
		QTY:             util.ParseInt(data["qty"]),
		UserID:          uint(util.ParseInt(data["user_id"])),
		StoreID:         uint(util.ParseInt(data["store_id"])),
		CategoriesID:    uint(util.ParseInt(data["categories_id"])),
		CurrencyID:      uint(util.ParseInt(data["currency_id"])),
	}
	if err := database.DBConn.Create(&product).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err,
		})
	}

	return c.JSON(product)
}
