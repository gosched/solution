SELECT Person.FirstName, Person.LastName, Address.City, Address.State
FROM Person LEFT JOIN Address
ON Person.PersonId = Address.PersonId;

SELECT P.FirstName, P.LastName, A.City, A.State
FROM Person P LEFT JOIN Address A
ON P.PersonId = A.PersonId;

SELECT P.FirstName, P.LastName, A.City, A.State
FROM Person P LEFT JOIN (SELECT PersonId, City, State FROM Address) A
ON P.PersonId = A.PersonId;

SELECT P.FirstName, P.LastName, A.City, A.State
FROM Person P LEFT JOIN (SELECT DISTINCT PersonId, City, State FROM Address) A
ON P.PersonId = A.PersonId;