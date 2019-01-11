package models

//var (
//	db *gorm.DB
//)
//
//func InitModel(config ivankastd.ConfigMysql) {
//	db = toolkit.CreateDB(config)
//	db.AutoMigrate(
//		&InServiceUser{},
//	)
//}
//
//func CloseDB() {
//	if db != nil {
//		db.Close()
//	}
//}
//
//func DB() *gorm.DB {
//	return db
//}



type Pagination struct {
	Page  int64 `json:"page" query:"page"`
	Limit int64 `json:"limit" query:"limit"`
}

