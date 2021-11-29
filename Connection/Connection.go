package Connection;

import (
	"log"
	"github.com/bwmarrin/discordgo"

	"os"
	"os/signal"
	"syscall"
)

func Start() *discordgo.Session {
	dg, err := discordgo.New("Bot " + getToken())
	if err != nil {
		log.Fatal("error creating Discord session,", err)
	}
	
	connectBotToDiscord(dg)
	return dg
}

func connectBotToDiscord(dg *discordgo.Session){
	err := dg.Open()
	if err != nil {
		log.Println("error opening connection,", err)
		return
	}
}

func Wait(dg *discordgo.Session,){
	log.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	dg.Close()
}