// Code generated by ent, DO NOT EDIT.

package ent

import (
	"go-entgo/ent/card"
	"go-entgo/ent/group"
	"go-entgo/ent/pet"
	"go-entgo/ent/user"
	"go-entgo/schema"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	cardFields := schema.Card{}.Fields()
	_ = cardFields
	// cardDescExpired is the schema descriptor for expired field.
	cardDescExpired := cardFields[0].Descriptor()
	// card.DefaultExpired holds the default value on creation for the expired field.
	card.DefaultExpired = cardDescExpired.Default.(time.Time)
	groupFields := schema.Group{}.Fields()
	_ = groupFields
	// groupDescAge is the schema descriptor for age field.
	groupDescAge := groupFields[0].Descriptor()
	// group.AgeValidator is a validator for the "age" field. It is called by the builders before save.
	group.AgeValidator = groupDescAge.Validators[0].(func(int) error)
	// groupDescName is the schema descriptor for name field.
	groupDescName := groupFields[1].Descriptor()
	// group.DefaultName holds the default value on creation for the name field.
	group.DefaultName = groupDescName.Default.(string)
	petFields := schema.Pet{}.Fields()
	_ = petFields
	// petDescAge is the schema descriptor for age field.
	petDescAge := petFields[0].Descriptor()
	// pet.AgeValidator is a validator for the "age" field. It is called by the builders before save.
	pet.AgeValidator = petDescAge.Validators[0].(func(int) error)
	// petDescName is the schema descriptor for name field.
	petDescName := petFields[1].Descriptor()
	// pet.DefaultName holds the default value on creation for the name field.
	pet.DefaultName = petDescName.Default.(string)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescAge is the schema descriptor for age field.
	userDescAge := userFields[0].Descriptor()
	// user.AgeValidator is a validator for the "age" field. It is called by the builders before save.
	user.AgeValidator = userDescAge.Validators[0].(func(int) error)
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[1].Descriptor()
	// user.DefaultName holds the default value on creation for the name field.
	user.DefaultName = userDescName.Default.(string)
}
