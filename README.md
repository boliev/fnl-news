# fnl-news
## The story
As a big fan of the Alania Vladikavkaz football club, which plays in the second Russian league (FNL), I want to get all news about the team from all sports sites in one place. I guess, there are other fans, who want to get news about their teams in FNL. So the FNL-news is a small service, which parses several most popular Russian sports sites, added team-tags to the articles (#Alania if FC Alania mentioned in an article, or #Torpedo if FC Torpedo mentioned), and posts the links in a telegram channel.
## Run
### Config
The service needs PostgreSQL database DSN, telegram bot access token, and telegram channel id. Default values are in the `.env` files, which can be replaced by values from `.env.local`. Just copy the file and change the values
```
cp .env .env.local
```
### Build
```shell
go build -o fnl_news cmd/main.go
```
### Run
```shell
./fnl_news
```
## Database
I was trying to keep it light and simple, so there is only one table in the database: `articles`. It will be created when you first start the service if the database DSN is correct.

If you don't have installed PostgreSQL, you can install it from docker.
```shell
docker compose up
```
will create a container with PostgreSQL database. The DSN for the database
```
DATABASE_DSN=host=localhost user=fnluser password=123456 dbname=fnl port=5432 sslmode=disable
```
Note that there is no service container, only database. Use it only for developing, not in production.
## Modules and packages
### Modules
 - [Resty](https://github.com/go-resty/resty) for HTTP requests. To get articles and send links to the channel.
 - [Logrus](https://github.com/sirupsen/logrus) as a logger
 - [Viper](https://github.com/spf13/viper) for dealing with config files
 - [dig](https://github.com/uber-go/dig) for dependency injection
 - [GORM](https://github.com/go-gorm/gorm/) as an ORM
 
### Packages
 - `pkg/config` - Wrapper for the Viper module
 - `pkg/httpclient` - Wrapper for the Resty module

### Internal packages
 - `internal/domain` - Domain entities. Only article entity for now.
 - `internal/parser` - The parser. Parses sources and returns the articles list. Also contains `Tag Matcher` which defines team-tags for an article
 - `internal/publisher` - The publishers. There is only the Telegram publisher for now. All Publishers must implement the `publisher.Publisher` interface.
 - `internal/repository` - Article Repository. Knows how to query articles from the database.
 - `internal/source` - The sport sites sources. All must implement `source.Source` interface. Each Source contains regular expressions for extracting data from its HTML, and methods to modify and check the data.

## Result
[Telegram channel](https://t.me/FNL_News)

## Todo
 - Domain entity `Article` depends on `Gorm` module as far as contains gorm descriptions for the fields.
 - Parsers for different sources are described manually (`internal/parser/ParsersFactory`), maybe it is a good idea to automate it somehow
 - ArticleRepository knows how to mark an article as sent for a concrete publisher(Telegram publisher). I guess we don't want to edit the article repository every time when a new publisher was added.
 - `httpClientResty` and `httpClientResty1251` have similar Post method.
 - Add tests