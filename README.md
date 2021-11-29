# Go Terminal Chat

Basic Got chat program using Ably for networking.
Taken from [GopherCon UK 2021: Tom Camp - Creating a basic chat app](https://www.youtube.com/watch?v=StNWEEfV_1s).

## Setup
Replace the **YOUR_API_KEY** text in `.env.example` with your Ably API key and re-name the file to `.env`.
Sign up for an account with Ably and access your API key from the app dashboard.
We keep the API key in .env and ignore it in .gitignore to avoid accidentally sharing the API key.


Next, run the application with the following line, specifying your ClientID as an argument, and you're ready to start communicating over Ably!

```
~ $ go run chat.go YOUR_USERNAME
```