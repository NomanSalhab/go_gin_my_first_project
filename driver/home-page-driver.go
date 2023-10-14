package driver

import (
	"errors"

	"github.com/NomanSalhab/go_gin_my_first_project/entity"
)

type HomePageDriver interface {
	GetHomePage(limit int, appVersion float32, wantedAreaId int) (entity.HomePage, error)
	GetLatestVersion() (float32, string)
}

type homePageDriver struct {
}

var (
	sld SliderDriver = NewSliderDriver()
)

func NewHomePageDriver() HomePageDriver {
	return &homePageDriver{}
}

func (driver *homePageDriver) GetHomePage(limit int, appVersion float32, wantedAreaId int) (entity.HomePage, error) {
	version, link := driver.GetLatestVersion()
	if version > appVersion {
		return entity.HomePage{
			Sliders:             make([]entity.Slider, 0),
			BestSellingProducts: make([]entity.Product, 0),
			BestSellingStores:   make([]entity.Store, 0),
			OffersProducts:      make([]entity.Product, 0),
			IsUpdatedVersion:    false,
			VersionLink:         link,
		}, errors.New("old app version")
	}
	bestSellingStores, err := sd.FindBestSellingStores(limit, wantedAreaId)
	if err != nil {
		return entity.HomePage{
			Sliders:             make([]entity.Slider, 0),
			BestSellingProducts: make([]entity.Product, 0),
			BestSellingStores:   make([]entity.Store, 0),
			OffersProducts:      make([]entity.Product, 0),
			IsUpdatedVersion:    true,
			VersionLink:         "",
		}, err
	}
	bestSellingProducts, err := pd.FindBestSellingProducts(limit)
	if err != nil {
		return entity.HomePage{
			Sliders:             make([]entity.Slider, 0),
			BestSellingProducts: make([]entity.Product, 0),
			BestSellingStores:   make([]entity.Store, 0),
			OffersProducts:      make([]entity.Product, 0),
			IsUpdatedVersion:    true,
			VersionLink:         "",
		}, err
	}
	offersProducts, err := pd.FindOffersProducts()
	if err != nil {
		return entity.HomePage{
			Sliders:             make([]entity.Slider, 0),
			BestSellingProducts: make([]entity.Product, 0),
			BestSellingStores:   make([]entity.Store, 0),
			OffersProducts:      make([]entity.Product, 0),
			IsUpdatedVersion:    true,
			VersionLink:         "",
		}, err
	}
	sliders, err := sld.FindActiveSliders()
	if err != nil {
		return entity.HomePage{
			Sliders:             make([]entity.Slider, 0),
			BestSellingProducts: make([]entity.Product, 0),
			BestSellingStores:   make([]entity.Store, 0),
			OffersProducts:      make([]entity.Product, 0),
			IsUpdatedVersion:    true,
			VersionLink:         "",
		}, err
	}
	homePage := entity.HomePage{
		Sliders:             sliders,
		BestSellingProducts: bestSellingProducts,
		BestSellingStores:   bestSellingStores,
		OffersProducts:      offersProducts,
		IsUpdatedVersion:    true,
		VersionLink:         link,
	}
	return homePage, nil
}

func (driver *homePageDriver) GetLatestVersion() (float32, string) {
	query := `select id, version, link from latest_version where id = $1`
	var id int
	var version float32
	var link string
	row := dbConn.SQL.QueryRow(query, 1)
	err := row.Scan(&id, &version, &link)
	// fmt.Println("User Data:", id, name, phone)
	if err != nil {
		return 100.0, ""
	}
	return version, link
}
