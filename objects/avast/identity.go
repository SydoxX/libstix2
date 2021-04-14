package avast

import (
	"time"

	"github.com/avast/libstix2/stixid"

	"github.com/avast/libstix2/objects/identity"
	"github.com/google/uuid"
)

const OrganizationNane = "Avast Software s.r.o."

var (
	Namespace = uuid.MustParse("e8f01b05-e1c1-48f9-a53d-5bcd2e822cab")

	IdentityUUID    = uuid.NewSHA1(Namespace, []byte(OrganizationNane))
	IdentityId      = stixid.BuildId("identity", IdentityUUID)
	IdentityCreated = time.Date(2020, 10, 20, 0, 0, 0, 0, time.UTC)
)

func Identity() *identity.Identity {
	idt := identity.New()
	idt.SetID(IdentityId)
	idt.SetCreatedByRef(IdentityId)
	idt.SetCreated(IdentityCreated)
	idt.SetModified(IdentityCreated)
	idt.SetName(OrganizationNane)
	idt.SetIdentityClass("organization")
	return idt
}
