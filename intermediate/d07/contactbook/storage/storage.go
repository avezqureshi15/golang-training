package storage

import (
	"contactbook/models"
	"encoding/json"
	"os"
)

const FileName = "contacts.json"

func LoadContacts() ([]models.Contact,error) {
	var contacts []models.Contact

	file,err := os.ReadFile(FileName)
	if err != nil{
		if os.IsNotExist(err) {
			return contacts, nil 
		}
		return nil ,err
	}
	err = json.Unmarshal(file,&contacts)
	return contacts,err 
}

func SaveContacts(contacts []models.Contact) error{
	data,err := json.MarshalIndent(contacts,""," ")
	if err != nil{
		return err 
	}
	return os.WriteFile(FileName,data,0644)
}

