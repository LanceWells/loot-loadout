import React, {
  MouseEventHandler,
} from 'react';
import useWebSocket from 'react-use-websocket';
export type ChatRoomProps = {}

/**
 * sdf
 * @param props sdf
 * @returns sdf
 */
export default function ChatRoom(props: ChatRoomProps) {
  // const socket: WebSocket = useMemo(() => {
  //   const socket = new WebSocket('ws://localhost/ws');
  //   socket.addEventListener('open', (event) => {
  //     socket.send('Hello!');
  //   });
  //   return socket;
  // }, []);

  const socket = useWebSocket('ws://localhost:80/ws');

  const onButtonClick: MouseEventHandler = (e) => {
    socket.sendMessage('Test');
    // socket.send('TestMessage');
  };

  return (
    <div>
      <button onClick={onButtonClick}>
        CLICK ME
      </button>
    </div>
  );
}
