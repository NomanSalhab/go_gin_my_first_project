package controller

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/NomanSalhab/go_gin_my_first_project/entity"
	"github.com/NomanSalhab/go_gin_my_first_project/service"
	"github.com/gin-gonic/gin"
)

type ProductController interface {
	FindAllProducts() []entity.Product
	AddProduct(ctx *gin.Context, cst ProductCategoryController, sc StoreController, fc FileController) error
	FindActiveProducts() []entity.Product
	FindNotActiveProducts() []entity.Product
	GetProductById(ctx *gin.Context) (entity.Product, error)
	GetProductByCategory(ctx *gin.Context) ([]entity.Product, error)
	EditProduct(ctx *gin.Context, cst ProductCategoryController, sc StoreController, fc FileController) error
	ActivateProduct(ctx *gin.Context) error
	DeactivateProduct(ctx *gin.Context) error
	DeleteProduct(ctx *gin.Context, fc FileController) error
	OrderProduct(ctx *gin.Context) error
}

type productController struct {
	service service.ProductService
}

func NewProductController(service service.ProductService) ProductController {
	return &productController{
		service: service,
	}
}

func (c *productController) FindAllProducts() []entity.Product {
	return c.service.FindAllProducts()
}

func (c *productController) AddProduct(ctx *gin.Context, cst ProductCategoryController, sc StoreController, fc FileController) error {
	var product entity.Product
	// err := ctx.ShouldBindJSON(&product)
	// if err != nil {
	// 	return err
	// }
	// err = validate.Struct(product)
	// if err != nil {
	// 	return err
	// }
	product, image, err := SetProduct(ctx)
	if err != nil {
		return err
	}
	productCategories := cst.FindAllProductCategories()
	stores := sc.FindAllStores()
	for i := 0; i < len(productCategories); i++ {
		if productCategories[i].ID == product.ProductCategoryId {
			for i := 0; i < len(stores); i++ {
				if stores[i].ID == product.StoreId {

					err = c.service.AddProduct(product)
					if err != nil {
						return err
					}
					err = fc.AddFile(image)
					return err
				}
			}
		}
	}
	return errors.New("product category or store does not exist")
}

func (c *productController) FindActiveProducts() []entity.Product {
	return c.service.FindActiveProducts()
}

func (c *productController) FindNotActiveProducts() []entity.Product {
	return c.service.FindNotActiveProducts()
}

