package dbtools


func SubmitDeleteRequest(url string, reqBody []byte) int {

	resp, _ := MakeHttpRequestWrapper(url, "DELETE", nil)

	return resp.StatusCode
}
