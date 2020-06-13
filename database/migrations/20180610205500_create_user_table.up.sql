create table `user` (
    `id` int(11) PRIMARY KEY AUTO_INCREMENT,
    `first_name` varchar(256) NOT NULL,
    `last_name` varchar(256) NOT NULL,
    `date_of_birth` datetime NOT NULL
);