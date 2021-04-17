#!/bin/sh
docker run --name=mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root -d mysql/mysql-server
sleep 10
docker exec mysql mysql -u root -proot mysql -e "UPDATE user SET host='%' WHERE user='root';FLUSH PRIVILEGES;SELECT host, user FROM user; CREATE DATABASE IF NOT EXISTS users_db; USE users_db; CREATE TABLE IF NOT EXISTS users(id INT NOT NULL AUTO_INCREMENT PRIMARY KEY, first_name VARCHAR(100), last_name VARCHAR(100), email VARCHAR(200) NOT NULL, date_created DATE); INSERT INTO users_db.users (first_name, last_name, email, date_created) VALUES ('test', 'test', 'test@gmail.com', '2021-04-17'); SELECT * FROM users;"