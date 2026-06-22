package db

import (
    "strings"
    "testing"

    "github.com/stretchr/testify/assert"
)

// Compile-time type assertion — no runtime test needed.
var _ DBName = DBName("").CompanyDBName("")

func TestCompanyDBName(t *testing.T) {
    t.Parallel()

    tests := []struct {
        name        string
        dbName      DBName
        companyCode string
        expected    DBName
    }{
        // Valid company code
        {
            name:        "valid company code with IAMDBName",
            dbName:      IAMDBName,
            companyCode: "acme",
            expected:    DBName("acme:tagsamurai_iam"),
        },
        {
            name:        "valid company code with AdminDBName",
            dbName:      AdminDBName,
            companyCode: "corp",
            expected:    DBName("corp:admin_tagsamurai"),
        },
        {
            name:        "valid company code with GlobalDBName",
            dbName:      GlobalDBName,
            companyCode: "mycompany",
            expected:    DBName("mycompany:tagsamurai_global"),
        },

        // Empty inputs
        {
            name:        "empty company code produces colon prefix",
            dbName:      IAMDBName,
            companyCode: "",
            expected:    DBName(":tagsamurai_iam"),
        },
        {
            name:        "empty db name produces colon suffix",
            dbName:      DBName(""),
            companyCode: "acme",
            expected:    DBName("acme:"),
        },
        {
            name:        "both empty produces single colon",
            dbName:      DBName(""),
            companyCode: "",
            expected:    DBName(":"),
        },

        // Special characters in company code
        {
            name:        "company code with hyphen",
            dbName:      IAMDBName,
            companyCode: "my-company",
            expected:    DBName("my-company:tagsamurai_iam"),
        },
        {
            name:        "company code with underscore",
            dbName:      IAMDBName,
            companyCode: "my_company",
            expected:    DBName("my_company:tagsamurai_iam"),
        },
        {
            name:        "company code with numbers",
            dbName:      NotificationDBName,
            companyCode: "company123",
            expected:    DBName("company123:tagsamurai_notification"),
        },
        {
            name:        "company code with dots",
            dbName:      FileManagerDBName,
            companyCode: "company.name",
            expected:    DBName("company.name:tagsamurai_file_manager"),
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            t.Parallel()

            result := tt.dbName.CompanyDBName(tt.companyCode)
            assert.Equal(t, tt.expected, result)

            assert.Equal(t, 1, strings.Count(string(result), ":"),
                "result %q should contain exactly one colon separator", result)
        })
    }
}