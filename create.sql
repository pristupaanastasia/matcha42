CREATE TABLE chat (
        id_chat         SERIAL         NOT NULL,
        user_one_id         int         NOT NULL,
        user_two_id         int         NOT NULL,
        PRIMARY KEY (id_chat)
);
CREATE TABLE message (
        id_chat       int  REFERENCES chat NOT NULL,
        id_message    SERIAL         NOT NULL,
        sender_id         int         NOT NULL,
        receiver_id         int         NOT NULL,
        message         char(250)   NOT NULL,
        time         time           NOT NULL,
        receiver_viewed int         NOT NULL,
        PRIMARY KEY (id_message)
);

CREATE TABLE users (
        id_user         SERIAL         NOT NULL,
        login         char(64)         NOT NULL,
        password         char(80)         NOT NULL,
        first_name   char(150)         NOT NULL,
        last_name    char(150)         NOT NULL,
        PRIMARY KEY (id_user)
);

CREATE TABLE profile (
        id_user   int      REFERENCES users NOT NULL,
        age         int,
        image         char(150)[] ,
        description   char(400) ,
        sex    int NOT NULL,
        tags char(50)[],
        gps char(150) NOT NULL,
        fame_rating int NOT NULL,
        online int,
        PRIMARY KEY (id_user)
);
CREATE TABLE connect (
        id_user   int      REFERENCES users NOT NULL,
        like_user int[],
        liked int[],
        block int[],
        blocked int[],
        history int[],
        PRIMARY KEY (id_user)
);
