package pay_center_client

func NewClient() *PayClient {
	client := NewPayClient(APPID, APIKEY, "http://127.0.0.1:8128")
	return client
}
