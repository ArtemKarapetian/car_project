package converter

import (
	servicemodel "car_project/internal/model"
	"car_project/internal/repository/car/model"
)

func toPeopleFromRepo(people *model.People) *servicemodel.People {
	return &servicemodel.People{
		Name:       people.Name,
		Surname:    people.Surname,
		Patronymic: people.Patronymic,
	}
}

func toPeopleFromService(people *servicemodel.People) *model.People {
	return &model.People{
		Name:       people.Name,
		Surname:    people.Surname,
		Patronymic: people.Patronymic,
	}
}
