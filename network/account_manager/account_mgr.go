package account_manager

import (
	"bnet-mock/network/client/authentication_service"
	"bnet-mock/network/client/entity_types"
	"crypto/rand"
	"errors"

	"google.golang.org/protobuf/proto"
)

type AccountMgr struct{}

// Returns error determining if the given credentials are valid and whether the account is restricted
func (m *AccountMgr) Authenticate(email, password string) error {
	logonResult, err := m.GetAccountByEmail(email)
	if err != nil {
		return err
	}

	if *logonResult.RestrictedMode {
		return errors.New("restricted account. Geoip country US, and email test@restricted.com are banned")
	}
	if email == "test@test.com" && password == "test" {
		return nil // Authentication successful
	}
	return errors.New("bad password, valid email is " + email + " and valid password is test")
}

// Returns data for a specific (currently hardcoded) email
func (m *AccountMgr) GetAccountByEmail(email string) (*authentication_service.LogonResult, error) {
	if email == "test@test.com" {
		high1 := uint64(72057594037927936)
		low1 := uint64(379775058)
		high2 := uint64(144115192376095343)
		low2 := uint64(559865145)

		return &authentication_service.LogonResult{
			ErrorCode: proto.Uint32(0),
			AccountId: &entity_types.EntityId{
				High: &high1,
				Low:  &low1,
			},
			GameAccountId: []*entity_types.EntityId{
				{
					High: &high2,
					Low:  &low2,
				},
			},
			Email:           proto.String(email),
			AvailableRegion: []uint32{1},
			ConnectedRegion: proto.Uint32(1),
			BattleTag:       proto.String("Test#0001"),
			GeoipCountry:    proto.String("US"),
			SessionKey:      GenerateSessionKey(), // []byte{0xAB, 0xA8, 0xF2, 0x94, 0x99, 0x77, 0x65, 0x55, 0x58, 0x6B, 0xEB, 0xC1, 0xEE, 0xF0, 0xB4, 0xCE, 0x24, 0xCC, 0xAF, 0xDD, 0x53, 0x37, 0xF7, 0x78, 0x62, 0xC1, 0xA0, 0x54, 0xF9, 0xA5, 0x36, 0x15, 0x0D, 0xE6, 0x78, 0xFD, 0x16, 0x49, 0x39, 0xDC, 0xE7, 0xB4, 0x48, 0x5B, 0xA2, 0xC4, 0x73, 0x05, 0x44, 0xB2, 0x00, 0xAD, 0xA8, 0x90, 0x6E, 0x74, 0x15, 0xCC, 0xF2, 0x65, 0x73, 0x21, 0x23, 0x99},
			RestrictedMode:  proto.Bool(IsUserRestricted(email, "EU")),
		}, nil
	} else if email == "test@restricted.com" {
		high1 := uint64(72057594037927937)
		low1 := uint64(379775058)
		high2 := uint64(144115192376095344)
		low2 := uint64(559865145)

		return &authentication_service.LogonResult{
			ErrorCode: proto.Uint32(0),
			AccountId: &entity_types.EntityId{
				High: &high1,
				Low:  &low1,
			},
			GameAccountId: []*entity_types.EntityId{
				{
					High: &high2,
					Low:  &low2,
				},
			},
			Email:           proto.String(email),
			AvailableRegion: []uint32{1},
			ConnectedRegion: proto.Uint32(1),
			BattleTag:       proto.String("Test#0001"),
			GeoipCountry:    proto.String("US"),
			SessionKey:      GenerateSessionKey(), // []byte{0xAB, 0xA8, 0xF2, 0x94, 0x99, 0x77, 0x65, 0x55, 0x58, 0x6B, 0xEB, 0xC1, 0xEE, 0xF0, 0xB4, 0xCE, 0x24, 0xCC, 0xAF, 0xDD, 0x53, 0x37, 0xF7, 0x78, 0x62, 0xC1, 0xA0, 0x54, 0xF9, 0xA5, 0x36, 0x15, 0x0D, 0xE6, 0x78, 0xFD, 0x16, 0x49, 0x39, 0xDC, 0xE7, 0xB4, 0x48, 0x5B, 0xA2, 0xC4, 0x73, 0x05, 0x44, 0xB2, 0x00, 0xAD, 0xA8, 0x90, 0x6E, 0x74, 0x15, 0xCC, 0xF2, 0x65, 0x73, 0x21, 0x23, 0x99},
			RestrictedMode:  proto.Bool(IsUserRestricted(email, "US")),
		}, nil
	}
	return nil, errors.New("account not found")
}

// Returns a 64-byte session key
func GenerateSessionKey() []byte {
	key := make([]byte, 64)
	_, _ = rand.Read(key)
	return key
}

// Returns whether user should be restricted
// Example: Restrict all users from a specific country or specific email
func IsUserRestricted(userEmail string, geoipCountry string) bool {
	restrictedCountries := map[string]bool{
		"US": true, // Explicitly restricted
	}
	return restrictedCountries[geoipCountry] || userEmail == "test@restricted.com"
}
