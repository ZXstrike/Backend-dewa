package devicesession

import (
	"time"
	"zxsttm/database"
	"zxsttm/database/models"
)

func ElectricalExpenseCal() {
	// Calculate electrical expense
	// This function will calculate the electrical expense for each device
	// and update the database with the new value

	db := database.DB

	// Get all devices

	for {
		var devices []models.ESP
		db.Find(&devices)

		// Loop through all devices

		for _, device := range devices {
			var devicePzem []models.PZEM
			db.Where("es_pid = ?", device.ID).Where("created_at > ?", time.Now().Add(30*time.Second)).Find(&devicePzem)

			// Calculate electrical expense
			// The formula for calculating electrical expense is:
			// electrical expense = Energy * cost

			// where:
			// Energy = devicePzem.Energy
			// time = devicePzem.Time
			// cost = 0.1

			for _, pzem := range devicePzem {
				// Calculate electrical expense

				electricalExpense := pzem.Energy * 0.1

				history := models.History{
					ESPcode: device.ESPcode,
					KwH:     pzem.Energy,
					Price:   electricalExpense,
				}

				db.Create(&history)
			}

			// Update the database with the new electrical expense

		}

		time.Sleep(30 * time.Second)
	}

}
