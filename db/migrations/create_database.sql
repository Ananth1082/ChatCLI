CREATE TABLE
  sessions (
    session_id CHAR(10),
    created_at TIMESTAMP,
    last_active_at TIMESTAMP,
    PRIMARY KEY (session_id)
  );

CREATE TABLE
  chatrooms (
    chatroom_id CHAR(10) PRIMARY KEY,
    name TEXT,
    created_at TIMESTAMP
  );

CREATE TABLE
  members (
    chatroom_id CHAR(10),
    session_id CHAR(10),
    joined_at TIMESTAMP,
    PRIMARY KEY (chatroom_id, session_id),
    FOREIGN KEY (chatroom_id) REFERENCES chatrooms (chatroom_id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (session_id) REFERENCES sessions (session_id) ON UPDATE CASCADE ON DELETE CASCADE
  );

CREATE TABLE
  messages (
    message_id CHAR(10),
    chatroom_id CHAR(10),
    session_id CHAR(10),
    content TEXT,
    sent_at TIMESTAMP,
    PRIMARY KEY (message_id),
    FOREIGN KEY (chatroom_id) REFERENCES chatrooms (chatroom_id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (session_id) REFERENCES sessions (session_id) ON UPDATE CASCADE ON DELETE CASCADE
  );