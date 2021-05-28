package config

func init() {
	Config.SetDefault("log.level", "info")

	Config.SetDefault("agentinet.listen",
		[]map[string]string{map[string]string{"address": ":2021"}})

	Config.SetDefault("agentinet.devmode", false)

	Config.SetDefault("repo.dev.interval", 60)
	Config.SetDefault("repo.prod.interval", 14400)
	Config.SetDefault("repo.dev.location", "https://artifactory.company.com/artifactory/myproject/agent/dev/")
	Config.SetDefault("repo.prod.location", "https://artifactory.company.com/artifactory/myproject/agent/prod/")
}
