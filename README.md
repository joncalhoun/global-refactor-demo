# Global Refactor Demo

This requires a database and in that database it needs a table like the following:

```sql
CREATE TABLE tweets (
  id SERIAL,
  creator_id INT,
  content VARCHAR(255)
);
```

This was used as part of an email to my mailing list at <https://signup.calhoun.io/>. I'll probably also publish it to my website <https://www.calhoun.io/> at some point, but it isn't there at this time.
