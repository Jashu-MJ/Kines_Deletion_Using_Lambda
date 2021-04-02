package api

import (
	"log"
	"strconv"
)

//ValidateRequestBody is used for validating the body and all the params
//are valid or not using IsReqBodyValid()

// IsReqBodyValid function is used to validate the request body.
func isReqBodyValid(apiReq RequestBody) bool {
	// checking for nil values in the apireq body
	_, err := strconv.Atoi(apiReq.ClientSize)
	if err != nil || apiReq.ClientName == "" {
		log.Println("the entered values are inappropriate")
		return false
	}
	return true
}
