package team

import "errors"

type Team struct {
	Name    string
	Country string
	Players []Player
}

func (t Team) CalculateTotalSkillsByTeam() (int, int, int, error) {
	var totalTechnique, totalMental, totalPhisique int

	for _, player := range t.Players {
		totalTechnique += player.Technique
		totalMental += player.Mental
		totalPhisique += player.Physique
	}
	if totalTechnique < 0 || totalMental < 0 || totalPhisique < 0 {
		return -1, -1, -1, errors.New("invalid player attributes: skill values cannot be negative")
	}
	return totalTechnique, totalMental, totalPhisique, nil
}
