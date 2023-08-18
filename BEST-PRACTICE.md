# Memar Best Practice

### Developing hints
- Don't think about network when you develope a business service logic. Use `sk protocol.Socket` as stream data not network stream and don't remove it even don't need it from `Process` method arguments.

## Protocols
- [IETF working groups](https://datatracker.ietf.org/wg/)
- You can find protocol descriptions in its directory as now [protocol](./protocol/), [ISO](./iso/)    
Read more about each protocol or library in its [RFC](https://github.com/GeniusesGroup/RFCs)
As [suggest here](https://github.com/golang/go/issues/48087) to comply with the standards we add [protocol](./protocol) package and all other memar packages implement this package. You can implement these protocols in your own way if our packages can't satisfied you or your requirements.   
A standard is a formalized protocol accepted by most of the parties that implement it. A protocol is not a set of rules. A protocol is the thing those rules describe the rules of. This is why programs implement a protocol and comply with a standard.

## Optimize memory allocation and de-allocation
- use `go build -gcflags=-m {{file-name}}.go` to determine `heap escape`

### build tags
Some functionality in files that have build tags `//go:build tag_name`

## Contribute Rules
- Write benchmarks and tests codes in different files as `{{file-name}}_test.go` and `{{file-name}}_bench_test.go`
- Naming by [common naming convention](https://en.wikipedia.org/wiki/Naming_convention_(programming))

## Structure methods
All structure declare in memar should have these methods:
- **Init**: initialize call just after an object allocate.
- **Reinit**: re-initialize call when allocated object want to reuse immediately or pass to a pool to reuse later. It will prevent memory leak by remove any references in the object.
- **Deinit**: de-initialize call just before an object want to de-allocated (GC).

All structure declare in memar can have these methods too:
- **Alloc**: call to allocate the object. It must be called on pointer to the struct not direct use e.g. `t *embed`.
- **Dealloc**: call to de-allocate the object. It must be called on pointer to the struct not direct use e.g. `t *embed`.
- **clone**:
- **open**:
- **reset**:
- **close**:

## Rules
- Use `var` keyword when declare new variable instead of `:=`. Just use `:=` in `for`, `range`, `switch .(type)`
- Write one logic per line and avoid to use `;` to write multiple logic in one line e.g.
```go
    if err := Logic(); err != nil {
        return
    }

    var err = Logic()
    if err != nil {
        return
    }
```
- Don't access fields directly (`structName.field1`) in any way, just use methods to access structure fields even inside the encapsulation.
- Don't use `panic` and `defer` in any way. return `protocol.Error` if you think caller need to get any error.

## GIT
Git is not the best version control mechanism for a software project, but it is the most use one.

### Some useful commands
- Make project version control by `git init`
- Clone exiting repo by `git clone ${repository path}`
- **get just the latest revision of a project:** `git clone --depth 1 https://github.com/user/repo.git`
- Add memar to project as submodule by ```git submodule add -b master https://github.com/GeniusesGroup/memar-go```
- Clone existing project with just latest commits not all one ```git clone ${repository path} --recursive --shallow-submodules```
- Change memar version by ```git checkout tag/${tag}``` or update by ```git submodule update -f --init --remote --checkout --recursive``` if needed.

## Methods name in relation to SQL(Structure Query Language)
You can read all [SQL syntax here](https://www.w3schools.com/sql/).

- `SELECT`:
    - `GET`: get primary key and return full record data
    - `FIND`: get an index key and return primary keys
    - `FILTER`: get at least one key that one of them is not index and must do some logic in runtime and usually(not always) return primary keys.
    - `Search`: text search

- `SELECT DISTINCT`:
    - `LIST`: is same as `DISTINCT` but force storage engine to index to speed up quires. Service get an index key and usually return list of other keys.

- `INSERT INTO`:
    - `SAVE` because storage layer don't know about the structure of the data model.

```sql
INSERT INTO table_name (column1, column2, column3, ...)
VALUES (value1, value2, value3, ...);
```

- `UPDATE`:
```sql
UPDATE table_name
SET column1 = value1, column2 = value2, ...
WHERE condition;
```

- `DELETE`:
```sql
DELETE FROM table_name WHERE condition;
```

- `LOCK`:
- `UNLOCK`:

- `WHERE`: We use `BY` because the service conditions declared and just conditions values change in each service call.
    - `NOT` Operator:
        ```sql
        SELECT column1, column2, ...
        FROM table_name
        WHERE NOT condition;

        SELECT * FROM Customers
        WHERE NOT Country='Germany';
        ```
    - `IN`
    - `BETWEEN`:
        ```sql
        SELECT column_name(s)
        FROM table_name
        WHERE column_name BETWEEN value1 AND value2;
        ``` 
    - `Range`: min{fieldName}, max{fieldName}
    - 
    - `=`	Equal
    - `>`	Greater than	
    - `<`	Less than	
    - `>=`	Greater than or equal	
    - `<=`	Less than or equal	
    - `<>`	Not equal. Note: In some versions of SQL this operator may be written as !=	
    - `BETWEEN`   Between a certain range	
    - `LIKE` use to search for a pattern
        ```sql
        SELECT column1, column2, ...
        FROM table_name
        WHERE columnN LIKE pattern;
        ```
    - `IN`	To specify multiple possible values for a column
    - `AND` Operator:
        ```sql
        SELECT column1, column2, ...
        FROM table_name
        WHERE condition1 AND condition2 AND condition3 ...;

        SELECT * FROM Customers
        WHERE Country='Germany' AND City='Berlin';
        ```
    - `OR` Operator: 
        ```sql
        SELECT column1, column2, ...
        FROM table_name
        WHERE condition1 OR condition2 OR condition3 ...;

        SELECT * FROM Customers
        WHERE City='Berlin' OR City='München';

        SELECT * FROM Customers
        WHERE Country='Germany' OR Country='Spain';
        ```
    - Combining `AND`, `OR` and `NOT` Operators:
        ```sql
        SELECT * FROM Customers
        WHERE Country='Germany' AND (City='Berlin' OR City='München');

        SELECT * FROM Customers
        WHERE NOT Country='Germany' AND NOT Country='USA';
        ```
    - `ORDER BY`
        ```sql
        SELECT column1, column2, ...
        FROM table_name
        ORDER BY column1, column2, ... ASC|DESC;
    - `TOP`: 
    - `LIMIT`: We use `offset`, `limit` as uint64
        ```sql
        -- SQL Server / MS Access Syntax:
        SELECT TOP number|percent column_name(s)
        FROM table_name
        WHERE condition;

        -- MySQL Syntax:
        SELECT column_name(s)
        FROM table_name
        WHERE condition
        LIMIT number;

        -- Oracle 12 Syntax:
        SELECT column_name(s)
        FROM table_name
        ORDER BY column_name(s)
        FETCH FIRST number ROWS ONLY;

        -- Older Oracle Syntax:
        SELECT column_name(s)
        FROM table_name
        WHERE ROWNUM <= number;

        -- Older Oracle Syntax (with ORDER BY):
        SELECT *
        FROM (SELECT column_name(s) FROM table_name ORDER BY column_name(s))
        WHERE ROWNUM <= number;
        ```

- `CURRENT`: in time series data we need some way to get last data with some condition.

- `MIN` & `MAX`
```sql
-- MIN() Syntax
SELECT MIN(column_name)
FROM table_name
WHERE condition;

-- MAX() Syntax
SELECT MAX(column_name)
FROM table_name
WHERE condition;
```

## Abbreviations & Definitions
- **UI**: (any) User Interface
    - **GUI**: Graphic User Interface
    - **VUI**: Voice User Interface
    - **CLI**: Command Line Interface
    - **CLA**: Command Line Arguments
- **Modules**: a kind of collection of packages
- **Packages**: a kind of collection of files
- **dp**: domain protocol
- **init**: initialize an object
- **reinit**: re-initialize to reuse an object
- **deinit**: de-initialize an object to ready for de allocation.
- **open**:
- **reset**:
- **close**:
- **SCM**: Source control management (SCM) systems provide a running history of code development and help to resolve conflicts when merging contributions from multiple sources.
- **DSL**: Data Storage Layer
- **DAL**: Data Access Layer
- **PL**: Persistence Layer
- **repo**: Repository
- **RDBMS**: Relational Database Management System
- **dm**: Data Model
- **dt**: Data Type
- **ss**: Storage Service
- **bs**: Business Service
- **BLL**(Business-Logic Layer) act as an intermediate between the Presentation Layer and the Data Access Layer (DAL).
- **ce**: character encoding
- **wg**: working group
- **da**: Data analytics is the process of analyzing raw data in order to draw out meaningful, actionable insights, which are then used to inform. Analysis is the division of a whole into small components, and analytics is the science of logical analysis. While analysis looks backward over time and works on the facts and figures of what has happened, analytics work towards modeling the future or predicting a result.
