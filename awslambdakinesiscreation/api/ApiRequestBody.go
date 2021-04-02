//Package api contains API request body
package api

//RequestBody is the form in which the body of the request has to be.
type RequestBody struct {
	ClientName string `json:"name"`
	ClientSize string `json:"size"`
}
