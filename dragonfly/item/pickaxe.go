package item

import (
	"git.jetbrains.space/dragonfly/dragonfly.git/dragonfly/item/tool"
	"git.jetbrains.space/dragonfly/dragonfly.git/dragonfly/world"
)

// Pickaxe is a tool generally used for mining stone-like blocks and ores at a higher speed and to obtain
// their drops.
type Pickaxe struct {
	// Tier is the tier of the pickaxe.
	Tier tool.Tier
}

// ToolType returns the type for pickaxes.
func (p Pickaxe) ToolType() tool.Type {
	return tool.TypePickaxe
}

// HarvestLevel returns the level that this pickaxe is able to harvest. If a block has a harvest level above
// this one, this pickaxe won't be able to harvest it.
func (p Pickaxe) HarvestLevel() int {
	return p.Tier.HarvestLevel
}

// BaseMiningEfficiency is the base efficiency of the pickaxe, when it comes to mining blocks. This decides
// the speed with which blocks can be mined.
func (p Pickaxe) BaseMiningEfficiency(world.Block) float64 {
	return p.Tier.BaseMiningEfficiency
}

// MaxCount returns 1.
func (p Pickaxe) MaxCount() int {
	return 1
}

// AttackDamage returns the attack damage of the pickaxe.
func (p Pickaxe) AttackDamage() float64 {
	return p.Tier.BaseAttackDamage + 1
}

// DurabilityInfo ...
func (p Pickaxe) DurabilityInfo() DurabilityInfo {
	return DurabilityInfo{
		MaxDurability:    p.Tier.Durability,
		BrokenItem:       simpleItem(Stack{}),
		AttackDurability: 2,
		BreakDurability:  1,
	}
}

// EncodeItem ...
func (p Pickaxe) EncodeItem() (id int32, meta int16) {
	switch p.Tier {
	case tool.TierWood:
		return 270, 0
	case tool.TierGold:
		return 285, 0
	case tool.TierStone:
		return 274, 0
	case tool.TierIron:
		return 257, 0
	case tool.TierDiamond:
		return 278, 0
	}
	panic("invalid pickaxe tier")
}
