package main

// 메시지 허브 만들기
// 모든 들어오는 메시지를 받아서 연결된 Client로 쏜다!
type hub struct {
	clients    map[*client]bool
	broadcast  chan string
	register   chan *client
	unregister chan *client

	content string
}

// Hub에 몇몇 채널을 정의하는데, 이경우 비동기적으로 처리가능하다
// register와 unregister 이벤트 둘다, 또는 브로드캐스팅 메시지와도
// 끝으로, 인스턴스를 생성한다.
var h = hub{
	broadcast:  make(chan string),
	register:   make(chan *client),
	unregister: make(chan *client),
	clients:    make(map[*client]bool),
	content:    "",
}

// 여기서 go 채널의 이점을 갖는다.
// 채널은 FIFO스택과 같다. Client는 요청을 이들 채널로부터 하나에 저장하고
//  고루틴은 가능한 빨리 이들을 언스택한다.
func (h *hub) run() {
	for {
		select {
		// 이는 h.register 채널로 부터 들어오는 값을 받는다.
		case c := <-h.register:
			h.clients[c] = true
			// []byte는 캐스팅을 위해 사용한것이다.
			c.send <- []byte(h.content)
			break
		// 만약 고객이 unregister 이벤트를 전송하면,
		// 우리는 그 채널에 접근해 Hub 커넥션으로 부터제거하면 된다
		case c := <-h.unregister:
			_, ok := h.clients[c]
			if ok {
				delete(h.clients, c)
				close(c.send)
			}
			break
		// 최종적으로 만약 우리가 한 고객으로 부터 broadcast채널을 통해
		// 메시지를 받는다면 우리는단지 Hub 컨텐트를 업데이트한다.
		// 그리고 브로드 캐스트 메시지는 모든 클라이언트로 전송된다.
		case m := <-h.broadcast:
			h.content = m
			h.broadcastMessage()
			break
		}
	}
}

// Do not forget the break here.
// This is the reason why we exceeeded the deadline of our hackday,
// a missing break. Thus, after broadcasting first message,
// websocket closed itself unexpectedly, with a very vague Chrome error message.
// This is the disadvantage of extreme programming:
// wanting to go still faster cause big time loss on trivial errors.

func (h *hub) broadcastMessage() {
	for c := range h.clients {
		select {
		case c.send <- []byte(h.content):
			// go에서는 break하는 것을 잊지말자!
			break

		// 만약 클라이언트에 도달할 수 없을 경우
		default:
			close(c.send)
			delete(h.clients, c)
		}
	}
}
