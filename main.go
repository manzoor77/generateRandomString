package main

import (
	"crypto/rand"
	"encoding/base32"
	"fmt"
	"strings"

	"github.com/pborman/uuid"
)

var reservedName = []string{
	"admin",
	"api",
	"channel",
	"claim",
	"error",
	"files",
	"help",
	"landing",
	"login",
	"mfa",
	"oauth",
	"plug",
	"plugins",
	"post",
	"signup",
	"boards",
	"playbooks",
}
var encoding = base32.NewEncoding("ybndrfg8ejkmcpqxot1uwisza345h769").WithPadding(base32.NoPadding)

func main() {
	randomstring:=NewRandomString(9)
	fmt.Println("************* randomstring:", randomstring)
	randomstringname:=NewRandomTeamName()
	fmt.Println("************* randomstringname:", randomstringname)
	id:=NewId()
	fmt.Println("************* id:", id)
}

// NewId is a globally unique identifier.  It is a [A-Z0-9] string 26
// characters long.  It is a UUID version 4 Guid that is zbased32 encoded
// without the padding.
func NewId() string {
	return encoding.EncodeToString(uuid.NewRandom())
}

// NewRandomTeamName is a NewId that will be a valid team name.
func NewRandomTeamName() string {
	teamName := NewId()
	for IsReservedTeamName(teamName) {
		teamName = NewId()
	}
	return teamName
}

// NewRandomString returns a random string of the given length.
// The resulting entropy will be (5 * length) bits.
func NewRandomString(length int) string {
	data := make([]byte, 1+(length*5/8))
	rand.Read(data)
	return encoding.EncodeToString(data)[:length]
}

func IsReservedTeamName(s string) bool {
	s = strings.ToLower(s)

	for _, value := range reservedName {
		if strings.Index(s, value) == 0 {
			return true
		}
	}

	return false
}
