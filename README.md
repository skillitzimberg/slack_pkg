# Slack Step Package

Steps for formatting, fetching, & posting messages to Slack.  
[Messsage Composition](https://api.slack.com/messaging/composing)
[Messsage Formatting](https://api.slack.com/messaging/composing/formatting)
[Messsage Sending](https://api.slack.com/messaging/sending)

### Message payload

**Section Block**
A highly customizable block that accepts a plaintext or markdown element, and also supports an accessory that can be an image or action element.

- type: string - must be "section"

- block_id: string

- text: object

- accessory: object

- fields: array

#### Approval Composition

```
[
	{
		"type": "section",
		"text": {
			"type": "mrkdwn",
			"text": "Portland Junior Developers:\n*<fakeLink.toMeetup.com|Juniper Marques: Open Source - What even is? How even to?>*"
		}
	},
	{
		"type": "section",
		"fields": [
			{
				"type": "mrkdwn",
				"text": "*Type:*\nComputer (laptop)"
			},
			{
				"type": "mrkdwn",
				"text": "*When:*\nSubmitted Aut 10"
			},
			{
				"type": "mrkdwn",
				"text": "*Last Update:*\nMar 10, 2015 (3 years, 5 months)"
			},
			{
				"type": "mrkdwn",
				"text": "*Reason:*\nAll vowel keys aren't working."
			},
			{
				"type": "mrkdwn",
				"text": "*Specs:*\n\"Cheetah Pro 15\" - Fast, really fast\""
			}
		]
	},
	{
		"type": "actions",
		"elements": [
			{
				"type": "button",
				"text": {
					"type": "plain_text",
					"emoji": true,
					"text": "Going"
				},
				"value": "click_me_123",
				"style": "primary"
			},
			{
				"type": "button",
				"text": {
					"type": "plain_text",
					"emoji": true,
					"text": "Not Going"
				},
				"value": "click_me_123",
				"style": "danger"
			}
		]
	}
]
```

#### Notification Composition

```
[
	{
		"type": "section",
		"text": {
			"type": "plain_text",
			"emoji": true,
			"text": "Looks like you have a scheduling conflict with this event:"
		}
	},
	{
		"type": "divider"
	},
	{
		"type": "section",
		"text": {
			"type": "mrkdwn",
			"text": "*<fakeLink.toUserProfiles.com|Iris / Zelda 1-1>*\nTuesday, January 21 4:00-4:30pm\nBuilding 2 - Havarti Cheese (3)\n2 guests"
		},
		"accessory": {
			"type": "image",
			"image_url": "https://api.slack.com/img/blocks/bkb_template_images/notifications.png",
			"alt_text": "calendar thumbnail"
		}
	},
	{
		"type": "context",
		"elements": [
			{
				"type": "image",
				"image_url": "https://api.slack.com/img/blocks/bkb_template_images/notificationsWarningIcon.png",
				"alt_text": "notifications warning icon"
			},
			{
				"type": "mrkdwn",
				"text": "*Conflicts with Team Huddle: 4:15-4:30pm*"
			}
		]
	},
	{
		"type": "divider"
	},
	{
		"type": "section",
		"text": {
			"type": "mrkdwn",
			"text": "*Propose a new time:*"
		}
	},
	{
		"type": "section",
		"text": {
			"type": "mrkdwn",
			"text": "*Today - 4:30-5pm*\nEveryone is available: @iris, @zelda"
		},
		"accessory": {
			"type": "button",
			"text": {
				"type": "plain_text",
				"emoji": true,
				"text": "Choose"
			},
			"value": "click_me_123"
		}
	},
	{
		"type": "section",
		"text": {
			"type": "mrkdwn",
			"text": "*Tomorrow - 4-4:30pm*\nEveryone is available: @iris, @zelda"
		},
		"accessory": {
			"type": "button",
			"text": {
				"type": "plain_text",
				"emoji": true,
				"text": "Choose"
			},
			"value": "click_me_123"
		}
	},
	{
		"type": "section",
		"text": {
			"type": "mrkdwn",
			"text": "*Tomorrow - 6-6:30pm*\nSome people aren't available: @iris, ~@zelda~"
		},
		"accessory": {
			"type": "button",
			"text": {
				"type": "plain_text",
				"emoji": true,
				"text": "Choose"
			},
			"value": "click_me_123"
		}
	},
	{
		"type": "section",
		"text": {
			"type": "mrkdwn",
			"text": "*<fakelink.ToMoreTimes.com|Show more times>*"
		}
	}
]
```

### Authors

Scott Bergler  
Ngan Nguyen

## Table of contents

**[Resources](#resources)**<br>
**[Setup & Installation](#setup-and-installation)**<br>
**[Known Bugs](#known-bugs)**<br>
**[Acknowledgements](#acknowledgements)**<br>
**[Support](#support)**<br>
**[Technologies Used](#technologies-used)**<br>
**[License](#license)**<br>
**[Copyright](#copyright)**<br>

## Resources

## Setup and Installation

## Known Bugs

## Acknowledgements

## Support

Scott Bergler :: commandinghands@gmail.com  
Ngan Nguyen :: ngan@nguyen.com

## Technologies Used

| Dependency  | Description                    |
| ----------- | ------------------------------ |
| apptree CLI | CLI tools for AppTree platform |

## License

## Copyright

<!-- Copyright (c) 2019 **Scott Bergler; Ngan Nguyen** -->
