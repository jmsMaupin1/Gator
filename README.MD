<p align="center">
<img src="assets/gator_logo.webp" width="200"/>
</p>

# gator

A multi-user command line tool for aggregating RSS feeds and browsing posts

# Prerequisites

- [Go toolchain](https://golang.org/dl/) 
- [local PostgreSQL database](https://www.postgresql.org/download/)
- [Goose database migration tool](https://github.com/pressly/goose) 

As a note for windows users, I highly recommend you set up `WSL2`, you can find more info [here](https://learn.microsoft.com/en-us/windows/wsl/install) going forward I will assume you have set up `WSL2`, you should follow the Linux portion of the instructions during setup. If you dont want to use `WSL2` you will likely just need to do a bit of research on how to do each of these steps independently.

## MacOS with Brew

```bash
brew install postgresql@15
```

## Linux / WSL (Debian). [Here are the docs from Microsoft](https://learn.microsoft.com/en-us/windows/wsl/tutorials/wsl-database#install-postgresql), but simply:

```bash
sudo apt update
sudo apt install postgresql postgresql-contrib
```

## Ensure the installation worked. The psql command-line utility is the default client for Postgres. Use it to make sure you're on version 15+ of Postgres:

```bash
psql --version
```

## (Linux only) Update postgres password:

```bash
sudo passwd postgres
```

# Start the Postgres server in the background

## For Mac:

```bash
brew services start postgresql
```

## For Linux

```bash
sudo service postgresql start
```
# Setting up the Database

Enter the psql shell

## For Mac:

```bash
psql postgres
```

## For Linux

```bash
sudo -u postgres psql
```

You should see a new prompt that looks like:

```bash
postgres=#
```

# Now create a new database

Note: you can choose any name you want for the database, it doesnt need to be gator, make sure you keep track of the name for connecting to the database as well as construcing your psql connection string (more on this further below)

```bash
CREATE DATABASE gator;
```

# Connect to the database

```bash
\c gator
```

After connecting you should see a prompt that looks like this:

```bash
gator=#
```

## Linux Users

For linux users we need to set the user password 
You can use whatever password you like here, just make sure to keep track of this for the connection string (more on this further below)

```bash
ALTER USER postgres PASSWORD 'postgres';
```

# Constructing your connection string
#
A connection string is just a URL with all of the information needed to connect to a database. The format is:

`protocol://username:password@host:port/database`

- protocal: we will use `postgres`
- username: this is the username you set up for your database when you created it earlier, if you dont remember it, connect to the database and type `\conninfo`
- password: this is the password you might have created in the previous step, if you did not then we will just leave this blank.
- host: for local dev this will likely be set to `localhost`
- port: the port the database is set up on, typically postgres defaults to `5432`
- database: this will be your database name you created earlier.

Example of what your connection string might look like:

## For Mac users (no password set)

```
"postgres://jmsGears:@localhost:5432/gator"
```

## For Linux users (or people who added a password)

```
"postgres://jmsGears:postgres@localhost:5432/gator"
```

# Database Migrations

This will set up our database tables and relations, you will need the connection string you created in the last step

Open up your terminal and navigate to `sql/schema` and run:

```bash
goose postgres "<Connection String>" up
```

If you want to reset the database you can (make sure you are in the correct directory):

```bash
goose postgres "<Connection String>" down
```

# Installation (optional step)

Install `gator` with (this might take a minute or two depending on your system): 

```bash
go install ...
```

Note: if you dont want to install, and you have the go toolchain you can run this program by being in the root directory of the program by doing:
I will go into the commands and arguments available towards the end of this readme.

```bash
go run . <commands> <arguments>
```

# Configuration

We will need to do a bit of configuration before actually running this application. We need to create a file named `.gatorconfig.json` in your home directory:

## For Linux and Mac useres

```bash
touch ~/.gatorconfig.json
```

Next edit that file and put:

```json
{
	"db_url": "<Connection String>?sslmode=disable"
}
```

## Examaple for Mac users (or people who dont have a password)

```json
{
	"db_url": "postgres://jmsGears:@localhost:5432/gator?sslmode=disable"
}
```

## Example for Linux users (or people who have set a password)

```json
{
	"db_url": "postgres://jmsGears:postgresql@localhost:5432/gator?sslmode=disable"
}
```

# Usage

Remember if you skipped the installation step above, make sure you are in the root directory of this project and substitue `gator` with `go run .`

## Create a new user (this will also log the created user in)

```bash
gator register <user_name>
```

## Log a user in (certain commands require a user to be logged in)

```bash
gator login <user_name>
```

## List all users currently registered

```bash
gator users
```

## Deletes all current users

```bash
gator reset
```

## List all feeds currently subscribed to by all current users

```bash
gator feeds
```

## Start the aggregator 

```bash
gator agg <interval>
```
this will grab new posts by all users feeds once every `interval`

Intervals can be entered in the format `00h00m00s` You dont need the entire string, just the parts that are non 0

Example, run the aggregator to grab posts every 30 minutes:

```bash
gator agg 30m
```

## Add a feed (this will automatically set the currently logged in users to follow this feed)

```bash
gator addfeed <name> <feed_url>
```

- name: the name of the feed, this can be whatever you want
- feed_url: this is the url to the RSS feed you want to subscribe to

## View all of your posts (Requires a user to be logged in)

```bash
gator browse <limit>
```

limit is a number, this will return the most recent <limit> number of posts from all feeds the currently logged in user is subscribed to

## Follow a feed (Requires a user to be logged in)

```bash
gator follow <feed_url>
```

- feed_url: this is the url to the RSS feed you want to follow

## Unfollow a feed (Requires a user to be logged in)

```bash
gator unfollow <feed_url>
```
- feed_url: this is the url to the RSS feed you want to unfollow

## Get a list of feeds current user is following (Requires a user to be logged in)

```bash
gator following
```
