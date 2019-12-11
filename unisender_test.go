package unisender_test

import (
	"github.com/alexeyco/unisender/contacts"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/alexeyco/unisender"
	"github.com/alexeyco/unisender/test"
)

func TestUniSender_ApiKey(t *testing.T) {
	apiKeyExpected := test.RandomString(12, 36)
	var apiKeyRequested string

	c := test.NewClient(func(req *http.Request) (res *http.Response, err error) {
		apiKeyRequested = req.FormValue("api_key")

		res = &http.Response{
			StatusCode: http.StatusOK,
		}

		return
	})

	usndr := unisender.New(apiKeyExpected)
	usndr.SetClient(c)

	err := usndr.DeleteList(123).Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if apiKeyExpected != apiKeyRequested {
		t.Fatalf(`API key should be "%s", "%s" given`, apiKeyExpected, apiKeyRequested)
	}
}

func TestUniSender_SetLanguageEnglish(t *testing.T) {
	expectedLanguage := "en"
	var givenLanguage string

	c := test.NewClient(func(req *http.Request) (res *http.Response, err error) {
		givenLanguage = req.FormValue("lang")

		res = &http.Response{
			StatusCode: http.StatusOK,
		}

		return
	})

	usndr := unisender.New(test.RandomString(12, 36))
	usndr.SetLanguageEnglish()
	usndr.SetClient(c)

	err := usndr.DeleteList(123).Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedLanguage != givenLanguage {
		t.Fatalf(`Language should be "%s", "%s" given`, expectedLanguage, givenLanguage)
	}
}

func TestUniSender_SetLanguageItalian(t *testing.T) {
	expectedLanguage := "it"
	var givenLanguage string

	c := test.NewClient(func(req *http.Request) (res *http.Response, err error) {
		givenLanguage = req.FormValue("lang")

		res = &http.Response{
			StatusCode: http.StatusOK,
		}

		return
	})

	usndr := unisender.New(test.RandomString(12, 36))
	usndr.SetLanguageItalian()
	usndr.SetClient(c)

	err := usndr.DeleteList(123).Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedLanguage != givenLanguage {
		t.Fatalf(`Language should be "%s", "%s" given`, expectedLanguage, givenLanguage)
	}
}

func TestUniSender_SetLanguageRussian(t *testing.T) {
	expectedLanguage := "ru"
	var givenLanguage string

	c := test.NewClient(func(req *http.Request) (res *http.Response, err error) {
		givenLanguage = req.FormValue("lang")

		res = &http.Response{
			StatusCode: http.StatusOK,
		}

		return
	})

	usndr := unisender.New(test.RandomString(12, 36))
	usndr.SetLanguageRussian()
	usndr.SetClient(c)

	err := usndr.DeleteList(123).Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedLanguage != givenLanguage {
		t.Fatalf(`Language should be "%s", "%s" given`, expectedLanguage, givenLanguage)
	}
}

func TestUniSender_Format(t *testing.T) {
	apiKeyExpected := test.RandomString(12, 36)

	formatExpected := "json"
	var formatRequested string

	c := test.NewClient(func(req *http.Request) (res *http.Response, err error) {
		formatRequested = req.FormValue("format")

		res = &http.Response{
			StatusCode: http.StatusOK,
		}

		return
	})

	usndr := unisender.New(apiKeyExpected)
	usndr.SetClient(c)

	err := usndr.DeleteList(123).Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if formatExpected != formatRequested {
		t.Fatalf(`Format should be "%s", "%s" given`, formatExpected, formatRequested)
	}
}

func ExampleUniSender_GetCurrencyRates() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	currencyRates, err := usndr.GetCurrencyRates().
		Execute()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(currencyRates)
}

func ExampleUniSender_CancelCampaign() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	var campaignID int64 = 123
	if err := usndr.CancelCampaign(123).Execute(); err != nil {
		log.Fatalln(err)
	}

	log.Printf("Campaign (id=%d) cancelled", campaignID)
}

