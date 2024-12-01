package devicesession

import (
	"time"
	"zxsttm/database/models"

	"gorm.io/gorm"
)

func DeviceLiveCheck(db *gorm.DB) {

	for {
		// Get all devices
		var devices []models.ESP
		db.Find(&devices)

		// Loop through all devices
		for _, device := range devices {
			// Check if device is offline
			if device.LiveTimeOut != nil && device.LiveTimeOut.Add(5*time.Second).Before(time.Now()) {
				// Set device status to offline
				device.IsActive = false
				db.Save(&device)
			}
		}

		// Sleep for 5 seconds
		time.Sleep(5 * time.Second)
	}

}
