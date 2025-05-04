package database

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"gorm.io/gorm"
)

// Hàm MigrateAll thực hiện migration theo yêu cầu "up" hoặc "down"
func MigrateAll(db *gorm.DB, direction string) {
	migrationsPath := "./database/migrations"

	// Đọc tất cả các file trong thư mục migrations
	files, err := ioutil.ReadDir(migrationsPath)
	if err != nil {
		log.Fatal("Failed to read migration directory:", err)
	}

	// Sắp xếp các file theo tên (nếu cần thiết)
	for _, file := range files {
		// Lọc chỉ các file có đuôi .sql
		if strings.HasSuffix(file.Name(), ".sql") {
			filePath := filepath.Join(migrationsPath, file.Name())
			migrationSQL, err := ioutil.ReadFile(filePath)
			if err != nil {
				log.Fatal("Failed to read migration file:", err)
			}

			// Kiểm tra và chỉ thực thi các phần tương ứng với "Up" hoặc "Down"
			if direction == "up" && strings.Contains(string(migrationSQL), "-- +migrate Up") {
				// Chạy phần Up
				upSQL := extractSQLSection(string(migrationSQL), "-- +migrate Up", "-- +migrate Down")
				if err := db.Exec(upSQL).Error; err != nil {
					log.Fatal("Failed to apply migration:", err)
				}
				fmt.Printf("Applied migration UP: %s\n", file.Name())
			} else if direction == "down" && strings.Contains(string(migrationSQL), "-- +migrate Down") {
				// Chạy phần Down
				downSQL := extractSQLSection(string(migrationSQL), "-- +migrate Down", "-- +migrate Up")
				if err := db.Exec(downSQL).Error; err != nil {
					log.Fatal("Failed to roll back migration:", err)
				}
				fmt.Printf("Applied migration DOWN: %s\n", file.Name())
			}
		}
	}
}

// Hàm extractSQLSection dùng để lấy phần SQL giữa 2 comment chỉ định
func extractSQLSection(migrationSQL, startComment, endComment string) string {
	startIndex := strings.Index(migrationSQL, startComment)
	endIndex := strings.Index(migrationSQL, endComment)

	// Nếu tìm thấy cả 2 phần comment, trích xuất SQL giữa chúng
	if startIndex != -1 && endIndex != -1 {
		return migrationSQL[startIndex+len(startComment) : endIndex]
	}
	return ""
}
