func RunMigrations(db *gorm.DB) {
	createTables(db)
}
