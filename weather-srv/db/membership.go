package db

import "gorm.io/gorm"

// UpdateMemberShipStatus updates the membership status for the given address in the database.
func UpdateMemberShipStatus(DB *gorm.DB, address string, status MembershipStatus) error {
	return DB.Model(&Membership{}).Where("address = ?", address).
		Update("status", status).Error
}

// FindMemberShip finds the membership with the given address in the database.
func FindMemberShip(DB *gorm.DB, address string) (*Membership, error) {
	membership := &Membership{}
	err := DB.First(membership, "address = ?", address).Error
	if err != nil {
		return nil, err
	}
	return membership, nil
}

// CreateMembership creates a new membership in the database.
func CreateMembership(DB *gorm.DB, membership *Membership) error {
	return DB.Create(membership).Error
}
