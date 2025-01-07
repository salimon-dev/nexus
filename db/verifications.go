package db

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

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

func generateRandomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, n)
	for i := range result {
		letter := letters[rand.Int()%len(letters)]
		result[i] = letter
	}
	return string(result)
}

func InserVerification(userId string, vType VerificationType, domain VerificationDomain, expiresAt time.Time) (*Verification, error) {
	token := generateRandomString(16)
	query := "INSERT INTO verifications (id, user_id, type, domain, token, expires_at) VALUES ($1, $2, $3, $4, $5, $6)"
	id := uuid.New().String()

	_, err := DB.Exec(query, id, userId, vType, domain, token, expiresAt)
	if err != nil {
		return nil, err
	}
	verification := Verification{
		Id:        id,
		UserId:    userId,
		Type:      vType,
		Domain:    domain,
		Token:     token,
		ExpiresAt: expiresAt,
	}
	return &verification, nil
}

// gets active verification record based on token and expire time
func GetVerificationRecord(token string) (*Verification, error) {
	query := "SELECT * FROM verifications WHERE token = $1 AND expires_at < $2"
	rows, err := DB.Query(query, token, time.Now())
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var verification Verification
	if rows.Next() {
		err := rows.Scan(&verification.Id, &verification.UserId, &verification.Type, &verification.Domain, &verification.Token, &verification.ExpiresAt)
		if err != nil {
			return nil, err
		}
		return &verification, nil
	}
	return nil, nil
}

func InsertRegisterEmailVerification(userId string) (*Verification, error) {
	return InserVerification(userId, VerificationTypeEmail, VerificationDomainRegister, time.Now().Add(time.Hour*24))
}
