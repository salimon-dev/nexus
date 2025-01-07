package types

import "time"

type VerificationType int8

const (
	VerificationTypeEmail VerificationType = 0
	VerificationTypeSMS   VerificationType = 1
)

type VerificationDomain int8

const (
	VerificationDomainRegister       VerificationDomain = 0
	VerificationDomainPasswordReset  VerificationDomain = 1
	VerificationDomainEmailUpdate    VerificationDomain = 2
	VerificationDomainUsernameUpdate VerificationDomain = 3
)

type Verification struct {
	Id        string             `json:"id"`
	UserId    string             `json:"user_id"`
	Type      VerificationType   `json:"type"`
	Domain    VerificationDomain `json:"domain"`
	Token     string             `json:"token"`
	ExpiresAt time.Time          `json:"expires_at"`
}
