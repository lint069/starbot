package permissions

// https://discord.com/developers/docs/topics/permissions#permissions-bitwise-permission-flags

var (
	KickMembers                      int64 = 1 << 1
	BanMembers                       int64 = 1 << 2
	Administrator                    int64 = 1 << 3
	ManageChannels                   int64 = 1 << 4
	ManageGuild                      int64 = 1 << 5
	AddReactions                     int64 = 1 << 6
	ViewAuditLog                     int64 = 1 << 7
	PrioritySpeaker                  int64 = 1 << 8
	Stream                           int64 = 1 << 9
	ViewChannel                      int64 = 1 << 10
	SendMessages                     int64 = 1 << 11
	SendTTSMessages                  int64 = 1 << 12
	ManageMessages                   int64 = 1 << 13
	EmbedLinks                       int64 = 1 << 14
	AttachFiles                      int64 = 1 << 15
	ReadMessageHistory               int64 = 1 << 16
	MentionEveryone                  int64 = 1 << 17
	UseExternalEmojis                int64 = 1 << 18
	ViewGuildInsights                int64 = 1 << 19
	Connect                          int64 = 1 << 20
	Speak                            int64 = 1 << 21
	MuteMembers                      int64 = 1 << 22
	DeafenMembers                    int64 = 1 << 23
	MoveMembers                      int64 = 1 << 24
	UseVAD                           int64 = 1 << 25
	ChangeNickname                   int64 = 1 << 26
	ManageNicknames                  int64 = 1 << 27
	ManageRoles                      int64 = 1 << 28
	ManageWebhooks                   int64 = 1 << 29
	ManageGuildExpressions           int64 = 1 << 30
	UseApplicationCommands           int64 = 1 << 31
	RequestToSpeak                   int64 = 1 << 32
	ManageEvents                     int64 = 1 << 33
	ManageThreads                    int64 = 1 << 34
	CreatePublicThreads              int64 = 1 << 35
	CreatePrivateThreads             int64 = 1 << 36
	UseExternalStickers              int64 = 1 << 37
	SendMessagesInThreads            int64 = 1 << 38
	UseEmbeddedActivities            int64 = 1 << 39
	ModerateMembers                  int64 = 1 << 40
	ViewCreatorMonetizationAnalytics int64 = 1 << 41
	UseSoundboard                    int64 = 1 << 42
	CreateGuildExpressions           int64 = 1 << 43
	CreateEvents                     int64 = 1 << 44
	UseExternalSounds                int64 = 1 << 45
	SendVoiceMessages                int64 = 1 << 46
	SendPolls                        int64 = 1 << 49
	UseExternalApps                  int64 = 1 << 50
)
