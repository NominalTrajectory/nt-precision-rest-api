package okr

import (
	"net/http"

	"github.com/NominalTrajectory/nt-precision-rest-api/database"
	"github.com/NominalTrajectory/nt-precision-rest-api/models"
	"github.com/NominalTrajectory/nt-precision-rest-api/utils"
)

//

func GetAllObjectives(w http.ResponseWriter, r *http.Request) {
	var objectives []models.Objective
	if err := database.DB.Find(&objectives).Error; err != nil {
		http.Error(w, err.Error(), utils.ErrToStatusCode(err))
	} else {
		utils.WriteJSONResult(w, objectives)
	}
}

// Get objective by id
func GetObjective(w http.ResponseWriter, r *http.Request) {

}

// Create objective
func CreateObjective(w http.ResponseWriter, r *http.Request) {

}

// Update objective
func UpdateObjective(w http.ResponseWriter, r *http.Request) {

}

// Delete objective
func DeleteObjective(w http.ResponseWriter, r *http.Request) {

}

/* KEY RESULTS */
