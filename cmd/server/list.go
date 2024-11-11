package main

func FetchQuestData() Quest {
	return Quest{
		Chars: []string{
			"Ace", "Magna", "Gin", "Killer Bee", "Sato", "Air",
			"Buchwald", "Bojack", "Walter", "Mikami", "Asada",
			"Bulla", "Jose Porla", "Marlene", "Kotetsu",
		},
		Animes: []string{
			"One Piece", "Naruto", "Bleach", "X-men",
			"Black Clover", "Full Metal", "Attack on Titan", "HxH",
			"HxH", "Naruto", "Yugi-Yo", "Full metal",
			"Naruto", "One Piece", "Attack on Titan", "Bleach",
			"My Hero Academia", "Dragon Ball Z", "Demon Slayer", "Naruto",
			"One Punch Man", "Hunter x Hunter", "Naruto", "Death Note",
			"Attack on Titan", "Demon Slayer", "Fullmetal Alchemist", "Naruto",
			"Dragon Ball Z", "My Hero Academia", "Bleach", "Naruto",
			"Black Clover", "One Punch Man", "Hunter x Hunter", "Naruto",
			"Death Note", "One Piece", "Attack on Titan", "Fullmetal Alchemist",
			"Sword Art Online", "Demon Slayer", "My Hero Academia", "Naruto",
			"Dragon Ball Z", "Tokyo Ghoul", "One Punch Man", "Naruto",
			"Fairy Tail", "Death Note", "Naruto", "Hunter x Hunter",
			"Attack on Titan", "Naruto", "My Hero Academia", "Sword Art Online",
			"Demon Slayer", "Tokyo Ghoul", "Bleach", "Dragon Ball Z",
		},
	}
}

const (
	OnePiece           = "One Piece"
	Naruto             = "Naruto"
	Bleach             = "Bleach"
	BlackClover        = "Black Clover"
	FullMetal          = "Full Metal"
	AttackOnTitan      = "Attack on Titan"
	Yugioh             = "Yugi-Yo"
	MyHeroAcademia     = "My Hero Academia"
	DragonBallZ        = "Dragon Ball Z"
	DemonSlayer        = "Demon Slayer"
	OnePunchMan        = "One Punch Man"
	HunterXHunter      = "Hunter x Hunter"
	DeathNote          = "Death Note"
	FullmetalAlchemist = "Fullmetal Alchemist"
	SwordArtOnline     = "Sword Art Online"
	TokyoGhoul         = "Tokyo Ghoul"
	FairyTail          = "Fairy Tail"
)

func FetchData() Data {
	return Data{
		Char: []map[string]string{
			{"name": "Ace", "anime": OnePiece, "image": "ace.jpg"},
			{"name": "Magna", "anime": BlackClover, "image": "magna.jpg"},
			{"name": "Gin", "anime": HunterXHunter, "image": "gin.jpg"},
			{"name": "Killer Bee", "anime": Naruto, "image": "killer-bee.jpg"},
			{"name": "Sato", "anime": MyHeroAcademia, "image": "sato.jpg"},
			{"name": "Air", "anime": OnePunchMan, "image": "air.jpg"},
			{"name": "Buchwald", "anime": AttackOnTitan, "image": "buchwald.jpg"},
			{"name": "Bojack", "anime": DragonBallZ, "image": "bojack.jpg"},
			{"name": "Walter", "anime": BlackClover, "image": "walter.jpg"},
			{"name": "Mikami", "anime": DeathNote, "image": "mikami.jpg"},
			{"name": "Asada", "anime": SwordArtOnline, "image": "asada.jpg"},
			{"name": "Bulla", "anime": DragonBallZ, "image": "bulla.jpg"},
			{"name": "Jose Porla", "anime": FairyTail, "image": "jose-porla.jpg"},
			{"name": "Marlene", "anime": AttackOnTitan, "image": "marlene.jpg"},
			{"name": "Kotetsu", "anime": DemonSlayer, "image": "kotetsu.jpg"},
		},
		AnimeList: []string{
			OnePiece,
			Naruto,
			Bleach,
			BlackClover,
			FullMetal,
			AttackOnTitan,
			Yugioh,
			MyHeroAcademia,
			DragonBallZ,
			DemonSlayer,
			OnePunchMan,
			HunterXHunter,
			DeathNote,
			FullmetalAlchemist,
			SwordArtOnline,
			TokyoGhoul,
			FairyTail,
		},
	}

}
