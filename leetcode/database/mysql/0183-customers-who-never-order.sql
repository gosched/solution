select name as Customers
from customers
where customers.Id not in (select CustomerId from Orders);