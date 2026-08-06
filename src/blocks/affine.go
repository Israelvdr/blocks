package blocks

const (
	RESOLVABLE_GROUP_DIVISIBLE_DESIGN_NAME        string = "Resolvable Group Divisible (Block) Design"
	RESOLVABLE_GROUP_DIVISIBLE_DESIGN_DESCRIPTION string = `
		A resolvable group divisible design is a perfect solution to the social golfer problem. It
		groups n players into g blocks of p players for each of s rounds. Each players appears in
		exactly one block per round, and each pair of players among n appears exactly once across
		all rounds.`
)

type resolvableGroupDivisibleDesign struct {
	blockDesignBase
}

func (abd *resolvableGroupDivisibleDesign) designFamilyName() string {
	return RESOLVABLE_GROUP_DIVISIBLE_DESIGN_NAME
}

func (abd *resolvableGroupDivisibleDesign) designFamilyDescription() string {
	return RESOLVABLE_GROUP_DIVISIBLE_DESIGN_DESCRIPTION
}

func (abd *resolvableGroupDivisibleDesign) solve() blockDesign {
	return abd
}

func (abd *resolvableGroupDivisibleDesign) optimise() blockDesign {
	return abd
}
