package logic

import messages "github.com/mikkoryynanen/real-time/generated/proto"

type Vector2 struct {
	X	float32
	Y	float32
}
type Logic struct {
	clientPositions		map[string]Vector2
}

func NewLogic() *Logic {
	return &Logic{}
}

func (l *Logic) HandleLogic(input *messages.ClientInputRequest, sessionId string) {
	if position, exists := l.clientPositions[sessionId]; exists {
		// TODO Add the actual moving logic
		// TODO there should also be constantly running loop that would send the client every X seconds/milliseconds all of the client positions
		position.X = input.Input.X
		position.Y = input.Input.Y
	}
}

func (l *Logic) UpdateState() {
	
}