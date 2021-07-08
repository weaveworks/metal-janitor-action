package action_test

import (
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/ghttp"

	"github.com/weaveworks/metal-janitor-action/action"
)

var _ = Describe("TestMetalJanitor", func() {
	var (
		err error
		s   *Server
		a   action.MetalJanitorAction
	)

	BeforeEach(func() {
		s = NewServer()

		a, err = action.NewWithURL("12345678", s.HTTPTestServer.Client(), s.HTTPTestServer.URL)
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		s.Close()
	})

	Describe("Cleaning up a single project", func() {
		BeforeEach(func() {
			s.AppendHandlers(
				CombineHandlers(
					VerifyRequest("GET", "/projects"),
					RespondWith(http.StatusOK, getTestData("list_projects_2_resp.json"), http.Header{"Content-Type": []string{"application/json; charset=utf-8"}}),
				),
				CombineHandlers(
					VerifyRequest("GET", "/projects/1e499db8-5803-47fa-989a-4d5c7dee1dae/devices"),
					RespondWith(http.StatusOK, getTestData("list_devices_resp.json"), http.Header{"Content-Type": []string{"application/json; charset=utf-8"}}),
				),
				CombineHandlers(
					VerifyRequest("DELETE", "/devices/321c8340-dc80-40a9-a2e6-e0efdb9a59a6"),
					RespondWith(http.StatusAccepted, nil, http.Header{"Content-Type": []string{"application/json; charset=utf-8"}}),
				),
				CombineHandlers(
					VerifyRequest("GET", "/projects/1e499db8-5803-47fa-989a-4d5c7dee1dae/storage"),
					RespondWith(http.StatusOK, getTestData("list_volumes_empty_resp.json"), http.Header{"Content-Type": []string{"application/json; charset=utf-8"}}),
				),
				CombineHandlers(
					VerifyRequest("DELETE", "/projects/1e499db8-5803-47fa-989a-4d5c7dee1dae"),
					RespondWith(http.StatusAccepted, nil, http.Header{"Content-Type": []string{"application/json; charset=utf-8"}}),
				),
			)

			err = a.Cleanup("ProjectA", false)
		})

		It("should cleanup the projects", func() {
			Expect(err).NotTo(HaveOccurred())
			Expect(len(s.ReceivedRequests())).Should(Equal(5))
		})
	})

	Describe("Cleaning up a single project in dry run mode", func() {
		BeforeEach(func() {
			s.AppendHandlers(
				CombineHandlers(
					VerifyRequest("GET", "/projects"),
					RespondWith(http.StatusOK, getTestData("list_projects_2_resp.json"), http.Header{"Content-Type": []string{"application/json; charset=utf-8"}}),
				),
				CombineHandlers(
					VerifyRequest("GET", "/projects/1e499db8-5803-47fa-989a-4d5c7dee1dae/devices"),
					RespondWith(http.StatusOK, getTestData("list_devices_resp.json"), http.Header{"Content-Type": []string{"application/json; charset=utf-8"}}),
				),
				CombineHandlers(
					VerifyRequest("GET", "/projects/1e499db8-5803-47fa-989a-4d5c7dee1dae/storage"),
					RespondWith(http.StatusOK, getTestData("list_volumes_empty_resp.json"), http.Header{"Content-Type": []string{"application/json; charset=utf-8"}}),
				),
			)

			err = a.Cleanup("ProjectA", true)
		})

		It("should cleanup the projects", func() {
			Expect(err).NotTo(HaveOccurred())
			Expect(len(s.ReceivedRequests())).Should(Equal(3))
		})
	})

	Describe("Cleaning up all projects", func() {
		BeforeEach(func() {
			s.AppendHandlers(
				CombineHandlers(
					VerifyRequest("GET", "/projects"),
					RespondWith(http.StatusOK, getTestData("list_projects_1_resp.json"), http.Header{"Content-Type": []string{"application/json; charset=utf-8"}}),
				),
				CombineHandlers(
					VerifyRequest("GET", "/projects/1e499db8-5803-47fa-989a-4d5c7dee1dae/devices"),
					RespondWith(http.StatusOK, getTestData("list_devices_resp.json"), http.Header{"Content-Type": []string{"application/json; charset=utf-8"}}),
				),
				CombineHandlers(
					VerifyRequest("DELETE", "/devices/321c8340-dc80-40a9-a2e6-e0efdb9a59a6"),
					RespondWith(http.StatusAccepted, nil, http.Header{"Content-Type": []string{"application/json; charset=utf-8"}}),
				),
				CombineHandlers(
					VerifyRequest("GET", "/projects/1e499db8-5803-47fa-989a-4d5c7dee1dae/storage"),
					RespondWith(http.StatusOK, getTestData("list_volumes_empty_resp.json"), http.Header{"Content-Type": []string{"application/json; charset=utf-8"}}),
				),
				CombineHandlers(
					VerifyRequest("DELETE", "/projects/1e499db8-5803-47fa-989a-4d5c7dee1dae"),
					RespondWith(http.StatusAccepted, nil, http.Header{"Content-Type": []string{"application/json; charset=utf-8"}}),
				),
			)

			err = a.Cleanup("DELETEALL", false)
		})

		It("should cleanup the projects", func() {
			Expect(err).NotTo(HaveOccurred())
			Expect(len(s.ReceivedRequests())).Should(Equal(5))
		})
	})
})
