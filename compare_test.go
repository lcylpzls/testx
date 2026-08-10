package testx

import "testing"

func TestCompareAssertions(t *testing.T) {
	checkRequirePasses(t, func(t *fakeTB) { RequireGreater(t, 3, 2) })
	checkRequirePasses(t, func(t *fakeTB) { RequireGreater(t, 3.5, 3) })
	checkRequirePasses(t, func(t *fakeTB) { RequireGreaterOrEqual(t, 3, 3) })
	checkRequirePasses(t, func(t *fakeTB) { RequireLess(t, 2, 3) })
	checkRequirePasses(t, func(t *fakeTB) { RequireLessOrEqual(t, 3, 3) })
	checkRequirePasses(t, func(t *fakeTB) { RequireLessOrEqual(t, uint8(2), uint8(3)) })
}

func TestCompareAssertionsFailures(t *testing.T) {
	checkRequireFails(t, func(t *fakeTB) { RequireGreater(t, 2, 3) })
	checkRequireFails(t, func(t *fakeTB) { RequireGreater(t, 2, 2) })
	checkRequireFails(t, func(t *fakeTB) { RequireGreaterOrEqual(t, 2, 3) })
	checkRequireFails(t, func(t *fakeTB) { RequireLess(t, 3, 2) })
	checkRequireFails(t, func(t *fakeTB) { RequireLess(t, 3, 3) })
	checkRequireFails(t, func(t *fakeTB) { RequireLessOrEqual(t, 3, 2) })
	checkRequireFails(t, func(t *fakeTB) { RequireGreater(t, "a", "b") })
	checkRequireFails(t, func(t *fakeTB) { RequireGreater(t, 1, "b") })
}
