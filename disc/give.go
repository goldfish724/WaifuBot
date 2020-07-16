package disc

import (
	"bot/database"
	"bot/query"
	"fmt"

	"github.com/andersfylling/disgord"
)

func giveChar(data *disgord.MessageCreate, args []string) {
	// Verify if give is valid, also deletes the character from User1's database if valid
	desc, valid := validGive(data, args)

	// Get avatar
	avatar, err := data.Message.Author.AvatarURL(128, false)
	if err != nil {
		fmt.Println(err)
	}

	if valid == true {
		// Get char
		resp, err := query.CharSearch(ParseArgToSearch(args))
		if err != nil {
			fmt.Println(err)
		}

		// Add the char to the mentionned user's database
		database.AddChar(database.InputChar{
			UserID: data.Message.Mentions[0].ID,
			WaifuList: database.CharLayout{
				ID:    resp.Character.ID,
				Image: resp.Character.Image.Large,
				Name:  resp.Character.Name.Full,
			}})

		// Send confirmation Message
		client.CreateMessage(
			ctx,
			data.Message.ChannelID,
			&disgord.CreateMessageParams{
				Embed: &disgord.Embed{
					Title:       "Give Waifu Succeded",
					Thumbnail:   &disgord.EmbedThumbnail{URL: avatar},
					Description: fmt.Sprintf("%s gave %s to %s", data.Message.Author.Username, resp.Character.Name.Full, data.Message.Mentions[0].Username),
					Image:       &disgord.EmbedImage{URL: resp.Character.Image.Large},
					Color:       0x43e99a,
				},
			})
	} else {
		// Send message
		client.CreateMessage(
			ctx,
			data.Message.ChannelID,
			&disgord.CreateMessageParams{
				Embed: &disgord.Embed{
					Title:       "Give Waifu Failed",
					Thumbnail:   &disgord.EmbedThumbnail{URL: avatar},
					Description: desc,
					Color:       0xcc0000,
				},
			})
	}
}

// Verify if give is valid, also deletes the character from User1's database
func validGive(data *disgord.MessageCreate, arg []string) (desc string, isValid bool) {
	if len(arg) > 0 {
		resp := ParseArgToSearch(arg)
		switch {
		case resp.ID == 0:
			return fmt.Sprintf("Error, %d is not a valid WaifuID,\nRefer to %shelp to see this command's syntax", resp.ID, conf.Prefix), false
		case data.Message.Mentions == nil:
			return fmt.Sprintf("Error, please tag a discord user,\nRefer to %shelp to see this command's syntax", conf.Prefix), false
		case database.DelChar(database.DelWaifuStruct{UserID: data.Message.Author.ID, CharID: resp.ID}) == false:
			return fmt.Sprintf("You do not own the character ID %d,\nVerify if the ID you entered is correct", resp.ID), false
		default:
			return "", true
		}
	}
	return "Please enter arguments,\nRefer to help to see how to use this command", false
}