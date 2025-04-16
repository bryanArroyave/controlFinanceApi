package subscribers

import (
	"fmt"
	"strconv"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/category/application/usecases/saveCategory/dtos"
)

func (h *handler) OnUserRegister(msg *message.Message) error {

	idStr := string(msg.Payload)

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return fmt.Errorf("invalid user ID: %w", err)
	}
	defaultCategories := []*dtos.SaveCategoryParam{
		{Name: "💵 Salario", Color: "#FF5733", Type: "income"},
		{Name: "🚕 Transporte", Color: "#33FF57", Type: "expense"},
		{Name: "🍔 Comida", Color: "#3357FF", Type: "expense"},
		{Name: "🏠 Alquiler", Color: "#FF33A1", Type: "expense"},
		{Name: "📱 Teléfono", Color: "#FF8C33", Type: "expense"},
		{Name: "💻 Internet", Color: "#33FFA1", Type: "expense"},
		{Name: "🛍️ Compras", Color: "#A133FF", Type: "expense"},
		{Name: "🎉 Entretenimiento", Color: "#FF33D4", Type: "expense"},
		{Name: "🏥 Salud", Color: "#33D4FF", Type: "expense"},
		{Name: "Otros", Color: "#D4FF33", Type: "expense"},
	}

	for _, category := range defaultCategories {
		h.saveCategoryUsecase.SaveCategory(msg.Context(), id, category)
	}

	return nil
}
