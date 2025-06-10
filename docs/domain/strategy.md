# STRATEGY OPTIONS

type Strategy struct {
	StrategyTeam         Team
	Formation            string  4-4-2 4-3-3 4-5-1 5-4-1 5-3-2 3-4-3 3-5-2
	PlayingStyle         string possession counter_attack direct_play high_press low_block
	GameTempo            string fast_tempo balanced_tempo slow_tempo
	PassingStyle         string short long
	DefensivePositioning string zonal_marking man_marking
	BuildUpPlay          string play_from_back long_clearance
	AttackFocus          string wide_play central_play
	KeyPlayerUsage       string reference_player free_role_player
}
