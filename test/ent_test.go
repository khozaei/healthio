package test

import (
	"context"
	"testing"

	"github.com/khozaei/healthio/ent"
	"github.com/khozaei/healthio/ent/enttest"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

func TestSavePatient(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name   string                        // name of the test
		before func(*testing.T, *ent.Client) // work before test
		expect func(*testing.T, *ent.Client) // actual exceptions
	}{
		{
			name: "create/already_exist/NationalCode",
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
			name: "create/already_exist/IdentityCode",
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
