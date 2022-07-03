package api

import (
	"encoding/json"
	"net/http"

	"github.com/sofiukl/oms-core/models"
	"github.com/sofiukl/oms-core/utils"
)

// FindCart - find cart
func FindCart(id string, w http.ResponseWriter, r *http.Request) {

	cart := models.CartModel{
		ID: "c2",
		Products: []models.ProductModel{
			{ID: "p1", Quantity: 1},
		},
	}

	o1 := models.OrderModel{
		OrderID:      "01",
		OrderAddress: "order address",
	}

	// Return response
	var inInterface map[string]interface{}
	inrec, _ := json.Marshal(cart)
	json.Unmarshal(inrec, &inInterface)

	inrec, _ = json.Marshal(o1)
	json.Unmarshal(inrec, &inInterface)

	utils.RespondWithJSON(w, http.StatusOK, "Cart found successfully", "", inInterface)
}
