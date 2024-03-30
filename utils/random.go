package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var (
	alphabet     = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	firstNames   = []string{"John", "Michael", "Emily", "Sarah", "David", "Jessica", "Matthew", "Jennifer", "Christopher", "Linda"}
	lastNames    = []string{"Smith", "Johnson", "Brown", "Jones", "Garcia", "Miller", "Davis", "Rodriguez", "Martinez", "Wilson"}
	emailDomains = []string{"gmail.com", "yahoo.com", "hotmail.com", "outlook.com", "icloud.com"}
	categoryType = []string{"debit"}
	minDate      = time.Date(1900, time.January, 1, 0, 0, 0, 0, time.UTC)
	maxDate      = time.Date(2010, time.December, 31, 23, 59, 59, 0, time.UTC)
)

func RandomString(length int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < length; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomName() (string, string) {
	firstName := firstNames[rand.Intn(len(firstNames))]
	lastName := lastNames[rand.Intn(len(lastNames))]
	return firstName, lastName
}

func RandomEmail(firstName, lastName string) string {
	return fmt.Sprintf("%s.%s@%s",
		strings.ToLower(firstName),
		strings.ToLower(lastName),
		emailDomains[rand.Intn(len(emailDomains))],
	)
}

func RandomBirthDate() time.Time {
	minUnix := minDate.Unix()
	maxUnix := maxDate.Unix()

	randomUnix := rand.Int63n(maxUnix-minUnix) + minUnix
	randomBirth := time.Unix(randomUnix, 0)

	return randomBirth
}

func RandomCategoryType() string {
	categoryType := categoryType[rand.Intn(len(categoryType))]
	return categoryType
}

func RandomAccountValue() int32 {
	value := rand.Intn(100)
	return int32(value)
}
