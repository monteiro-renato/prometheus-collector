package querymetrics

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"prometheus-collector/otelcollector/test/utils"
)

var _ = Describe("Query Metrics Test Suite", func() {
	Context("When querying metrics", func() {
		It("should return the expected results", func() {
			_, err := utils.InstantQuery(PrometheusQueryClient, "up")
			Expect(err).NotTo(HaveOccurred())
		})
	})
})
