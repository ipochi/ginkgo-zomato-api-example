package tests

import (
	"github.com/ipochi/ginkgo-zomato-api-example/api"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Zomato API - Restaurants", func() {

	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	Context("Bad Api Key", func() {
		apiKey := "badKey"

		It("Should return error for bad api", func() {
			cityID := "4"
			count := "2"
			entityType := "city"
			statusCode, _, err := api.GetRestaurantsInACity(apiKey, cityID, entityType, count)
			Expect(err).NotTo(HaveOccurred())
			Expect(statusCode).To(BeEquivalentTo(403))
		})
	})

	Context("Get Resturants", func() {
		apiKey := "512b7be2cd07d5291d12d39629b6d23c"

		It("Should return restaurants of a given city id", func() {
			cityID := "4"
			count := "2"
			entityType := "city"
			statusCode, restaurants, err := api.GetRestaurantsInACity(apiKey, cityID, entityType, count)
			Expect(err).NotTo(HaveOccurred())
			Expect(statusCode).To(BeEquivalentTo(200))
			Expect(len(restaurants)).To(BeEquivalentTo(2))
		})

		It("Should Failed to return restaurants of a given city id", func() {
			cityID := "40"
			count := "2"
			entityType := "city"
			statusCode, restaurants, err := api.GetRestaurantsInACity(apiKey, cityID, entityType, count)
			Expect(err).NotTo(HaveOccurred())
			Expect(statusCode).To(BeEquivalentTo(200))
			Expect(len(restaurants)).To(BeEquivalentTo(2))
		})

	})
})
