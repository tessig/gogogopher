package resources

// images
//go:generate go run github.com/hajimehoshi/file2byteslice/cmd/file2byteslice -package=images -input=./images/gopher-sprite.png -output=./images/gopher-sprite.go -var=GopherSprite
//go:generate go run github.com/hajimehoshi/file2byteslice/cmd/file2byteslice -package=images -input=./images/forest.png -output=./images/forest.go -var=ForestTiles
//go:generate go run github.com/hajimehoshi/file2byteslice/cmd/file2byteslice -package=images -input=./images/plains.png -output=./images/plains.go -var=PlainsTiles
//go:generate go run github.com/hajimehoshi/file2byteslice/cmd/file2byteslice -package=images -input=./images/gopher-emojis.png -output=./images/gopher-emojis.go -var=GopherEmojis
//go:generate go run github.com/hajimehoshi/file2byteslice/cmd/file2byteslice -package=images -input=./images/elephpant.png -output=./images/elephpant.go -var=Elephpant
//go:generate go run github.com/hajimehoshi/file2byteslice/cmd/file2byteslice -package=images -input=./images/alien_elephpant.png -output=./images/alien_elephpant.go -var=AlienElephpant
//go:generate go run github.com/hajimehoshi/file2byteslice/cmd/file2byteslice -package=images -input=./images/python.png -output=./images/python.go -var=Python
//go:generate go run github.com/hajimehoshi/file2byteslice/cmd/file2byteslice -package=images -input=./images/alien_python.png -output=./images/alien_python.go -var=AlienPython
//go:generate go run github.com/hajimehoshi/file2byteslice/cmd/file2byteslice -package=images -input=./images/duke.png -output=./images/duke.go -var=Duke
//go:generate go run github.com/hajimehoshi/file2byteslice/cmd/file2byteslice -package=images -input=./images/alien_duke.png -output=./images/alien_duke.go -var=AlienDuke
//go:generate go run github.com/hajimehoshi/file2byteslice/cmd/file2byteslice -package=images -input=./images/ferris.png -output=./images/ferris.go -var=Ferris
//go:generate go run github.com/hajimehoshi/file2byteslice/cmd/file2byteslice -package=images -input=./images/alien_ferris.png -output=./images/alien_ferris.go -var=AlienFerris
//go:generate go run github.com/hajimehoshi/file2byteslice/cmd/file2byteslice -package=images -input=./images/alien.png -output=./images/alien.go -var=Alien
//go:generate go run github.com/hajimehoshi/file2byteslice/cmd/file2byteslice -package=images -input=./images/coin.png -output=./images/coin.go -var=Coin
//go:generate go run github.com/hajimehoshi/file2byteslice/cmd/file2byteslice -package=images -input=./images/chest.png -output=./images/chest.go -var=Chest

// sounds
//go:generate go run github.com/hajimehoshi/file2byteslice/cmd/file2byteslice -package=sounds -input=./sounds/jump22.ogg -output=./sounds/jump22.go -var=JumpSound
//go:generate go run github.com/hajimehoshi/file2byteslice/cmd/file2byteslice -package=sounds -input=./sounds/hurt.ogg -output=./sounds/hurt.go -var=HurtSound
//go:generate go run github.com/hajimehoshi/file2byteslice/cmd/file2byteslice -package=sounds -input=./sounds/coin9.ogg -output=./sounds/coin9.go -var=CoinSound
//go:generate go run github.com/hajimehoshi/file2byteslice/cmd/file2byteslice -package=sounds -input=./sounds/powerup02.ogg -output=./sounds/powerup02.go -var=LifeSound

// music
//go:generate go run github.com/hajimehoshi/file2byteslice/cmd/file2byteslice -package=music -input=./music/GreenHills.ogg -output=./music/GreenHills.go -var=GreenHillsMusic
//go:generate go run github.com/hajimehoshi/file2byteslice/cmd/file2byteslice -package=music -input=./music/Funny_Chase.ogg -output=./music/Funny_Chase.go -var=FunnyChase
//go:generate go run github.com/hajimehoshi/file2byteslice/cmd/file2byteslice -package=music -input=./music/spring_thing.ogg -output=./music/spring_thing.go -var=SpringThing
//go:generate go run github.com/hajimehoshi/file2byteslice/cmd/file2byteslice -package=music -input=./music/Birthday_Cake.ogg -output=./music/Birthday_Cake.go -var=BirthdayCake
//go:generate go run github.com/hajimehoshi/file2byteslice/cmd/file2byteslice -package=music -input=./music/Retro_No_hope.ogg -output=./music/Retro_No_hope.go -var=RetroNoHope
//go:generate go run github.com/hajimehoshi/file2byteslice/cmd/file2byteslice -package=music -input=./music/Victory.ogg -output=./music/Victory.go -var=Victory
//go:generate go run github.com/hajimehoshi/file2byteslice/cmd/file2byteslice -package=music -input=./music/Theme_Song_full.ogg -output=./music/Theme_Song_full.go -var=ThemeSongFull
//go:generate go run github.com/hajimehoshi/file2byteslice/cmd/file2byteslice -package=music -input=./music/under_the_sun.ogg -output=./music/under_the_sun.go -var=UnderTheSun
//go:generate go run github.com/hajimehoshi/file2byteslice/cmd/file2byteslice -package=music -input=./music/Platform.ogg -output=./music/Platform.go -var=Platform
//go:generate go run github.com/hajimehoshi/file2byteslice/cmd/file2byteslice -package=music -input=./music/a_little_journey.ogg -output=./music/a_little_journey.go -var=ALittleJourney
//go:generate go run github.com/hajimehoshi/file2byteslice/cmd/file2byteslice -package=music -input=./music/proper_summer.ogg -output=./music/proper_summer.go -var=ProperSummer

// fonts
//go:generate file2byteslice -package=fonts -input=./fonts/pressstart2p.ttf -output=./fonts/pressstart2p.go -var=PressStart2P_ttf
