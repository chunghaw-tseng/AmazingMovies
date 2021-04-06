package persistence

import(
	"example.com/amazingmovies/src/pkg/db"
	models "example.com/amazingmovies/src/pkg/models/movies"

)

type PeopleRepository struct{}
var peopleRepository *PeopleRepository

func GetPeopleRepository() *PeopleRepository {
	if peopleRepository == nil {
		peopleRepository = &PeopleRepository{}
	}
	return peopleRepository
}

func (r *PeopleRepository) Add(people *models.People) (*models.People, error) {
	err := Create(&people)
	err = Save(&people)
	return people, err
}

func (r *PeopleRepository) GetFromName(name string) (*models.People, error){
	var ppl models.People
	where := models.People{}
	where.Name = name
	_, err := First(&where, &ppl, []string{})
	if err != nil {
		return nil, err
	}
	return &ppl, err
}

func (r *PeopleRepository) Update(people *models.People) error { return db.GetDB().Save(&people).Error}

func (r *PeopleRepository) Delete(people *models.People) error { return db.GetDB().Unscoped().Delete(&people).Error }
