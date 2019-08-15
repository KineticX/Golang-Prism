package models

/*
__________        .__                
\______   \_______|__| ______ _____  
 |     ___/\_  __ \  |/  ___//     \ 
 |    |     |  | \/  |\___ \|  Y Y  \
 |____|     |__|  |__/____  >__|_|  /
                          \/      \/ 

---==-=---=---===----=-=-=-=----=-=-=-=-----=-=-=*--=-=-=-=-------

 => File: sample.go
 => Version: X.X.X.X
 => Author: Jonathan McAllister
 => Purpose: 

    	Model example for the sample route. The aim of this is to
		simply provide a simple API solution for the MVC framework
		The calling entity is ./api/v1/sample.go
*/


// Imports
// ----------------------------------
import (

	"github.com/jinzhu/gorm"
)

// STRUCT: Sample
// ----------------------------------
type Sample struct {

	Model

	TagID int `json:"tag_id" gorm:"index"`
	
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	CreatedBy     string `json:"created_by"`
	ModifiedBy    string `json:"modified_by"`
	State         int    `json:"state"`

}


// FUNC: GetArticle Get a single sample based on ID
// ----------------------------------
func GetSample(id int) (*Sample, error) {

	var sample Sample

	err := db.Where("id = ? AND deleted_on = ? ", id, 0).First(&sample).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}


	return &sample, nil
}
