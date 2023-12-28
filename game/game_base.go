package game

func (g *Game) GetPlayerById(userId uint64) (*Player, bool) {
	for _, v := range g.players {

		if v.user.userID == userId {
			return v, true
		}
	}

	return nil, false
}

func (g *Game) JoinPlayer(data *EventPlayerLoginData) bool {
	if _, ok := g.GetPlayerById(data.userId); ok {
		return false
	}

	if len(g.players) >= len(g.allocateUsers) {
		return false
	}

	p := &Player{
		seat: g.seatIndex,
		user: &User{
			session:     data.session,
			userID:      data.userId,
			nickname:    data.nickname,
			accountType: 0,
			headImgUrl:  "",
			offline:     false,
		},
	}

	g.seatIndex++
	g.players = append(g.players, p)
	return true
}
