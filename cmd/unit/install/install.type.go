package install

type InstallStructResponseType struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Route   string `json:"route"`
}

type InstallBodyStruct struct {
	Pid    string `json:"production_identifier"`
	Tid    string `json:"tls_guid"`
	System struct {
		Product      string `json:"product"`
		Catalog_num  string `json:"catalog_num"`
		HW           string `json:"hw"`
		Board_rev    string `json:"board_rev"`
		Firmware     string `json:"firmware"`
		Tls_subject  string `json:"tls_subject"`
		Iot_base_url string `json:"iot_base_url"`
	} `json:"system"`
	Installer struct {
		Company_name string `json:"company_name"`
		Country      string `json:"country"`
		Site         string `json:"site"`
		User_name    string `json:"user_name"`
		First_name   string `json:"first_name"`
		Last_name    string `json:"last_name"`
	} `json:"installer"`
	Vehicle struct {
		Manufacturer       string `json:"manufacturer"`
		Vehicle_model      string `json:"vehicle_model"`
		License_plate_type string `json:"license_plate_type"`
		License_plate      string `json:"license_plate"`
		Chassis            string `json:"chassis"`
		Fleet              string `json:"fleet"`
	} `json:"vehicle"`
	Installation struct {
		Start_date      string `json:"start_date"`
		End_date        string `json:"end_date"`
		Finish_status   string `json:"finish_status"`
		Location        string `json:"location"`
		Ic_sw_ver       string `json:"ic_sw_ver"`
		Ic_tool_type    string `json:"ic_tool_type"`
		Installation_id string `json:"installation_id"`
	} `json:"installation"`
}

var bodyBytes = []byte(`{
  "production_identifier": "0000000000000000",
  "tls_guid": "00000000-0000-0000-0000-000000000000",
  "system": {
    "product": "XXY",
    "catalog_num": "XXY000000000004G",
    "hw": "Something lite",
    "board_rev": "Rev 1.0",    
    "firmware": "FW1.1.1.1.1.1.1.1.1",
    "tls_subject": "CN=CN, OU=OU, O=O, L=L, S=S, C=C",
    "iot_base_url": "some.place.com"
  },  
  "installer": {
    "company_name": "test",
    "country": "IL",
    "site": "site",
    "user_name": "someone@someplace.com",
    "first_name": "John",
    "last_name": "Smith"
  },  
  "vehicle": {
    "manufacturer": "manufacturer",
    "vehicle_model": "1234",
    "license_plate_type": null,
    "license_plate": null,
    "chassis": null,
    "fleet": null
  },  
  "installation": {
    "start_date": "2019-09-10T15:05:35.857",
    "end_date": "2019-09-10T15:10:00.107",
    "finish_status": "FINISHED",
    "location": null,
    "ic_sw_ver": "1.0",
    "ic_tool_type": "type",
    "installation_id": "1234"
  }
}`)
