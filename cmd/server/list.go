package main

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
			{"name": "Ace", "anime": OnePiece, "image": "https://i.ibb.co/ZXchsTH/image.png"},
			{"name": "Magna", "anime": BlackClover, "image": "https://i.ibb.co/cFWnD5P/image.png"},
			{"name": "Gin", "anime": HunterXHunter, "image": "https://i.ibb.co/GJdMbCv/image.png"},
			{"name": "Killer Bee", "anime": Naruto, "image": "https://i.ibb.co/fYtC0jf/image.png"},
			{"name": "Sato", "anime": MyHeroAcademia, "image": "https://i.ibb.co/TLKD3xF/image.png"},
			{"name": "Air", "anime": OnePunchMan, "image": "https://i.ibb.co/2K8mRY1/image.png"},
			{"name": "Buchwald", "anime": AttackOnTitan, "image": "https://i.ibb.co/4jXZYNv/image.png"},
			{"name": "Bojack", "anime": DragonBallZ, "image": "https://i.ibb.co/vZ9tsPg/image.png"},
			{"name": "Walter", "anime": BlackClover, "image": "https://i.ibb.co/7CQhn9z/image.png"},
			{"name": "Mikami", "anime": DeathNote, "image": "https://i.ibb.co/6yY465G/image.png"},
			{"name": "Asada", "anime": SwordArtOnline, "image": "https://i.ibb.co/WtNsq8D/image.png"},
			{"name": "Bulla", "anime": DragonBallZ, "image": "https://i.ibb.co/xC5ypP1/image.png"},
			{"name": "Jose Porla", "anime": FairyTail, "image": "https://i.ibb.co/61k2Nxg/image.png"},
			{"name": "Marlene", "anime": AttackOnTitan, "image": "https://i.ibb.co/7SrncFR/image.png"},
			{"name": "Kotetsu", "anime": DemonSlayer, "image": "https://i.ibb.co/WBXxGBQ/image.png"},
		},
		AnimeList: []string{
			OnePiece,
			Naruto,
			Bleach,
			BlackClover,
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
