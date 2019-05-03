package test

import (
	"github.com/mvmaasakkers/certificates/database"
	"testing"
)

var certificateTests = []struct {
	Id          string
	Error       error
	Certificate *database.Certificate
}{
	{
		Id:          "test.id",
		Error:       nil,
		Certificate: &database.Certificate{CommonName: "test.id", SerialNumber: "testserial"},
	},
	{
		Id:          "testnotfound",
		Error:       database.ErrorObjectNotFound,
		Certificate: &database.Certificate{},
	},
}

func TestCertificate_Certificate(t *testing.T, certificateRepository database.CertificateRepository) {
	for _, test := range certificateTests {
		_, err := certificateRepository.GetBySerialNumber(test.Certificate.SerialNumber)
		if err != test.Error {
			t.Errorf("%s: expected error %+v, got error %+v", test.Id, test.Error, err)
			t.Fail()
		}
	}
}

var createCertificateTests = []struct {
	Error       error
	DeleteError error
	Certificate *database.Certificate
}{
	{
		Certificate: &database.Certificate{CommonName: "test.id", SerialNumber: "testserial"},
		Error:       database.ErrorDuplicateObject,
	},
	{
		Certificate: &database.Certificate{CommonName: "testid_2", SerialNumber: "two"},
		Error:       nil,
	},
}

func TestCertificate_CreateCertificate(t *testing.T, certificateRepository database.CertificateRepository) {
	for _, test := range createCertificateTests {
		err := certificateRepository.Create(test.Certificate)
		if err != test.Error {
			t.Errorf("%s: expected error %+v, got error %+v", test.Certificate.SerialNumber, test.Error, err)
			t.Fail()
		}

		if err := certificateRepository.DeleteBySerialNumber(test.Certificate.SerialNumber); err != test.DeleteError {
			t.Errorf("%s: expected delete error %+v, got error %+v", test.Certificate.SerialNumber, test.Error, err)
			t.Fail()
		}
	}
}
