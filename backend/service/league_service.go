package service

type ILeagueService interface {
	GetMatchData(summoner string, region string)
}
