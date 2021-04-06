package persistence

import(
	"github.com/jinzhu/gorm"
	"example.com/amazingmovies/src/pkg/db"
)

// Create
func Create(value interface{}) error {
	return db.GetDB().Create(value).Error
}

// Save
func Save(value interface{}) error {
	return db.GetDB().Save(value).Error
}

// Updates
func Updates(where interface{}, value interface{}) error {
	return db.GetDB().Model(where).Updates(value).Error
}


// GET First Element by ID
func FirstByID(out interface{}, id string) (notFound bool, err error) {
	err = db.GetDB().First(out, id).Error
	if err != nil {
		notFound = gorm.IsRecordNotFoundError(err)
	}
	return
}

// GET First element
func First(where interface{}, out interface{}, associations []string) (notFound bool, err error) {
	db := db.GetDB()
	for _, a := range associations {
		db = db.Preload(a)
	}
	err = db.Where(where).First(out).Error
	if err != nil {
		notFound = gorm.IsRecordNotFoundError(err)
	}
	return
}

// Find
func Find(where interface{}, out interface{}, associations []string, orders ...string) error {
	db := db.GetDB()
	for _, a := range associations {
		db = db.Preload(a)
	}
	db = db.Where(where)
	if len(orders) > 0 {
		for _, order := range orders {
			db = db.Order(order)
		}
	}
	return db.Find(out).Error
}




// Find similar to 
func FindLike(where string, query string, out interface{}, associations []string, orders ...string) error {
	db := db.GetDB()
	for _, a := range associations {
		db = db.Preload(a)
	}
	db = db.Where(where , query)
	if len(orders) > 0 {
		for _, order := range orders {
			db = db.Order(order)
		}
	}
	return db.Find(out).Error
}

// Delete Model
func DeleteByModel(model interface{}) (count int64, err error) {
	db := db.GetDB().Delete(model)
	err = db.Error
	if err != nil {
		return
	}
	count = db.RowsAffected
	return
}

// Delete
func DeleteByWhere(model, where interface{}) (count int64, err error) {
	db := db.GetDB().Where(where).Delete(model)
	err = db.Error
	if err != nil {
		return
	}
	count = db.RowsAffected
	return
}

// Delete by ID
func DeleteByID(model interface{}, id uint64) (count int64, err error) {
	db := db.GetDB().Where("id=?", id).Delete(model)
	err = db.Error
	if err != nil {
		return
	}
	count = db.RowsAffected
	return
}
