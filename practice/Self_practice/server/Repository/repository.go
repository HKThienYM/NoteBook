package repo

// "github.com/jinzhu/gorm"
// _ "github.com/jinzhu/gorm/dialects/mysql"

type Brand struct {
	ID        uint `gorm:"primary_key"`
	Name      string
	cars      []Vehicle
	motocycle []Vehicle
}

type Vehicle struct {
	ID         uint `gorm:"primary_key"`
	BrandID    int
	Name       string
	NumInStock int
}

var db = []Brand{{ID: 1, Name: "Yamaha", motocycle: []Vehicle{{ID: 1, BrandID: 1, Name: "FZ25", NumInStock: 10}}}}

func GetBrand() []Brand {
	return db
}

func GetAllVehicle() []Vehicle {
	var result []Vehicle
	for _, brand := range db {
		result = append(result, brand.motocycle...)
	}
	return result
}

func GetBrandVehicle(brand string) []Vehicle {
	for _, b := range db {
		if b.Name == brand {
			return b.motocycle
		}
	}
	return nil
}
