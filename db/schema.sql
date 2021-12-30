-- Create tables --
DROP TABLE IF EXISTS InterestsToUsers;
DROP TABLE IF EXISTS Forums;
DROP TABLE IF EXISTS Interests;
DROP TABLE IF EXISTS Users;

CREATE TABLE Interests
(
  id INT PRIMARY KEY,
  Name VARCHAR(50) NOT NULL UNIQUE,
);

CREATE TABLE Forums
(
  id INT PRIMARY KEY,
  Name NVARCHAR(50) NOT NULL UNIQUE,
  InterestId INT NOT NULL FOREIGN KEY REFERENCES Interests(id)
);

CREATE TABLE Users
(
  id INT PRIMARY KEY,
  UserName VARCHAR(50) NOT NULL UNIQUE,
);

CREATE TABLE InterestsToUsers
(
  UserId INT NOT NULL FOREIGN KEY REFERENCES Users(id),
  InterestId INT NOT NULL FOREIGN KEY REFERENCES Interests(id),
  PRIMARY KEY (UserId, InterestId)
);
