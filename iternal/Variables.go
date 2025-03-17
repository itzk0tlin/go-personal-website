package iternal

var VisitorAmount int = 0
var PossibleBotsAmount int = 0

var RecentConnections map[string]int = make(map[string]int)
var BotsAndTheirTimeout map[string]int = make(map[string]int)

var CronJobsRunning int = 0
