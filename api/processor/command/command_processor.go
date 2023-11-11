package command_processor

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/takuya-okada-01/badminist/api/domain/community"
	"github.com/takuya-okada-01/badminist/api/domain/community/member"
	"github.com/takuya-okada-01/badminist/api/domain/community/player"
	"github.com/takuya-okada-01/badminist/api/domain/user"
	command_repository_if "github.com/takuya-okada-01/badminist/api/interface_adaptor_if/repository_if/command"
	"github.com/takuya-okada-01/badminist/api/utils"
)

type CommandProcessor interface {
	CreateCommunity(
		name community.CommunityName,
		description community.CommunityDescription,
		executorId user.UserId,
	) error
	RenameCommunity(
		communityId community.CommunityId,
		name community.CommunityName,
		executorId user.UserId,
	) error
	EditCommunityDescription(
		communityId community.CommunityId,
		description community.CommunityDescription,
		executorId user.UserId,
	) error
	DeleteCommunity(
		communityId community.CommunityId,
		executorId user.UserId,
	) error
	AddMember(
		communityId community.CommunityId,
		userId user.UserId,
		memberRole member.MemberRole,
		executorId user.UserId,
	) error
	RemoveMember(
		communityId community.CommunityId,
		userId user.UserId,
		executorId user.UserId,
	) error
	ChangeMemberRole(
		communityId community.CommunityId,
		userId user.UserId,
		role member.MemberRole,
		executorId user.UserId,
	) error
	AddPlayer(
		communityId community.CommunityId,
		playerName player.PlayerName,
		playerGender player.PlayerGender,
		playerAge player.PlayerAge,
		playerLevel player.PlayerLevel,
		playerNumGames player.PlayerNumGames,
		playerStatus player.PlayerStatus,
		executorId user.UserId,
	) error
	RemovePlayer(
		communityId community.CommunityId,
		playerId player.PlayerId,
		executorId user.UserId,
	) error
	ChangePlayerProperty(
		communityId community.CommunityId,
		playerId player.PlayerId,
		playerName player.PlayerName,
		playerGender player.PlayerGender,
		playerAge player.PlayerAge,
		playerLevel player.PlayerLevel,
		playerNumGames player.PlayerNumGames,
		playerStatus player.PlayerStatus,
		executorId user.UserId,
	) error
	ChangePlayerNumGames(
		communityId community.CommunityId,
		playerId player.PlayerId,
		numGames player.PlayerNumGames,
		executorId user.UserId,
	) error
	ResetPlayerNumGames(
		communityId community.CommunityId,
		playerId player.PlayerId,
		executorId user.UserId,
	) error
	TemporaryRegistration(
		name user.UserName,
		email user.UserEmail,
		password user.UserPassword,
	) (string, error)
	ActivateUser(
		email user.UserEmail,
		confirmPass user.UserConfirmPass,
	) error
	Login(
		email user.UserEmail,
		password user.UserPassword,
	) (string, error)
	ReissueConfirmPass(
		id user.UserId,
	) error
}

type commandProcessor struct {
	communityRepo command_repository_if.CommunityRepository
	userRepo      command_repository_if.UserRepository
	emailServer   utils.EmailServer
}

func NewCommandProcessor(
	communityRepo command_repository_if.CommunityRepository,
	userRepo command_repository_if.UserRepository,
	emailServer utils.EmailServer,
) CommandProcessor {
	return &commandProcessor{
		communityRepo: communityRepo,
		userRepo:      userRepo,
		emailServer:   emailServer,
	}
}

func (c *commandProcessor) CreateCommunity(
	name community.CommunityName,
	description community.CommunityDescription,
	executorId user.UserId,
) error {
	members := member.NewMembers(executorId)
	community := community.NewCommunity(
		community.NewCommunityId(),
		name,
		description,
		player.NewPlayers([]player.Player{}),
		members,
	)
	if err := c.communityRepo.CreateCommunity(
		community.BreachEncapsulationOfId(),
		community.BreachEncapsulationOfName(),
		community.BreachEncapsulationOfDescription(),
		community.BreachEncapsulationOfMembers(),
	); err != nil {
		return err
	}
	return nil
}

func (c *commandProcessor) RenameCommunity(
	communityId community.CommunityId,
	name community.CommunityName,
	executorId user.UserId,
) error {
	community, err := c.communityRepo.FindCommunityById(communityId)
	if err != nil {
		return err
	}
	event, err := community.RenameCommunity(name, executorId)
	if err != nil {
		return err
	}
	if err := c.communityRepo.RenameCommunity(
		event.CommunityId,
		event.Name,
	); err != nil {
		return err
	}
	return nil
}

func (c *commandProcessor) EditCommunityDescription(
	communityId community.CommunityId,
	description community.CommunityDescription,
	executorId user.UserId,
) error {
	community, err := c.communityRepo.FindCommunityById(communityId)
	if err != nil {
		return err
	}
	event, err := community.EditDescription(description, executorId)
	if err != nil {
		return err
	}
	if err := c.communityRepo.EditCommunityDescription(
		event.CommunityId,
		community.BreachEncapsulationOfName(),
		event.Description,
	); err != nil {
		return err
	}
	return nil
}

