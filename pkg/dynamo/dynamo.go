package dynamo

type HardcodedFeedLister struct {
}

func (h *HardcodedFeedLister) ListAll() ([]string, error) {
	return []string{
		"https://marioleone.me/index.xml",
		"https://www.guildwars2.com/en/feed/",
	}, nil
}
