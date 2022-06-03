
package main

import (
	d "github.com/bwmarrin/discordgo"
)

var handler = CommandHandler{}

func (self *CommandHandler) Create() {
	self.New("test", "This is a test command", func(s *d.Session, i *d.InteractionCreate) {
		s.InteractionRespond(i.Interaction, &d.InteractionResponse{
			Type: d.InteractionResponseChannelMessageWithSource,
			Data: &d.InteractionResponseData{
				Content: "This a response",
			},
		})
	})
}