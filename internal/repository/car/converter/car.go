package converter

import (
	servicemodel "car_project/internal/model"
	"car_project/internal/repository/car/model"
)

func ToCarFromRepo(car *model.Car) *servicemodel.Car {
	return &servicemodel.Car{
		RegNum: car.RegNum,
		Mark:   car.Mark,
		Model:  car.Model,
		Year:   car.Year,
		Owner:  toPeopleFromRepo(car.Owner),
	}
}

func ToRepoFromCar(car *servicemodel.Car) *model.Car {
	return &model.Car{
		RegNum: car.RegNum,
		Mark:   car.Mark,
		Model:  car.Model,
		Year:   car.Year,
		Owner:  toPeopleFromService(car.Owner),
	}
}
