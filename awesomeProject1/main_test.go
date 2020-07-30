package main
import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func dotest(arr []int, exp int) {
	var ans = Solve(arr)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Example tests", func() {
	It("It should work for basic tests", func() {
		dotest([] int{1,-1,2,-2,3}, 3)
		dotest([] int{-3,1,2,3,-1,-4,-2}, -4)
		dotest([] int{1,-1,2,-2,3,3}, 3)
		dotest([] int{-110,110,-38,-38,-62,62,-38,-38,-38}, -38)
		dotest([] int{-9,-105,-9,-9,-9,-9,105}, -9)
	})
})
