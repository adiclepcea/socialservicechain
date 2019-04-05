package sscstate

import "fmt"

//MemoryStateStorage implements a simple in memory state storage
type MemoryStateStorage struct {
	StateStorage
}

//GetNGO returns the NGO by name
func (mss *MemoryStateStorage) GetNGO(name string) (*NGO, error) {
	return nil, fmt.Errorf("Not implemented yet")
}

//GetNGOByID returns the NGO by id
func (mss *MemoryStateStorage) GetNGOByID(id int64) (*NGO, error) {
	return nil, fmt.Errorf("Not implemented yet")
}

//SaveNGO will try to save the NGO
func (mss *MemoryStateStorage) SaveNGO(*NGO) error {
	return fmt.Errorf("Not implemented yet")
}

//DeleteNGO will delete the NGO
func (mss *MemoryStateStorage) DeleteNGO(NGO) error {
	return fmt.Errorf("Not implemented yet")
}

//GetSocialCase will return the social case by name
func (mss *MemoryStateStorage) GetSocialCase(name string) (*SocialCase, error) {
	return nil, fmt.Errorf("Not implemented yet")
}

//GetSocialCaseByID will return the social case by id
func (mss *MemoryStateStorage) GetSocialCaseByID(id int64) (*SocialCase, error) {
	return nil, fmt.Errorf("Not implemented yet")
}

//SaveSocialCase will save the social case
func (mss *MemoryStateStorage) SaveSocialCase(*SocialCase) error {
	return fmt.Errorf("Not implemented yet")
}

//DeleteSocialCase will delete the social case
func (mss *MemoryStateStorage) DeleteSocialCase(*SocialCase) error {
	return fmt.Errorf("Not implemented yet")
}

//GetDonor will return a donor by name
func (mss *MemoryStateStorage) GetDonor(name string) (*Donor, error) {
	return nil, fmt.Errorf("Not implemented yet")
}

//GetDonorByID will return a donor by id
func (mss *MemoryStateStorage) GetDonorByID(id int64) (*Donor, error) {
	return nil, fmt.Errorf("Not implemented yet")
}

//SaveDonor will save the passed in donor into memory
func (mss *MemoryStateStorage) SaveDonor(*Donor) error {
	return fmt.Errorf("Not implemented yet")
}

//DeleteDonor will remove the passed in donor
func (mss *MemoryStateStorage) DeleteDonor(*Donor) error {
	return fmt.Errorf("Not implemented yet")
}

//GetDonationByID will return a donation by id
func (mss *MemoryStateStorage) GetDonationByID(id int64) (*Donation, error) {
	return nil, fmt.Errorf("Not implemented yet")
}

//SaveDonation will save the passed in donation
func (mss *MemoryStateStorage) SaveDonation(*Donation) error {
	return fmt.Errorf("Not implemented yet")
}

//DeleteDonation will delete a donation
func (mss *MemoryStateStorage) DeleteDonation(Donation) error {
	return fmt.Errorf("Not implemented yet")
}

//GetDonationAssignmentByID will return the donation with the given ID
func (mss *MemoryStateStorage) GetDonationAssignmentByID(id int64) (*DonationAssignment, error) {
	return nil, fmt.Errorf("Not implemented yet")
}

//SaveDonationAssignment will save a donation assignement
func (mss *MemoryStateStorage) SaveDonationAssignment(*DonationAssignment) error {
	return fmt.Errorf("Not implemented yet")
}
