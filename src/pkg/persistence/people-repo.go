package persistence

import(
	"strconv"
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

func (r *PeopleRepository) Get(id string) (*models.People, error) {
	var people models.People
	where := models.People{}
	where.ID, _ = strconv.ParseUint(id, 10, 64)
	_, err := First(&where, &people, []string{})
	if err != nil {
		return nil, err
	}
	return &people, err
}


func (r *PeopleRepository) Query(q *models.People) (*[]models.People, error) {
	var ppl []models.People
	err := Find(&q, &ppl, []string{}, "id asc")
	return &ppl, err
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

func (r *PeopleRepository) Delete(people *models.People) error { 
	db.GetDB().Model(people).Association("Movies").Clear()
	return db.GetDB().Unscoped().Delete(&people).Error }
