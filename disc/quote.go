package disc

import (
	"bot/database"
	"fmt"
	"strings"

	"github.com/andersfylling/disgord"
)

func quote(data *disgord.MessageCreate, args CmdArguments) {
	if len(args) > 0 {
		// Transform args into a full string
		argStr := strings.Join(args, " ")

		// Set quote in database
		database.NewQuote{
			UserID: data.Message.Author.ID,
			Quote:  argStr,
		}.SetQuote()

		// Get avatar
		avatarURL, err := data.Message.Author.AvatarURL(64, false)
		if err != nil {
			fmt.Println(err)
		}

		// Send confirmation message
		client.CreateMessage(
			ctx,
			data.Message.ChannelID,
			&disgord.CreateMessageParams{
				Embed: &disgord.Embed{
					Title:       "New quote set",
					Thumbnail:   &disgord.EmbedThumbnail{URL: avatarURL},
					Description: fmt.Sprintf("Your new quote is %s", argStr),
					Color:       0xffe2fe,
				}})
	} else {
		client.CreateMessage(
			ctx,
			data.Message.ChannelID,
			&disgord.CreateMessageParams{
				Embed: &disgord.Embed{
					Title: "Error, quote requires at least 1 argument",
					Color: 0xcc0000,
				}})
	}
}