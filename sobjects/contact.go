package sobjects

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"regexp"

	"github.com/grokify/gotilla/io/ioutilmore"
	"github.com/grokify/gotilla/net/httputilmore"
)

type ContactSet struct {
	IdSet      IdSet              `xml:"-"`
	Records    []Contact          `json:"records,omitempty" xml:"records"`
	RecordsMap map[string]Contact `xml:"-"`
}

func NewContactSet() ContactSet {
	set := ContactSet{
		IdSet:      NewIdSet(),
		Records:    []Contact{},
		RecordsMap: map[string]Contact{}}
	return set
}

func NewContactSetSetFromXml(bytes []byte) (ContactSet, error) {
	set := ContactSet{IdSet: NewIdSet()}
	err := xml.Unmarshal(bytes, &set)
	set.Inflate()
	return set, err
}

func NewContactSetFromXmlFile(filepath string) (ContactSet, error) {
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return ContactSet{}, err
	}
	return NewContactSetSetFromXml(bytes)
}

func NewContactSetFromJSONResponse(resp *http.Response) (ContactSet, error) {
	set := NewContactSet()
	bytes, err := httputilmore.ResponseBody(resp)
	if err != nil {
		return set, err
	}
	err = json.Unmarshal(bytes, &set)
	return set, err
}

func (set *ContactSet) ReadJsonFilesFromDir(dir string) error {
	files, err := ioutilmore.DirEntriesReSizeGt0(dir, regexp.MustCompile(`(?i)\.json$`))
	if err != nil {
		return err
	}
	for _, fi := range files {
		filepath := path.Join(dir, fi.Name())
		contact, err := NewContactFromJsonFile(filepath)
		if err == nil && len(contact.Id) > 0 {
			set.Records = append(set.Records, contact)
		}
	}
	return nil
}

func (set *ContactSet) Inflate() {
	for _, record := range set.Records {
		if len(record.Id) > 0 {
			set.IdSet.AddId(record.Id)
			set.RecordsMap[record.Id] = record
		}
		if len(record.AccountId) > 0 {
			set.IdSet.AddId(record.AccountId)
		}
	}
}

func (set *ContactSet) GetContactByName(name string) (Contact, error) {
	for _, contact := range set.Records {
		if contact.Name == name {
			return contact, nil
		}
	}
	return Contact{}, errors.New(fmt.Sprintf("Could not found Contact by name [%v]", name))
}

func (set *ContactSet) GetContactById(id string) (Contact, error) {
	for _, contact := range set.Records {
		if contact.Id == id {
			return contact, nil
		}
	}
	return Contact{}, errors.New(fmt.Sprintf("Could not found Contact by id [%v]", id))
}

type Contact struct {
	Id         string
	AccountId  string
	Department string
	Email      string
	Fax        string
	FirstName  string
	LastName   string
	Name       string
}

func NewContactFromJson(bytes []byte) (Contact, error) {
	obj := Contact{}
	err := json.Unmarshal(bytes, &obj)
	return obj, err
}

func NewContactFromJsonFile(filepath string) (Contact, error) {
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return Contact{}, err
	}
	return NewContactFromJson(bytes)
}
