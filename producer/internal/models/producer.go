package models

type MainProducer struct {
	ProducerID          string          `json:"producer_id"`
	ProducerMessage     ProducerMessage `json:"producer_message"`
	ProducerServiceArea string          `json:"producer_service_area"`
	ProducerCreatedAt   string          `json:"producer_created_at"`
}

type ProducerMessage struct {
	ProducerMessageID string       `json:"producer_message_id"`
	Host              string       `json:"host"`
	Client            string       `json:"client"`
	Ip                string       `json:"ip"`
	Port              string       `json:"port"`
	DataProducer      DataProducer `json:"data_producer"`
	CreatedAt         string       `json:"information_created_at"`
}

type DataProducer struct {
	DataProducerID string  `json:"data_producer_id"`
	Product        string  `json:"product"`
	Name           string  `json:"name"`
	Category       string  `json:"category"`
	SubCategory    string  `json:"sub_category"`
	Price          float64 `json:"price"`
	Quantity       int32   `json:"quantity"`
	Supplier       string  `json:"supplier"`
	Description    string  `json:"description"`
	Gender         string  `json:"gender"`
}
