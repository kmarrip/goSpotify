# goSpotify
[![GitHub license](https://img.shields.io/github/license/chaithanyaMarripati/goSpotify)](https://github.com/chaithanyaMarripati/goSpotify/blob/main/LICENSE)
[![GitHub issues](https://img.shields.io/github/issues/chaithanyaMarripati/goSpotify)](https://github.com/chaithanyaMarripati/goSpotify/issues)
### goSpotify uses Oauth2 to authorize the users with their spotify account and returns the song/album currently being played

## join [goSpotify](https://discord.gg/vQUaPhbp) discord server

# QuickStart

This project makes use of git hooks, before committing execute this line: 
```
git config core.hooksPath .githooks
```

1. [Install Docker](https://docs.docker.com/engine/install/).
2. Go to https://developer.spotify.com/ and login/register
3. In dashboard add a new app, so in `edit settings` add this redirect url http://localhost:8080/callback
4. Get your client id and client secret and create .env file like this: 
   ```
      baseUrl=https://accounts.spotify.com/authorize?
      clientId=your-client-id
      clientSecret=your-client-secret
      redirectUrl=http://localhost:8080/callback
      scopes=user-read-playback-state
      tokenUrl=https://accounts.spotify.com/api/token
      getMeSpotify=https://api.spotify.com/v1/me
      currentlyPlaying=https://api.spotify.com/v1/me/player/
   ```
5. Run the build command
   ```
   docker build -t gospotify:latest .
   ```  
6. Run the start command
   ```
   docker run -p 8080:8080 gospotify:latest
   ```


# Todo list 
these are the pending action items, list will be updated overtime
- [x] Add state in the url parmas when the user is redirected, to protect against csrf attacks
- [ ] After successful authentication return a html page instead of text
- [x] Add error page for errors during /authorize step 
- [x] Add more test cases
- [ ] Add more functionality to the app
- [ ] Buy a domain and add it to the cloud run

# Deployment
CI/CD pipeline is configured in gcp, google cloud build processes the builds and deploys it to the cloud run as per cloudbuild.yaml file.
This app is currently deployed to cloud run
[app](https://gospotify-sjskww6rpa-el.a.run.app/)
