package utils_test

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/jghiloni/go-commonutils/utils"
)

type mapTest struct {
	Val    int
	IsEven *bool
}

var _ = Describe("Slice", func() {
	DescribeTable("AnySlice", func(actual []any, expected []any) {
		Expect(actual).To(HaveLen(len(expected)))
		for i, a := range actual {
			Expect(a).To(BeEquivalentTo(expected[i]))
		}
	},
		Entry("full", utils.AnySlice([]string{"a", "b", "c"}), []any{"a", "b", "c"}),
		Entry("empty", utils.AnySlice([]bool{}), []any{}),
	)

	DescribeTable("FromAnySlice", func(actual []string, expected []string) {
		Expect(actual).To(HaveLen(len(expected)))
		for i, a := range actual {
			Expect(a).To(BeEquivalentTo(expected[i]))
		}
	},
		Entry("full", utils.FromAnySlice[string]([]any{"a", "b", "c"}), []string{"a", "b", "c"}),
		Entry("empty", utils.FromAnySlice[string]([]any{}), []string{}),
	)

	DescribeTable("Reverse", func(actual []int, expected []int) {
		Expect(actual).To(Equal(expected))
	},
		Entry("odd", utils.Reverse([]int{1, 2, 3, 4, 5}), []int{5, 4, 3, 2, 1}),
		Entry("even", utils.Reverse([]int{0, 2, 4, 6}), []int{6, 4, 2, 0}),
		Entry("empty", utils.Reverse([]int{}), []int{}),
	)

	DescribeTable("SubsliceUntil", func(actual []bool, expected []bool) {
		Expect(actual).To(Equal(expected))
	},
		Entry("full", utils.SubsliceUntil([]bool{false, false, true}, func(b bool) bool { return b }), []bool{false, false}),
		Entry("empty", utils.SubsliceUntil([]bool{}, func(b bool) bool { return !b }), []bool{}),
	)

	DescribeTable("Map", func(actual []mapTest, expected []mapTest) {
		Expect(actual).To(Equal(expected))
	},
		Entry("full", utils.Map([]int{1, 2, 4, 6}, func(i int) mapTest {
			return mapTest{Val: i, IsEven: utils.Ref(i%2 == 0)}
		}), []mapTest{{1, utils.Ref(false)}, {2, utils.Ref(true)}, {4, utils.Ref(true)}, {6, utils.Ref(true)}}),
		Entry("empty", utils.Map([]int{}, func(int) mapTest { return mapTest{} }), []mapTest{}),
	)

	DescribeTable("Filter", func(actual []time.Duration, expected []time.Duration) {
		Expect(actual).To(Equal(expected))
	},
		Entry("full", utils.Filter([]time.Duration{time.Minute, time.Millisecond, 3601 * time.Second, time.Hour, 59_000 * time.Millisecond}, func(d time.Duration) bool {
			return d < time.Hour
		}), []time.Duration{time.Minute, time.Millisecond, 59_000_000_000 * time.Nanosecond}),
		Entry("empty result", utils.Filter([]time.Duration{61 * time.Minute, 2 * time.Hour}, func(d time.Duration) bool {
			return d < time.Hour
		}), []time.Duration{}),
		Entry("empty src", utils.Filter([]time.Duration{}, nil), []time.Duration{}),
	)

	DescribeTable("Reduce", func(x int, y int) { Expect(x).To(Equal(y)) },
		Entry("fac", utils.Reduce([]int{1, 2, 3, 4, 5}, func(cur int, val int) int {
			return cur * val
		}, 1), 120),
		Entry("empty factorial", utils.Reduce([]int{}, func(c int, v int) int { return c * v }, 1), 1),
	)
})
