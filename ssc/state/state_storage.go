package sscstate

//StateStorage stores the state changes and enables the states reads
type StateStorage interface {
	GetNGO(name string) (*NGO, error)
	GetNGOByID(id int64) (*NGO, error)
	SaveNGO(*NGO) error
	DeleteNGO(NGO) error

	GetSocialCase(name string) (*SocialCase, error)
	GetSocialCaseByID(id int64) (*SocialCase, error)
	SaveSocialCase(*SocialCase) error
	DeleteSocialCase(*SocialCase) error

	GetDonor(name string) (*Donor, error)
	GetDonorByID(id int64) (*Donor, error)
	SaveDonor(*Donor) error
	DeleteDonor(*Donor) error

	GetDonationByID(id int64) (*Donation, error)
	SaveDonation(*Donation) error
	DeleteDonation(Donation) error

	GetDonationAssignmentByID(id int64) (*DonationAssignment, error)
	SaveDonationAssignment(*DonationAssignment) error
}
