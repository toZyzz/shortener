package url

import "gorm.io/gorm"

type Url struct {
	gorm.Model
	Short string
	Full  string
}
