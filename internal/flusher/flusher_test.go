package flusher_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	"github.com/ozoncp/ocp-solution-api/internal/mocks"
)

var _ = Describe("Flusher", func() {
	var (
		ctrl     *gomock.Controller
		mockRepo *mocks.MockRepo
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockRepo(ctrl)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("", func() {
		BeforeEach(func() {
		})
		It("", func() {
			mockRepo.EXPECT().
				AddSolutions(gomock.Any()).
				Return(nil)
		})
	})
})
