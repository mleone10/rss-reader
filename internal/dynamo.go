package internal

type HardcodedFeedLister struct {
}

func listFeeds() ([]string, error) {
	return []string{
		"https://marioleone.me/index.xml",
		"https://www.guildwars2.com/en/feed/",
	}, nil
}
