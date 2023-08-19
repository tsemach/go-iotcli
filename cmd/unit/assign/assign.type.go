package assign

type AssignResponseStruct struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Route   string `json:"route"`
}

type AssignBodyStruct struct {
	Pid              string `json:"production_identifier"`
	Project_name     string `json:"project_name"`
	Fleet_name       string `json:"fleet_name"`
	Fleet_id         string `json:"fleet_id"`
	Fleet_assignment bool   `json:"fleet_assignment"`
}

var bodyBytes = []byte(`{
	"production_identifier": "0000000000000000",
	"project_name": "project_name",
	"fleet_name": "fleet_name",
	"project_id": "project_id",
	"fleet_id": "fleet_id",
	"fleet_assignment": true
}`)
