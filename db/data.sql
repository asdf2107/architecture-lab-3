-- Insert demo data --
insert into Interests (id, Name) values
(0, 'it'),
(1, 'sport'),
(2, 'food'),
(3, 'ukraine-politics'),
(4, 'world-politics');

insert into Forums (id, Name, InterestId) values
(0, N'Сучасні IT технології', 0),
(1, N'Здорова їжа', 2),
(2, N'Політика в Україні', 3);

insert into Users (id, UserName) values
(0, 'petro1977'),
(1, 'lenskiy_the_great');

insert into InterestsToUsers (UserId, InterestId) values
(0, 0),
(0, 1),
(0, 3),
(1, 0),
(1, 1),
(1, 2),
(1, 3),
(1, 4);

insert into UsersToForums (UserId, ForumId) values
(0, 0),
(0, 1),
(0, 2),
(1, 0),
(1, 1),
(1, 2);
