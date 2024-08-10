CREATE TABLE
  sessions (
    session_id CHAR(36),
    user_name VARCHAR(50),
    local_ip VARCHAR(20),
    created_at TIMESTAMP,
    last_active_at TIMESTAMP,
    PRIMARY KEY (session_id)
  );

CREATE TABLE
  chatrooms (
    chatroom_id CHAR(36) PRIMARY KEY,
    name TEXT,
    created_at TIMESTAMP
  );

CREATE TABLE
  members (
    chatroom_id CHAR(36),
    session_id CHAR(36),
    joined_at TIMESTAMP,
    PRIMARY KEY (chatroom_id, session_id),
    FOREIGN KEY (chatroom_id) REFERENCES chatrooms (chatroom_id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (session_id) REFERENCES sessions (session_id) ON UPDATE CASCADE ON DELETE CASCADE
  );

CREATE TABLE
  messages (
    message_id CHAR(36),
    chatroom_id CHAR(36),
    session_id CHAR(36),
    content TEXT,
    sent_at TIMESTAMP,
    PRIMARY KEY (message_id),
    FOREIGN KEY (chatroom_id) REFERENCES chatrooms (chatroom_id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (session_id) REFERENCES sessions (session_id) ON UPDATE CASCADE ON DELETE CASCADE
  );

INSERT INTO chatrooms values("000000000000000000000000000000000000","Lobby",CURRENT_TIMESTAMP);