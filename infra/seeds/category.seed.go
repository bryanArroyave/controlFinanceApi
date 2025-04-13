package seeds

import (
	"log"

	"github.com/bryanArroyave/eventsplit/back/user-service/infra/models"
	"gorm.io/gorm"
)

func seedCategories(db *gorm.DB) {
	categories := []models.Category{
		// Ingresos
		{ID: 1, Name: "💵 Salario", Color: "#FF5733", Type: "income"},

		// Gastos

		{ID: 2, Name: "🚕 Transporte", Color: "#33FF57", Type: "expense"},
		{ID: 3, Name: "🍔 Comida", Color: "#3357FF", Type: "expense"},
		{ID: 4, Name: "🏠 Alquiler", Color: "#FF33A1", Type: "expense"},
		{ID: 5, Name: "📱 Teléfono", Color: "#FF8C33", Type: "expense"},
		{ID: 6, Name: "💻 Internet", Color: "#33FFA1", Type: "expense"},
		{ID: 7, Name: "🛍️ Compras", Color: "#A133FF", Type: "expense"},
		{ID: 8, Name: "🎉 Entretenimiento", Color: "#FF33D4", Type: "expense"},
		{ID: 9, Name: "🏥 Salud", Color: "#33D4FF", Type: "expense"},
		{ID: 10, Name: "Otros", Color: "#D4FF33", Type: "expense"},
	}

	for _, category := range categories {
		var existing models.Category
		err := db.Where("id = ?", category.ID).First(&existing).Error
		if err == gorm.ErrRecordNotFound {
			if err := db.Create(&category).Error; err != nil {
				log.Printf("Error creando categoría %s: %v", category.Name, err)
				panic(err)
			} else {
				log.Printf("Categoría creada: %s", category.Name)
			}
		}
	}
}
