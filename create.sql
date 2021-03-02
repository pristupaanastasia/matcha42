
CREATE TABLE chat (
        id_chat         char(50)         NOT NULL,
        user_one_id         int         NOT NULL,
        user_two_id         int         NOT NULL,
        PRIMARY KEY (id_chat)
);

CREATE TABLE message (
        id_chat       char(50)  REFERENCES chat NOT NULL,
        id_message    SERIAL         NOT NULL,
        sender_id         int         NOT NULL,
        receiver_id         int         NOT NULL,
        message         char(250)   NOT NULL,
        time         time           NOT NULL,
        receiver_viewed int         NOT NULL,
        PRIMARY KEY (id_message)
);

CREATE TABLE users (
        id_user         char(50)         NOT NULL,
        email         char(60)         NOT NULL,
        login         char(64)         NOT NULL,
        password         char(80)         NOT NULL,
        first_name   char(150)         NOT NULL,
        last_name    char(150)         NOT NULL,
        verify         boolean,
        PRIMARY KEY (id_user)
);

CREATE TABLE profile (
        id_user  char(50)     NOT NULL,
        age         int,
        gender      boolean NOT NULL ,
        description   char(400) ,
        sex    int NOT NULL,
        gps char(150) NOT NULL,
        fame_rating int NOT NULL,
        online boolean,
        FOREIGN KEY (id_user) REFERENCES users (id_user) ON DELETE CASCADE
);

CREATE TABLE image (
    id_user  char(50)     NOT NULL,
    main_image char(150) NOT NULL,
    image_two char(150) ,
    image_tree char(150) ,
    image_four char(150) ,
    image_five char(150) ,
    FOREIGN KEY (id_user) REFERENCES users (id_user) ON DELETE CASCADE
);


CREATE TABLE block (
    id SERIAL NOT NULL,
    user_one_id char(50) REFERENCES users (id_user),
    user_two_id char(50) REFERENCES users (id_user),
    PRIMARY KEY (id)
);

CREATE TABLE like_user (
    id SERIAL NOT NULL,
    user_one_id char(50) REFERENCES users (id_user),
    user_two_id char(50) REFERENCES users (id_user),
    PRIMARY KEY (id)
);

CREATE TABLE history (
    id SERIAL NOT NULL,
    user_one_id char(50) REFERENCES users (id_user),
    user_two_id char(50) REFERENCES users (id_user),
    PRIMARY KEY (id)
);

CREATE TABLE user_session (
    id_user   char(50)      REFERENCES users NOT NULL,
    session_key char(500) NOT NULL,
    login_time time NOT NULL,
    last_seen_time time NOT NULL,
    PRIMARY KEY (id_user)
);

CREATE TABLE tag_and_profile (
    id SERIAL NOT NULL,
    id_user char(50) REFERENCES users NOT NULL,
    id_tag int REFERENCES tag NOT NULL,
    PRIMARY KEY (id)
);


CREATE TABLE tag (
    id_tag SERIAL NOT NULL,
    tag_name char(50) NOT NULL,
    PRIMARY KEY (id_tag)
);
