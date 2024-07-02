package aiven

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UserProfileHandler", func() {
	var (
		client   *Client
		tearDown func()
		ctx      context.Context
		handler  *UserProfileHandler
	)

	BeforeEach(func() {
		client, tearDown = setupUserProfileHandlerTestCase()
		ctx = context.Background()
		handler = &UserProfileHandler{client: client}
	})

	AfterEach(func() {
		tearDown()
	})

	Context("Me", func() {
		It("should return the current user's profile", func() {
			expectedUser := &User{
				Auth:                   []string{"auth1", "auth2"},
				City:                   ptrToString("Helsinki"),
				Country:                "FI",
				CreateTime:             "2022-01-01T00:00:00Z",
				Department:             ptrToString("Engineering"),
				Features:               Features{FreeTierEnabled: true, ReferralEnabled: true, ShowConfigDetailsStep: true},
				Intercom:               Intercom{AppID: "app_id", HMAC: "hmac_value"},
				Invitations:            []interface{}{},
				JobTitle:               ptrToString("Software Engineer"),
				ManagedBySCIM:          true,
				ManagingOrganizationID: ptrToString("org_id"),
				ProjectMembership:      map[string]interface{}{"project1": "admin"},
				ProjectMemberships:     map[string]interface{}{"project1": "admin"},
				Projects:               []interface{}{"project1"},
				RealName:               "John Doe",
				State:                  "active",
				TokenValidityBegin:     ptrToString("2022-01-01T00:00:00Z"),
				User:                   "john.doe",
				UserID:                 "user_id",
				ViewedIndicators:       []string{"indicator1", "indicator2"},
			}

			user, err := handler.Me(ctx)
			Expect(err).NotTo(HaveOccurred())
			Expect(compareUsers(user, expectedUser)).To(BeTrue())
		})
	})
})

func setupUserProfileHandlerTestCase() (*Client, func()) {
	const (
		UserName     = "test@aiven.io"
		UserPassword = "testabcd"
		AccessToken  = "some-random-token"
	)

	// Mock server to simulate Aiven API responses
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/me" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			err := json.NewEncoder(w).Encode(CurrentUserProfileResponse{
				APIResponse: APIResponse{},
				User: User{
					Auth:                   []string{"auth1", "auth2"},
					City:                   ptrToString("Helsinki"),
					Country:                "FI",
					CreateTime:             "2022-01-01T00:00:00Z",
					Department:             ptrToString("Engineering"),
					Features:               Features{FreeTierEnabled: true, ReferralEnabled: true, ShowConfigDetailsStep: true},
					Intercom:               Intercom{AppID: "app_id", HMAC: "hmac_value"},
					Invitations:            []interface{}{},
					JobTitle:               ptrToString("Software Engineer"),
					ManagedBySCIM:          true,
					ManagingOrganizationID: ptrToString("org_id"),
					ProjectMembership:      map[string]interface{}{"project1": "admin"},
					ProjectMemberships:     map[string]interface{}{"project1": "admin"},
					Projects:               []interface{}{"project1"},
					RealName:               "John Doe",
					State:                  "active",
					TokenValidityBegin:     ptrToString("2022-01-01T00:00:00Z"),
					User:                   "john.doe",
					UserID:                 "user_id",
					ViewedIndicators:       []string{"indicator1", "indicator2"},
				},
			})

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
			return
		}
	}))

	apiUrl = ts.URL
	c, err := NewUserClient(UserName, UserPassword, "aiven-go-client-test/"+Version())
	if err != nil {
		Fail("user authentication error: " + err.Error())
	}

	return c, func() {
		ts.Close()
	}
}

func ptrToString(s string) *string {
	return &s
}

func derefStringPtr(ptr *string) string {
	if ptr != nil {
		return *ptr
	}
	return ""
}

func compareUsers(got, want *User) bool {
	if !reflect.DeepEqual(got.Auth, want.Auth) ||
		derefStringPtr(got.City) != derefStringPtr(want.City) ||
		got.Country != want.Country ||
		got.CreateTime != want.CreateTime ||
		derefStringPtr(got.Department) != derefStringPtr(want.Department) ||
		!reflect.DeepEqual(got.Features, want.Features) ||
		!reflect.DeepEqual(got.Intercom, want.Intercom) ||
		!reflect.DeepEqual(got.Invitations, want.Invitations) ||
		derefStringPtr(got.JobTitle) != derefStringPtr(want.JobTitle) ||
		got.ManagedBySCIM != want.ManagedBySCIM ||
		derefStringPtr(got.ManagingOrganizationID) != derefStringPtr(want.ManagingOrganizationID) ||
		!reflect.DeepEqual(got.ProjectMembership, want.ProjectMembership) ||
		!reflect.DeepEqual(got.ProjectMemberships, want.ProjectMemberships) ||
		!reflect.DeepEqual(got.Projects, want.Projects) ||
		got.RealName != want.RealName ||
		got.State != want.State ||
		derefStringPtr(got.TokenValidityBegin) != derefStringPtr(want.TokenValidityBegin) ||
		got.User != want.User ||
		got.UserID != want.UserID ||
		!reflect.DeepEqual(got.ViewedIndicators, want.ViewedIndicators) {
		return false
	}
	return true
}
