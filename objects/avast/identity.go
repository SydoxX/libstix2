package avast

import (
	"time"

	"github.com/avast/libstix2/objects"

	"github.com/avast/libstix2/vocabs"

	"github.com/avast/libstix2/datatypes/timestamp"

	"github.com/avast/libstix2/datatypes/stixid"

	"github.com/avast/libstix2/objects/identity"
	"github.com/google/uuid"
)

const (
	OrganizationNane = "Avast Software s.r.o."

	FeedVulnerabilityName = "Avast Vulnerability Feed"
)

var (
	Namespace = uuid.MustParse("e8f01b05-e1c1-48f9-a53d-5bcd2e822cab")

	IdentityUUID    = uuid.NewSHA1(Namespace, []byte(OrganizationNane))
	IdentityId      = stixid.BuildId(objects.TypeIdentity, IdentityUUID)
	IdentityCreated = time.Date(2020, 10, 20, 0, 0, 0, 0, time.UTC)

	FeedVulnerabilityUUID    = uuid.NewSHA1(Namespace, []byte(FeedVulnerabilityName))
	FeedVulnerabilityId      = stixid.BuildId(objects.TypeIdentity, FeedVulnerabilityUUID)
	FeedVulnerabilityCreated = time.Date(2021, 04, 23, 0, 0, 0, 0, time.UTC)
)

func Identity() *identity.Identity {
	idt := identity.New()
	idt.ID = IdentityId
	idt.CreatedByRef = IdentityId
	idt.Created = timestamp.NewPtr(IdentityCreated)
	idt.Modified = timestamp.NewPtr(IdentityCreated)
	idt.Name = OrganizationNane
	idt.IdentityClass = vocabs.IdentityOrganization
	return idt
}

func FeedVulnerability() *identity.Identity {
	idt := identity.New()
	idt.ID = FeedVulnerabilityId
	idt.CreatedByRef = IdentityId
	idt.Created = timestamp.NewPtr(FeedVulnerabilityCreated)
	idt.Modified = timestamp.NewPtr(FeedVulnerabilityCreated)
	idt.Name = FeedVulnerabilityName
	idt.IdentityClass = "feed"
	return idt
}