func (c *commandProcessor) DeleteCommunity(
	communityId community.CommunityId,
	executorId user.UserId,
) error {
	community, err := c.communityRepo.FindCommunityById(communityId)
	if err != nil {
		return err
	}
	event, err := community.DeleteCommunity(executorId)
	if err != nil {
		return err
	}
	if err := c.communityRepo.DeleteCommunity(event.CommunityId); err != nil {
		return err
	}
	return nil
}

func (c *commandProcessor) AddMember(
	communityId community.CommunityId,
	userId user.UserId,
	memberRole member.MemberRole,
	executorId user.UserId,
) error {
	community, err := c.communityRepo.FindCommunityById(communityId)
	if err != nil {
		return err
	}
	member := member.NewMember(
		member.NewMemberId(),
		memberRole,
		userId,
	)

	event, err := community.AddMember(
		member,
		executorId,
	)
	if err != nil {
		return err
	}
	if err := c.communityRepo.AddMember(
		event.Member.BreachEncapsulationOfId(),
		event.CommunityId,
		event.Member.BreachEncapsulationOfUserId(),
		event.Member.BreachEncapsulationOfRole(),
	); err != nil {
		return err
	}
	return nil
}

func (c *commandProcessor) RemoveMember(
	communityId community.CommunityId,
	userId user.UserId,
	executorId user.UserId,
) error {
	community, err := c.communityRepo.FindCommunityById(communityId)
	if err != nil {
		return err
	}
	event, err := community.RemoveMember(
		userId,
		executorId,
	)
	if err != nil {
		return err
	}
	if err := c.communityRepo.RemoveMember(
		event.CommunityId,
		event.UserId,
	); err != nil {
		return err
	}
	return nil
}

func (c *commandProcessor) ChangeMemberRole(
	communityId community.CommunityId,
	userId user.UserId,
	role member.MemberRole,
	executorId user.UserId,
) error {
	community, err := c.communityRepo.FindCommunityById(communityId)
	if err != nil {
		return err
	}
	event, err := community.ChangeMemberRole(
		userId,
		role,
		executorId,
	)
	if err != nil {
		return err
	}
	if err := c.communityRepo.ChangeMemberRole(
		event.CommunityId,
		event.Member.BreachEncapsulationOfUserId(),
		event.Member.BreachEncapsulationOfRole(),
	); err != nil {
		return err
	}
	return nil
}

func (c *commandProcessor) AddPlayer(
	communityId community.CommunityId,
	playerName player.PlayerName,
	playerGender player.PlayerGender,
	playerAge player.PlayerAge,
	playerLevel player.PlayerLevel,
	playerNumGames player.PlayerNumGames,
	playerStatus player.PlayerStatus,
	executorId user.UserId,

) error {
	community, err := c.communityRepo.FindCommunityById(communityId)
	if err != nil {
		return err
	}
	playerId := player.NewPlayerId()

	player := player.NewPlayer(
		playerId,
		playerName,
		playerGender,
		playerAge,
		playerLevel,
		playerNumGames,
		playerStatus,
	)
	event, err := community.AddPlayer(
		player,
		executorId,
	)
	if err != nil {
		return err
	}

	if err := c.communityRepo.AddPlayer(
		event.CommunityId,
		event.Player.BreachEncapsulationOfId(),
		event.Player.BreachEncapsulationOfName(),
		event.Player.BreachEncapsulationOfGender(),
		event.Player.BreachEncapsulationOfAge(),
		event.Player.BreachEncapsulationOfLevel(),
		event.Player.BreachEncapsulationOfNumGames(),
		event.Player.BreachEncapsulationOfStatus(),
	); err != nil {
		return err
	}
	return nil
}

func (c *commandProcessor) RemovePlayer(
	communityId community.CommunityId,
	playerId player.PlayerId,
	executorId user.UserId,
) error {
	community, err := c.communityRepo.FindCommunityById(communityId)
	if err != nil {
		return err
	}
	event, err := community.RemovePlayer(
		playerId,
		executorId,
	)
	if err != nil {
		return err
	}
	if err := c.communityRepo.RemovePlayer(
		event.CommunityId,
		event.PlayerId,
	); err != nil {
		return err
	}
	return nil
}

