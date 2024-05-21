# Showdown Match Data

## Description
Small API to fetch results based on user queries from [Pokemon Showdown](https://pokemonshowdown.com), built using [Go Fiber](https://gofiber.io/)

## Routes

User:
 - `/api/user?user=user1`: Fetch user specified.

Replays:
 - `/api/replays?user=user1`: Fetch specified user's replays.
 - `/api/replays/users?user=user1&user2=user`: Fetch replays of matches between given users.
 - `/api/replays/format?format=gen9ou`: Fetch replays of matches in given format. 
 - `/api/replays/user-format?user=user1&format=format=gen9ou`: Fetches replays of user in given format.

Ladder:
 - `/api/ladder?format=gen9ou`: Fetch given format's ladder data.

Dex:
 - `/api/dex`: Fetch Pokedex information.
 - `/api/moves`: Fetch Moves information.
