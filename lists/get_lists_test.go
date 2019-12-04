package lists_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"reflect"
	"testing"

	"github.com/alexeyco/unisender/api"
	"github.com/alexeyco/unisender/lists"
)

func TestGetListsRequest_Execute(t *testing.T) {
	expectedLists := getRandomListsSlice()
	j := listsSliceToJson(expectedLists)

	req := newRequest(func(req *http.Request) (res *http.Response, err error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBufferString(j)),
		}, nil
	})

	givenLists, err := lists.GetLists(req).Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if len(givenLists) != len(expectedLists) {
		t.Fatalf(`Lists slice should have length %d, %d given`, len(expectedLists), len(givenLists))
	}

	if !reflect.DeepEqual(expectedLists, givenLists) {
		t.Fatal("Expected and given lists should be equal")
	}
}

func getRandomListsSlice() (slice []lists.List) {
	l := randomInt(12, 36)
	for i := 0; i < l; i++ {
		slice = append(slice, lists.List{
			ID:    int64(randomInt(9999, 999999)),
			Title: fmt.Sprintf("Title #%d", randomInt(9999, 999999)),
		})
	}

	return
}

func listsSliceToJson(slice []lists.List) string {
	b, _ := json.Marshal(&api.Response{
		Result: slice,
	})

	return string(b)
}

func randomInt(min, max int) int {
	return rand.Intn(max-min) + min
}
