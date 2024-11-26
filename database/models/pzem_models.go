package models

import (
	"gorm.io/gorm"
)

type PZEM struct {
	gorm.Model
	ESPID   string
	Current float64
	Power   float64
	Voltage float64
	Energy  float64 //ntar bikin hitungan energy * harga kwh lalu dicatat dan dihitung setiap jam (setiap 30 hari muncul langsung historynya jadi bukan live counting)
}
