package helpers

import "github.com/AmaraNecib/Borhan-backend/types"

func IsValidSlopeType(slope types.SlopeType) bool {
	switch slope {
	case types.Downsloping, types.Flat, types.Upsloping:
		return true
	}
	return false
}

func IsValidCaType(ca types.CaType) bool {
	switch ca {
	case types.Zero, types.One, types.Two, types.Three:
		return true
	}
	return false
}

func IsValidThalType(thal types.ThalType) bool {
	switch thal {
	case types.FixedDefect, types.NormalThal, types.ReversableDefect, types.Reversable:
		return true
	}
	return false
}

func IsValidLogic(logic types.Logic) bool {
	switch logic {
	case types.False, types.True:
		return true
	}
	return false
}

func IsValidRestecgType(restecg types.RestecgType) bool {
	switch restecg {
	case types.Hypertrophy, types.NormalRestecg, types.STTWaveAbnormality:
		return true
	}
	return false
}

func IsValidCpType(cp types.CpType) bool {
	switch cp {
	case types.TypicalAngina, types.Asymptomatic, types.NonAnginal, types.AtypicalAngina:
		return true
	}
	return false
}
