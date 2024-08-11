CREATE TABLE
  sessions (
    user_name VARCHAR(50) PRIMARY KEY,
    local_ip VARCHAR(20),
    created_at TIMESTAMP,
    is_active INT CHECK (is_active in (0,1)) DEFAULT 1,
    last_active_at TIMESTAMP
  );

CREATE TABLE
  chatrooms (
    chatroom_name TEXT PRIMARY KEY,
    created_at TIMESTAMP
  );

CREATE TABLE
  members (
    chatroom_id CHAR(36),
    session_id CHAR(36),
    joined_at TIMESTAMP,
    PRIMARY KEY (chatroom_id, session_id),
    FOREIGN KEY (chatroom_id) REFERENCES chatrooms (chatroom_name) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (session_id) REFERENCES sessions (user_name) ON UPDATE CASCADE ON DELETE CASCADE
  );

CREATE TABLE
  messages (
    message_id CHAR(36),
    chatroom_name CHAR(36),
    user_name CHAR(36),
    content TEXT,
    sent_at TIMESTAMP,
    PRIMARY KEY (message_id),
    FOREIGN KEY (chatroom_name) REFERENCES chatrooms (chatroom_name) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (user_name) REFERENCES sessions (user_name) ON UPDATE CASCADE ON DELETE CASCADE
  );

INSERT INTO chatrooms values("lobby",CURRENT_TIMESTAMP);