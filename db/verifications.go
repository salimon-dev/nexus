package db

import (
	"math/rand"
	"salimon/nexus/types"
	"time"

	"github.com/google/uuid"
)

func generateRandomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, n)
	for i := range result {
		letter := letters[rand.Int()%len(letters)]
		result[i] = letter
	}
	return string(result)
}

func InserVerification(userId string, vType types.VerificationType, domain types.VerificationDomain, expiresAt time.Time) (*types.Verification, error) {
	token := generateRandomString(16)
	query := "INSERT INTO verifications (id, user_id, type, domain, token, expires_at) VALUES ($1, $2, $3, $4, $5, $6)"
	id := uuid.New().String()

	_, err := DB.Exec(query, id, userId, vType, domain, token, expiresAt)
	if err != nil {
		return nil, err
	}
	verification := types.Verification{
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
func GetVerificationRecord(token string) (*types.Verification, error) {
	query := "SELECT * FROM verifications WHERE token = $1 AND expires_at < $2"
	rows, err := DB.Query(query, token, time.Now())
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var verification types.Verification
	if rows.Next() {
		err := rows.Scan(&verification.Id, &verification.UserId, &verification.Type, &verification.Domain, &verification.Token, &verification.ExpiresAt)
		if err != nil {
			return nil, err
		}
		return &verification, nil
	}
	return nil, nil
}

func InsertRegisterEmailVerification(userId string) (*types.Verification, error) {
	return InserVerification(userId, types.VerificationTypeEmail, types.VerificationDomainRegister, time.Now().Add(time.Hour*24))
}

func InsertPasswordResetVerification(userId string) (*types.Verification, error) {
	return InserVerification(userId, types.VerificationTypeEmail, types.VerificationDomainPasswordReset, time.Now().Add(time.Hour*24))
}

func InsertEmailUpdateVerification(userId string) (*types.Verification, error) {
	return InserVerification(userId, types.VerificationTypeEmail, types.VerificationDomainEmailUpdate, time.Now().Add(time.Hour*24))
}

func DeleteVerification(verification *types.Verification) error {
	query := "DELETE FROM verifications WHERE id=$1"
	_, err := DB.Exec(query, verification.Id)
	return err
}
