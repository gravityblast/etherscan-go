package etherscan

type AccountResponse struct {
	Status  string  `json:"status"`
	Message string  `json:"message"`
	Result  *BigInt `json:"result"`
}
