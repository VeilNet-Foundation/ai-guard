// ai_resistance/bce.go
type UserProfile struct {
	Name      string
	Times     []int     // активные часы (0-23)
	Delays    []int     // задержки между действиями
	Sites     []string  // типичные запросы
	Device    string    // "mobile", "desktop"
}

var Profiles = []UserProfile{
	{
		Name:   "night-owl",
		Times:  []int{22, 23, 0, 1, 2, 3, 4},
		Delays: []int{100, 500, 1000},
		Sites:  []string{"news.veil", "chat.veil"},
		Device: "mobile",
	},
	{
		Name:   "office-worker",
		Times:  []int{9, 10, 11, 12, 13, 14, 15, 16},
		Delays: []int{2000, 5000},
		Sites:  []string{"docs.veil", "mail.veil"},
		Device: "desktop",
	},
}

func EmulateProfile(profile UserProfile) {
	if !isActiveHour(profile.Times) {
		return
	}
	
	for i := 0; i < rand.Intn(5); i++ {
		site := profile.Sites[rand.Intn(len(profile.Sites))]
		SendFakeRequest(site)
		time.Sleep(time.Duration(randomChoice(profile.Delays)) * time.Millisecond)
	}
}