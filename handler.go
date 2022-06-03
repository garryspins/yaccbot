package main

import (
	"log"
	"os"

	d "github.com/bwmarrin/discordgo"
)

type CommandResponseFunc func(s *d.Session, i *d.InteractionCreate)
type CommandHandler struct {
	session 	*d.Session
	commands 	[]*d.ApplicationCommand
	registered 	[]*d.ApplicationCommand
	handlers 	map[string]CommandResponseFunc
}

func (self *CommandHandler) New(name string, description string, response CommandResponseFunc) *d.ApplicationCommand {
	if self.handlers == nil {
		self.handlers = map[string]CommandResponseFunc{}
	}
	
	manifest := d.ApplicationCommand{
		Name: name,
		Description: description,
	}
	
	self.handlers[name] = func(s *d.Session, i *d.InteractionCreate){
		s.InteractionRespond(i.Interaction, &d.InteractionResponse{
			Type: d.InteractionResponseChannelMessageWithSource,
			Data: &d.InteractionResponseData{
				Content: "Pong",
			},
		})
	}

	self.commands = append(self.commands, &manifest)
	return &manifest
}

func (self *CommandHandler) BotHandler(s *d.Session, i *d.InteractionCreate) {
	h, ok := self.handlers[i.ApplicationCommandData().Name]

	if ok {
		h(s, i)
	}
}

func (self *CommandHandler) InitializeCommands() {
	guild := os.Getenv("DISCORD_GUILD_ID")

	self.registered = make([]*d.ApplicationCommand, len(self.commands))
	for k, v := range self.commands {
		cmd, err := self.session.ApplicationCommandCreate(self.session.State.User.ID, guild, v)

		if err != nil {
			log.Panicf("Cannot create command '%v', %v", v.Name, err)
		}

		self.registered[k] = cmd
	}
}

func (self *CommandHandler) CleanupCommands() {
	log.Printf("Removing commands")
	guild := os.Getenv("DISCORD_GUILD_ID")

	for _, v := range self.registered {
		err := self.session.ApplicationCommandDelete(self.session.State.User.ID, guild, v.ID)

		if err != nil {
			log.Panicf("Cannot delete '%v', %v", v.Name, err)
		}
	}
}