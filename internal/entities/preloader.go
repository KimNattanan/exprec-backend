package entities

import "gorm.io/gorm"

func PreloadDepth(s string, depth int) func(*gorm.DB) *gorm.DB {
	return func(d *gorm.DB) *gorm.DB {
		if depth <= 0 {
			return d
		}
		return d.Preload(s, PreloadDepth(s, depth-1))
	}
}
