package satker

import (
	"net/http"
	"web-http/utils"

	"github.com/jmoiron/sqlx"
)

type DBHandler struct {
	DB *sqlx.DB
}

func Handler(db *sqlx.DB) *DBHandler {
	return &DBHandler{DB: db}
}

type SatkerResponse struct {
	Satker string `json:"satker"`
	Singkatan *string `json:"singkatan"`
}

func (h *DBHandler) SatkerHandler(res http.ResponseWriter, req *http.Request)  {
  utils.ResponseSetup(res, req)

  var masterSatker []SatkerResponse
  query := `SELECT satker, singkatan FROM master_satker`
  err := sqlx.Select(h.DB, &masterSatker, query)
  if err != nil {
    utils.SendResponse(res, err.Error(), http.StatusInternalServerError, nil)
  }

  utils.SendResponse(res, "Success get all Master Satker", http.StatusOK, masterSatker)
}