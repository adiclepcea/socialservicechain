package main

const (
	//FamilyName is the name of the transaction family for this client
	FamilyName string = "ssc"
	//FamilyVersion is the version of the transaction family for this client
	FamilyVersion string = "1.0"
	//DistributionName is the name of this client
	DistributionName string = "sawtooth-ssc"
	//DefaultURL will be used if no specific url is given
	DefaultURL string = "http://127.0.0.1:8008"

	////Verbs

	//VerbCreateNGO is the name of the action to be used for creating a NGO
	VerbCreateNGO string = "createNGO"
	//VerbCreateSocialCase is the name of the action to be used for creating a SocialCase
	VerbCreateSocialCase string = "createSocialCase"
	//VerbCreateDonor is the name of the action to be used for creating a Donor
	VerbCreateDonor string = "createDonor"
	//VerbDeleteNGO is the name of the action to be used for deleting a NGO
	VerbDeleteNGO string = "deleteNGO"
	//VerbDeleteSocialCase is the name of the action to be used for deleting a SocialCase
	VerbDeleteSocialCase string = "deleteSocialCase"
	//VerbDeleteDonor is the name of the action to be used for deleting a Donor
	VerbDeleteDonor string = "deleteDonor"
	//VerbDonate is the name of the action to be used for causing a donor to donate
	VerbDonate string = "donate"
	//VerbDoHelp is the name of the action to be used for allocating a donation to a social case
	VerbDoHelp string = "doHelp"

	//APIs

	//BatchSubmitAPI is the validator endpoint to call for batch transaction submission
	BatchSubmitAPI string = "batches"
	//BatchStatusAPI is the validator endpoint to call for a batch status check
	BatchStatusAPI string = "batch_statuses"
	//StateAPI is the validator endpoint for checking for state
	StateAPI string = "state"

	// Content types

	//ContentTypeOctetStream is used to represent an octet stream/binary transaction over http
	ContentTypeOctetStream string = "application/octet-stream"

	//Family related variables

	//FamilyNamespaceAddressLen represents the length of the family address
	FamilyNamespaceAddressLen uint = 6
	//FamilyVerbAddressLen is the length of the verb as it be encoded
	FamilyVerbAddressLen uint = 64
)
