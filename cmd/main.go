// main.go

package main

import (
	"lixIQ/backend/internal/db"
)

func main() {
	// MongoDB bağlantısını yapın
	// db.Init() otomatik olarak yapılacaktır

	// Bağlantıyı kapatın (işiniz bittiğinde)
	defer db.Close()

	// MongoDB veritabanı ve koleksiyon işlemleri burada yapılabilir
	db.GetDatabase().Collection("user")

	// fmt.Println("MongoDB bağlantısı başarıyla yapıldı!")
}
