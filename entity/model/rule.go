package model

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/context"
)

type ReactionCalculator interface {
	// ReactionCalculate 计算当前角色身上附着的元素之间能否发生反应，返回反应类型和剩余元素
	ReactionCalculate([]enum.ElementType) (reaction enum.Reaction, elementRemains []enum.ElementType)

	// DamageCalculate 根据反应类型计算对应的伤害修正
	DamageCalculate(reaction enum.Reaction, targetCharacter Character, ctx *context.DamageContext)

	// EffectCalculate 根据反应类型计算对应的反应效果
	EffectCalculate(reaction enum.Reaction, targetPlayer Player) (ctx *context.CallbackContext)

	// Attach 尝试让新元素附着在现有元素集合内，此时不触发元素反应，返回尝试附着后的元素集合
	Attach(originalElements []enum.ElementType, newElement enum.ElementType) (resultElements []enum.ElementType)

	// Relative 判断某种反应是否是某元素的相关反应
	Relative(reaction enum.Reaction, relativeElement enum.ElementType) bool
}

type VictorCalculator interface {
	// CalculateVictors 计算游戏的胜利者
	CalculateVictors(players []Player) (has bool, victors []Player)
}

type GameOptions struct {
	ReRollTimes uint                      // ReRollTimes 所有玩家的基础可重掷次数
	StaticCost  map[enum.ElementType]uint // StaticCost 所有玩家的基础固定持有骰子
	RollAmount  uint                      // RollAmount 所有玩家的投掷阶段生成元素骰子数量
	GetCards    uint                      // GetCards 所有玩家在回合开始时可以获得的卡牌数量
	SwitchCost  map[enum.ElementType]uint // SwitchCost 切换角色所需要的消费
}

var (
	nullReactionCalculator ReactionCalculator = nil
	nullVictorCalculator   VictorCalculator   = nil
)

type RuleSet struct {
	GameOptions        *GameOptions
	ReactionCalculator ReactionCalculator
	VictorCalculator   VictorCalculator
}

func (r RuleSet) ImplementationCheck() bool {
	if r.GameOptions == nil {
		return false
	}

	if r.ReactionCalculator == nullReactionCalculator {
		return false
	}

	if r.VictorCalculator == nullVictorCalculator {
		return false
	}

	return true
}
