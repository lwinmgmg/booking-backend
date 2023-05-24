package schemas

type ProductResponse struct {
	ID       uint    `json:"id"`
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	Price    float64 `json:"price"`
	UomID    uint    `json:"uom_id"`
	UomName  string  `json:"uom_name"`
	ImageUrl string  `json:"image_url"`
}
