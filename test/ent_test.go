package test

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/khozaei/healthio/ent"
	"github.com/khozaei/healthio/ent/enttest"
	"github.com/khozaei/healthio/ent/patient"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

func TestSavePatient(t *testing.T) {
	ctx := context.Background()
	var savedPatient *ent.Patient
	tests := []struct {
		name   string                        // name of the test
		before func(*testing.T, *ent.Client) // work before test
		expect func(*testing.T, *ent.Client) // actual exceptions
	}{
		{
			name: "create/patient/already_exist/NationalCode",
			before: func(t *testing.T, client *ent.Client) {
				_, err := client.Patient.Create().SetFirstName("gholi").
					SetNationalCode("123").Save(ctx)
				require.NoError(t, err)
			},
			expect: func(t *testing.T, client *ent.Client) {
				_, err := client.Patient.Create().SetFirstName("gholi").
					SetNationalCode("123").Save(ctx)
				require.Error(t, err)
			},
		},
		{
			name: "create/patient/already_exist/IdentityCode",
			before: func(t *testing.T, client *ent.Client) {
				_, err := client.Patient.Create().SetFirstName("gholi").
					SetIdentityCode("123").Save(ctx)
				require.NoError(t, err)
			},
			expect: func(t *testing.T, client *ent.Client) {
				_, err := client.Patient.Create().SetFirstName("gholi").
					SetIdentityCode("123").Save(ctx)
				require.Error(t, err)
			},
		},
		{
			name: "update/patient/FirstName",
			before: func(t *testing.T, client *ent.Client) {
				var err error
				savedPatient, err = client.Patient.Create().
					SetFirstName("gholi").Save(ctx)
				require.NoError(t, err)
			},
			expect: func(t *testing.T, client *ent.Client) {
				updated, err := savedPatient.Update().SetFirstName("Ali").
					Save(ctx)
				if err != nil {
					require.NoError(t, err)
				}
				if updated.FirstName != "Ali" {
					err := errors.New("firstname not updated")
					require.NoError(t, err)
				}
				require.NoError(t, nil)
			},
		},
		{
			name: "update/patient/FirstName/by-id",
			before: func(t *testing.T, client *ent.Client) {
				var err error
				savedPatient, err = client.Patient.Create().
					SetFirstName("gholi").Save(ctx)
				require.NoError(t, err)
			},
			expect: func(t *testing.T, client *ent.Client) {
				updated, err := client.Patient.UpdateOneID(savedPatient.ID).
					SetFirstName("Ali").Save(ctx)
				if err != nil {
					require.NoError(t, err)
				}
				if updated.FirstName != "Ali" {
					err := errors.New("firstname not updated")
					require.NoError(t, err)
				}
				require.NoError(t, nil)
			},
		},
		{
			name: "select/patient/FirstName",
			before: func(t *testing.T, client *ent.Client) {
				var err error
				savedPatient, err = client.Patient.Create().
					SetFirstName("amin").Save(ctx)
				require.NoError(t, err)
			},
			expect: func(t *testing.T, client *ent.Client) {
				patients, err := client.Patient.Query().Where(
					patient.FirstNameContains("min")).All(ctx)
				if err != nil {
					require.NoError(t, err)
				}
				for _, p := range patients {
					if !strings.Contains(p.FirstName, "min") {
						err := errors.New("contains not working properly")
						require.NoError(t, err)
					}
				}
				require.NoError(t, nil)
			},
		},
		{
			name: "delete/patient/id",
			before: func(t *testing.T, client *ent.Client) {
				var err error
				savedPatient, err = client.Patient.Create().
					SetFirstName("azar").Save(ctx)
				require.NoError(t, err)
			},
			expect: func(t *testing.T, client *ent.Client) {
				err := client.Patient.DeleteOneID(savedPatient.ID).
					Exec(ctx)
				require.NoError(t, err)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
			defer client.Close()
			tt.before(t, client)
			tt.expect(t, client)
		})
	}
}