func (c *commandProcessor) ChangePlayerProperty(
	communityId community.CommunityId,
	playerId player.PlayerId,
	playerName player.PlayerName,
	playerGender player.PlayerGender,
	playerAge player.PlayerAge,
	playerLevel player.PlayerLevel,
	playerNumGames player.PlayerNumGames,
	playerStatus player.PlayerStatus,
	executorId user.UserId,
) error {
	community, err := c.communityRepo.FindCommunityById(communityId)
	if err != nil {
		return err
	}
	event, err := community.ChangePlayerProperty(
		playerId,
		playerName,
		playerGender,
		playerAge,
		playerLevel,
		playerNumGames,
		playerStatus,
		executorId,
	)
	if err != nil {
		return err
	}
	if err := c.communityRepo.ChangePlayerProperty(
		event.CommunityId,
		event.Player.BreachEncapsulationOfId(),
		event.Player.BreachEncapsulationOfName(),
		event.Player.BreachEncapsulationOfGender(),
		event.Player.BreachEncapsulationOfAge(),
		event.Player.BreachEncapsulationOfLevel(),
		event.Player.BreachEncapsulationOfNumGames(),
		event.Player.BreachEncapsulationOfStatus(),
	); err != nil {
		return err
	}
	return nil
}

func (c *commandProcessor) ResetPlayerNumGames(
	communityId community.CommunityId,
	playerId player.PlayerId,
	executorId user.UserId,
) error {
	community, err := c.communityRepo.FindCommunityById(communityId)
	if err != nil {
		return err
	}
	event, err := community.ResetPlayerNumGames(
		playerId,
		executorId,
	)
	if err != nil {
		return err
	}

	if err := c.communityRepo.ChangePlayerNumGames(
		event.CommunityId,
		event.PlayerId,
		event.NumGames,
	); err != nil {
		return err
	}
	return nil
}

func (c *commandProcessor) ChangePlayerNumGames(
	communityId community.CommunityId,
	playerId player.PlayerId,
	numGames player.PlayerNumGames,
	executorId user.UserId,
) error {
	community, err := c.communityRepo.FindCommunityById(communityId)
	if err != nil {
		return err
	}
	event, err := community.ChangePlayerNumGames(
		playerId,
		numGames,
		executorId,
	)
	if err != nil {
		return err
	}

	if err := c.communityRepo.ChangePlayerNumGames(
		event.CommunityId,
		event.PlayerId,
		numGames,
	); err != nil {
		return err
	}
	return nil
}

func (c *commandProcessor) TemporaryRegistration(
	name user.UserName,
	email user.UserEmail,
	password user.UserPassword,
) (string, error) {
	confirmPass := user.NewUserConfirmPass()
	status, _ := user.NewUserStatus(user.Inactive)
	id := user.NewUserId()
	password.Encrypt()
	if err := c.userRepo.CreateUser(
		id,
		name,
		email,
		password,
		confirmPass,
		status,
	); err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id.Value(),
		"exp":     time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	c.emailServer.SendEmail(
		[]string{email.Value()},
		"確認コードのお知らせ",
		`
		<p>Badministへ登録いただきありがとうございます</p>
		<p>以下の確認コードをアプリに入力して登録を完了してください</p>
		<p>確認コード: `+confirmPass.Value()+`</p>
		`,
	)

	return tokenString, nil
}

func (c *commandProcessor) ActivateUser(
	email user.UserEmail,
	confirmPass user.UserConfirmPass,
) error {
	user, err := c.userRepo.FindUserByEmail(email)
	if err != nil {
		return err
	}
	if !user.CompareConfirmPass(confirmPass) {
		return errors.New("確認用パスワードが間違っています")
	}
	event, err := user.Activate()
	if err != nil {
		return err
	}

	if err := c.userRepo.ActivateUser(event.UserId); err != nil {
		return err
	}

	return nil
}

func (c *commandProcessor) Login(
	email user.UserEmail,
	password user.UserPassword,
) (string, error) {
	userDomain, err := c.userRepo.FindUserByEmail(email)
	if err != nil {
		return "", err
	}
	if !userDomain.Authenticate(password) {
		return "", errors.New("メールアドレスかパスワードが間違っています")
	}
	// jwtの発行
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userDomain.BreachEncapsulationOfId().Value(),
		"exp":     time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	if userDomain.BreachEncapsulationOfStatus().Value() == user.Inactive.String() {
		c.emailServer.SendEmail(
			[]string{email.Value()},
			"確認コードのお知らせ",
			`
			<p>Badministへ登録いただきありがとうございます</p>
			<p>以下の確認コードをアプリに入力して登録を完了してください</p>
			<p>確認コード: `+userDomain.BreachEncapsulationOfConfirmPass().Value()+`</p>
			`,
		)
	}
	return tokenString, nil
}

func (c *commandProcessor) ReissueConfirmPass(
	id user.UserId,
) error {
	user, err := c.userRepo.FindUserById(id)
	if err != nil {
		return err
	}
	event, err := user.ReissueConfirmPass()
	if err != nil {
		return err
	}
	if err := c.userRepo.ReissueConfirmPass(
		event.UserId,
		event.ConfirmPass,
	); err != nil {
		return err
	}
	c.emailServer.SendEmail(
		[]string{user.BreachEncapsulationOfEmail().Value()},
		"確認コードのお知らせ",
		`
		<p>Badministへ登録いただきありがとうございます</p>
		<p>以下の確認コードをアプリに入力して登録を完了してください</p>
		<p>確認コード: `+user.BreachEncapsulationOfConfirmPass().Value()+`</p>
		`,
	)
	return nil
}
