package message_types

type MessageType int

type BaseMessage struct {
	Type MessageType `json:"type"`
}

type PlayerInput struct {
	BaseMessage

	X int16 `json:"x"`
	Y int16 `json:"y"`
}

type GameState struct {
	BaseMessage

	PlayerPositionX int16 `json:"player_position_x"`
	PlayerPositionY int16 `json:"player_position_y"`
}
