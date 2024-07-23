package modals


type DltResponse struct {
	BatchID  string `json:"batchId"`
	//FarmerID string `json:"farmerId"`
	Variety  string `json:"variety"`
	Quantity string    `json:"quantity"`
	Origin   string `json:"origin"`
	Status string `json:"status"`
	Price  string `json:"price"`
	ProducedBy string `json:"producedBy"`
	ProcessedBy string `json:"processedBy"`
	DistributedBy string `json:"distributedBy"`
	ReceivedBy string `json:"receivedBy"`
	CurentOwner string `json:"currentowner"`
	NewOwner   string   `json:"newowner"`
	Traceability Event `json:"traceability"`

}

type Event struct {
	// Timestamp time.Time `json:"timestamp"`
	Org    string `json:"org"`
	Status string `json:"status"`
}