func (c *productController) GetProductById(ctx *gin.Context) (entity.Product, error) {
	var productId entity.ProductInfoRequest
	var product entity.Product
	err := ctx.ShouldBindJSON(&productId)
	if err != nil {
		return product, err
	}
	product, err = c.service.FindProduct(productId)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (c *productController) EditProduct(ctx *gin.Context, cst ProductCategoryController, sc StoreController, fc FileController) error {
	var productEditInfo entity.ProductEditRequest
	// err := ctx.ShouldBindJSON(&productEditInfo)
	// if err != nil {
	// 	return err
	// }
	// err = validate.Struct(productEditInfo)
	// if err != nil {
	// 	return err
	// }

	productEditInfo, err := SetProductForEditing(ctx)
	fmt.Println("Product:", productEditInfo)
	if err != nil {
		return err
	}
	if productEditInfo.Image != "" {
		oldProduct, err := c.service.FindProduct(entity.ProductInfoRequest{ID: productEditInfo.ID})
		if err != nil {
			return err
		}
		err = fc.DeleteFile(oldProduct.Image)
		if err != nil {
			return err
		}
		file, err := GetImage(ctx)
		if err != nil {
			return err
		}
		err = fc.AddFile(file)
		if err != nil {
			return err
		}
		productEditInfo.Image = file.UUID
	}

	// productCategories := cst.FindAllProductCategories()
	// stores := sc.FindAllStores()
	// for i := 0; i < len(productCategories); i++ {
	// 	if productCategories[i].ID == productEditInfo.ProductCategoryId {
	// 		for j := 0; j < len(stores); j++ {
	// 			//!!! TODO::: Make The Product Category Be Related To The Required Store
	// 			if stores[j].ID == productEditInfo.StoreId /*&& (stores[j].ID == productCategories[i].StoreId)*/ {
	err = c.service.EditProduct(productEditInfo)
	return err
	// 			}
	// 		}
	// 	}
	// }
	// return errors.New("product category or store does not exist")
}

func (c *productController) ActivateProduct(ctx *gin.Context) error {
	var productEditInfo entity.ProductInfoRequest
	err := ctx.ShouldBindJSON(&productEditInfo)
	if err != nil {
		return err
	}
	err = validate.Struct(productEditInfo)
	if err != nil {
		return err
	}

	err = c.service.ActivateProduct(productEditInfo)
	return err
}

func (c *productController) DeactivateProduct(ctx *gin.Context) error {
	var productEditInfo entity.ProductInfoRequest
	err := ctx.ShouldBindJSON(&productEditInfo)
	if err != nil {
		return err
	}
	err = validate.Struct(productEditInfo)
	if err != nil {
		return err
	}

	err = c.service.DeactivateProduct(productEditInfo)
	return err
}

func (c *productController) DeleteProduct(ctx *gin.Context, fc FileController) error {
	var productId entity.ProductInfoRequest
	err := ctx.ShouldBindJSON(&productId)
	if err != nil {
		return err
	}
	err = validate.Struct(productId)
	if err != nil {
		return err
	}

	oldProduct, err := c.service.FindProduct(productId)
	if err != nil {
		return err
	}
	err = c.service.DeleteProduct(productId)
	if err != nil {
		return err
	}
	err = fc.DeleteFile(oldProduct.Image)
	if err != nil {
		return err
	}
	return nil
}

func (c *productController) GetProductByCategory(ctx *gin.Context) ([]entity.Product, error) {
	var productId entity.ProductByCategoryRequest
	var products []entity.Product
	err := ctx.ShouldBindJSON(&productId)
	if err != nil {
		return products, err
	}
	err = validate.Struct(productId)
	if err != nil {
		return products, err
	}

	products, err = c.service.FindProductsByCategory(productId)
	if err != nil {
		return products, err
	}
	return products, nil
}

func (c *productController) OrderProduct(ctx *gin.Context) error {
	var productId entity.OrderProductRequest
	err := ctx.ShouldBindJSON(&productId)
	if err != nil {
		return err
	}
	err = validate.Struct(productId)
	if err != nil {
		return err
	}

	err = c.service.OrderProduct(productId)
	if err != nil {
		return err
	}
	return nil
}

func SetProduct(ctx *gin.Context) (entity.Product, entity.File, error) {
	addons, addonsBool := ctx.GetPostForm("addons")
	if !addonsBool {
		return entity.Product{}, entity.File{}, errors.New("addons are a required field")
	}
	fmt.Println("Addons:", addons)
	addonsValue, addonsError := StringToSlice(addons[1 : len(addons)-1])
	if addonsError != nil {
		return entity.Product{}, entity.File{}, addonsError
	}
	flavors, flavorsBool := ctx.GetPostForm("flavors")
	if !flavorsBool {
		return entity.Product{}, entity.File{}, errors.New("flavors are a required field")
	}
	fmt.Println("Flavors:", flavors)
	flavorsValue, flavorsError := StringToSlice(flavors[1 : len(flavors)-1])
	if flavorsError != nil {
		return entity.Product{}, entity.File{}, flavorsError
	}
	volumes, volumesBool := ctx.GetPostForm("volumes")
	if !volumesBool {
		return entity.Product{}, entity.File{}, errors.New("volumes are a required field")
	}
	fmt.Println("Volumes:", volumes)
	volumesValue, volumesError := StringToSlice(volumes[1 : len(volumes)-1])
	if volumesError != nil {
		return entity.Product{}, entity.File{}, volumesError
	}
	fileMetadata, err := GetImage(ctx)
	if err != nil {
		return entity.Product{}, entity.File{}, err
	}

	storeId, storeIdBool := ctx.GetPostForm("store_id")
	if !storeIdBool {
		return entity.Product{}, entity.File{}, errors.New("store id is a required field")
	}
	storeIdValue, storeIdError := strconv.Atoi(storeId)
	if storeIdError != nil {
		return entity.Product{}, entity.File{}, errors.New("an unacceptable store id value")
	}
	productCategoryId, productCategoryIdBool := ctx.GetPostForm("product_category_id")
	if !productCategoryIdBool {
		return entity.Product{}, entity.File{}, errors.New("product category id is a required field")
	}
	productCategoryIdValue, productCategoryIdError := strconv.Atoi(productCategoryId)
	if productCategoryIdError != nil {
		return entity.Product{}, entity.File{}, errors.New("an unacceptable product category id value")
	}
	price, priceBool := ctx.GetPostForm("price")
	if !priceBool {
		return entity.Product{}, entity.File{}, errors.New("price is a required field")
	}
	priceValue, priceError := strconv.Atoi(price)
	if priceError != nil {
		return entity.Product{}, entity.File{}, errors.New("an unacceptable price value")
	}
	summary, summaryError := ctx.GetPostForm("summary")
	if !summaryError || len(summary) == 0 {
		return entity.Product{}, entity.File{}, errors.New("summary is a required field")
	}
	discountRatio, discountRatioBool := ctx.GetPostForm("discount_ratio")
	if !discountRatioBool {
		discountRatio = "0.0"
	}
	discountRatio64Value, discountRatioError := strconv.ParseFloat(discountRatio, 32)
	if discountRatioError != nil {
		return entity.Product{}, entity.File{}, errors.New("an unacceptable discount value")
	}
	discountRatioValue := float32(discountRatio64Value)
	active, _ := ctx.GetPostForm("active")
	var activeValue bool
	if active == "true" {
		activeValue = true
	} else {
		activeValue = false
	}
	name, nameBool := ctx.GetPostForm("name")
	if !nameBool || len(name) == 0 {
		return entity.Product{}, entity.File{}, errors.New("name is a required field")
	}

	product := entity.Product{
		Active:            activeValue,
		Name:              name,
		Image:             fileMetadata.UUID,
		StoreId:           storeIdValue,
		ProductCategoryId: productCategoryIdValue,
		Summary:           summary,
		DiscountRatio:     discountRatioValue,
		Price:             priceValue,
		Addons:            addonsValue,
		Flavors:           flavorsValue,
		Volumes:           volumesValue,
	}
	return product, fileMetadata, nil
}

func SetProductForEditing(ctx *gin.Context) (entity.ProductEditRequest, error) {
	var product entity.ProductEditRequest
	fileMetadata, imageError := GetImage(ctx)
	addons, addonsBool := ctx.GetPostForm("addons")
	if !addonsBool {
		addons = "[]"
	}
	addonsValue, addonsError := StringToSlice(addons[1 : len(addons)-1])
	if addonsError != nil {
		return entity.ProductEditRequest{}, addonsError
	}
	flavors, flavorsBool := ctx.GetPostForm("flavors")
	if !flavorsBool {
		flavors = "[]"
	}
	fmt.Println("Flavors:", flavors)
	flavorsValue, flavorsError := StringToSlice(flavors[1 : len(flavors)-1])
	if flavorsError != nil {
		return entity.ProductEditRequest{}, flavorsError
	}
	volumes, volumesBool := ctx.GetPostForm("volumes")
	if !volumesBool {
		volumes = "[]"
	}
	fmt.Println("Volumes:", volumes)
	volumesValue, volumesError := StringToSlice(volumes[1 : len(volumes)-1])
	if volumesError != nil {
		return entity.ProductEditRequest{}, volumesError
	}

	id, idBool := ctx.GetPostForm("id")
	if !idBool {
		return entity.ProductEditRequest{}, errors.New("id is a required field")
	}
	idValue, idError := strconv.Atoi(id)
	if idError != nil {
		return entity.ProductEditRequest{}, errors.New("an unacceptable id value")
	}
	storeId, storeIdBool := ctx.GetPostForm("store_id")
	if !storeIdBool {
		storeId = "0"
	}
	storeIdValue, storeIdError := strconv.Atoi(storeId)
	if storeIdError != nil {
		return entity.ProductEditRequest{}, errors.New("an unacceptable store id value")
	}
	productCategoryId, productCategoryIdBool := ctx.GetPostForm("product_category_id")
	if !productCategoryIdBool {
		productCategoryId = "0"
	}
	productCategoryIdValue, productCategoryIdError := strconv.Atoi(productCategoryId)
	if productCategoryIdError != nil {
		return entity.ProductEditRequest{}, errors.New("an unacceptable product category id value")
	}
	price, priceBool := ctx.GetPostForm("price")
	if !priceBool {
		price = "0"
	}
	priceValue, priceError := strconv.Atoi(price)
	if priceError != nil {
		return entity.ProductEditRequest{}, errors.New("an unacceptable price value")
	}
	summary, summaryBool := ctx.GetPostForm("summary")
	if !summaryBool {
		summary = ""
	}
	discountRatio, discountRatioBool := ctx.GetPostForm("discount_ratio")
	if !discountRatioBool {
		discountRatio = "0.0"
	}
	discountRatio64Value, discountRatioError := strconv.ParseFloat(discountRatio, 32)
	if discountRatioError != nil {
		return entity.ProductEditRequest{}, errors.New("an unacceptable discount value")
	}
	discountRatioValue := float32(discountRatio64Value)
	active, _ := ctx.GetPostForm("active")
	var activeValue bool
	if active == "true" {
		activeValue = true
	} else {
		activeValue = false
	}
	name, _ := ctx.GetPostForm("name")
	if imageError != nil {
		product = entity.ProductEditRequest{
			ID:                idValue,
			Active:            activeValue,
			Name:              name,
			StoreId:           storeIdValue,
			ProductCategoryId: productCategoryIdValue,
			Summary:           summary,
			DiscountRatio:     discountRatioValue,
			Price:             priceValue,
			Addons:            addonsValue,
			Flavors:           flavorsValue,
			Volumes:           volumesValue,
		}
	} else {
		product = entity.ProductEditRequest{
			ID:                idValue,
			Active:            activeValue,
			Name:              name,
			Image:             fileMetadata.UUID,
			StoreId:           storeIdValue,
			ProductCategoryId: productCategoryIdValue,
			Summary:           summary,
			DiscountRatio:     discountRatioValue,
			Price:             priceValue,
			Addons:            addonsValue,
			Flavors:           flavorsValue,
			Volumes:           volumesValue,
		}
	}

	return product, nil
}

func StringToSlice(sliceString string) ([]entity.DetailEditRequest, error) {
	if len(sliceString) == 0 {
		return make([]entity.DetailEditRequest, 0), nil
	}
	s := strings.Split(sliceString, ",")
	fmt.Println("s:", s)
	slice := make([]entity.DetailEditRequest, 0)
	for i := 0; i < len(s); i++ {
		fmt.Println("item:", s[i])
		id, err := strconv.Atoi(s[i])
		if err != nil {
			return slice, err
		}
		slice = append(slice, entity.DetailEditRequest{ID: id})
	}
	fmt.Println("slice:", slice)

	return slice, nil
}
