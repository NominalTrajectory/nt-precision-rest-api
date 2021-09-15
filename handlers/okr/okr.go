package okr

import (
	"net/http"

	"github.com/NominalTrajectory/nt-precision-rest-api/database"
	"github.com/NominalTrajectory/nt-precision-rest-api/model"
	"github.com/NominalTrajectory/nt-precision-rest-api/utils"
)

func GetAllObjectives(w http.ResponseWriter, r *http.Request) {
	var objectives []model.Objective
	if err := database.DB.Find(&objectives).Error; err != nil {
		http.Error(w, err.Error(), utils.ErrToStatusCode(err))
	} else {
		utils.WriteJSONResult(w, objectives)
	}

}