func ExampleUniSender_CreateCampaign() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	var messageID int64 = 123
	res, err := usndr.CreateCampaign(messageID).
		StartTime(time.Now().Add(3 * time.Hour)).
		TrackRead().
		TrackLinks().
		Execute()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
}

func ExampleUniSender_GetCampaignCommonStats() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	var campaignID int64 = 123
	res, err := usndr.GetCampaignCommonStats(campaignID).
		Execute()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
}

func ExampleUniSender_GetCampaigns() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	res, err := usndr.GetCampaigns().
		From(time.Now().Add(-time.Hour)).
		To(time.Now()).
		Limit(100).
		Offset(0).
		Execute()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
}

func ExampleUniSender_GetVisitedLinks() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	var campaignID int64 = 123
	res, err := usndr.GetVisitedLinks(campaignID).
		Execute()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
}

func ExampleUniSender_CreateField() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	fieldID, err := usndr.CreateField("SomeField").
		TypeString().
		Execute()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(fieldID)
}

func ExampleUniSender_DeleteField() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	var fieldID int64 = 123
	err := usndr.DeleteField(fieldID).
		Execute()

	if err != nil {
		log.Fatalln(err)
	}
}

func ExampleUniSender_DeleteTag() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	var tagID int64 = 123
	err := usndr.DeleteTag(tagID).
		Execute()

	if err != nil {
		log.Fatalln(err)
	}
}

func ExampleUniSender_Exclude() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	email := "foo@bar.example"
	err := usndr.Exclude(email).
		ContactTypeEmail().
		ListIDs(1, 2, 3).
		Execute()

	if err != nil {
		log.Fatalln(err)
	}
}

func ExampleUniSender_ExportContacts() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	res, err := usndr.ExportContacts().
		EmailStatusBlocked().
		NotifyUrl("https://foo.bar/example").
		Execute()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
}

func ExampleUniSender_GetContact() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	res, err := usndr.GetContact("foo@bar.example").
		IncludeLists().
		IncludeFields().
		IncludeDetails().
		Execute()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
}

func ExampleUniSender_GetContactCount() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	res, err := usndr.GetContactCount(1).
		ParamsTypeAddress().
		Execute()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
}

func ExampleUniSender_GetContactFieldValues() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	res, err := usndr.GetContactFieldValues("foo@bar.example", 1, 2, 3).
		Execute()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
}

func ExampleUniSender_GetFields() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	res, err := usndr.GetFields().
		Execute()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
}

func ExampleUniSender_GetTags() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	res, err := usndr.GetTags().
		Execute()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
}

func ExampleUniSender_GetTotalContactsCount() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	res, err := usndr.GetTotalContactsCount("my-login").
		Execute()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
}

func ExampleUniSender_ImportContacts() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	collection := contacts.NewImportContactsCollection()
	collection.Email("foo1@bar.example").
		AddListID(1, time.Now()).
		SetAvailabilityAvailable().
		SetStatusActive()

	collection.Email("foo2@bar.example").
		Delete()

	res, err := usndr.ImportContacts(collection).
		Execute()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
}

func ExampleUniSender_IsContactInList() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	res, err := usndr.IsContactInList("foo@bar.example", 1, 2, 3).
		ConditionOr().
		Execute()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
}

func ExampleUniSender_Subscribe() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	res, err := usndr.Subscribe().
		Email("foo@bar.example").
		Tags("foo", "bar", "example").
		Execute()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
}

func ExampleUniSender_Unsubscribe() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	err := usndr.Unsubscribe("foo@bar.example").
		ContactTypeEmail().
		Execute()

	if err != nil {
		log.Fatalln(err)
	}
}

func ExampleUniSender_CheckEmail() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	res, err := usndr.CheckEmail(123).
		Execute()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
}

func ExampleUniSender_CheckSMS() {
	usndr := unisender.New("your-api-key").
		SetLanguageEnglish()

	res, err := usndr.CheckSMS(123).
		Execute()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
}
